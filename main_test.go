package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func isValidJSON(input string) bool {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(input), "", "  ")
	return err == nil
}

func TestIsValidJSON(t *testing.T) {
	type testCase struct {
		input string
	}

	runCases := []testCase{
		{issueList},
	}

	submitCases := append(runCases, []testCase{
		{userObject},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passed := 0
	failed := 0

	for _, test := range testCases {
		if output := isValidJSON(test.input); !output {
			failed++
			t.Errorf(`---------------------------------
Test Failed. Input:
%v
  =>
expected isValidJSON: %v
actual isValidJSON: %v
`,
				test.input, true, output)
		} else {
			passed++
			fmt.Printf(`---------------------------------
Test Passed. Input:
%v
  =>
expected isValidJSON: %v
actual isValidJSON: %v
`,
				test.input, true, output)
		}
	}
	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passed, failed, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passed, failed)
	}

}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
