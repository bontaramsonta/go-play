package main

import "fmt"

func bubbleSort(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < (len(numbers)-i)-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
}

func main() {
	var numbers = []int{6, 1, 7, 3, 5, 9}
	bubbleSort(numbers)
	fmt.Println(numbers)
}
