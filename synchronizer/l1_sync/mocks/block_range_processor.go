// Code generated by mockery. DO NOT EDIT.

package mock_l1sync

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	etherman "github.com/0xPolygonHermez/zkevm-synchronizer-l1/etherman"

	mock "github.com/stretchr/testify/mock"
)

// BlockRangeProcessor is an autogenerated mock type for the BlockRangeProcessor type
type BlockRangeProcessor struct {
	mock.Mock
}

type BlockRangeProcessor_Expecter struct {
	mock *mock.Mock
}

func (_m *BlockRangeProcessor) EXPECT() *BlockRangeProcessor_Expecter {
	return &BlockRangeProcessor_Expecter{mock: &_m.Mock}
}

// ProcessBlockRange provides a mock function with given fields: ctx, blocks, order, finalizedBlockNumber
func (_m *BlockRangeProcessor) ProcessBlockRange(ctx context.Context, blocks []etherman.Block, order map[common.Hash][]etherman.Order, finalizedBlockNumber uint64) error {
	ret := _m.Called(ctx, blocks, order, finalizedBlockNumber)

	if len(ret) == 0 {
		panic("no return value specified for ProcessBlockRange")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []etherman.Block, map[common.Hash][]etherman.Order, uint64) error); ok {
		r0 = rf(ctx, blocks, order, finalizedBlockNumber)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BlockRangeProcessor_ProcessBlockRange_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessBlockRange'
type BlockRangeProcessor_ProcessBlockRange_Call struct {
	*mock.Call
}

// ProcessBlockRange is a helper method to define mock.On call
//   - ctx context.Context
//   - blocks []etherman.Block
//   - order map[common.Hash][]etherman.Order
//   - finalizedBlockNumber uint64
func (_e *BlockRangeProcessor_Expecter) ProcessBlockRange(ctx interface{}, blocks interface{}, order interface{}, finalizedBlockNumber interface{}) *BlockRangeProcessor_ProcessBlockRange_Call {
	return &BlockRangeProcessor_ProcessBlockRange_Call{Call: _e.mock.On("ProcessBlockRange", ctx, blocks, order, finalizedBlockNumber)}
}

func (_c *BlockRangeProcessor_ProcessBlockRange_Call) Run(run func(ctx context.Context, blocks []etherman.Block, order map[common.Hash][]etherman.Order, finalizedBlockNumber uint64)) *BlockRangeProcessor_ProcessBlockRange_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]etherman.Block), args[2].(map[common.Hash][]etherman.Order), args[3].(uint64))
	})
	return _c
}

func (_c *BlockRangeProcessor_ProcessBlockRange_Call) Return(_a0 error) *BlockRangeProcessor_ProcessBlockRange_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BlockRangeProcessor_ProcessBlockRange_Call) RunAndReturn(run func(context.Context, []etherman.Block, map[common.Hash][]etherman.Order, uint64) error) *BlockRangeProcessor_ProcessBlockRange_Call {
	_c.Call.Return(run)
	return _c
}

// NewBlockRangeProcessor creates a new instance of BlockRangeProcessor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlockRangeProcessor(t interface {
	mock.TestingT
	Cleanup(func())
}) *BlockRangeProcessor {
	mock := &BlockRangeProcessor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
