package feistal

import (
	"crypto/sha256"
	"encoding/binary"
	"math/bits"
	"testing"
)

func feistel(msg []byte, roundKeys [][]byte) []byte {
	lhs := msg[:len(msg)/2]
	rhs := msg[len(msg)/2:]
	for _, key := range roundKeys {
		lhs, rhs = rhs, xor(lhs, hash(rhs, key, len(key)))
	}
	return append(rhs, lhs...)
}

func reverse[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func xor(lhs, rhs []byte) []byte {
	res := []byte{}
	for i := range lhs {
		res = append(res, lhs[i]^rhs[i])
	}
	return res
}

// outputLength should be equal to the key length
// when used in feistel so that the XOR operates on
// inputs of the same size
func hash(first, second []byte, outputLength int) []byte {
	h := sha256.New()
	h.Write(append(first, second...))
	return h.Sum(nil)[:outputLength]
}

func TestFeistel(t *testing.T) {
	type testCase struct {
		msg      []byte
		key      []byte
		rounds   int
		expected string
	}

	testCases := []testCase{
		{[]byte("General Kenobi!!!!"), []byte("thesecret"), 8, "General Kenobi!!!!"},
		{[]byte("Hello there!"), []byte("@n@kiN"), 16, "Hello there!"},
		{[]byte("Goodbye!"), []byte("roundkey"), 8, "Goodbye!"},
	}

	passed, failed := 0, 0

	for _, test := range testCases {
		roundKeys := generateRoundKeys(test.key, test.rounds)
		encrypted := feistel(test.msg, roundKeys)
		decrypted := feistel(encrypted, reverse(roundKeys))

		if string(decrypted) != test.expected {
			failed++
			t.Errorf(`---------------------------------
Inputs:      msg: %v, key: %v, rounds: %d
Expecting:   decrypted: %s
Actual:      decrypted: %s
Fail
`, test.msg, test.key, test.rounds, test.expected, string(decrypted))
		} else {
			passed++
			t.Logf(`---------------------------------
Inputs:      msg: %v, key: %v, rounds: %d
Expecting:   decrypted: %s
Actual:      decrypted: %s
Pass
`, test.msg, test.key, test.rounds, test.expected, string(decrypted))
		}
	}

	t.Logf("%d passed, %d failed\n", passed, failed)

}

func generateRoundKeys(key []byte, rounds int) [][]byte {
	roundKeys := [][]byte{}
	for i := range rounds {
		ui := binary.BigEndian.Uint32(key)
		rotated := bits.RotateLeft32(uint32(ui), i)
		finalRound := make([]byte, len(key))
		binary.LittleEndian.PutUint32(finalRound, uint32(rotated))
		roundKeys = append(roundKeys, finalRound)
	}
	return roundKeys
}
