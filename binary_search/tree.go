package binary_search

import "github.com/KevinAnthony/go-tree"

type binarySearchTree struct {
	count int64
	root  *binaryNode
}

func NewTree(data ...tree.Data) tree.Tree {
	tree := binarySearchTree{}
	for _, d := range data {
		tree.Insert(d)
	}
	return &tree
}

func (b binarySearchTree) Count() int64 {
	return b.count
}

func (b binarySearchTree) Insert(value tree.Data) {
	node := &binaryNode{
		data:  value,
		left:  nil,
		right: nil,
	}
	b.count++
	if b.root == nil {
		b.root = node
		return
	}
	if b.root.GetData().GreaterThan(value) {
		b.insert(b.root.left, node)
	} else {
		b.insert(b.root.right, node)
	}

}

func (b binarySearchTree) Delete(value tree.Data, allInstances bool) bool {
	panic("implement me")
}

func (b binarySearchTree) Search(tree.Data) tree.Node {
	panic("implement me")
}

func (b binarySearchTree) Asc() <-chan tree.Node {
	panic("implement me")
}

func (b binarySearchTree) Desc() <-chan tree.Node {
	panic("implement me")
}

func (b binarySearchTree) insert(treeNode *binaryNode, node *binaryNode) {
	if treeNode == nil {
		treeNode = node
		return
	}
	if treeNode.data.GreaterThan(node.data) {
		b.insert(b.root.left, node)
	} else {
		b.insert(b.root.right, node)
	}
}
