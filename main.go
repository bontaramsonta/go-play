package main

import (
	"fmt"
)

func candy(ratings []int) int {
	result := make([]int, len(ratings))

	for i := range ratings {
		result[i] = 1

		if i > 0 && ratings[i] > ratings[i-1] {
			result[i] = result[i-1] + 1
		}
	}
	fmt.Printf("1st Pass Result: %v\n", result)

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			result[i] = max(result[i], result[i+1]+1)
		}
	}
	fmt.Printf("2nd Pass Result: %v\n", result)

	sum := 0
	for _, v := range result {
		sum += v
	}
	return sum
}

func trap(height []int) int {
	prefix := make([]int, len(height))
	suffix := make([]int, len(height))

	current := height[0]

	for i := 1; i < len(height); i++ {
		if height[i] > current {
			current = height[i]
		}
		prefix[i] = current
	}

	fmt.Printf("Prefix: %v\n", prefix)

	current = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		if height[i] > current {
			current = height[i]
		}
		suffix[i] = current
	}

	fmt.Printf("Suffix: %v\n", suffix)

	sum := 0
	for i := 0; i < len(height); i++ {
		sum += max(0, min(prefix[i], suffix[i])-height[i])
	}

	return sum
}

type RomanNumeralValue struct {
	order      int
	decimalVal int
}

func romanToInt(s string) int {
	romanValue := map[byte]RomanNumeralValue{
		'I': RomanNumeralValue{order: 1, decimalVal: 1},
		'V': RomanNumeralValue{order: 2, decimalVal: 5},
		'X': RomanNumeralValue{order: 3, decimalVal: 10},
		'L': RomanNumeralValue{order: 4, decimalVal: 50},
		'C': RomanNumeralValue{order: 5, decimalVal: 100},
		'D': RomanNumeralValue{order: 6, decimalVal: 500},
		'M': RomanNumeralValue{order: 7, decimalVal: 1000},
	}

	sum := 0
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && romanValue[s[i]].order < romanValue[s[i+1]].order {
			sum -= romanValue[s[i]].decimalVal
		} else {
			sum += romanValue[s[i]].decimalVal
		}
	}
	return sum
}

func main() {
	var ratings []int

	ratings = []int{1, 0, 2}
	fmt.Printf("For %v : %d\n", ratings, candy(ratings))

	ratings = []int{1, 2, 2}
	fmt.Printf("For %v : %d\n", ratings, candy(ratings))

	var height []int

	height = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Printf("For %v : %d\n", height, trap(height))

	height = []int{4, 2, 0, 3, 2, 5}
	fmt.Printf("For %v : %d\n", height, trap(height))

	var s string

	s = "III"
	fmt.Printf("For %s Roman: %d\n", s, romanToInt(s))

	s = "LVIII"
	fmt.Printf("For %s Roman: %d\n", s, romanToInt(s))

	s = "MCMXCIV"
	fmt.Printf("For %s Roman: %d\n", s, romanToInt(s))
}
