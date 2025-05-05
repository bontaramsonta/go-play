package main

import (
	"fmt"
	"strings"
)

func getHexString(b []byte) string {
	var builder strings.Builder
	for i, v := range b {
		builder.WriteString(fmt.Sprintf("%02x", v))
		if i != len(b)-1 {
			builder.WriteString(":")
		}
	}
	return builder.String()
}

func getBinaryString(b []byte) string {
	var builder strings.Builder
	for i, v := range b {
		builder.WriteString(fmt.Sprintf("%08b", v))
		if i != len(b)-1 {
			builder.WriteString(":")
		}
	}
	return builder.String()
}

func main() {
	hex := getHexString([]byte("Hello, World!"))
	bin := getBinaryString([]byte("Hello, World!"))
	fmt.Println(hex)
	fmt.Println(bin)
}
