package tree

import "math"

type Tree interface {
	Root() *Node[int]
	InsertLeft(parent *Node[int], value int) *Node[int]
	InsertRight(parent *Node[int], value int) *Node[int]
	levelOrder(traversal func(value int, level int)) []int
	GetSize() int
	GetHeight() int
	GetMaxWidth() int
}

type PrintableTree interface {
	PrintLeftView() []int
	PrintRightView() []int
	PrintBottomView() []int
}

type SumPropertyTree interface {
	HasSumProperty() bool
}

type BalancedTree interface {
	isBalanced() bool
}

type Node[T any] struct {
	value       T
	left, right *Node[T]
}

type BinaryTree struct {
	root    *Node[int]
	count   int
	current int
}

func (b *BinaryTree) levelOrder(traversal func(value int, level int)) []int {
	var result []int
	if b.count != -1 {
		result = make([]int, 0, b.count)
	} else {
		result = make([]int, 0)
	}

	if b.root == nil {
		return result
	}

	type QueueItem struct {
		node  *Node[int]
		level int
	}
	var queue []QueueItem
	if b.count != -1 {
		queue = make([]QueueItem, 0, b.count)
	} else {
		queue = make([]QueueItem, 0)
	}

	queue = append(queue, QueueItem{b.root, 0})
	for i := 0; i < len(queue); i++ {
		current := queue[i]
		traversal(current.node.value, current.level)
		result = append(result, current.node.value)

		if current.node.left != nil {
			queue = append(queue, QueueItem{current.node.left, current.level + 1})
		}
		if current.node.right != nil {
			queue = append(queue, QueueItem{current.node.right, current.level + 1})
		}
	}
	return result
}

func (b *BinaryTree) GetSize() int {
	if b.root == nil {
		return 0
	}
	queue := []*Node[int]{b.root}
	for i := 0; i < len(queue); i++ {
		current := queue[i]
		if current.left != nil {
			queue = append(queue, current.left)
		}
		if current.right != nil {
			queue = append(queue, current.right)
		}
	}
	return len(queue)
}

func (b *BinaryTree) GetHeight() int {
	if b.root == nil {
		return 0
	}
	type QueueItem struct {
		node  *Node[int]
		level int
	}
	queue := []QueueItem{{b.root, 1}}
	maxHeight := 0
	for i := 0; i < len(queue); i++ {
		current := queue[i]
		if current.level > maxHeight {
			maxHeight = current.level
		}
		if current.node.left != nil {
			queue = append(queue, QueueItem{current.node.left, current.level + 1})
		}
		if current.node.right != nil {
			queue = append(queue, QueueItem{current.node.right, current.level + 1})
		}
	}
	return maxHeight
}

func (b *BinaryTree) PrintLeftView() []int {
	var result []int
	if b.root == nil {
		return result
	}

	type QueueItem struct {
		node  *Node[int]
		level int
	}
	var queue []QueueItem
	queue = append(queue, QueueItem{b.root, 0})
	previousLevel := -1
	for i := 0; i < len(queue); i++ {
		current := queue[i]
		if current.level > previousLevel {
			result = append(result, current.node.value)
			previousLevel = current.level
		}

		if current.node.left != nil {
			queue = append(queue, QueueItem{current.node.left, current.level + 1})
		}
		if current.node.right != nil {
			queue = append(queue, QueueItem{current.node.right, current.level + 1})
		}
	}
	return result
}

func (b *BinaryTree) PrintRightView() []int {
	var result []int
	if b.root == nil {
		return result
	}

	type QueueItem struct {
		node  *Node[int]
		level int
	}
	var queue []QueueItem
	queue = append(queue, QueueItem{b.root, 0})
	previousLevel := -1
	for i := 0; i < len(queue); i++ {
		current := queue[i]
		if current.level > previousLevel {
			result = append(result, current.node.value)
			previousLevel = current.level
		}

		if current.node.right != nil {
			queue = append(queue, QueueItem{current.node.right, current.level + 1})
		}
		if current.node.left != nil {
			queue = append(queue, QueueItem{current.node.left, current.level + 1})
		}
	}
	return result
}

