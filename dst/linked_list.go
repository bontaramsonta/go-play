package dst

import (
	"fmt"
	"strings"
)

// errors

type EmptyListError struct{}

func (e *EmptyListError) Error() string {
	return "linked list is empty"
}

type ValueNotFoundError[T comparable] struct {
	Value T
}

func (e *ValueNotFoundError[T]) Error() string {
	return fmt.Sprintf("value %v not found in linked list", e.Value)
}

func (e *ValueNotFoundError[T]) Is(target error) bool {
	if v, ok := target.(*ValueNotFoundError[T]); ok {
		return e.Value == v.Value
	}
	return false
}

// LinkedList Interface

type LinkedList[T comparable] interface {
	Length() int
	String() string
	InsertAtTheEnd(newData T)
	InsertAtTheBeginning(newData T)
	DeleteByValue(value T) error

	// reverse
	// merge
	// hasLoop
	// common_point
}

// SingleLinkedList Implementation

type Node[T any] struct {
	data T
	next *Node[T]
}

type SingleLinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
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
		ll.tail = newNode
	} else {
		ll.tail.next = newNode
		ll.tail = newNode
	}
}

func (ll *SingleLinkedList[T]) InsertAtTheBeginning(newData T) {
	newNode := &Node[T]{data: newData}
	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.next = ll.head
		ll.head = newNode
	}
}

func (ll *SingleLinkedList[T]) DeleteByValue(value T) error {
	if ll.head == nil {
		return &EmptyListError{}
	}
	if ll.head.data == value {
		ll.head = ll.head.next
		if ll.head == nil {
			ll.tail = nil
		}
		return nil
	}
	found := false
	ll.traverseByNode(func(current *Node[T]) {
		if current.next != nil && current.next.data == value {
			current.next = current.next.next
			if current.next == nil {
				ll.tail = current
			}
			found = true
			return
		}
	})
	if !found {
		return &ValueNotFoundError[T]{Value: value}
	}
	return nil
}

// For printing the linked list
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

func (ll *SingleLinkedList[T]) traverseByValue(fn func(T)) {
	ll.traverseByNode(func(node *Node[T]) {
		fn(node.data)
	})
}

func (ll *SingleLinkedList[T]) traverseByNode(fn func(*Node[T])) {
	current := ll.head
	for current != nil {
		fn(current)
		current = current.next
	}
}

func (ll *SingleLinkedList[T]) ForEach(fn func(T)) {
	ll.traverseByValue(func(data T) {
		fn(data)
	})
}

func (ll *SingleLinkedList[T]) Filter(fn func(T) bool) *SingleLinkedList[T] {
	newList := new(SingleLinkedList[T])
	ll.traverseByValue(func(data T) {
		if fn(data) {
			newList.InsertAtTheEnd(data)
		}
	})
	return newList
}

func (ll *SingleLinkedList[T]) Map(fn func(T) T) *SingleLinkedList[T] {
	newList := new(SingleLinkedList[T])
	ll.traverseByValue(func(data T) {
		newList.InsertAtTheEnd(fn(data))
	})
	return newList
}

func (ll *SingleLinkedList[T]) ToSlice() []T {
	var slice []T
	ll.traverseByValue(func(data T) {
		slice = append(slice, data)
	})
	return slice
}

func (ll *SingleLinkedList[T]) Concat(other *SingleLinkedList[T]) *SingleLinkedList[T] {
	newList := new(SingleLinkedList[T])
	ll.traverseByValue(func(data T) {
		newList.InsertAtTheEnd(data)
	})
	other.traverseByValue(func(data T) {
		newList.InsertAtTheEnd(data)
	})
	return newList
}

func (ll *SingleLinkedList[T]) Reduce(fn func(T, T) T, initial T) T {
	result := initial
	ll.traverseByValue(func(data T) {
		result = fn(result, data)
	})
	return result
}

func (ll *SingleLinkedList[T]) FindNode(fn func(T) bool) *Node[T] {
	var result *Node[T]
	ll.traverseByNode(func(node *Node[T]) {
		if fn(node.data) {
			result = node
		}
	})
	return result
}
