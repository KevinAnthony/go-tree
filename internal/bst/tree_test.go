package bst_test

import (
	"fmt"
	"testing"

	"github.com/KevinAnthony/go-tree/internal/bst"
	"github.com/KevinAnthony/go-tree/types"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewTree(t *testing.T) {
	Convey("NewTree", t, func() {
		Convey("should return valid tree", func() {
			f := func() {
				bst.NewTree()
			}
			So(f, ShouldNotPanic)
		})
		Convey("should accept unordered list and return ordered list", func() {
			f := func() {
				t := bst.NewTree(unordered()...)
				c := t.Asc()
				actual := drain(c)
				So(actual, ShouldResemble, asc())
			}
			So(f, ShouldNotPanic)
		})
	})
}

func TestBinarySearchTree_Asc(t *testing.T) {
	Convey("Asc", t, func() {
		Convey("should return asc ordered list", func() {
			t := bst.NewTree(unordered()...)
			c := t.Asc()
			actual := drain(c)
			So(actual, ShouldResemble, asc())
		})
		Convey("should return no values when tree is empty", func() {
			t := bst.NewTree()
			c := t.Asc()
			actual := drain(c)
			So(actual, ShouldBeEmpty)
		})
	})
}

func TestBinarySearchTree_Desc(t *testing.T) {
	Convey("Desc", t, func() {
		Convey("should return asc ordered list", func() {
			t := bst.NewTree(unordered()...)
			c := t.Desc()
			actual := drain(c)
			So(actual, ShouldResemble, desc())
		})
		Convey("should return no values when tree is empty", func() {
			t := bst.NewTree()
			c := t.Desc()
			actual := drain(c)
			So(actual, ShouldBeEmpty)
		})
	})
}

func TestBinarySearchTree_Search(t *testing.T) {
	Convey("Search", t, func() {
		expected := []types.Data{
			unordered()[0],
			types.NewInt(3),
			types.NewInt(3),
			types.NewInt(3),
		}
		Convey("when tree is empty", func() {
			t := bst.NewTree()
			Convey("should return no data", func() {
				c := t.Search(types.NewInt(3))
				actual := drain(c)
				So(actual, ShouldBeEmpty)
			})
		})
		Convey("when tree has values", func() {
			t := bst.NewTree(unordered()...)
			t.InsertMany(expected[1], expected[2], expected[3])
			Convey("should return 4 values when value is in the tree", func() {
				c := t.Search(types.NewInt(3))
				actual := drain(c)
				So(actual, ShouldResemble, expected)
			})
			Convey("should return 0 values when value is not in tree", func() {
				c := t.Search(types.NewInt(77))
				actual := drain(c)
				So(actual, ShouldHaveLength, 0)
			})
		})
	})
}

func TestBinarySearchTree_Contains(t *testing.T) {
	Convey("Contains", t, func() {
		Convey("when tree is empty", func() {
			t := bst.NewTree()
			Convey("should not contain anything", func() {
				So(t.Contains(types.NewInt(6)), ShouldBeFalse)
			})
		})
		Convey("when tree has data", func() {
			t := bst.NewTree(unordered()...)
			Convey("should return true when tree contains value", func() {
				So(t.Contains(types.NewInt(4)), ShouldBeTrue)
			})
			Convey("should return false when tree does not contain value", func() {
				So(t.Contains(types.NewInt(66)), ShouldBeFalse)
			})
		})
	})
}

func TestBinarySearchTree_AutoRebalance(t *testing.T) {
	//TODO the best way to test this seems to be to have a tree that needs balancing,
	// 		turn on autobalancing, and see if it's balanced
	// The 2nd convey is to have a tree that IS balnanced,
	// 		turn off authbalancing, unbalance it, then make sure it remains unbalanced.
}

func TestBinarySearchTree_Count(t *testing.T) {
	Convey("Count", t, func() {
		t := bst.NewTree()
		Convey("should return 0 on empty tree", func() {
			So(t.Count(), ShouldEqual, 0)
		})
		Convey("should return correct count", func() {
			for i := 0; i <= 1024; i++ {
				t.Insert(types.NewInt(i))
				So(t.Count(), ShouldEqual, i+1)
			}
		})
	})
}

func TestBinarySearchTree_Delete(t *testing.T) {
	Convey("Delete", t, func() {
		t := bst.NewTree(unordered()...)
		t.AutoRebalance(false)
		Convey("should return tree when root is nil", func() {
			t := bst.NewTree()
			t.Delete(types.NewInt(77))
			c := t.Asc()
			actual := drain(c)
			So(actual, ShouldHaveLength, 0)
			So(t.Count(), ShouldEqual, 0)
		})
		Convey("should keep other nodes when deleting root", func() {
			Convey("and count == 1", func() {
				t := bst.NewTree(types.NewInt(11))
				t.Delete(types.NewInt(11))
				c := t.Asc()
				actual := drain(c)
				So(actual, ShouldHaveLength, 0)
				So(t.Count(), ShouldEqual, 0)
			})
			Convey("and we try and delete every node", func() {
				for _, data := range unordered() {
					closure := data
					Convey(fmt.Sprintf("try and delete node: %s", closure), func() {
						t.Delete(closure)
						c := t.Asc()
						actual := drain(c)
						So(actual, ShouldNotContain, closure)
						So(actual, ShouldHaveLength, len(unordered())-1)
						So(t.Count(), ShouldEqual, len(unordered())-1)
					})
				}
			})
			Convey("should delete successfully when the tree is all left", func() {
				left := []types.Data{
					types.NewInt(5),
					types.NewInt(4),
					types.NewInt(3),
					types.NewInt(2),
					types.NewInt(1)}
				t := bst.NewTree(left...)
				for _, data := range left {
					closure := data
					Convey(fmt.Sprintf("try and delete node: %s", closure), func() {
						t.Delete(closure)
						c := t.Asc()
						actual := drain(c)
						So(actual, ShouldNotContain, closure)
						So(actual, ShouldHaveLength, len(left)-1)
						So(t.Count(), ShouldEqual, len(left)-1)
					})
				}
			})
			Convey("should delete successfully when the tree is all right", func() {
				right := []types.Data{
					types.NewInt(1),
					types.NewInt(2),
					types.NewInt(3),
					types.NewInt(4),
					types.NewInt(5)}
				t := bst.NewTree(right...)
				for _, data := range right {
					closure := data
					Convey(fmt.Sprintf("try and delete node: %s", closure), func() {
						t.Delete(closure)
						c := t.Asc()
						actual := drain(c)
						So(actual, ShouldNotContain, closure)
						So(actual, ShouldHaveLength, len(right)-1)
						So(t.Count(), ShouldEqual, len(right)-1)
					})
				}
			})
		})
		Convey("should keep other nodes when deleting nothing, called with missing type", func() {
			t.Delete(types.NewInt(99))
			c := t.Asc()
			actual := drain(c)
			So(actual, ShouldHaveLength, len(unordered()))
			So(t.Count(), ShouldEqual, len(unordered()))
		})
	})
}

func TestBinarySearchTree_Insert(t *testing.T) {

}

func TestBinarySearchTree_InsertMany(t *testing.T) {

}

func TestBinarySearchTree_IsBalanced(t *testing.T) {

}

func TestBinarySearchTree_Rebalance(t *testing.T) {

}

//unbalanced tree
func unordered() []types.Data {
	return []types.Data{
		types.NewInt(3),
		types.NewInt(6),
		types.NewInt(4),
		types.NewInt(9),
		types.NewInt(7),
		types.NewInt(8),
		types.NewInt(1),
		types.NewInt(5),
		types.NewInt(2),
		types.NewInt(10),
	}
}

func asc() []types.Data {
	return []types.Data{
		types.NewInt(1),
		types.NewInt(2),
		types.NewInt(3),
		types.NewInt(4),
		types.NewInt(5),
		types.NewInt(6),
		types.NewInt(7),
		types.NewInt(8),
		types.NewInt(9),
		types.NewInt(10),
	}
}

func desc() []types.Data {
	return []types.Data{
		types.NewInt(10),
		types.NewInt(9),
		types.NewInt(8),
		types.NewInt(7),
		types.NewInt(6),
		types.NewInt(5),
		types.NewInt(4),
		types.NewInt(3),
		types.NewInt(2),
		types.NewInt(1),
	}
}

func drain(c <-chan types.Node) []types.Data {
	actual := make([]types.Data, 0)
	for data := range c {
		actual = append(actual, data.GetData())
	}
	return actual
}
