package tree

import (
	"github.com/KevinAnthony/go-tree/internal/bst"
	"github.com/KevinAnthony/go-tree/types"
)

type Type int

const (
	BinarySearchTree Type = 0
)

func New(t Type, data ...types.Data) types.Tree {
	switch t {
	case BinarySearchTree:
		return bst.NewTree(data...)
	default:
		panic("invalid tree type")
	}
}
