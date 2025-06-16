package dst

import (
	"fmt"
	"strings"
)

// LinkedList Interface

type LinkedList[T any] interface {
	New() *LinkedList[T]
	Length() int
	InsertAtTheEnd(newData T)
	InsertAtTheBeginning(newData T)
	DeleteByValue(value T) error
	String() string
}

// errors

type EmptyListError struct{}

func (e *EmptyListError) Error() string {
	return "linked list is empty"
}

type ValueNotFoundError[T any] struct {
	Value T
}

func (e *ValueNotFoundError[T]) Error() string {
	return fmt.Sprintf("value %v not found in linked list", e.Value)
}

// SingleLinkedList Implementation

type Node[T any] struct {
	data T
	next *Node[T]
}

type SingleLinkedList[T comparable] struct {
	head *Node[T]
}

func (ll *SingleLinkedList[T]) New() *SingleLinkedList[T] {
	return &SingleLinkedList[T]{}
}

func (ll *SingleLinkedList[T]) Length() int {
	current := ll.head
	count := 0
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func (ll *SingleLinkedList[T]) InsertAtTheEnd(newData T) {
	newNode := &Node[T]{data: newData}
	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (ll *SingleLinkedList[T]) InsertAtTheBeginning(newData T) {
	newNode := &Node[T]{data: newData}
	newNode.next = ll.head
	ll.head = newNode
}

func (ll *SingleLinkedList[T]) DeleteByValue(value T) error {
	if ll.head == nil {
		return &EmptyListError{}
	}
	if ll.head.data == value {
		ll.head = ll.head.next
		return nil
	}
	current := ll.head
	for current.next != nil {
		if current.next.data == value {
			current.next = current.next.next
			return nil
		}
		current = current.next
	}
	return &ValueNotFoundError[T]{Value: value}
}

func (ll *SingleLinkedList[T]) String() string {
	if ll.head == nil {
		return "[]"
	}
	var sb strings.Builder
	current := ll.head
	for current != nil {
		sb.WriteString(fmt.Sprintf("%v", current.data))
		if current.next != nil {
			sb.WriteString(" -> ")
		}
		current = current.next
	}
	return sb.String()
}

// Next steps
// tail pointer optimisation for insert at the end
// add unit tests
// traversal utilities: ForEach, Filter, Map, Reduce, Find
// ToSlice
// implement Iterator
