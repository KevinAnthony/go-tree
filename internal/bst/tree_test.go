//nolint: funlen, gomnd
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
				actual := drain(t.Asc())
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
			actual := drain(t.Asc())
			So(actual, ShouldResemble, asc())
		})
		Convey("should return no values when tree is empty", func() {
			t := bst.NewTree()
			actual := drain(t.Asc())
			So(actual, ShouldBeEmpty)
		})
	})
}

func TestBinarySearchTree_Desc(t *testing.T) {
	Convey("Desc", t, func() {
		Convey("should return asc ordered list", func() {
			t := bst.NewTree(unordered()...)
			actual := drain(t.Desc())
			So(actual, ShouldResemble, desc())
		})
		Convey("should return no values when tree is empty", func() {
			t := bst.NewTree()
			actual := drain(t.Desc())
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
				actual := drain(t.Search(types.NewInt(3)))
				So(actual, ShouldBeEmpty)
			})
		})
		Convey("when tree has values", func() {
			t := bst.NewTree(unordered()...)
			t.InsertMany(expected[1], expected[2], expected[3])
			Convey("should return 4 values when value is in the tree", func() {
				actual := drain(t.Search(types.NewInt(3)))
				So(actual, ShouldResemble, expected)
			})
			Convey("should return 0 values when value is not in tree", func() {
				actual := drain(t.Search(types.NewInt(77)))
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
			actual := drain(t.Asc())
			So(actual, ShouldHaveLength, 0)
			So(t.Count(), ShouldEqual, 0)
		})
		Convey("should keep other nodes when deleting root", func() {
			Convey("and count == 1", func() {
				t := bst.NewTree(types.NewInt(11))
				t.Delete(types.NewInt(11))
				actual := drain(t.Asc())
				So(actual, ShouldHaveLength, 0)
				So(t.Count(), ShouldEqual, 0)
			})
			Convey("and we try and delete every node", func() {
				for _, data := range unordered() {
					closure := data
					Convey(fmt.Sprintf("try and delete node: %s", closure), func() {
						t.Delete(closure)
						actual := drain(t.Asc())
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
						actual := drain(t.Asc())
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
						actual := drain(t.Asc())
						So(actual, ShouldNotContain, closure)
						So(actual, ShouldHaveLength, len(right)-1)
						So(t.Count(), ShouldEqual, len(right)-1)
					})
				}
			})
		})
		Convey("should keep other nodes when deleting nothing, called with missing type", func() {
			t.Delete(types.NewInt(99))
			actual := drain(t.Asc())
			So(actual, ShouldHaveLength, len(unordered()))
			So(t.Count(), ShouldEqual, len(unordered()))
		})
	})
}

func TestBinarySearchTree_Insert(t *testing.T) {
	Convey("Insert", t, func() {
		t := bst.NewTree()
		start := drain(t.Asc())
		So(start, ShouldBeEmpty)
		Convey("should insert into tree", func() {
			t.Insert(types.NewInt(7))
			intermitten := drain(t.Asc())
			So(intermitten, ShouldHaveLength, 1)
			So(intermitten, ShouldContain, types.NewInt(7))
			t.Insert(types.NewInt(77))
			final := drain(t.Asc())
			So(final, ShouldHaveLength, 2)
			So(final, ShouldContain, types.NewInt(7))
			So(final, ShouldContain, types.NewInt(77))
		})
	})
}

func TestBinarySearchTree_InsertMany(t *testing.T) {
	Convey("Insert", t, func() {
		t := bst.NewTree()
		start := drain(t.Asc())
		So(start, ShouldBeEmpty)
		Convey("should insert one into tree", func() {
			t.InsertMany(types.NewInt(7))
			actual := drain(t.Asc())
			So(actual, ShouldHaveLength, 1)
			So(actual, ShouldContain, types.NewInt(7))
		})
		Convey("should insert many into tree", func() {
			expected := unordered()
			t.InsertMany(expected...)
			actual := drain(t.Asc())
			So(actual, ShouldHaveLength, len(expected))
			for _, exp := range expected {
				So(actual, ShouldContain, exp)
			}
		})
	})
}

