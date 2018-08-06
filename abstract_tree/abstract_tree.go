package abstractTree

import (
	"github.com/grandquista/data-structures-and-algorithms/queue"
)

// Empty optional value unpacked.
type emptyOptional struct{}

func (emptyOptional) Error() string {
	return "empty optional"
}

type poison struct{}

// Optional value container.
type optional struct {
	value interface{}
}

func makeOptional() optional {
	return optional{poison{}}
}

func makeOptionalValue(value interface{}) optional {
	return optional{value}
}

func (optional optional) Bool() bool {
	switch optional.value.(type) {
	case poison:
		return false
	}
	return true
}

func (optional optional) Visit(visit func(interface{})) {
	if optional.Bool() {
		visit(optional.value)
	}
}

func (optional optional) GetValue() (interface{}, error) {
	if optional.Bool() {
		return optional.value, nil
	}
	return nil, emptyOptional{}
}

func (optional optional) SetValue(value interface{}) {
	optional.value = value
}

func (optional optional) DeleteValue() {
	optional.value = poison{}
}

// AbstractBaseTree abstractTree.
type AbstractBaseTree interface {
	Len() int
	GetValue() (interface{}, error)
	SetValue(value interface{})
	DeleteValue()
	PostOrder(visit func(interface{}))
	PreOrder(visit func(interface{}))
}

type abstractBaseTree struct {
	value optional
	left  AbstractBaseTree
	right AbstractBaseTree
	size  int
}

// MakeAbstractBaseTree Initialize new tree with optional value and right child.
func MakeAbstractBaseTree() AbstractBaseTree {
	return abstractBaseTree{}
}

// MakeAbstractBaseTreeValue Initialize new tree with optional value and right child.
func MakeAbstractBaseTreeValue(value interface{}) AbstractBaseTree {
	return abstractBaseTree{value: makeOptionalValue(value), size: 1}
}

// Initialize new tree with optional value and right child.
func makeAbstractBaseTreeValue(value interface{}, right AbstractBaseTree) AbstractBaseTree {
	return abstractBaseTree{value: makeOptionalValue(value), right: right, size: 1 + right.Len()}
}

func (tree abstractBaseTree) GetValue() (interface{}, error) {
	return tree.value.GetValue()
}

func (tree abstractBaseTree) SetValue(value interface{}) {
	tree.value.SetValue(value)
}

func (tree abstractBaseTree) DeleteValue() {
	tree.value.DeleteValue()
}

// Indicate if the value is found in the tree.
func (tree abstractBaseTree) Contains(value interface{}) bool {
	if tree.size == 0 {
		return false
	}
	queue := queue.MakeQueue([]interface{}{tree})
	for queue.Len() > 0 {
		v, _ := queue.Dequeue()
		tree, _ := v.(*abstractBaseTree)
		for tree != nil {
			if tree.left != nil {
				queue.Enqueue(tree.left)
			}
			switch val, _ := tree.GetValue(); val {
			case value:
				return true
			}
			tree = tree.right.(*abstractBaseTree)
		}
	}
	return false
}

// Return the number of values currently in the tree.
func (tree abstractBaseTree) Len() int {
	return tree.size
}

// Visit each of the values in breadth first order.
func (tree abstractBaseTree) BreadthFirst(visit func(interface{})) {
	if tree.size == 0 {
		return
	}
	queue := queue.MakeQueue([]interface{}{tree})
	for queue.Len() > 0 {
		v, _ := queue.Dequeue()
		tree, _ := v.(*abstractBaseTree)
		for tree != nil {
			if tree.left != nil {
				queue.Enqueue(tree.left)
			}
			tree.value.Visit(visit)
			tree = tree.right.(*abstractBaseTree)
		}
	}
}

// Update size for count inserts.
func (tree abstractBaseTree) insert(count int) int {
	tree.size += count
	return count
}

// Visit each of the values in post order.
func (tree abstractBaseTree) PostOrder(visit func(interface{})) {
	if tree.left != nil {
		tree.left.PostOrder(visit)
	}
	if tree.right != nil {
		tree.right.PostOrder(visit)
	}
	tree.value.Visit(visit)
}

// Visit each of the values in pre order.
func (tree abstractBaseTree) PreOrder(visit func(interface{})) {
	tree.value.Visit(visit)
	if tree.left != nil {
		tree.left.PostOrder(visit)
	}
	if tree.right != nil {
		tree.right.PostOrder(visit)
	}
}
