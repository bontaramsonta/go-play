package main

import (
	"fmt"
	"testing"
)

func TestHIndex(t *testing.T) {
	tests := []struct {
		citations []int
		expected  int
	}{
		{
			citations: []int{3, 0, 6, 1, 5},
			expected:  3,
		},
		{
			citations: []int{1, 3, 1},
			expected:  1,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("For %v", tt.citations), func(t *testing.T) {
			result := hIndex(tt.citations)
			if result != tt.expected {
				t.Errorf("hIndex(%v) = %d, expected %d", tt.citations, result, tt.expected)
			}
		})
	}
}
