package types

type Node interface {
	GetData() Data
	IsLeaf() bool
}

type Tree interface {
	Count() int64
	Insert(value Data)
	InsertMany(value ...Data)
	Delete(value Data, allInstances bool) bool
	Search(Data) <-chan Node
	Contains(Data) bool
	Asc() <-chan Node
	Desc() <-chan Node
	Rebalance()
	IsBalanced() bool
	AutoRebalance(bool)
}

type Data interface {
	GreaterThan(Data) bool
	LessThan(Data) bool
	GreaterThanOrEqual(Data) bool
	LessThanOrEqual(Data) bool
	Equals(Data) bool
}
