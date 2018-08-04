package node_test

import (
	"testing"

	. "github.com/grandquista/data-structures-and-algorithms/node"
)

func newNode() *Node {
	return MakeNode(nil, nil)
}

func chainedNode() *Node {
	return MakeNode(1, MakeNode(2, MakeNode(3, nil)))
}

func loopNode() *Node {
	end := MakeNode(4, nil)
	node := MakeNode(1, MakeNode(2, MakeNode(3, MakeNode(2, MakeNode(1, end)))))
	end.Next = node.Next
	return node
}

func TestEmptyNode(t *testing.T) {
	node := newNode()
	if node.Value != nil {
		t.Fail()
	}
	if node.Next != nil {
		t.Fail()
	}
}

func TestChainedNode(t *testing.T) {
	node := chainedNode()
	if node.Value != 1 {
		t.Fail()
	}
	if node.Next == nil {
		t.Fail()
	}
}
