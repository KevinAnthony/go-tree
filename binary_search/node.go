package binary_search

import "github.com/KevinAnthony/tree"

type binaryNode struct {
	data tree.NodeData
	left *binaryNode
	right *binaryNode
}

func NewNode(data interface{}) tree.Node{
	return binaryNode{
		data:  data,
		left:  nil,
		right: nil,
	}
}

func (b binaryNode) GetData() tree.NodeData {
	return b.data
}

func (b binaryNode) IsLeaf() bool {
	return b.left == nil && b.right == nil
}