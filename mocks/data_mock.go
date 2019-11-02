package mocks

import (
	"github.com/KevinAnthony/go-tree/types"

	"github.com/stretchr/testify/mock"
)

var _ types.Data = &DataMock{}

type DataMock struct {
	mock.Mock
}

func (m *DataMock) GreaterThan(data types.Data) bool {
	return m.Called(data).Bool(0)
}

func (m *DataMock) LessThan(data types.Data) bool {
	return m.Called(data).Bool(0)
}

func (m *DataMock) GreaterThanOrEqual(data types.Data) bool {
	return m.Called(data).Bool(0)
}

func (m *DataMock) LessThanOrEqual(data types.Data) bool {
	return m.Called(data).Bool(0)
}

func (m *DataMock) Equals(data types.Data) bool {
	return m.Called(data).Bool(0)
}
