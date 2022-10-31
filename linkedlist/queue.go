package linkedlist

import (
	"fmt"
)

// Node : this is a struct type that contains all the characteristic of
// each value in a linkedlist
type Node struct {
	Value any
	Next  *Node
	Prev  *Node
}

// Linkedlist : struct type for the Linkedlist and it field properties
type Linkedlist struct {
	Lenght     int
	Head, Tail *Node
}

// Peak : this return the first node in the linkedList
func (lt *Linkedlist) Peak() interface{} {
	return lt.Head.Value
}

// Display: this helps to show the node available in the linkedlist
func (lt *Linkedlist) Display() error {
	if lt.Head == nil {
		return fmt.Errorf("Display List : List is empty")
	}
	newValue := lt.Head
	for newValue != nil {
		fmt.Printf("%v -> ", newValue.Value)
		newValue = newValue.Next
	}
	fmt.Println()
	return nil
}

// Size: this return the size of the linkedlist
func (lt *Linkedlist) Size() int {
	return lt.Lenght
}

// InsertAtBeginning : this method helps to insert new value at the beginning
// of the linkedlist
func (lt *Linkedlist) InsertAtBeginning(data any) error {
	node := &Node{
		Value: data,
		Next:  nil,
	}
	if lt.Head == nil {
		lt.Head = node
	} else {
		node.Next = lt.Head
		lt.Head = node
	}
	lt.Lenght++
	return nil
}

// InsertAtEnd : this helps to insert new value at the end of the linkedlist
func (lt *Linkedlist) InsertAtEnd(data any) error {
	node := &Node{
		Value: data,
	}
	if lt.Tail == nil {
		lt.Tail = node
	} else {
		currentTail := lt.Tail
		for currentTail != nil {
			currentTail = currentTail.Next
		}
		currentTail.Next = node

	}
	lt.Lenght++
	return nil
}

// InsertAtRandom: this method helps to insert newvalue at random index in the
// linkedlist
func (lt *Linkedlist) InsertAtRandom(index int, data any) error {
	if index < 1 || index > lt.Lenght+1 {
		return fmt.Errorf("insertAtRandom: invalid index to insert data")
	}
	node := &Node{
		Value: data,
		Next:  nil,
	}

	var (
		currentNode  *Node
		previousNode *Node
	)

	previousNode = nil
	currentNode = lt.Head
	if index > 1 {
		previousNode = currentNode
		currentNode = currentNode.Next
		index--
	}
	if previousNode != nil {
		previousNode.Next = node
		node.Next = currentNode
	} else {
		node.Next = currentNode
		lt.Head = node
	}
	lt.Lenght++
	return nil
}

// DeleteAtHead : this helps to delete the value at the beginning of the linkedlist
func (lt *Linkedlist) DeleteAtHead() (any, error) {
	if lt.Head == nil {
		return nil, fmt.Errorf("deletion at the head : No value in the list")
	} else {
		tempNode := lt.Head
		lt.Head = tempNode.Next
		lt.Lenght--
		return tempNode, nil
	}
}

// DeleteAtTail : this method helps to remove or delete the last value in
// in linkedlist with respect to the node
func (lt *Linkedlist) DeleteAtTail() (any, error) {
	if lt.Head == nil {
		return nil, fmt.Errorf("deletion at the tail : No value in list")
	}

	var previousNode *Node
	currentNode := lt.Head
	for currentNode.Next != nil {
		previousNode = currentNode
		currentNode = currentNode.Next
	}
	if previousNode != nil {
		previousNode.Next = nil
	} else {
		lt.Head = nil
	}
	lt.Lenght++
	return previousNode, nil
}
