package list

import (
	"reflect"
	"testing"
)

func TestEmptyList(t *testing.T) {
	l := New()
	if l.length != 0 {
		t.Errorf("Expected 0, got %v", l.length)
	}
	if l.tail.next != nil {
		t.Errorf("Expected nil, got %v", l.tail.next)
	}
}

func TestInsertFront(t *testing.T) {
	var goldSlice = [...]int{0}
	l := New()
	l.InsertFront(0)
	if l.length != len(goldSlice) {
		t.Errorf("Expected %v, got %v", len(goldSlice), l.length)
	}
	if !reflect.DeepEqual(l.ToSlice(), goldSlice[:]) {
		t.Errorf("Expected %v, got %v", goldSlice, l.ToSlice())
	}
}

func TestInsertFrontTwice(t *testing.T) {
	var goldSlice = [...]int{1, 0}
	l := New()
	l.InsertFront(0)
	l.InsertFront(1)
	if l.length != len(goldSlice) {
		t.Errorf("Expected %v, got %v", len(goldSlice), l.length)
	}
	if !reflect.DeepEqual(l.ToSlice(), goldSlice[:]) {
		t.Errorf("Expected %v, got %v", goldSlice, l.ToSlice())
	}
}

func TestInsertBack(t *testing.T) {
	var goldSlice = [...]int{0}
	l := New()
	l.InsertBack(0)
	if l.length != len(goldSlice) {
		t.Errorf("Expected %v, got %v", len(goldSlice), l.length)
	}
	if !reflect.DeepEqual(l.ToSlice(), goldSlice[:]) {
		t.Errorf("Expected %v, got %v", goldSlice, l.ToSlice())
	}
}

func TestInsertBackTwice(t *testing.T) {
	var goldSlice = [...]int{0, 1}
	l := New()
	l.InsertBack(0)
	l.InsertBack(1)
	if l.length != len(goldSlice) {
		t.Errorf("Expected %v, got %v", len(goldSlice), l.length)
	}
	if !reflect.DeepEqual(l.ToSlice(), goldSlice[:]) {
		t.Errorf("Expected %v, got %v", goldSlice, l.ToSlice())
	}
}

func TestPop(t *testing.T) {
	l := New()
	l.InsertBack(0)
	val, err := l.PopFront()
	if err != nil {
		t.Errorf("List is empty but should not be")
	}
	if val != 0 {
		t.Errorf("Expected %v, got %v", 0, val)
	}
}

func TestFifo(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	l := New()

	for _, num := range numbers {
		l.InsertBack(num)
	}
	for _, num := range numbers {
		next, err := l.PopFront()
		if err != nil {
			t.Errorf("List is empty but should not be")
		}
		if next != num {
			t.Errorf("Expected %v, got %v", num, next)
		}
	}
}
