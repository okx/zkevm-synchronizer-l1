// Code generated by mockery. DO NOT EDIT.

package mock_dataavailability

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// SequenceSender is an autogenerated mock type for the SequenceSender type
type SequenceSender struct {
	mock.Mock
}

type SequenceSender_Expecter struct {
	mock *mock.Mock
}

func (_m *SequenceSender) EXPECT() *SequenceSender_Expecter {
	return &SequenceSender_Expecter{mock: &_m.Mock}
}

// PostSequence provides a mock function with given fields: ctx, batchesData
func (_m *SequenceSender) PostSequence(ctx context.Context, batchesData [][]byte) ([]byte, error) {
	ret := _m.Called(ctx, batchesData)

	if len(ret) == 0 {
		panic("no return value specified for PostSequence")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, [][]byte) ([]byte, error)); ok {
		return rf(ctx, batchesData)
	}
	if rf, ok := ret.Get(0).(func(context.Context, [][]byte) []byte); ok {
		r0 = rf(ctx, batchesData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, [][]byte) error); ok {
		r1 = rf(ctx, batchesData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SequenceSender_PostSequence_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PostSequence'
type SequenceSender_PostSequence_Call struct {
	*mock.Call
}

// PostSequence is a helper method to define mock.On call
//   - ctx context.Context
//   - batchesData [][]byte
func (_e *SequenceSender_Expecter) PostSequence(ctx interface{}, batchesData interface{}) *SequenceSender_PostSequence_Call {
	return &SequenceSender_PostSequence_Call{Call: _e.mock.On("PostSequence", ctx, batchesData)}
}

func (_c *SequenceSender_PostSequence_Call) Run(run func(ctx context.Context, batchesData [][]byte)) *SequenceSender_PostSequence_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([][]byte))
	})
	return _c
}

func (_c *SequenceSender_PostSequence_Call) Return(_a0 []byte, _a1 error) *SequenceSender_PostSequence_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SequenceSender_PostSequence_Call) RunAndReturn(run func(context.Context, [][]byte) ([]byte, error)) *SequenceSender_PostSequence_Call {
	_c.Call.Return(run)
	return _c
}

// NewSequenceSender creates a new instance of SequenceSender. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSequenceSender(t interface {
	mock.TestingT
	Cleanup(func())
}) *SequenceSender {
	mock := &SequenceSender{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
