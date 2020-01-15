package bst

import (
	"math"

	"github.com/KevinAnthony/go-tree/types"
)

type binaryNode struct {
	data  types.Data
	left  *binaryNode
	right *binaryNode
}

func (b binaryNode) GetData() types.Data {
	return b.data
}

func (b binaryNode) IsLeaf() bool {
	return b.left == nil && b.right == nil
}

func (b *binaryNode) insert(node *binaryNode) {
	if b.data.GreaterThan(node.data) {
		if b.left == nil {
			b.left = node
		} else {
			b.left.insert(node)
		}
	} else {
		if b.right == nil {
			b.right = node
		} else {
			b.right.insert(node)
		}
	}
}

func (b *binaryNode) preorder(c chan<- types.Node) {
	if b == nil {
		return
	}
	b.left.preorder(c)
	c <- b
	b.right.preorder(c)
}

func (b *binaryNode) postorder(c chan<- types.Node) {
	if b == nil {
		return
	}
	b.right.postorder(c)
	c <- b
	b.left.postorder(c)
}

func (b *binaryNode) search(data types.Data, c chan<- types.Node) {
	if b.left != nil && b.left.data.LessThanOrEqual(data) {
		b.left.search(data, c)
	}
	if b.data.Equals(data) {
		c <- b
	}
	if b.right != nil && b.right.data.GreaterThanOrEqual(data) {
		b.right.search(data, c)
	}
}

func (b *binaryNode) contains(data types.Data) bool {
	if b.data.Equals(data) {
		return true
	}
	if b.left != nil && b.left.data.LessThanOrEqual(data) && b.left.contains(data) {
		return true
	}
	if b.right != nil && b.right.data.GreaterThanOrEqual(data) && b.right.contains(data) {
		return true
	}
	return false
}

//returns is balanced, and it's height. this allows us O(n) vs O(n^2)
func (b *binaryNode) isBalanced() (bool, float64) {
	leftHeight := 0.0
	leftBalanced := true
	rightHeight := 0.0
	rightBalanced := true
	if b.left != nil {
		leftBalanced, leftHeight = b.left.isBalanced()
	}
	if b.right != nil {
		rightBalanced, rightHeight = b.right.isBalanced()
	}
	return math.Abs(leftHeight-rightHeight) <= 1 && leftBalanced && rightBalanced,
		math.Max(leftHeight, rightHeight) + 1
}

func (b *binaryNode) delete(data types.Data) *binaryNode {
	if b.data.GreaterThan(data) {
		b.left = b.left.delete(data)
		return b
	}
	if b.data.LessThan(data) {
		b.right = b.right.delete(data)
		return b
	}
	if b.left == nil {
		return b.right
	}
	if b.right == nil {
		return b.left
	}
	successorParent := b.right
	successor := b.right
	for successor.left != nil {
		successorParent = successor
		successor = successor.left
	}
	successorParent.left = successor.right

	// Copy Successor Data to root
	b.data = successor.data
	return b
}
