package tree

type intData int

func NewInt(i int) Data {
	d := intData(i)
	return &d
}

func (i intData) GreaterThan(d Data) bool {
	return i > d.(intData)
}

func (i intData) LessThan(d Data) bool {
	return i < d.(intData)
}

func (i intData) GreaterThanOrEqual(d Data) bool {
	return i >= d.(intData)
}

func (i intData) LessThanOrEqual(d Data) bool {
	return i <= d.(intData)
}

func (i intData) Equals(d Data) bool {
	return i == d.(intData)
}
