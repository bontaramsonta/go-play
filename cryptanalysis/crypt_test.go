package cryptanalysis

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"sync"
	"testing"
)

// It looks like if the task is CPU bound using goroutines is bad

// faster: 3.07s user 0.56s system 103% cpu 3.506 total
func findKeySimple(encrypted []byte, decrypted string) ([]byte, error) {
	for i := range int(math.Pow(2, 24)) {
		key := intToBytes(i)
		if string(crypt(encrypted, key)) == decrypted {
			return key, nil
		}
	}
	return nil, fmt.Errorf("unable to find key")
}

// slower: 11.43s user 5.29s system 465% cpu 3.596 total
func findKey(encrypted []byte, decrypted string) ([]byte, error) {
	var wg sync.WaitGroup
	keyChan := make(chan []byte, 1) // Buffer of 1 to avoid goroutine leaks
	errChan := make(chan error, 1)

	maxKey := int(math.Pow(2, 24))
	const parts = 4
	partSize := maxKey / parts

	// Launch [parts] goroutines for parallel search
	for t := range parts {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for i := start; i < end; i++ {
				// fmt.Printf("%d Processing key: %d\n", t, i)
				key := intToBytes(i)
				if string(crypt(encrypted, key)) == decrypted {
					select {
					case keyChan <- key:
					default:
					}
					return
				}
			}
		}(t*partSize, (t+1)*partSize)
	}

	// Wait in a goroutine
	go func() {
		wg.Wait()
		close(keyChan)
		errChan <- fmt.Errorf("unable to find key")
	}()

	// Wait for either a key or all goroutines to finish
	select {
	case key := <-keyChan:
		return key, nil
	case err := <-errChan:
		return nil, err
	}
}

// Helper function: crypt performs XOR-based encryption/decryption
func crypt(dat, key []byte) []byte {
	final := []byte{}
	for i, d := range dat {
		final = append(final, d^key[i])
	}
	return final
}

// Helper function: intToBytes converts an integer to a 3-byte slice (little-endian)
func intToBytes(num int) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, int64(num))
	if err != nil {
		return nil
	}
	bs := buf.Bytes()
	if len(bs) > 3 {
		return bs[:3]
	}
	return bs
}

func TestFindKey(t *testing.T) {
	type testCase struct {
		encrypted  []byte
		decrypted  string
		expected   []byte
		shouldFail bool
	}

	testCases := []testCase{
		{[]byte{0x1b, 0x2c, 0x3d}, "yes", []byte{0x62, 0x49, 0x4e}, false}, // Correct key for "yes" is 62 49 4e
		{[]byte{0x2a, 0xff, 0xea}, "car", []byte{0x49, 0x9e, 0x98}, false}, // Correct key for "car" is 49 9e 98
		{[]byte{0x7d, 0x31, 0x32}, "she", []byte{0x0e, 0x59, 0x57}, false}, // Correct key for "she" is 0e 59 57
		{[]byte{0x2b, 0xff, 0xaa}, "top", []byte{0x5f, 0x90, 0xda}, false}, // Correct key for "top" is 5f 90 da
		{[]byte{0x1c, 0x4d, 0x5e}, "win", []byte{0x6b, 0x24, 0x30}, false}, // Correct key for "win" is 6b 24 30
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		key, err := findKeySimple(test.encrypted, test.decrypted)
		if (err != nil) != test.shouldFail {
			failCount++
			t.Errorf(`---------------------------------
Inputs:      encrypted: %v, decrypted: %v
Expecting:   Error: %v
Actual:      Error: %v
Fail
`, test.encrypted, test.decrypted, test.shouldFail, err)
		} else if !test.shouldFail && string(key) != string(test.expected) {
			failCount++
			t.Errorf(`---------------------------------
Inputs:      encrypted: %v, decrypted: %v
Expecting:   Key: %x
Actual:      Key: %x
Fail
`, test.encrypted, test.decrypted, test.expected, key)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:      encrypted: %v, decrypted: %v
Expecting:   Key: %x
Actual:      Key: %x
Pass
`, test.encrypted, test.decrypted, test.expected, key)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)

}
