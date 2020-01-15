package bst

import (
	"sync"

	"github.com/KevinAnthony/go-tree/types"
)

type binarySearchTree struct {
	isBalanced    bool
	autoRebalance bool
	count         int64
	root          *binaryNode
	mutex         sync.RWMutex
}

func NewTree(data ...types.Data) types.Tree {
	t := binarySearchTree{
		autoRebalance: true,
	}
	if len(data) > 0 {
		t.InsertMany(data...)
	}
	return &t
}

func (b *binarySearchTree) Insert(value types.Data) {
	node := &binaryNode{
		data:  value,
		left:  nil,
		right: nil,
	}
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.count++
	if b.root == nil {
		b.root = node
		return
	}
	b.root.insert(node)
}

// NOTE: we do not acquire the mutex here because insert uses it.
// if you mutex on both insert  and  insertmany then you will deadlock
func (b *binarySearchTree) InsertMany(values ...types.Data) {
	for _, value := range values {
		b.Insert(value)
	}
	b.autoRebalanceMaybe()
}

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
		if parent.left == node {
			parent.left = nil
		} else {
			parent.right = nil
		}
		return
	case node.left == nil:
		if node == b.root {
			b.root = node.right
			return
		}
		if parent.left == node {
			parent.left = node.right
		} else {
			parent.right = node.right
		}
		return

	case node.right == nil:
		if node == b.root {
			b.root = node.left
			return
		}
		if parent.left == node {
			parent.left = node.left
		} else {
			parent.right = node.left
		}
		return
	case node.right.left == nil:
		node.data = node.right.data
		node.right = node.right.right
	default:
		successor := node.right
		successorRent := node
		for successor.left != nil {
			successorRent = successor
			successor = successor.left
		}
		node.data = successor.data
		successorRent.left = successor.right
	}
}

func (b *binarySearchTree) Search(data types.Data) <-chan types.Node {
	c := make(chan types.Node)
	go func(c chan<- types.Node) {
		defer close(c)
		b.mutex.RLock()
		defer b.mutex.RUnlock()
		if b.count == 0 {
			return
		}
		b.root.search(data, c)
	}(c)
	return c
}

func (b *binarySearchTree) Contains(data types.Data) bool {
	if b.root == nil {
		return false
	}
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	return b.root.contains(data)
}

func (b *binarySearchTree) Count() int64 {
	return b.count
}

func (b *binarySearchTree) Asc() <-chan types.Node {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	return b.traverse(b.root.preorder)
}

func (b *binarySearchTree) Desc() <-chan types.Node {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	return b.traverse(b.root.postorder)
}

func (b *binarySearchTree) Rebalance() {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	//TODO implement
}

func (b *binarySearchTree) IsBalanced() bool {
	b.detectBalance()
	return b.isBalanced
}

func (b *binarySearchTree) AutoRebalance(rebalance bool) {
	b.autoRebalance = rebalance
}

func (b *binarySearchTree) traverse(f func(c chan<- types.Node)) <-chan types.Node {
	c := make(chan types.Node)
	go func(c chan<- types.Node) {
		defer close(c)
		if b.count == 0 {
			return
		}
		f(c)
	}(c)
	return c
}

func (b *binarySearchTree) detectBalance() {
	if b.root == nil {
		b.isBalanced = true
	}
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	// we ignore the height, since it is only used for calculations
	b.isBalanced, _ = b.root.isBalanced()
}

//TODO i would love for this to be async
func (b *binarySearchTree) autoRebalanceMaybe() {
	if b.autoRebalance {
		if !b.IsBalanced() {
			b.Rebalance()
		}
	}
}
