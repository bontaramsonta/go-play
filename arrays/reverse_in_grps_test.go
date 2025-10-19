package arrays

import (
	"slices"
	"testing"
)

func reverseInGroups[T any](arr []T, k int) {
	if k <= 1 {
		return
	}
	for i := 0; i < len(arr); i += k {
		end := min(i+k, len(arr))
		slices.Reverse(arr[i:end])
	}
}

func Test(t *testing.T) {
	testCases := []struct {
		arr      []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 5, 4}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{5, 4, 3, 2, 1}},
		{[]int{}, 3, []int{}},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			reverseInGroups(tc.arr, tc.k)
			if !slices.Equal(tc.arr, tc.expected) {
				t.Errorf("got %v, want %v", tc.arr, tc.expected)
			}
		})
	}
}
