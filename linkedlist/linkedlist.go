package linkedlist

import "fmt"

// Node structure
type Node struct {
	Value int
	Next  *Node
}

// List structure
type List struct {
	Head *Node
	Tail *Node
}

// Append a node in a linked list
// Complexity: O(1)
func (list *List) Append(node *Node) {
	node.Next = nil

	if list.Head == nil {
		list.Head = node
	} else {
		(*list.Tail).Next = node
	}
	list.Tail = node
}

// InsertFirst adds a new node at the beginning of a linked list
// Complexity: O(1)
func (list *List) InsertFirst(node *Node) {
	if list.Head == nil {
		list.Head = node
	} else {
		node.Next = list.Head
		list.Head = node
	}
}

// RemoveFirst node from a linked list
// Complexity: O(1)
func (list *List) RemoveFirst() {
	if list.Head != nil {
		list.Head = (*list.Head).Next
		(*list.Head).Next = nil
	}
}

// RemoveLast node from a linked list
// Complexity: O(1)
func (list *List) RemoveLast() {
	currentNode := *list.Head
	previousNode := currentNode

	for currentNode.Next != nil {
		previousNode = currentNode
		currentNode = *currentNode.Next
	}
	previousNode.Next = nil
}

// Length returns the list length
// Complexity: O(n)
func (list *List) Length() int {
	var length int

	if list.Head != nil {
		currentNode := *list.Head
		length++

		for currentNode.Next != nil {
			length++
			currentNode = *currentNode.Next
		}
	}
	return length
}

// Print the linked list in order, starting from head
// Complexity: O(n)
func (list *List) Print() {
	if list.Head != nil {
		currentNode := *list.Head
		currentNode.Print()

		for currentNode.Next != nil {
			currentNode = *currentNode.Next
			currentNode.Print()
		}
	} else {
		fmt.Println("The list is empty")
	}
}

// Print the content (value, next) of a node
// Complexity: O(1)
func (node *Node) Print() {
	fmt.Println(node.Value, "->", node.Next)
}

// Reverse a linked list
// Complexity: O(n)
func (list *List) Reverse() {
	currentNode := list.Head
	var nextNode *Node
	var previousNode *Node

	for currentNode != nil {
		nextNode = currentNode.Next
		currentNode.Next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}

	list.Head, list.Tail = list.Tail, list.Head
}
