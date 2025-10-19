package tree

type Tree interface {
	Root() *Node[int]
	levelOrder() []int
	InsertLeft(parent *Node[int], value int) *Node[int]
	InsertRight(parent *Node[int], value int) *Node[int]
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

func (b *BinaryTree) levelOrder() []int {
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
	previousLevel := 0
	for i := 0; i < len(queue); i++ {
		current := queue[i]
		if current.level > previousLevel {
			result = append(result, -1)
			previousLevel = current.level
		}
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
