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

func TestDetectIndents(t *testing.T) {
	testcase := []string{
		"nomatch",
		"nomatch",
		"    nomatch",
		"    nomatch",
		"    nomatch",
		"    nomatch",
		"	match",
		"	match",
	}
	_, majorityIndentStyle := DetectIndents(testcase)
	if majorityIndentStyle != " " {
		t.Errorf("Majority indent style reported incorrectly, expected SPACE and got TAB")
	}
}
