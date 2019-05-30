package list

import (
	"errors"
	"fmt"
	"testing"
)

type listNode struct {
	value int
	next *listNode
}

type LinkedList struct {
	sentinel *listNode
	tail *listNode
	length uint
}

func NewLinkedList() LinkedList {
	sentinel := listNode{}
	return LinkedList{&sentinel, &sentinel, 0}
}

func (l *LinkedList) InsertBack(val int) {
	newNode := listNode{val, nil}
	l.tail.next = &newNode
	l.tail = &newNode
	l.length++
}

func (l *LinkedList) InsertFront(val int) {
	newNode := listNode{val, nil}
	newNode.next = l.sentinel.next
	l.sentinel.next = &newNode
	l.length++
}

func (l *LinkedList) PopFront() (int, error) {
	return 0, errors.New("not implemented")
}

func (l *LinkedList) Println() {
	for node := l.sentinel.next; node != nil; node = node.next {
		fmt.Printf("(%v)", node.value)
		if node.next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}

func TestEmptyList(t *testing.T) {
	l := NewLinkedList()
	if l.length != 0 {
		t.Error("empty list is not 0 length")
	}
	if l.tail.next != nil {
		t.Error("tail does not point to nil")
	}
}
