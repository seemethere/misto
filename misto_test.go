package main

import "testing"

func TestDetectMixedIndent(t *testing.T) {
	testcases := []struct {
		input    string
		expected int
	}{
		{"hello", 0},
		{"	hello", 0},
		{"    hello", 0},
		{"	  hello", 1},
		{"    	hello", 2},
	}
	for _, testcase := range testcases {
		errorCode := DetectMixedIndent(testcase.input)
		if errorCode != testcase.expected {
			t.Errorf("Error code incorrect for '%s', got: %d, expected: %d", formatLine(testcase.input), errorCode, testcase.expected)
		}
	}
}
