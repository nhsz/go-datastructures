package queue

// Node structure
type Node struct {
	Value int
	Next  *Node
}

// Queue structure
type Queue struct {
	Head *Node
	Tail *Node
}

// Enqueue an element
// Complexity: O(1)
func (queue *Queue) Enqueue(node *Node) {
	node.Next = nil

	if queue.Head == nil {
		queue.Head = node
	} else {
		(*queue.Tail).Next = node
	}
	queue.Tail = node
}

// Dequeue first element
// Complexity: O(1)
func (queue *Queue) Dequeue() {
	if queue.Head != nil {
		queue.Head = (*queue.Head).Next
		(*queue.Head).Next = nil
	}
}

// Peek returns first element value
// Complexity: O(1)
func (queue *Queue) Peek() int {
	return (*queue.Head).Value
}
