package dst

import (
	"fmt"
	"strings"
)

type Queue[T comparable] interface {
	Enqueue(val T)
	Dequeue() *Node[T]
}

type SimpleQueue[T comparable] struct {
	head     *Node[T]
	tail     *Node[T]
	Length   uint
	Capacity uint
}

func (sq *SimpleQueue[T]) isEmpty() bool {
	return sq.head == nil && sq.tail == nil
}

func (sq *SimpleQueue[T]) isFull() bool {
	return sq.Length == sq.Capacity
}

func (sq *SimpleQueue[T]) Enqueue(val T) {
	if sq.isFull() {
		panic("enqueue failed: queue is already full")
	}
	newNode := Node[T]{data: val}
	if sq.isEmpty() {
		sq.head = &newNode
		sq.tail = &newNode
	} else {
		sq.tail.next = &newNode
		sq.tail = &newNode
	}
	sq.Length++
}

func (sq *SimpleQueue[T]) Dequeue() {
	if sq.isEmpty() {
		panic("dequeue failed: queue is already empty")
	}
	nodeToDelete := sq.head
	if sq.Length == 1 {
		sq.head = nil
		sq.tail = nil
	} else {
		sq.head = sq.head.next
		nodeToDelete.next = nil
	}
	sq.Length--
}

func (sq *SimpleQueue[T]) String() string {
	var sb strings.Builder
	sb.WriteString("[ ")

	for current := sq.head; current != nil; current = current.next {
		sb.WriteString(fmt.Sprintf("%v ", current.data))
		if current.next != nil {
			sb.WriteString("-> ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}
