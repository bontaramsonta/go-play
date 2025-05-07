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

func TestDeriveRoundKey(t *testing.T) {
	type testCase struct {
		masterKey   [4]byte
		roundNumber int
		expected    [4]byte
	}

	testCases := []testCase{
		{[4]byte{0xAA, 0xFF, 0x11, 0xBC}, 1, [4]byte{0xAB, 0xFE, 0x10, 0xBD}},
		{[4]byte{0xEB, 0xCD, 0x13, 0xFC}, 2, [4]byte{0xE9, 0xCF, 0x11, 0xFE}},
		{[4]byte{0xAA, 0xFF, 0x11, 0xBC}, 5, [4]byte{0xAF, 0xFA, 0x14, 0xB9}},
		{[4]byte{0xEB, 0xCD, 0x13, 0xFC}, 7, [4]byte{0xEC, 0xCA, 0x14, 0xFB}},
	}

	passed, failed := 0, 0

	for _, test := range testCases {
		result := deriveRoundKey(test.masterKey, test.roundNumber)
		if result != test.expected {
			failed++
			t.Errorf(`---------------------------------
Inputs:    masterKey: %X, roundNumber: %d
Expecting: roundKey: %X
Actual:    roundKey: %X
Fail
`, test.masterKey, test.roundNumber, test.expected, result)
		} else {
			passed++
			t.Logf(`---------------------------------
Inputs:    masterKey: %X, roundNumber: %d
Expecting: roundKey: %X
Actual:    roundKey: %X
Pass
`, test.masterKey, test.roundNumber, test.expected, result)
		}
	}

	t.Logf("%d passed, %d failed\n", passed, failed)

}
