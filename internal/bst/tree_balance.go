package bst

import "github.com/KevinAnthony/go-tree/types"

func (b *binarySearchTree) Rebalance() {
	if b == nil || b.root == nil || b.root.IsLeaf() || b.isBalanced {
		return
	}
	b.mutex.Lock()
	defer b.mutex.Unlock()
	list := b.getAscList()
	b.root = b.buildTree(list, 0, len(list)-1)
}

func (b *binarySearchTree) IsBalanced() bool {
	b.detectBalance()
	return b.isBalanced
}

func (b *binarySearchTree) AutoRebalance(rebalance bool) {
	b.autoRebalance = rebalance
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

func (b *binarySearchTree) getAscList() []types.Data {
	actual := make([]types.Data, 0)
	c := b.traverse(b.root.preorder)
	for data := range c {
		actual = append(actual, data.GetData())
	}
	return actual
}

func (b *binarySearchTree) buildTree(list []types.Data, start, end int) *binaryNode {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	data := list[mid]
	return &binaryNode{
		data:  data,
		left:  b.buildTree(list, start, mid-1),
		right: b.buildTree(list, mid+1, end),
	}
}
