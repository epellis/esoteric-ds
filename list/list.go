package list

import (
	"fmt"
	"github.com/pkg/errors"
)

type listNode struct {
	value interface{}
	next  *listNode
}

type LinkedList struct {
	sentinel *listNode
	tail     *listNode
	length   int
}

func New() LinkedList {
	sentinel := listNode{}
	return LinkedList{&sentinel, &sentinel, 0}
}

func (l *LinkedList) InsertBack(val interface{}) {
	newNode := listNode{val, nil}
	l.tail.next = &newNode
	l.tail = &newNode
	l.length++
}

func (l *LinkedList) InsertFront(val interface{}) {
	newNode := listNode{val, nil}
	newNode.next = l.sentinel.next
	l.sentinel.next = &newNode
	l.length++
}

func (l *LinkedList) PopFront() (interface{}, error) {
	if l.length == 0 {
		return 0, errors.New("empty list")
	}
	value := l.sentinel.next.value
	l.sentinel.next = l.sentinel.next.next
	return value, nil
}

func (l *LinkedList) ToSlice() []int {
	var slice []int

	for node := l.sentinel.next; node != nil; node = node.next {
		slice = append(slice, node.value.(int))
	}

	return slice
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
