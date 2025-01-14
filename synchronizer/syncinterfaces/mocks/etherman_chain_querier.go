// Code generated by mockery. DO NOT EDIT.

package mock_syncinterfaces

import mock "github.com/stretchr/testify/mock"

// EthermanChainQuerier is an autogenerated mock type for the EthermanChainQuerier type
type EthermanChainQuerier struct {
	mock.Mock
}

type EthermanChainQuerier_Expecter struct {
	mock *mock.Mock
}

func (_m *EthermanChainQuerier) EXPECT() *EthermanChainQuerier_Expecter {
	return &EthermanChainQuerier_Expecter{mock: &_m.Mock}
}

// GetL1ChainID provides a mock function with given fields:
func (_m *EthermanChainQuerier) GetL1ChainID() uint64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetL1ChainID")
	}

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// EthermanChainQuerier_GetL1ChainID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetL1ChainID'
type EthermanChainQuerier_GetL1ChainID_Call struct {
	*mock.Call
}

// GetL1ChainID is a helper method to define mock.On call
func (_e *EthermanChainQuerier_Expecter) GetL1ChainID() *EthermanChainQuerier_GetL1ChainID_Call {
	return &EthermanChainQuerier_GetL1ChainID_Call{Call: _e.mock.On("GetL1ChainID")}
}

func (_c *EthermanChainQuerier_GetL1ChainID_Call) Run(run func()) *EthermanChainQuerier_GetL1ChainID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EthermanChainQuerier_GetL1ChainID_Call) Return(_a0 uint64) *EthermanChainQuerier_GetL1ChainID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EthermanChainQuerier_GetL1ChainID_Call) RunAndReturn(run func() uint64) *EthermanChainQuerier_GetL1ChainID_Call {
	_c.Call.Return(run)
	return _c
}

// GetRollupID provides a mock function with given fields:
func (_m *EthermanChainQuerier) GetRollupID() uint {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetRollupID")
	}

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// EthermanChainQuerier_GetRollupID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRollupID'
type EthermanChainQuerier_GetRollupID_Call struct {
	*mock.Call
}

// GetRollupID is a helper method to define mock.On call
func (_e *EthermanChainQuerier_Expecter) GetRollupID() *EthermanChainQuerier_GetRollupID_Call {
	return &EthermanChainQuerier_GetRollupID_Call{Call: _e.mock.On("GetRollupID")}
}

func (_c *EthermanChainQuerier_GetRollupID_Call) Run(run func()) *EthermanChainQuerier_GetRollupID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EthermanChainQuerier_GetRollupID_Call) Return(_a0 uint) *EthermanChainQuerier_GetRollupID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EthermanChainQuerier_GetRollupID_Call) RunAndReturn(run func() uint) *EthermanChainQuerier_GetRollupID_Call {
	_c.Call.Return(run)
	return _c
}

// NewEthermanChainQuerier creates a new instance of EthermanChainQuerier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEthermanChainQuerier(t interface {
	mock.TestingT
	Cleanup(func())
}) *EthermanChainQuerier {
	mock := &EthermanChainQuerier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
