package random

import (
	"fmt"
	"math/rand"
	"testing"
)

func generateRandomKey(length int) (string, error) {
	key := make([]byte, length)
	if _, err := rand.Read(key); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", key), nil
}

func TestGenerateRandomKey(t *testing.T) {
	rand.Seed(0)
	type testCase struct {
		length     int
		shouldFail bool
		expected   string
	}

	cases := []testCase{
		{8, false, "a178892ee285ece1"},                                                  // Expected output for 8 bytes
		{16, false, "0194fdc2fa2ffcc041d3ff12045b73c8"},                                 // Expected output for 16 bytes
		{32, false, "6e4ff95ff662a5eee82abdf44a2d0b75fb180daf48a79ee0b10d394651850fd4"}, // Expected output for 32 bytes
		{64, false, "511455780875d64ee2d3d0d0de6bf8f9b44ce85ff044c6b1f83b8e883bbf857aab99c5b252c7429c32f3a8aeb79ef856f659c18f0dcecc77c75e7a81bfde275f"}, // Expected output for 64 bytes
	}

	passCount := 0
	failCount := 0

	for _, tc := range cases {
		key, err := generateRandomKey(tc.length)
		if err != nil && !tc.shouldFail {
			t.Errorf("generateRandomKey(%d) failed with error: %v", tc.length, err)
			failCount++
			continue
		}
		if err == nil && tc.shouldFail {
			t.Errorf("generateRandomKey(%d) should have failed but didn't", tc.length)
			failCount++
			continue
		}
		if key != tc.expected {
			t.Errorf("generateRandomKey(%d) = %s, want %s", tc.length, key, tc.expected)
			failCount++
			continue
		}
		passCount++
	}

	t.Logf("TestGenerateRandomKey passed %d times and failed %d times\n", passCount, failCount)
}
