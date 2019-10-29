package binary_search_test

import (
	"testing"

	"github.com/KevinAnthony/go-tree/internal/binary_search"
	"github.com/KevinAnthony/go-tree/types"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewTree(t *testing.T) {
	Convey("NewTree", t, func() {
		Convey("should return valid tree", func() {

		})
		Convey("should accept unordered list and return ordered list", func() {})
	})
}

func TestBinarySearchTree_Asc(t *testing.T) {
	Convey("Asc", t, func() {
		Convey("should return asc ordered list", func() {
			t := binary_search.NewTree(unordered()...)
			c := t.Asc()
			actual := make([]types.Data, 0, len(asc()))
			for data := range c {
				actual = append(actual, data.GetData())
			}
			So(actual, ShouldResemble, asc())
		})
	})
}

func TestBinarySearchTree_Desc(t *testing.T) {
	Convey("Desc", t, func() {
		Convey("should return asc ordered list", func() {
			t := binary_search.NewTree(unordered()...)
			c := t.Desc()
			actual := make([]types.Data, 0, len(asc()))
			for data := range c {
				actual = append(actual, data.GetData())
			}
			So(actual, ShouldResemble, desc())
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
		t := binary_search.NewTree(unordered()...)
		t.InsertMany(expected[1], expected[2], expected[3])
		Convey("should return 4 values when value is in the tree", func() {
			c := t.Search(types.NewInt(3))
			actual := make([]types.Data, 0, len(expected))
			for data := range c {
				actual = append(actual, data.GetData())
			}
			So(actual, ShouldResemble, expected)
		})
		Convey("should return 0 values when value is not in tree", func() {
			c := t.Search(types.NewInt(7))
			actual := make([]types.Data, 0, len(expected))
			for data := range c {
				actual = append(actual, data.GetData())
			}
			So(actual, ShouldHaveLength, 0)
		})
	})
}

func TestBinarySearchTree_Contains(t *testing.T) {
	Convey("Contains", t, func() {
		t := binary_search.NewTree(unordered()...)
		Convey("should return true when tree contains value", func() {
			So(t.Contains(types.NewInt(4)), ShouldBeTrue)
		})
		Convey("should return false when tree contains value", func() {
			So(t.Contains(types.NewInt(6)), ShouldBeFalse)
		})
	})
}

func unordered() []types.Data {
	return []types.Data{
		types.NewInt(3),
		types.NewInt(4),
		types.NewInt(1),
		types.NewInt(5),
		types.NewInt(2),
	}
}
func asc() []types.Data {
	return []types.Data{
		types.NewInt(1),
		types.NewInt(2),
		types.NewInt(3),
		types.NewInt(4),
		types.NewInt(5),
	}
}
func desc() []types.Data {
	return []types.Data{
		types.NewInt(5),
		types.NewInt(4),
		types.NewInt(3),
		types.NewInt(2),
		types.NewInt(1),
	}
}
