package types

import "fmt"

type Node interface {
	GetData() Data
	IsLeaf() bool
}

type Tree interface {
	Count() int64
	Insert(value Data)
	InsertMany(values ...Data)
	Delete(value Data)
	Search(Data) <-chan Node
	Contains(Data) bool
	Asc() <-chan Node
	Desc() <-chan Node
	Rebalance()
	IsBalanced() bool
	AutoRebalance(bool)
}

type Data interface {
	fmt.Stringer
	GreaterThan(Data) bool
	LessThan(Data) bool
	GreaterThanOrEqual(Data) bool
	LessThanOrEqual(Data) bool
	Equals(Data) bool
	Get() interface{}
}
