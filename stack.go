package main

import (
	"fmt"
)

func main() {
	fmt.Println("Testing Go")
}


type PersistentList interface {
	// Head returns the head of the list. The bool will be false if the list is
	// empty.
	Head() (interface{}, bool)

	// Tail returns the tail of the list. The bool will be false if the list is
	// empty.
	Tail() (PersistentList, bool)

	// IsEmpty indicates if the list is empty.
	IsEmpty() bool

	// Length returns the number of items in the list.
	Length() uint
	
	// Add will add the item to the list, retuning the new list.
	Add(head interface{}) PersistentList

	// Insert will insert the item at the given position, returning the new
	// list or an error if the position is invalid.
	Insert(val interface{}, pos uint)

	// Get returns the item at the given position.
	Get(pos uint) (interface{}, bool)
}

