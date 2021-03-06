// Copyright 2015-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/amazon-ecs-agent/agent/wsclient (interfaces: ClientServer,WebsocketConn)

package mock_wsclient

import (
	time "time"

	wsclient "github.com/aws/amazon-ecs-agent/agent/wsclient"
	gomock "github.com/golang/mock/gomock"
)

// Mock of ClientServer interface
type MockClientServer struct {
	ctrl     *gomock.Controller
	recorder *_MockClientServerRecorder
}

// Recorder for MockClientServer (not exported)
type _MockClientServerRecorder struct {
	mock *MockClientServer
}

func NewMockClientServer(ctrl *gomock.Controller) *MockClientServer {
	mock := &MockClientServer{ctrl: ctrl}
	mock.recorder = &_MockClientServerRecorder{mock}
	return mock
}

func (_m *MockClientServer) EXPECT() *_MockClientServerRecorder {
	return _m.recorder
}

func (_m *MockClientServer) AddRequestHandler(_param0 wsclient.RequestHandler) {
	_m.ctrl.Call(_m, "AddRequestHandler", _param0)
}

func (_mr *_MockClientServerRecorder) AddRequestHandler(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddRequestHandler", arg0)
}

func (_m *MockClientServer) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientServerRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockClientServer) Connect() error {
	ret := _m.ctrl.Call(_m, "Connect")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientServerRecorder) Connect() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Connect")
}

func (_m *MockClientServer) Disconnect(_param0 ...interface{}) error {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Disconnect", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientServerRecorder) Disconnect(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Disconnect", arg0...)
}

func (_m *MockClientServer) IsConnected() bool {
	ret := _m.ctrl.Call(_m, "IsConnected")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockClientServerRecorder) IsConnected() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsConnected")
}

func (_m *MockClientServer) MakeRequest(_param0 interface{}) error {
	ret := _m.ctrl.Call(_m, "MakeRequest", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientServerRecorder) MakeRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MakeRequest", arg0)
}

func (_m *MockClientServer) Serve() error {
	ret := _m.ctrl.Call(_m, "Serve")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientServerRecorder) Serve() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Serve")
}

func (_m *MockClientServer) SetAnyRequestHandler(_param0 wsclient.RequestHandler) {
	_m.ctrl.Call(_m, "SetAnyRequestHandler", _param0)
}

func (_mr *_MockClientServerRecorder) SetAnyRequestHandler(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetAnyRequestHandler", arg0)
}

func (_m *MockClientServer) SetConnection(_param0 wsclient.WebsocketConn) {
	_m.ctrl.Call(_m, "SetConnection", _param0)
}

func (_mr *_MockClientServerRecorder) SetConnection(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetConnection", arg0)
}

func (_m *MockClientServer) SetReadDeadline(_param0 time.Time) error {
	ret := _m.ctrl.Call(_m, "SetReadDeadline", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientServerRecorder) SetReadDeadline(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetReadDeadline", arg0)
}

func (_m *MockClientServer) WriteMessage(_param0 []byte) error {
	ret := _m.ctrl.Call(_m, "WriteMessage", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientServerRecorder) WriteMessage(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteMessage", arg0)
}

// Mock of WebsocketConn interface
type MockWebsocketConn struct {
	ctrl     *gomock.Controller
	recorder *_MockWebsocketConnRecorder
}

// Recorder for MockWebsocketConn (not exported)
type _MockWebsocketConnRecorder struct {
	mock *MockWebsocketConn
}

func NewMockWebsocketConn(ctrl *gomock.Controller) *MockWebsocketConn {
	mock := &MockWebsocketConn{ctrl: ctrl}
	mock.recorder = &_MockWebsocketConnRecorder{mock}
	return mock
}

func (_m *MockWebsocketConn) EXPECT() *_MockWebsocketConnRecorder {
	return _m.recorder
}

func (_m *MockWebsocketConn) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockWebsocketConnRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockWebsocketConn) ReadMessage() (int, []byte, error) {
	ret := _m.ctrl.Call(_m, "ReadMessage")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockWebsocketConnRecorder) ReadMessage() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadMessage")
}

func (_m *MockWebsocketConn) SetReadDeadline(_param0 time.Time) error {
	ret := _m.ctrl.Call(_m, "SetReadDeadline", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockWebsocketConnRecorder) SetReadDeadline(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetReadDeadline", arg0)
}

func (_m *MockWebsocketConn) SetWriteDeadline(_param0 time.Time) error {
	ret := _m.ctrl.Call(_m, "SetWriteDeadline", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockWebsocketConnRecorder) SetWriteDeadline(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetWriteDeadline", arg0)
}

func (_m *MockWebsocketConn) WriteMessage(_param0 int, _param1 []byte) error {
	ret := _m.ctrl.Call(_m, "WriteMessage", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockWebsocketConnRecorder) WriteMessage(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteMessage", arg0, arg1)
}
