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
	b.insert(value)
	b.autoRebalanceMaybe()
}

// NOTE: we do not acquire the mutex here because insert uses it.
// if you mutex on both insert  and  insertmany then you will deadlock
func (b *binarySearchTree) InsertMany(values ...types.Data) {
	for _, value := range values {
		b.insert(value)
	}
	b.autoRebalanceMaybe()
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

func (b *binarySearchTree) insert(value types.Data) {
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
