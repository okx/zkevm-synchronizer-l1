// Code generated by mockery. DO NOT EDIT.

package mock_storage

import (
	context "context"

	entities "github.com/0xPolygonHermez/zkevm-synchronizer-l1/state/entities"
	mock "github.com/stretchr/testify/mock"
)

// reorgStorer is an autogenerated mock type for the reorgStorer type
type reorgStorer struct {
	mock.Mock
}

type reorgStorer_Expecter struct {
	mock *mock.Mock
}

func (_m *reorgStorer) EXPECT() *reorgStorer_Expecter {
	return &reorgStorer_Expecter{mock: &_m.Mock}
}

// ResetToL1BlockNumber provides a mock function with given fields: ctx, firstBlockNumberToKeep, dbTx
func (_m *reorgStorer) ResetToL1BlockNumber(ctx context.Context, firstBlockNumberToKeep uint64, dbTx entities.Tx) error {
	ret := _m.Called(ctx, firstBlockNumberToKeep, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for ResetToL1BlockNumber")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, entities.Tx) error); ok {
		r0 = rf(ctx, firstBlockNumberToKeep, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// reorgStorer_ResetToL1BlockNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResetToL1BlockNumber'
type reorgStorer_ResetToL1BlockNumber_Call struct {
	*mock.Call
}

// ResetToL1BlockNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - firstBlockNumberToKeep uint64
//   - dbTx entities.Tx
func (_e *reorgStorer_Expecter) ResetToL1BlockNumber(ctx interface{}, firstBlockNumberToKeep interface{}, dbTx interface{}) *reorgStorer_ResetToL1BlockNumber_Call {
	return &reorgStorer_ResetToL1BlockNumber_Call{Call: _e.mock.On("ResetToL1BlockNumber", ctx, firstBlockNumberToKeep, dbTx)}
}

func (_c *reorgStorer_ResetToL1BlockNumber_Call) Run(run func(ctx context.Context, firstBlockNumberToKeep uint64, dbTx entities.Tx)) *reorgStorer_ResetToL1BlockNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(entities.Tx))
	})
	return _c
}

func (_c *reorgStorer_ResetToL1BlockNumber_Call) Return(_a0 error) *reorgStorer_ResetToL1BlockNumber_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *reorgStorer_ResetToL1BlockNumber_Call) RunAndReturn(run func(context.Context, uint64, entities.Tx) error) *reorgStorer_ResetToL1BlockNumber_Call {
	_c.Call.Return(run)
	return _c
}

// newReorgStorer creates a new instance of reorgStorer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newReorgStorer(t interface {
	mock.TestingT
	Cleanup(func())
}) *reorgStorer {
	mock := &reorgStorer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}