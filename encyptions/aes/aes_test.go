package aes

import (
	"testing"
)

func TestDebugEncryptDecrypt(t *testing.T) {
	type testCase struct {
		masterKey string
		iv        string
		password  string
		expectedE string
		expectedD string
	}

	const masterKey = "kjhgfdsaqwertyuioplkjhgfdsaqwert"
	const iv = "1234567812345678"

	testCases := []testCase{
		{masterKey, iv, "k33pThisPasswordSafe", encrypt("k33pThisPasswordSafe", masterKey, iv), "k33pThisPasswordSafe"},
		{masterKey, iv, "12345", encrypt("12345", masterKey, iv), "12345"},
		{masterKey, iv, "thePasswordOnMyLuggage", encrypt("thePasswordOnMyLuggage", masterKey, iv), "thePasswordOnMyLuggage"},
		{masterKey, iv, "pizza_the_HUt", encrypt("pizza_the_HUt", masterKey, iv), "pizza_the_HUt"},
		{masterKey, iv, "another_password", encrypt("another_password", masterKey, iv), "another_password"},
		{masterKey, iv, "1234567890", encrypt("1234567890", masterKey, iv), "1234567890"},
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		actualE := encrypt(test.password, test.masterKey, test.iv)
		actualD := decrypt(actualE, test.masterKey, test.iv)
		if actualE == test.expectedE && actualD == test.expectedD {
			passCount++
			t.Logf("Test passed for password %s", test.password)
		} else {
			failCount++
			t.Errorf("Expected encrypted: %s, decrypted: %s, got encrypted: %s, decrypted: %s", test.expectedE, test.expectedD, actualE, actualD)
		}
	}

	t.Logf("Passed: %d Failed: %d", passCount, failCount)
}
