package stack

// Node structure
type Node struct {
	Value int
	Next  *Node
}

// Stack structure
type Stack struct {
	Head *Node
}

// Push an element to a stack
// Complexity: O(1)
func (stack *Stack) Push(node *Node) {
	node.Next = stack.Head
	stack.Head = node
}

// Pop first element from a stack
// Complexity: O(1)
func (stack *Stack) Pop() {
	if stack.Head != nil {
		stack.Head = (*stack.Head).Next
		(*stack.Head).Next = nil
	}
}
