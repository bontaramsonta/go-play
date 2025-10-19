package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	b := []string{"Penn", "Teller"}
	slices.SortFunc(b, func(a, b string) int {
		return cmp.Compare(a, b)
	})
	for i, v := range slices.All(b) {
		fmt.Printf("%d: %v\n", i, v)
	}
}
