package main

import (
	"math/rand"
	"sort"
)

func hIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	for i, citation := range citations {
		if citation < i+1 {
			return i
		}
	}
	return len(citations)
}

type RandomizedSet struct {
	set map[int]int8
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		set: make(map[int]int8),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; ok {
		return false
	} else {
		this.set[val] = 1
		return true
	}
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.set[val]; ok {
		delete(this.set, val)
		return true
	} else {
		return false
	}
}

func (this *RandomizedSet) GetRandom() int {
	if len(this.set) == 0 {
		return 0
	}
	keys := make([]int, 0, len(this.set))
	for k := range this.set {
		keys = append(keys, k)
	}
	return keys[rand.Intn(len(keys))]
}

func productExceptSelf(nums []int) []int {
	product := 1
	countZero := 0
	for _, num := range nums {
		if num != 0 {
			product = product * num
		} else {
			countZero++
		}
	}
	answers := make([]int, len(nums))
	if countZero > 1 {
		return answers
	}
	for i, num := range nums {
		if num == 0 {
			answers[i] = product
			return append(make([]int, i+1), answers[i:]...)
		}
		answers[i] = product / num
	}
	return answers
}

func main() {
	// ["RandomizedSet","insert","remove","insert","getRandom","remove","insert","getRandom"]
	// [[],[1],[2],[2],[],[1],[2],[]]

	obj := Constructor()

	// insert(1)
	result1 := obj.Insert(1)
	println("insert(1):", result1)

	// remove(2)
	result2 := obj.Remove(2)
	println("remove(2):", result2)

	// insert(2)
	result3 := obj.Insert(2)
	println("insert(2):", result3)

	// getRandom()
	result4 := obj.GetRandom()
	println("getRandom():", result4)

	// remove(1)
	result5 := obj.Remove(1)
	println("remove(1):", result5)

	// insert(2)
	result6 := obj.Insert(2)
	println("insert(2):", result6)

	// getRandom()
	result7 := obj.GetRandom()
	println("getRandom():", result7)
}
