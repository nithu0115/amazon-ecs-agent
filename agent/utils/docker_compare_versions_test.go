package utils

import (
	"testing"
)

func TestParseVersion(t *testing.T) {
	testCases := []struct {
		input  string
		output Dockerversion
	}{
		{
			input: "1.27",
			output: Dockerversion{
				major: 1,
				minor: 27,
			},
		},
	}

	for i, testCase := range testCases {
		result, err := parseVersion(testCase.input)
		if err != nil {
			t.Errorf("#%v: Error parsing %v: %v", i, testCase.input, err)
		}
		if result != testCase.output{
			t.Errorf("#%v: Expected %v, got %v", i, testCase.output, result)
		}
	}

	invalidCases := []string{"foo", "", "bar", "x.y.z", "1.x.y", "1.1.z"}
	for i, invalidCase := range invalidCases {
		_, err := parseVersion(invalidCase)
		if err == nil {
			t.Errorf("#%v: Expected error, didn't get one. Input: %v", i, invalidCase)
		}
	}
}

func TestVersionMatches(t *testing.T) {
	testCases := []struct {
		version        string
		selector       string
		expectedOutput bool
	}{
		{
			version:        "1.19",
			selector:       "1.19",
			expectedOutput: true,
		},
		{
			version:        "1.19",
			selector:       ">=1.19",
			expectedOutput: true,
		},
		{
			version:        "1.21",
			selector:       ">=1.19",
			expectedOutput: true,
		},
		{
			version:        "1.21",
			selector:       "<=1.19",
			expectedOutput: false,
		},
		{
			version:        "1.19",
			selector:       "<1.21",
			expectedOutput: true,
		},
		{
			version:        "1.19",
			selector:       "<=1.21",
			expectedOutput: true,
		},
		{
			version:        "1.21",
			selector:       ">1.19",
			expectedOutput: true,
		},
		{
			version:        "1.20",
			selector:       ">1.19",
			expectedOutput: true,
		},
		{
			version:        "1.19",
			selector:       "1.19,1.20",
			expectedOutput: true,
		},
		{
			version:        "1.19",
			selector:       "1.21,1.20",
			expectedOutput: false,
		},
		{
			version:        "1.19",
			selector:       ">=1.19,<=1.19",
			expectedOutput: true,
		},
	}

	for i, testCase := range testCases {
		result, err := Version(testCase.version).matches(testCase.selector)
		if err != nil {
			t.Errorf("#%v: Unexpected error %v", i, err)
		}
		if result != testCase.expectedOutput {
			t.Errorf("#%v: %v(%v) expected %v but got %v", i, testCase.version, testCase.selector, testCase.expectedOutput, result)
		}
	}
}