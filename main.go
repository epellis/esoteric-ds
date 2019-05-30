package main

import "esoteric-ds/list"

func main() {
	myList := list.New()
	myList.InsertBack(0)
	myList.InsertBack(1)
	myList.InsertBack(2)
	myList.Println()
}
