package tree

import (
	"slices"
	"testing"
)

func TestBinaryTree_LevelOrder(t *testing.T) {
	tree := NewBinaryTree(1)
	left := tree.InsertLeft(tree.Root(), 2)
	right := tree.InsertRight(tree.Root(), 3)
	tree.InsertLeft(left, 4)
	tree.InsertRight(left, 5)
	tree.InsertLeft(right, 6)
	tree.InsertRight(right, 7)

	expected := []int{1, -1, 2, 3, -1, 4, 5, 6, 7}
	result := tree.levelOrder()

	if !slices.Equal(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestInsertWithMoreThanCount(t *testing.T) {
	tree := NewBinaryTreeWithCount(1, 3)
	left := tree.InsertLeft(tree.Root(), 2)
	last := tree.InsertRight(tree.Root(), 3)
	if last == nil {
		t.Errorf("Inserted second last but got nil")
	}
	moreThanLast := tree.InsertLeft(left, 4)
	if moreThanLast != nil {
		t.Errorf("Inserted more but still got %v, want %v", moreThanLast, nil)
	}
}
