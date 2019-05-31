package hashtable

import (
	"fmt"
	"github.com/pkg/errors"
)

const startingSize = 7
const resizeFactor = 2
const loadFactor = 0.7

// Standard Hash Table. Maps ints to ints using separate chaining
type HashTable struct {
	table []hashItem
	size  int
	capacity int
}

type hashItem struct {
	occupied bool
	key interface{}
	value interface{}
}

func New() HashTable {
	return HashTable{make([]hashItem, startingSize), 0, startingSize}
}

func (h *HashTable) Insert(key, value int) {
	idx := modHash(key, h.capacity)
	h.size++

	if float32(h.size) / float32(h.capacity) >= loadFactor {
		h.resize()
	}

	for idx := range h.iterIndices(idx) {
		item := h.table[idx]
		if !item.occupied {
			h.table[idx] = hashItem{true, key, value}
			return
		}
	}
}

func (h *HashTable) Get(key int) (int, error) {
	idx := modHash(key, len(h.table))

	for idx := range h.iterIndices(idx) {
		item := h.table[idx]
		if item.occupied && item.key == key {
			return item.value.(int), nil
		}
	}

	return 0, errors.New("key not in table")
}

func (h *HashTable) Println() {
	for i, v := range h.table {
		fmt.Printf("%v ", i)
		if v.occupied {
			fmt.Printf("[%v] ", v.value.(int))
		} else {
			fmt.Print("[ ] ")
		}
	}
	fmt.Println()
}


// Abstract away container traversal. Returns a sequence of indices to
// visit while searching for container to insert, remove or retrieve at.
func (h *HashTable) iterIndices(startIdx int) []int {
	var iterSlice []int

	var stopIdx int
	if startIdx == 0 {
		stopIdx = len(h.table) - 1
	} else {
		stopIdx = (startIdx - 1) % len(h.table)
	}

	for idx := startIdx; idx != stopIdx; idx = (idx + 1) % len(h.table) {
		iterSlice = append(iterSlice, idx)
	}

	return iterSlice
}

func (h *HashTable) resize() {
	oldTable := h.table

	h.capacity *= resizeFactor
	h.table = make([]hashItem, h.capacity)

	for _, item := range oldTable {
		key, value := item.key, item.value
		if key != nil {
			h.Insert(key.(int), value.(int))
		}
	}
}

// Uses the Knuth Multiplicative Hash
func modHash(key, mod int) int {
	return ((key * 2654435761) >> 32) % mod
}
