package dst_test

import (
	"errors"
	"testing"

	"github.com/bontaramsonta/main/dst"
)

func TestSingleLinkedList_Length(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *dst.SingleLinkedList[int]
		expected int
	}{
		{
			name: "empty list",
			setup: func() *dst.SingleLinkedList[int] {
				return new(dst.SingleLinkedList[int])
			},
			expected: 0,
		},
		{
			name: "single element",
			setup: func() *dst.SingleLinkedList[int] {
				list := new(dst.SingleLinkedList[int])
				list.InsertAtTheEnd(42)
				return list
			},
			expected: 1,
		},
		{
			name: "three elements",
			setup: func() *dst.SingleLinkedList[int] {
				list := new(dst.SingleLinkedList[int])
				list.InsertAtTheEnd(1)
				list.InsertAtTheEnd(2)
				list.InsertAtTheEnd(3)
				return list
			},
			expected: 3,
		},
		{
			name: "after deletion",
			setup: func() *dst.SingleLinkedList[int] {
				list := new(dst.SingleLinkedList[int])
				list.InsertAtTheEnd(1)
				list.InsertAtTheEnd(2)
				list.InsertAtTheEnd(3)
				list.DeleteByValue(2)
				return list
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := tt.setup()
			if got := list.Length(); got != tt.expected {
				t.Errorf("Length() = %d, want %d", got, tt.expected)
			}
		})
	}
}

