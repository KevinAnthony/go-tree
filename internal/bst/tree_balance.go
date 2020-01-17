package bst

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
