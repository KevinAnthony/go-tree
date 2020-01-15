package types

import "strconv"

type intData int

func NewInt(i int) Data {
	d := intData(i)
	return &d
}

func (i intData) String() string {
	return strconv.Itoa(int(i))
}

func (i intData) GreaterThan(d Data) bool {
	return i > getIntData(d)
}

func (i intData) LessThan(d Data) bool {
	return i < getIntData(d)
}

func (i intData) GreaterThanOrEqual(d Data) bool {
	return i >= getIntData(d)
}

func (i intData) LessThanOrEqual(d Data) bool {
	return i <= getIntData(d)
}

func (i intData) Equals(d Data) bool {
	return i == getIntData(d)
}

func getIntData(d Data) intData {
	switch i := d.(type) {
	case intData:
		return i
	case *intData:
		return *i
	default:
		panic("type not intData or *intData")
	}
}
