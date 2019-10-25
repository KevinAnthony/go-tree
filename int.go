package tree

type intData int

func NewInt(i int) NodeData {
	d := intData(i)
	return &d
}

func (i intData) GreaterThan(d NodeData) bool {
	return i > d.(intData)
}

func (i intData) LessThan(d NodeData) bool {
	return i < d.(intData)
}

func (i intData) GreaterThanOrEqual(d NodeData) bool {
	return i >= d.(intData)
}

func (i intData) LessThanOrEqual(d NodeData) bool {
	return i <= d.(intData)
}

func (i intData) Equals(d NodeData) bool {
	return i == d.(intData)
}
