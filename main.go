package main

func maxProfit(prices []int) int {
	mp := 0
	for i := range len(prices) - 1 {
		if prices[i+1] > prices[i] {
			mp += prices[i+1] - prices[i]
		}
	}
	return mp
}

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	mjp := 0
	for i := range len(nums) - 1 {
		mjp = max(mjp-1, nums[i])
		if mjp <= 0 {
			return false
		}
	}
	return true
}

func minJump(nums []int) int {
	fartest := 0
	jumps := 0
	current_end := 0

	for i := range len(nums) - 1 {
		fartest = max(fartest, nums[i]+i)

		if i == current_end {
			jumps++
			current_end = fartest
		}
	}
	return jumps
}

func main() {
	// Test maxProfit function
	println("=== Testing maxProfit function ===")
	maxProfitTests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{"Mixed prices", []int{7, 1, 5, 3, 6, 4}, 7},
		{"All increasing", []int{1, 2, 3, 4, 5}, 4},
		{"All decreasing", []int{5, 4, 3, 2, 1}, 0},
		{"Single price", []int{5}, 0},
		{"Two prices", []int{1, 5}, 4},
	}

	for _, test := range maxProfitTests {
		result := maxProfit(test.prices)
		status := "PASS"
		if result != test.expected {
			status = "FAIL"
		}
		println(status, "-", test.name, "-> Expected:", test.expected, "Got:", result)
	}

	// Test canJump function
	println("\n=== Testing canJump function ===")
	canJumpTests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{"Can jump with zeros", []int{2, 0, 0}, true},
		{"Standard success", []int{2, 3, 1, 1, 4}, true},
		{"Blocked by zero", []int{3, 2, 1, 0, 4}, false},
		{"Single element", []int{0}, true},
		{"Barely making it", []int{1, 1, 1, 1}, true},
		{"Large first jump", []int{5, 0, 0, 0, 0}, true},
	}

	for _, test := range canJumpTests {
		result := canJump(test.nums)
		status := "PASS"
		if result != test.expected {
			status = "FAIL"
		}
		println(status, "-", test.name, "-> Expected:", test.expected, "Got:", result)
	}

	// Test minJump function
	println("\n=== Testing minJump function ===")
	minJumpTests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"Standard case", []int{2, 3, 1, 1, 4}, 2},
		{"Another example", []int{2, 3, 0, 1, 4}, 2},
		{"Single element", []int{1}, 0},
		{"Multiple small jumps", []int{1, 1, 1, 1, 1}, 4},
		{"One big jump", []int{5, 1, 1, 1, 1}, 1},
		{"Mixed scenario", []int{1, 2, 3}, 2},
	}

	for _, test := range minJumpTests {
		result := minJump(test.nums)
		status := "PASS"
		if result != test.expected {
			status = "FAIL"
		}
		println(status, "-", test.name, "-> Expected:", test.expected, "Got:", result)
	}
}
