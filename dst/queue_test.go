package dst_test

import (
	"testing"

	"github.com/bontaramsonta/main/dst"
)

func TestSimpleQueue_Enqueue(t *testing.T) {
	sq := dst.SimpleQueue[int]{}
	sq.SetCap(3)
	sq.Enqueue(1)
	sq.Enqueue(2)
	sq.Enqueue(3)
	expected := "[ 1 -> 2 -> 3 ]"
	if got := sq.String(); got != expected {
		t.Errorf("queue output expected %s but got %s", expected, got)
	}
}

func TestSimpleQueue_Dequeue(t *testing.T) {
	sq := dst.SimpleQueue[int]{}
	sq.Enqueue(1)
	sq.Dequeue()
	expected := "[ ]"
	if got := sq.String(); got != expected {
		t.Errorf("queue output expected %s but got %s", expected, got)
	}
}

func TestSimpleQueue_Peek(t *testing.T) {
	sq := dst.SimpleQueue[int]{}
	sq.Enqueue(1)
	expected := 1
	if got := sq.Peek(); got != expected {
		t.Errorf("Peek: expected %v but got %v", expected, got)
	}
}
