package bst

import "github.com/KevinAnthony/go-tree/types"

func (b *binarySearchTree) Delete(data types.Data) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if b.root == nil {
		return
	}
	// root has no children
	if b.root.IsLeaf() {
		// if we match, delete
		if b.root.data.Equals(data) {
			b.root = nil
			b.count = 0
		}
		return
	}

	parent, node := b.root.find(data)
	if parent == nil || node == nil {
		return
	}
	b.count--
	switch {
	case node.IsLeaf():
		b.deleteIsLeaf(parent, node)
		return
	case node.left == nil:
		b.deleteLeftIsNil(parent, node)
		return

	case node.right == nil:
		b.deleteRightIsNil(parent, node)
		return
	case node.right.left == nil:
		b.deleteSuperSpecialCase(node)
	default:
		b.deleteHardCase(node)
	}
}

func (b *binarySearchTree) deleteRightIsNil(parent *binaryNode, node *binaryNode) {
	if node == b.root {
		b.root = node.left
		return
	}
	if parent.left == node {
		parent.left = node.left
	} else {
		parent.right = node.left
	}
}

func (b *binarySearchTree) deleteIsLeaf(parent *binaryNode, node *binaryNode) {
	if parent.left == node {
		parent.left = nil
	} else {
		parent.right = nil
	}
}

func (b *binarySearchTree) deleteLeftIsNil(parent *binaryNode, node *binaryNode) {
	if node == b.root {
		b.root = node.right
		return
	}
	if parent.left == node {
		parent.left = node.right
	} else {
		parent.right = node.right
	}
}

func (b *binarySearchTree) deleteSuperSpecialCase(node *binaryNode) {
	node.data = node.right.data
	node.right = node.right.right
}

func (b *binarySearchTree) deleteHardCase(node *binaryNode) {
	successor := node.right
	successorRent := node
	for successor.left != nil {
		successorRent = successor
		successor = successor.left
	}
	node.data = successor.data
	successorRent.left = successor.right
}