func TestSingleLinkedList_String(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *dst.SingleLinkedList[int]
		expected string
	}{
		{
			name: "empty list",
			setup: func() *dst.SingleLinkedList[int] {
				return new(dst.SingleLinkedList[int])
			},
			expected: "[]",
		},
		{
			name: "single element",
			setup: func() *dst.SingleLinkedList[int] {
				list := new(dst.SingleLinkedList[int])
				list.InsertAtTheEnd(42)
				return list
			},
			expected: "42",
		},
		{
			name: "multiple elements",
			setup: func() *dst.SingleLinkedList[int] {
				list := new(dst.SingleLinkedList[int])
				list.InsertAtTheEnd(1)
				list.InsertAtTheEnd(2)
				list.InsertAtTheEnd(3)
				return list
			},
			expected: "1 -> 2 -> 3",
		},
		{
			name: "mixed insertions",
			setup: func() *dst.SingleLinkedList[int] {
				list := new(dst.SingleLinkedList[int])
				list.InsertAtTheEnd(2)
				list.InsertAtTheBeginning(1)
				list.InsertAtTheEnd(3)
				return list
			},
			expected: "1 -> 2 -> 3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := tt.setup()
			if got := list.String(); got != tt.expected {
				t.Errorf("String() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestSingleLinkedList_DeleteByValue(t *testing.T) {
	tests := []struct {
		name          string
		setup         func() *dst.SingleLinkedList[int]
		deleteValue   int
		expectedError error
		expectedState string
	}{
		{
			name: "delete from empty list",
			setup: func() *dst.SingleLinkedList[int] {
				return new(dst.SingleLinkedList[int])
			},
			deleteValue:   1,
			expectedError: &dst.EmptyListError{},
			expectedState: "[]",
		},
		{
			name: "delete non-existent value",
			setup: func() *dst.SingleLinkedList[int] {
				list := new(dst.SingleLinkedList[int])
				list.InsertAtTheEnd(1)
				list.InsertAtTheEnd(2)
				return list
			},
			deleteValue:   3,
			expectedError: &dst.ValueNotFoundError[int]{Value: 3},
			expectedState: "1 -> 2",
		},
		{
			name: "delete first element",
			setup: func() *dst.SingleLinkedList[int] {
				list := new(dst.SingleLinkedList[int])
				list.InsertAtTheEnd(1)
				list.InsertAtTheEnd(2)
				list.InsertAtTheEnd(3)
				return list
			},
			deleteValue:   1,
			expectedError: nil,
			expectedState: "2 -> 3",
		},
		{
			name: "delete middle element",
			setup: func() *dst.SingleLinkedList[int] {
				list := new(dst.SingleLinkedList[int])
				list.InsertAtTheEnd(1)
				list.InsertAtTheEnd(2)
				list.InsertAtTheEnd(3)
				return list
			},
			deleteValue:   2,
			expectedError: nil,
			expectedState: "1 -> 3",
		},
		{
			name: "delete last element",
			setup: func() *dst.SingleLinkedList[int] {
				list := new(dst.SingleLinkedList[int])
				list.InsertAtTheEnd(1)
				list.InsertAtTheEnd(2)
				list.InsertAtTheEnd(3)
				return list
			},
			deleteValue:   3,
			expectedError: nil,
			expectedState: "1 -> 2",
		},
		{
			name: "delete only element",
			setup: func() *dst.SingleLinkedList[int] {
				list := new(dst.SingleLinkedList[int])
				list.InsertAtTheEnd(42)
				return list
			},
			deleteValue:   42,
			expectedError: nil,
			expectedState: "[]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := tt.setup()
			err := list.DeleteByValue(tt.deleteValue)

			if tt.expectedError != nil {
				if err == nil {
					t.Error("expected error but got nil")
				} else if !errors.Is(err, tt.expectedError) {
					t.Errorf("expected error %v, got %v", tt.expectedError, err)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}

			if got := list.String(); got != tt.expectedState {
				t.Errorf("final state = %q, want %q", got, tt.expectedState)
			}
		})
	}
}

func TestSingleLinkedList_Filter(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		predicate func(int) bool
		expected  string
	}{
		{
			name:  "empty list",
			input: []int{},
			predicate: func(n int) bool {
				return n > 0
			},
			expected: "[]",
		},
		{
			name:  "filter even numbers",
			input: []int{1, 2, 3, 4, 5, 6},
			predicate: func(n int) bool {
				return n%2 == 0
			},
			expected: "2 -> 4 -> 6",
		},
		{
			name:  "filter greater than 3",
			input: []int{1, 2, 3, 4, 5},
			predicate: func(n int) bool {
				return n > 3
			},
			expected: "4 -> 5",
		},
		{
			name:  "no matches",
			input: []int{1, 2, 3},
			predicate: func(n int) bool {
				return n > 10
			},
			expected: "[]",
		},
		{
			name:  "all match",
			input: []int{2, 4, 6},
			predicate: func(n int) bool {
				return n%2 == 0
			},
			expected: "2 -> 4 -> 6",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := new(dst.SingleLinkedList[int])
			for _, v := range tt.input {
				list.InsertAtTheEnd(v)
			}

			filtered := list.Filter(tt.predicate)
			if got := filtered.String(); got != tt.expected {
				t.Errorf("Filter() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestSingleLinkedList_Map(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		mapFunc  func(int) int
		expected string
	}{
		{
			name:  "empty list",
			input: []int{},
			mapFunc: func(n int) int {
				return n * 2
			},
			expected: "[]",
		},
		{
			name:  "double values",
			input: []int{1, 2, 3},
			mapFunc: func(n int) int {
				return n * 2
			},
			expected: "2 -> 4 -> 6",
		},
		{
			name:  "square values",
			input: []int{1, 2, 3, 4},
			mapFunc: func(n int) int {
				return n * n
			},
			expected: "1 -> 4 -> 9 -> 16",
		},
		{
			name:  "add constant",
			input: []int{1, 2, 3},
			mapFunc: func(n int) int {
				return n + 10
			},
			expected: "11 -> 12 -> 13",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := new(dst.SingleLinkedList[int])
			for _, v := range tt.input {
				list.InsertAtTheEnd(v)
			}

			mapped := list.Map(tt.mapFunc)
			if got := mapped.String(); got != tt.expected {
				t.Errorf("Map() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestSingleLinkedList_Reduce(t *testing.T) {
	tests := []struct {
		name       string
		input      []int
		reduceFunc func(int, int) int
		initial    int
		expected   int
	}{
		{
			name:  "empty list sum",
			input: []int{},
			reduceFunc: func(acc, val int) int {
				return acc + val
			},
			initial:  0,
			expected: 0,
		},
		{
			name:  "sum of numbers",
			input: []int{1, 2, 3, 4, 5},
			reduceFunc: func(acc, val int) int {
				return acc + val
			},
			initial:  0,
			expected: 15,
		},
		{
			name:  "product of numbers",
			input: []int{2, 3, 4},
			reduceFunc: func(acc, val int) int {
				return acc * val
			},
			initial:  1,
			expected: 24,
		},
		{
			name:  "find maximum",
			input: []int{3, 7, 2, 9, 1},
			reduceFunc: func(acc, val int) int {
				if val > acc {
					return val
				}
				return acc
			},
			initial:  0,
			expected: 9,
		},
		{
			name:  "count elements",
			input: []int{1, 2, 3, 4},
			reduceFunc: func(acc, val int) int {
				return acc + 1
			},
			initial:  0,
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := new(dst.SingleLinkedList[int])
			for _, v := range tt.input {
				list.InsertAtTheEnd(v)
			}

			result := list.Reduce(tt.reduceFunc, tt.initial)
			if result != tt.expected {
				t.Errorf("Reduce() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestSingleLinkedList_ComplexOperations(t *testing.T) {
	tests := []struct {
		name       string
		operations func() *dst.SingleLinkedList[int]
		expected   string
	}{
		{
			name: "insert at end then beginning",
			operations: func() *dst.SingleLinkedList[int] {
				list := new(dst.SingleLinkedList[int])
				list.InsertAtTheEnd(2)
				list.InsertAtTheEnd(3)
				list.InsertAtTheBeginning(1)
				return list
			},
			expected: "1 -> 2 -> 3",
		},
		{
			name: "multiple deletions",
			operations: func() *dst.SingleLinkedList[int] {
				list := new(dst.SingleLinkedList[int])
				for i := 1; i <= 5; i++ {
					list.InsertAtTheEnd(i)
				}
				list.DeleteByValue(2)
				list.DeleteByValue(4)
				return list
			},
			expected: "1 -> 3 -> 5",
		},
		{
			name: "alternating operations",
			operations: func() *dst.SingleLinkedList[int] {
				list := new(dst.SingleLinkedList[int])
				list.InsertAtTheEnd(1)
				list.InsertAtTheBeginning(0)
				list.InsertAtTheEnd(2)
				list.DeleteByValue(0)
				list.InsertAtTheBeginning(-1)
				return list
			},
			expected: "-1 -> 1 -> 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := tt.operations()
			if got := list.String(); got != tt.expected {
				t.Errorf("final state = %q, want %q", got, tt.expected)
			}
		})
	}
}
