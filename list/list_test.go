package list

import (
	"reflect"
	"testing"
)

func assertEqual(a, b interface{}, t *testing.T) {
	if a != b {
		t.Errorf("Got: %v Expected: %v", a, b)
	}
}

func TestEmptyList(t *testing.T) {
	l := New()
	assertEqual(l.length, 0, t)
}

func TestInsertFront(t *testing.T) {
	var goldSlice = [...]int{0}
	l := New()
	l.InsertFront(0)
	assertEqual(l.length, len(goldSlice), t)
	if !reflect.DeepEqual(l.ToSlice(), goldSlice[:]) {
		t.Errorf("Expected %v, got %v", goldSlice, l.ToSlice())
	}
}

func TestInsertFrontTwice(t *testing.T) {
	var goldSlice = [...]int{1, 0}
	l := New()
	l.InsertFront(0)
	l.InsertFront(1)
	assertEqual(l.length, len(goldSlice), t)
	if !reflect.DeepEqual(l.ToSlice(), goldSlice[:]) {
		t.Errorf("Expected %v, got %v", goldSlice, l.ToSlice())
	}
}

func TestInsertBack(t *testing.T) {
	var goldSlice = [...]int{0}
	l := New()
	l.InsertBack(0)
	assertEqual(l.length, len(goldSlice), t)
	if !reflect.DeepEqual(l.ToSlice(), goldSlice[:]) {
		t.Errorf("Expected %v, got %v", goldSlice, l.ToSlice())
	}
}

func TestInsertBackTwice(t *testing.T) {
	var goldSlice = [...]int{0, 1}
	l := New()
	l.InsertBack(0)
	l.InsertBack(1)
	assertEqual(l.length, len(goldSlice), t)
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
	assertEqual(val, 0, t)
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
		assertEqual(next, num, t)
	}
}
