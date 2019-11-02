package mocks

import (
	"github.com/KevinAnthony/go-tree/types"

	"github.com/stretchr/testify/mock"
)

var _ types.Node = &NodeMock{}

type NodeMock struct {
	mock.Mock
}

func (m *NodeMock) GetData() types.Data {
	args := m.Called()
	if item := args.Get(0); item != nil {
		return item.(types.Data)
	}
	return nil
}

func (m *NodeMock) IsLeaf() bool {
	return m.Called().Bool(0)
}
