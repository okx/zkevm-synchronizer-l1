// Code generated by mockery. DO NOT EDIT.

package mock_synchronizer

import (
	context "context"

	entities "github.com/0xPolygonHermez/zkevm-synchronizer-l1/state/entities"

	mock "github.com/stretchr/testify/mock"

	model "github.com/0xPolygonHermez/zkevm-synchronizer-l1/state/model"
)

// stateOnSequencedBatchesInterface is an autogenerated mock type for the stateOnSequencedBatchesInterface type
type stateOnSequencedBatchesInterface struct {
	mock.Mock
}

type stateOnSequencedBatchesInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *stateOnSequencedBatchesInterface) EXPECT() *stateOnSequencedBatchesInterface_Expecter {
	return &stateOnSequencedBatchesInterface_Expecter{mock: &_m.Mock}
}

// OnSequencedBatchesOnL1 provides a mock function with given fields: ctx, seq, dbTx
func (_m *stateOnSequencedBatchesInterface) OnSequencedBatchesOnL1(ctx context.Context, seq model.SequenceOfBatches, dbTx entities.Tx) error {
	ret := _m.Called(ctx, seq, dbTx)

	if len(ret) == 0 {
		panic("no return value specified for OnSequencedBatchesOnL1")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.SequenceOfBatches, entities.Tx) error); ok {
		r0 = rf(ctx, seq, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// stateOnSequencedBatchesInterface_OnSequencedBatchesOnL1_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnSequencedBatchesOnL1'
type stateOnSequencedBatchesInterface_OnSequencedBatchesOnL1_Call struct {
	*mock.Call
}

// OnSequencedBatchesOnL1 is a helper method to define mock.On call
//   - ctx context.Context
//   - seq model.SequenceOfBatches
//   - dbTx entities.Tx
func (_e *stateOnSequencedBatchesInterface_Expecter) OnSequencedBatchesOnL1(ctx interface{}, seq interface{}, dbTx interface{}) *stateOnSequencedBatchesInterface_OnSequencedBatchesOnL1_Call {
	return &stateOnSequencedBatchesInterface_OnSequencedBatchesOnL1_Call{Call: _e.mock.On("OnSequencedBatchesOnL1", ctx, seq, dbTx)}
}

func (_c *stateOnSequencedBatchesInterface_OnSequencedBatchesOnL1_Call) Run(run func(ctx context.Context, seq model.SequenceOfBatches, dbTx entities.Tx)) *stateOnSequencedBatchesInterface_OnSequencedBatchesOnL1_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.SequenceOfBatches), args[2].(entities.Tx))
	})
	return _c
}

func (_c *stateOnSequencedBatchesInterface_OnSequencedBatchesOnL1_Call) Return(_a0 error) *stateOnSequencedBatchesInterface_OnSequencedBatchesOnL1_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *stateOnSequencedBatchesInterface_OnSequencedBatchesOnL1_Call) RunAndReturn(run func(context.Context, model.SequenceOfBatches, entities.Tx) error) *stateOnSequencedBatchesInterface_OnSequencedBatchesOnL1_Call {
	_c.Call.Return(run)
	return _c
}

// newStateOnSequencedBatchesInterface creates a new instance of stateOnSequencedBatchesInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newStateOnSequencedBatchesInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *stateOnSequencedBatchesInterface {
	mock := &stateOnSequencedBatchesInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}