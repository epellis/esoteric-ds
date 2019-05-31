package hashtable

import (
	"testing"
)

func assertEqual(a, b interface{}, t *testing.T) {
	if a != b {
		t.Errorf("Got: %v Expected: %v", a, b)
	}
}

func TestEmptyHashTable(t *testing.T) {
	ht := New()
	_, err := ht.Get(0)
	if err == nil {
		t.Error("Hashtable should be empty")
	}
}

func TestInsert(t *testing.T) {
	ht := New()
	ht.Insert(0, 0)
	assertEqual(ht.size, 1, t)
}

func TestRetrieve(t *testing.T) {
	ht := New()
	ht.Insert(0, 0)
	val, err := ht.Get(0)
	assertEqual(err, nil, t)
	assertEqual(val, 0, t)
}

func TestRetrieveTwice(t *testing.T) {
	ht := New()
	ht.Insert(0, 1)
	ht.Insert(1, 2)

	val, err := ht.Get(0)
	assertEqual(err, nil, t)
	assertEqual(val, 1, t)

	val, err = ht.Get(1)
	assertEqual(err, nil, t)
	assertEqual(val, 2, t)
}

func TestRetrieveMany(t *testing.T) {
	ht := New()

	for i := 0; i < 300; i++ {
		ht.Insert(i, i+1)
	}
	for i := 0; i < 300; i++ {
		val, err := ht.Get(i)
		assertEqual(val, i+1, t)
		assertEqual(err, nil, t)
	}
}