func (b *BinaryTree) PrintBottomView() []int {
	if b.root == nil {
		return []int{}
	}

	type QueueItem struct {
		node       *Node[int]
		horizontal int
	}
	queue := []QueueItem{{b.root, 0}}
	horizontalMap := make(map[int]int)
	minHorizontal, maxHorizontal := 0, 0
	for i := 0; i < len(queue); i++ {
		current := queue[i]
		horizontalMap[current.horizontal] = current.node.value

		if current.node.left != nil {
			queue = append(queue, QueueItem{current.node.left, current.horizontal - 1})
		}
		if current.node.right != nil {
			queue = append(queue, QueueItem{current.node.right, current.horizontal + 1})
		}
		minHorizontal = min(minHorizontal, current.horizontal)
		maxHorizontal = max(maxHorizontal, current.horizontal)
	}

	var result []int
	for h := minHorizontal; h <= maxHorizontal; h++ {
		result = append(result, horizontalMap[h])
	}
	return result
}

func (b *BinaryTree) HasSumProperty() bool {
	var checkSumProperty func(node *Node[int]) bool
	checkSumProperty = func(node *Node[int]) bool {
		if node == nil {
			return true
		}
		if node.left == nil && node.right == nil {
			return true
		}

		sum := 0
		if node.left != nil {
			sum += node.left.value
		}
		if node.right != nil {
			sum += node.right.value
		}

		if node.value != sum {
			return false
		}

		return checkSumProperty(node.left) && checkSumProperty(node.right)
	}

	return checkSumProperty(b.root)
}

func (b *BinaryTree) isBalanced() bool {
	var checkBalance func(node *Node[int]) (bool, int)
	checkBalance = func(node *Node[int]) (bool, int) {
		if node == nil {
			return true, 0
		}

		leftBalanced, leftHeight := checkBalance(node.left)
		rightBalanced, rightHeight := checkBalance(node.right)

		currentBalanced := leftBalanced && rightBalanced && math.Abs(float64(leftHeight-rightHeight)) <= 1
		currentHeight := max(leftHeight, rightHeight) + 1

		return currentBalanced, currentHeight
	}

	balanced, _ := checkBalance(b.root)
	return balanced
}

func (b *BinaryTree) GetMaxWidth() int {
	if b.root == nil {
		return 0
	}
	var traversal func(value int, level int)
	currentLevel := 0
	currentLevelCount := 0
	maxLevelCount := 0
	traversal = func(value int, level int) {
		// same level
		if level == currentLevel {
			currentLevelCount++
		} else {
			// level increased
			if currentLevelCount > maxLevelCount {
				maxLevelCount = currentLevelCount
			}
			currentLevel = level
			currentLevelCount = 1
		}
	}

	b.levelOrder(traversal)

	// for the last level
	if currentLevelCount > maxLevelCount {
		maxLevelCount = currentLevelCount
	}

	return maxLevelCount
}

func NewBinaryTree(rootValue int) *BinaryTree {
	return &BinaryTree{
		root:    &Node[int]{value: rootValue},
		count:   -1,
		current: 1,
	}
}

func NewBinaryTreeWithCount(rootValue int, count int) *BinaryTree {
	return &BinaryTree{
		root:    &Node[int]{value: rootValue},
		count:   count,
		current: 1,
	}
}

func (b *BinaryTree) InsertLeft(parent *Node[int], value int) *Node[int] {
	if b.count != -1 && b.current >= b.count {
		return nil
	}
	b.current++
	newNode := &Node[int]{value: value}
	parent.left = newNode
	return newNode
}

func (b *BinaryTree) InsertRight(parent *Node[int], value int) *Node[int] {
	if b.count != -1 && b.current >= b.count {
		return nil
	}
	b.current++
	newNode := &Node[int]{value: value}
	parent.right = newNode
	return newNode
}

func (b *BinaryTree) Root() *Node[int] {
	return b.root
}
