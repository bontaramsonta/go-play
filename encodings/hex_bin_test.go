package encodings

import (
	"fmt"
	"strings"
	"testing"
)

func getHexString(b []byte) string {
	hex := []string{}
	for i, v := range b {
		hex = append(hex, fmt.Sprintf("%02x", v))
		if i != len(b)-1 {
			hex = append(hex, ":")
		}
	}
	return strings.Join(hex, "")
}

func getBinaryString(b []byte) string {
	binary := []string{}
	for i, v := range b {
		binary = append(binary, fmt.Sprintf("%08b", v))
		if i != len(b)-1 {
			binary = append(binary, ":")
		}
	}
	return strings.Join(binary, "")
}

func TestGetHexString(t *testing.T) {
	type testCase struct {
		input    []byte
		expected string
	}

	testCases := []testCase{
		{[]byte("Hello"), "48:65:6c:6c:6f"},     // Hex for "Hello"
		{[]byte("World"), "57:6f:72:6c:64"},     // Hex for "World"
		{[]byte("GoLang"), "47:6f:4c:61:6e:67"}, // Hex for "GoLang"
		{[]byte("Passly"), "50:61:73:73:6c:79"}, // Hex for "Passly"
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		actual := getHexString(test.input)
		if actual != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, actual)
			failCount++
		} else {
			t.Logf("Test passed: %s", test.input)
			passCount++
		}
	}
	t.Logf("Passed: %d, Failed: %d", passCount, failCount)
}

func TestGetBinaryString(t *testing.T) {
	type testCase struct {
		input    []byte
		expected string
	}

	testCases := []testCase{
		{[]byte("Hello"), "01001000:01100101:01101100:01101100:01101111"},           // Binary for "Hello"
		{[]byte("World"), "01010111:01101111:01110010:01101100:01100100"},           // Binary for "World"
		{[]byte("GoLang"), "01000111:01101111:01001100:01100001:01101110:01100111"}, // Binary for "GoLang"
		{[]byte("Passly"), "01010000:01100001:01110011:01110011:01101100:01111001"}, // Binary for "Passly"
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		actual := getBinaryString(test.input)
		if actual != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, actual)
			failCount++
		} else {
			t.Logf("Test passed: %s", test.input)
			passCount++
		}
	}
	t.Logf("Passed: %d, Failed: %d", passCount, failCount)
}
