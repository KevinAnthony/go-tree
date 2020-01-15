package mocks

import (
	"github.com/KevinAnthony/go-tree/types"

	"github.com/stretchr/testify/mock"
)

var _ types.Tree = &TreeMock{}

type TreeMock struct {
	mock.Mock
}

func (m *TreeMock) Count() int64 {
	return int64(m.Called().Int(0))
}

func (m *TreeMock) Insert(value types.Data) {
	m.Called(value)
}

func (m *TreeMock) InsertMany(values ...types.Data) {
	m.Called(values)
}

func (m *TreeMock) Delete(value types.Data) {
	m.Called(value)
}

func (m *TreeMock) Search(value types.Data) <-chan types.Node {
	args := m.Called(value)
	if item := args.Get(0); item != nil {
		return item.(chan types.Node)
	}
	return nil
}

func (m *TreeMock) Contains(value types.Data) bool {
	return m.Called(value).Bool(0)
}

func (m *TreeMock) Asc() <-chan types.Node {
	args := m.Called()
	if item := args.Get(0); item != nil {
		return item.(chan types.Node)
	}
	return nil
}

func (m *TreeMock) Desc() <-chan types.Node {
	args := m.Called()
	if item := args.Get(0); item != nil {
		return item.(chan types.Node)
	}
	return nil
}

func (m *TreeMock) Rebalance() {
	m.Called()
}

func (m *TreeMock) IsBalanced() bool {
	return m.Called().Bool(0)
}

func (m *TreeMock) AutoRebalance(auto bool) {
	m.Called(auto)
}
