package main

import "esoteric-ds/list"

func main() {
	myList := list.NewLinkedList()
	myList.InsertBack(0)
	myList.InsertBack(1)
	myList.InsertBack(2)
	myList.Println()
}
