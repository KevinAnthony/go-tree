package binary_search

import "github.com/KevinAnthony/go-tree"

type binaryNode struct {
	data  tree.Data
	left  *binaryNode
	right *binaryNode
}

func NewNode(data tree.Data) tree.Node {
	return binaryNode{
		data:  data,
		left:  nil,
		right: nil,
	}
}

func (b binaryNode) GetData() tree.Data {
	return b.data
}

func (b binaryNode) IsLeaf() bool {
	return b.left == nil && b.right == nil
}