func TestBinarySearchTree_IsBalanced(t *testing.T) {
	Convey("IsBalanced", t, func() {
		t := bst.NewTree()
		Convey("should be balanced if empty", func() {
			So(t.IsBalanced(), ShouldBeTrue)
		})
		Convey("should be balanced with", func() {
			Convey("one node", func() {
				t.Insert(types.NewInt(7))
				So(t.IsBalanced(), ShouldBeTrue)
			})
			Convey("one + left", func() {
				t.Insert(types.NewInt(7))
				t.Insert(types.NewInt(6))
				So(t.IsBalanced(), ShouldBeTrue)
			})
			Convey("one + right", func() {
				t.Insert(types.NewInt(7))
				t.Insert(types.NewInt(8))
				So(t.IsBalanced(), ShouldBeTrue)
			})
			Convey("even tree", func() {
				t.Insert(types.NewInt(7))
				t.Insert(types.NewInt(8))
				t.Insert(types.NewInt(6))
				So(t.IsBalanced(), ShouldBeTrue)
			})
		})
		Convey("should report as unbalanced", func() {
			Convey("when left heavy", func() {
				t.Insert(types.NewInt(3))
				t.Insert(types.NewInt(2))
				t.Insert(types.NewInt(1))
				So(t.IsBalanced(), ShouldBeFalse)
			})
			Convey("when right heavy", func() {
				t.Insert(types.NewInt(1))
				t.Insert(types.NewInt(2))
				t.Insert(types.NewInt(3))
				So(t.IsBalanced(), ShouldBeFalse)
			})
			Convey("when randomly unbalanced", func() {
				t.InsertMany(unordered()...)
				So(t.IsBalanced(), ShouldBeFalse)
			})
		})
	})
}

func TestBinarySearchTree_AutoRebalance(t *testing.T) {
	Convey("AutoRebalance", t, func() {
		t := bst.NewTree()
		Convey("should respect autorelance flag when it's turned on", func() {
			t.AutoRebalance(true)
			t.InsertMany(unordered()...)
			So(t.IsBalanced(), ShouldBeTrue)
		})
		Convey("should autobalance when flag is turned on", func() {
			t.AutoRebalance(false)
			t.InsertMany(unordered()...)
			So(t.IsBalanced(), ShouldBeFalse)
			t.AutoRebalance(true)
			t.Insert(types.NewInt(999))
			So(t.IsBalanced(), ShouldBeTrue)
		})
		Convey("should not autobalance once flag is turned off", func() {
			t.AutoRebalance(true)
			t.InsertMany(unordered()...)
			So(t.IsBalanced(), ShouldBeTrue)
			t.AutoRebalance(false)
			t.Insert(types.NewInt(997))
			t.Insert(types.NewInt(998))
			t.Insert(types.NewInt(999))
			So(t.IsBalanced(), ShouldBeFalse)
		})
	})
}

func TestBinarySearchTree_Rebalance(t *testing.T) {
	Convey("Rebalance", t, func() {
		t := bst.NewTree()
		t.AutoRebalance(false)
		Convey("should pass when root is nil", func() {
			t.Rebalance()
			So(t.IsBalanced(), ShouldBeTrue)
		})
		Convey("should pass when root is leaf", func() {
			t.Insert(types.NewInt(7))
			t.Rebalance()
			So(t.IsBalanced(), ShouldBeTrue)
		})
		Convey("should pass tree is balanced", func() {
			t.Insert(types.NewInt(7))
			t.Insert(types.NewInt(6))
			t.Insert(types.NewInt(8))
			t.Rebalance()
			So(t.IsBalanced(), ShouldBeTrue)
		})
		Convey("when tree is not balanced", func() {
			Convey("and is right heavy", func() {
				t.InsertMany(types.NewInt(3), types.NewInt(2), types.NewInt(1))
				So(t.IsBalanced(), ShouldBeFalse)
				t.Rebalance()
				So(t.IsBalanced(), ShouldBeTrue)
			})
			Convey("and is left heavy", func() {
				t.InsertMany(types.NewInt(1), types.NewInt(2), types.NewInt(3))
				So(t.IsBalanced(), ShouldBeFalse)
				t.Rebalance()
				So(t.IsBalanced(), ShouldBeTrue)
			})
			Convey("and is random", func() {
				t.InsertMany(unordered()...)
				So(t.IsBalanced(), ShouldBeFalse)
				t.Rebalance()
				So(t.IsBalanced(), ShouldBeTrue)
			})
		})
	})
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
