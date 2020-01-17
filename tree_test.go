package tree_test

import (
	"testing"

	"github.com/KevinAnthony/go-tree"
	"github.com/KevinAnthony/go-tree/internal/bst"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	Convey("New", t, func() {
		Convey("When passing Binary Tree should get a binary tree type", func() {
			t := tree.New(tree.BinarySearchTree)
			b := bst.NewTree()
			So(t, ShouldHaveSameTypeAs, b)
		})
		Convey("should panic when passing an unknown type", func() {
			f := func() {
				tree.New(tree.Type(99))
			}
			So(f, ShouldPanicWith, "invalid tree type")
		})
	})
}
