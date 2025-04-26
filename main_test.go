package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func Test(t *testing.T) {
	friendships := map[string][]string{
		"Dalinar": {"Kaladin", "Pattern", "Shallan"},
		"Kaladin": {"Dalinar", "Syl", "Teft", "Shallan"},
		"Pattern": {"Dalinar", "Teft", "Shallan"},
		"Syl":     {"Kaladin"},
		"Teft":    {"Kaladin", "Pattern"},
		"Moash":   {},
		"Shallan": {"Pattern", "Kaladin", "Dalinar"},
	}

	type testCase struct {
		username string
		expected []string
	}

	runCases := []testCase{
		{"Dalinar", []string{"Syl", "Teft"}},
		{"Kaladin", []string{"Pattern"}},
		{"Pattern", []string{"Kaladin"}},
		{"Syl", []string{"Dalinar", "Shallan", "Teft"}},
		{"Teft", []string{"Dalinar", "Shallan", "Syl"}},
		{"Moash", nil},
	}

	submitCases := append(runCases, []testCase{
		{
			"Odium", nil,
		},
		{
			"Shallan", []string{"Syl", "Teft"},
		},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)
	passCount := 0
	failCount := 0

	for _, test := range testCases {
		t.Run(test.username, func(t *testing.T) {
			result := findSuggestedFriends(test.username, friendships)
			sort.Strings(result)
			if !reflect.DeepEqual(result, test.expected) {
				failCount++
				t.Errorf(`---------------------------------
Test Failed for username %s:
Expecting:  %v
Actual:     %v
Fail
`, test.username, formatSlice(test.expected), formatSlice(result))
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Test Passed for username %s:
Expecting:  %v
Actual:     %v
Pass
`, test.username, formatSlice(test.expected), formatSlice(result))
			}
		})
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

func formatSlice(slice []string) string {
	if slice == nil {
		return "nil"
	}
	return fmt.Sprintf("%+v", slice)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
