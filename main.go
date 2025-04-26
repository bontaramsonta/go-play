package main

import (
	"fmt"
	"unicode/utf8"
)

const nihongo = "日本語"

func main() {
	// decode first rune
	first, _ := utf8.DecodeRuneInString(nihongo)
	s := string(first)
	println(s)
	// decode controlled
	for i, w := 0, 0; i < len(nihongo); i += w {
		// max length of utf8 rune is 4 bytes
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:min(i+4, len(nihongo))])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
	// standard way to decode runes
	runes := []rune(nihongo)
	println("Better loop")
	for _, r := range runes {
		fmt.Printf("%#U\n", r)
	}
}
