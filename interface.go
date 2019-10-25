package tree

type Node interface {
	GetData() Data
	IsLeaf() bool
}

type Tree interface {
	Count() int64
	Insert(value Data)
	Delete(value Data, allInstances bool) bool
	Search(Data) Node
	Asc() <-chan Node
	Desc() <-chan Node
}

type Data interface {
	GreaterThan(Data) bool
	LessThan(Data) bool
	GreaterThanOrEqual(Data) bool
	LessThanOrEqual(Data) bool
	Equals(Data) bool
}
