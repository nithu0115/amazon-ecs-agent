// +build windows

// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package app

import (
	"context"
	"sync"
	"time"

	"github.com/aws/amazon-ecs-agent/agent/engine"
	"github.com/aws/amazon-ecs-agent/agent/sighandlers"
	"github.com/aws/amazon-ecs-agent/agent/statemanager"
	"github.com/cihub/seelog"
	"golang.org/x/sys/windows/svc"
)

const (
	//EcsSvcName is the name of the service
	EcsSvcName = "AmazonECS"
)

// startWindowsService runs the ECS agent as a Windows Service
func (agent *ecsAgent) startWindowsService() int {
	svc.Run(EcsSvcName, newHandler(agent))
	return 0
}

// handler implements https://godoc.org/golang.org/x/sys/windows/svc#Handler
type handler struct {
	ecsAgent agent
}

func newHandler(agent agent) *handler {
	return &handler{agent}
}

// Execute implements https://godoc.org/golang.org/x/sys/windows/svc#Handler
// The basic way that this implementation works is through two channels (representing the requests from Windows and the
// responses we're sending to Windows) and two goroutines (one for message processing with Windows and the other to
// actually run the agent).  Once we've set everything up and started both goroutines, we wait for either one to exit
// (the Windows goroutine will exit based on messages from Windows while the agent goroutine exits if the agent exits)
// and then cancel the other.  Once everything has stopped running, this function returns and the process exits.
func (h *handler) Execute(args []string, requests <-chan svc.ChangeRequest, responses chan<- svc.Status) (bool, uint32) {
	// channels for communication between goroutines
	ctx, cancel := context.WithCancel(context.Background())
	agentDone := make(chan struct{})
	windowsDone := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer close(windowsDone)
		defer wg.Done()
		h.handleWindowsRequests(ctx, requests, responses)
	}()

	go func() {
		defer close(agentDone)
		defer wg.Done()
		h.runAgent(ctx)
	}()

	// Wait until one of the goroutines is either told to stop or fails spectacularly.  Under normal conditions we will
	// be waiting here for a long time.
	select {
	case <-windowsDone:
		// Service was told to stop by the Windows API.  This happens either through manual intervention (i.e.,
		// "Stop-Service ECS") or through system shutdown.  Regardless, this is a normal exit event and not an error.
		seelog.Info("Received normal signal from Windows to exit")
	case <-agentDone:
		// This means that the agent stopped on its own.  This is where it's appropriate to light the event log on fire
		// and set off all the alarms.
		seelog.Error("Exiting")
	}
	cancel()
	wg.Wait()
	return true, 0
}

// handleWindowsRequests is a loop intended to run in a goroutine.  It handles bidirectional communication with the
// Windows service manager.  This function works by pretty much immediately moving to running and then waiting for a
// stop or shut down message from Windows or to be canceled (which could happen if the agent exits by itself and the
// calling function cancels the context).
func (h *handler) handleWindowsRequests(ctx context.Context, requests <-chan svc.ChangeRequest, responses chan<- svc.Status) {
	// Immediately tell Windows that we are pending service start.
	responses <- svc.Status{State: svc.StartPending}
	seelog.Info("Starting Windows service")

	// TODO: Pre-start hooks go here (unclear if we need any yet)

	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms682108(v=vs.85).aspx
	// Not sure if a better link exists to describe what these values mean
	accepts := svc.AcceptStop | svc.AcceptShutdown

	// Announce that we are running and we accept the above-mentioned commands
	responses <- svc.Status{State: svc.Running, Accepts: accepts}

	defer func() {
		// Announce that we are stopping
		seelog.Info("Stopping Windows service")
		responses <- svc.Status{State: svc.StopPending}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case r := <-requests:
			switch r.Cmd {
			case svc.Interrogate:
				// Our status doesn't change unless we are told to stop or shutdown
				responses <- r.CurrentStatus
			case svc.Stop, svc.Shutdown:
				return
			default:
				continue
			}
		}
	}
}

// runAgent runs the ECS agent inside a goroutine and waits to be told to exit.
func (h *handler) runAgent(ctx context.Context) {
	agentCtx, cancel := context.WithCancel(ctx)
	wg := sync.WaitGroup{}
	running := true
	terminationHandler := func(saver statemanager.Saver, taskEngine engine.TaskEngine) {
		// We're using a waitgroup, a context, and a simple flag here.  The waitgroup gets added to as soon as this
		// handler is invoked (agent.start() ultimately invokes it in a goroutine) so that at the end of the outer
		// runAgent() function we know to wait for the handler to complete.  We then block on the context being
		// canceled; this is our signal that the handler should actually run (happens either when the parent context is
		// canceled because Windows told us to exit, or because the agent goroutine below exited unexpectedly).  The
		// flag gets evaluated so that we know whether to actually save state; if the agent isn't properly running, we
		// may not actually have any data to save.
		wg.Add(1)
		defer wg.Done()
		<-agentCtx.Done()
		if !running {
			return
		}

		seelog.Info("Termination handler received signal to stop")
		err := sighandlers.FinalSave(saver, taskEngine)
		if err != nil {
			seelog.Criticalf("Error saving state before final shutdown: %v", err)
		}
	}
	h.ecsAgent.setTerminationHandler(terminationHandler)

	go func() {
		h.ecsAgent.start() // should block forever, unless there is an error
		// TODO: distinguish between recoverable and unrecoverable errors
		running = false
		cancel()
	}()

	sleepCtx(agentCtx, time.Minute) // give the agent a minute to start and invoke terminationHandler

	// wait for the termination handler to run.  Once the termination handler runs, we can safely exit.  If the agent
	// exits by itself, the termination handler doesn't need to do anything and skips.  If the agent exits before the
	// termination handler is invoked, we can exit immediately.
	wg.Wait()
}

// sleepCtx provides a cancelable sleep
func sleepCtx(ctx context.Context, duration time.Duration) {
	done := make(chan struct{})
	time.AfterFunc(duration, func() {
		close(done)
	})
	select {
	case <-ctx.Done():
	case <-done:
	}
}
