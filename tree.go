package tree

import (
	"github.com/KevinAnthony/go-tree/internal/binary_search"
	"github.com/KevinAnthony/go-tree/types"
)

type Type int

const (
	BinaryTree Type = 0
)

func New(t Type, data ...types.Data) types.Tree {
	switch t {
	case BinaryTree:
		return binary_search.NewTree(data...)
	default:
		panic("invalid tree type")
	}
}
