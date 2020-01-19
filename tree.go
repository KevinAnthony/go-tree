// Package provides generic implementation of trees with a way to get a specific implementation
package tree

import (
	"github.com/KevinAnthony/go-tree/internal/bst"
	"github.com/KevinAnthony/go-tree/types"
)

// this is used to ask for a specific type of tree
type Type int

const (
	// used to ask for a Binary Search Tree from new
	BinarySearchTree Type = 0
)

// Create a new tree, optional variadic data.
// TODO once we have more then one implementation of Data, we should restrict all inserts to be the same type as root
func New(t Type, data ...types.Data) types.Tree {
	switch t {
	case BinarySearchTree:
		return bst.NewTree(data...)
	default:
		panic("invalid tree type")
	}
}
