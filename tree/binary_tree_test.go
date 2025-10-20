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

	expected := []int{1, 2, 3, 4, 5, 6, 7}
	c := 0
	var traversal func(value int, level int)
	traversal = func(value int, level int) {
		expectedValue := expected[c]
		if value != expectedValue {
			t.Errorf("at index %d, got %d, want %d", c, value, expectedValue)
		}
		c++
	}
	result := tree.levelOrder(traversal)

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

func TestBinaryTreeSize(t *testing.T) {
	tree := NewBinaryTree(1)
	left := tree.InsertLeft(tree.Root(), 2)
	tree.InsertRight(tree.Root(), 3)
	secondLeft := tree.InsertLeft(left, 4)      // second left
	thirdLeft := tree.InsertLeft(secondLeft, 5) // third left
	tree.InsertLeft(thirdLeft, 6)               // forth left

	expectedSize := 6
	resultSize := tree.GetSize()

	if resultSize != expectedSize {
		t.Errorf("got size %d, want %d", resultSize, expectedSize)
	}
}

func TestBinaryTreePrintLeftView(t *testing.T) {
	tree := NewBinaryTree(1)
	left := tree.InsertLeft(tree.Root(), 2)
	right := tree.InsertRight(tree.Root(), 3)
	tree.InsertLeft(left, 4)
	tree.InsertRight(left, 5)
	tree.InsertLeft(right, 6)
	tree.InsertRight(right, 7)

	expectedLeftView := []int{1, 2, 4}
	resultLeftView := tree.PrintLeftView()

	if !slices.Equal(resultLeftView, expectedLeftView) {
		t.Errorf("got left view %v, want %v", resultLeftView, expectedLeftView)
	}
}

func TestBinaryTreeHeight(t *testing.T) {
	tree := NewBinaryTree(1)
	left := tree.InsertLeft(tree.Root(), 2)
	tree.InsertRight(tree.Root(), 3)
	secondLeft := tree.InsertLeft(left, 4)      // second left
	thirdLeft := tree.InsertLeft(secondLeft, 5) // third left
	tree.InsertLeft(thirdLeft, 6)               // forth left

	expectedHeight := 5
	resultHeight := tree.GetHeight()

	if resultHeight != expectedHeight {
		t.Errorf("got height %d, want %d", resultHeight, expectedHeight)
	}
}

func TestBinaryTreePrintBottom(t *testing.T) {
	tree := NewBinaryTree(1)
	left := tree.InsertLeft(tree.Root(), 2)
	right := tree.InsertRight(tree.Root(), 3)
	tree.InsertLeft(left, 4)
	tree.InsertRight(left, 5)
	tree.InsertLeft(right, 6)
	tree.InsertRight(right, 7)

	expectedBottomView := []int{4, 5, 6, 7}
	resultBottomView := tree.PrintBottomView()

	if !slices.Equal(resultBottomView, expectedBottomView) {
		t.Errorf("got bottom view %v, want %v", resultBottomView, expectedBottomView)
	}
}

func TestBinaryTreeHasSumProperty(t *testing.T) {
	tree := NewBinaryTree(10)
	left := tree.InsertLeft(tree.Root(), 8)
	right := tree.InsertRight(tree.Root(), 2)
	tree.InsertLeft(left, 3)
	tree.InsertRight(left, 5)
	tree.InsertLeft(right, 2)

	if !tree.HasSumProperty() {
		t.Errorf("expected tree to have sum property")
	}

	// Modify tree to violate sum property
	right.left.value = 3

	if tree.HasSumProperty() {
		t.Errorf("expected tree to not have sum property")
	}
}

func TestBinaryTreeIsBalanced(t *testing.T) {
	tree := NewBinaryTree(1)
	left := tree.InsertLeft(tree.Root(), 2)
	right := tree.InsertRight(tree.Root(), 3)
	tree.InsertLeft(left, 4)
	tree.InsertRight(left, 5)
	lastLeft := tree.InsertLeft(right, 6)
	tree.InsertRight(right, 7)

	if !tree.isBalanced() {
		t.Errorf("expected tree to be balanced")
	}

	// Add extra nodes to make it unbalanced
	extraLeft := tree.InsertLeft(lastLeft, 8)
	tree.InsertLeft(extraLeft, 9)

	if tree.isBalanced() {
		t.Errorf("expected tree to be unbalanced")
	}
}

func TestBinaryTreeMaxWidth(t *testing.T) {
	tree := NewBinaryTree(1)
	left := tree.InsertLeft(tree.Root(), 2)
	right := tree.InsertRight(tree.Root(), 3)
	tree.InsertLeft(left, 4)
	tree.InsertRight(left, 5)
	tree.InsertLeft(right, 6)
	tree.InsertRight(right, 7)

	expectedMaxWidth := 4
	resultMaxWidth := tree.GetMaxWidth()

	if resultMaxWidth != expectedMaxWidth {
		t.Errorf("got max width %d, want %d", resultMaxWidth, expectedMaxWidth)
	}
}
