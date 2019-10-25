package tree


type Node interface{
	GetData() NodeData
	IsLeaf() bool
}

type Tree interface{
	Count() int64
	Insert(value NodeData)
	Delete(value NodeData,allInstances bool) bool
	Search(NodeData) Node
	Asc() <-chan Node
	Desc() <-chan Node
}

type NodeData interface{
	GreaterThan(NodeData) bool
	LessThan(NodeData) bool
	GreaterThanOrEqual(NodeData) bool
	LessThanOrEqual(NodeData) bool
	Equals(NodeData) bool
}