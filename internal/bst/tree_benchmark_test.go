//nolint: scopelint
package bst_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/KevinAnthony/go-tree/internal/bst"
	"github.com/KevinAnthony/go-tree/types"
)

const max = 256

func BenchmarkBinarySearchTree_Insert(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= max; i *= 2 {
		b.Run(fmt.Sprintf("Insert/%d", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				t := bst.NewTree()
				for j := 0; j < i; j++ {
					t.Insert(types.NewInt(rand.Intn(10*10 ^ 8)))
				}
			}
		})
	}
}

func BenchmarkBinarySearchTree_IsBalanced(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= max; i *= 2 {
		b.Run(fmt.Sprintf("IsBalanced/%d", i), func(b *testing.B) {
			t, _ := generateTree(b, i)
			for n := 0; n < b.N; n++ {
				t.IsBalanced()
			}
		})
	}
}

func BenchmarkBinarySearchTree_Delete(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= max; i *= 2 {
		b.Run(fmt.Sprintf("Delete/%d", i), func(b *testing.B) {
			t, values := generateTree(b, i)
			for n := 0; n < b.N; n++ {
				index := rand.Intn(i)
				t.Delete(values[index])
			}
		})
	}
}

func BenchmarkBinarySearchTree_Search(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= max; i *= 2 {
		b.Run(fmt.Sprintf("Search/%d", i), func(b *testing.B) {
			t, values := generateTree(b, i)
			for n := 0; n < b.N; n++ {
				index := rand.Intn(i)
				for range t.Search(values[index]) {
				}
			}
		})
	}
}

func BenchmarkBinarySearchTree_Rebalance(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= max; i *= 2 {
		b.Run(fmt.Sprintf("Rebalance/%d", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				t, _ := generateTree(b, i)
				t.Rebalance()
			}
		})
	}
}

func BenchmarkBinarySearchTree_Contains(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= max; i *= 2 {
		b.Run(fmt.Sprintf("Contains/%d", i), func(b *testing.B) {
			t, values := generateTree(b, i)
			for n := 0; n < b.N; n++ {
				index := rand.Intn(i)
				t.Contains(values[index])
			}
		})
	}
}

func generateTree(b *testing.B, count int) (types.Tree, []types.Data) {
	b.StopTimer()
	values := make([]types.Data, 0, count)
	t := bst.NewTree()
	t.AutoRebalance(false)
	for j := 0; j < count; j++ {
		d := types.NewInt(rand.Intn(10*10 ^ 8))
		t.Insert(d)
		values = append(values, d)
	}
	b.StartTimer()
	return t, values
}
