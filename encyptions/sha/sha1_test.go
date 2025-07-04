package sha

import (
	"crypto/sha1"
	"fmt"
	"testing"
)

func TestSha1Hash(t *testing.T) {
	result := sha1.Sum([]byte("Hello world"))

	fmt.Printf("hex string %x\n", result)
}
