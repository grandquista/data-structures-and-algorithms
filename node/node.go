package node

// Node node.
type Node struct {
	Value interface{}
	Next  *Node
}

// MakeNode Initialize new Node with optional next Node.
func MakeNode(Value interface{}, Next *Node) *Node {
	return &Node{Value, Next}
}

// HasLoop Return a boolean indicating if the list has a loop of nodes.
func (node *Node) HasLoop() bool {
	node1 := node
	node2 := node1.Next
	for node2 != nil {
		if node1 == node2 {
			return true
		}
		node2 = node2.Next
		if node2 == nil {
			return false
		}
		if node1 == node2 {
			return true
		}
		node1 = node1.Next
		node2 = node2.Next
	}
	return false
}
