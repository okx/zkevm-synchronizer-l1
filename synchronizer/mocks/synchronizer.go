// Code generated by mockery. DO NOT EDIT.

package mock_synchronizer

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	synchronizer "github.com/0xPolygonHermez/zkevm-synchronizer-l1/synchronizer"
)

// Synchronizer is an autogenerated mock type for the Synchronizer type
type Synchronizer struct {
	mock.Mock
}

type Synchronizer_Expecter struct {
	mock *mock.Mock
}

func (_m *Synchronizer) EXPECT() *Synchronizer_Expecter {
	return &Synchronizer_Expecter{mock: &_m.Mock}
}

// GetL1BlockByNumber provides a mock function with given fields: ctx, blockNumber
func (_m *Synchronizer) GetL1BlockByNumber(ctx context.Context, blockNumber uint64) (*synchronizer.L1Block, error) {
	ret := _m.Called(ctx, blockNumber)

	if len(ret) == 0 {
		panic("no return value specified for GetL1BlockByNumber")
	}

	var r0 *synchronizer.L1Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*synchronizer.L1Block, error)); ok {
		return rf(ctx, blockNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *synchronizer.L1Block); ok {
		r0 = rf(ctx, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synchronizer.L1Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Synchronizer_GetL1BlockByNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetL1BlockByNumber'
type Synchronizer_GetL1BlockByNumber_Call struct {
	*mock.Call
}

// GetL1BlockByNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - blockNumber uint64
func (_e *Synchronizer_Expecter) GetL1BlockByNumber(ctx interface{}, blockNumber interface{}) *Synchronizer_GetL1BlockByNumber_Call {
	return &Synchronizer_GetL1BlockByNumber_Call{Call: _e.mock.On("GetL1BlockByNumber", ctx, blockNumber)}
}

func (_c *Synchronizer_GetL1BlockByNumber_Call) Run(run func(ctx context.Context, blockNumber uint64)) *Synchronizer_GetL1BlockByNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *Synchronizer_GetL1BlockByNumber_Call) Return(_a0 *synchronizer.L1Block, _a1 error) *Synchronizer_GetL1BlockByNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Synchronizer_GetL1BlockByNumber_Call) RunAndReturn(run func(context.Context, uint64) (*synchronizer.L1Block, error)) *Synchronizer_GetL1BlockByNumber_Call {
	_c.Call.Return(run)
	return _c
}

// GetL1InfoRootPerIndex provides a mock function with given fields: ctx, L1InfoTreeIndex
func (_m *Synchronizer) GetL1InfoRootPerIndex(ctx context.Context, L1InfoTreeIndex uint32) (common.Hash, error) {
	ret := _m.Called(ctx, L1InfoTreeIndex)

	if len(ret) == 0 {
		panic("no return value specified for GetL1InfoRootPerIndex")
	}

	var r0 common.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) (common.Hash, error)); ok {
		return rf(ctx, L1InfoTreeIndex)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32) common.Hash); ok {
		r0 = rf(ctx, L1InfoTreeIndex)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, L1InfoTreeIndex)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Synchronizer_GetL1InfoRootPerIndex_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetL1InfoRootPerIndex'
type Synchronizer_GetL1InfoRootPerIndex_Call struct {
	*mock.Call
}

// GetL1InfoRootPerIndex is a helper method to define mock.On call
//   - ctx context.Context
//   - L1InfoTreeIndex uint32
func (_e *Synchronizer_Expecter) GetL1InfoRootPerIndex(ctx interface{}, L1InfoTreeIndex interface{}) *Synchronizer_GetL1InfoRootPerIndex_Call {
	return &Synchronizer_GetL1InfoRootPerIndex_Call{Call: _e.mock.On("GetL1InfoRootPerIndex", ctx, L1InfoTreeIndex)}
}

func (_c *Synchronizer_GetL1InfoRootPerIndex_Call) Run(run func(ctx context.Context, L1InfoTreeIndex uint32)) *Synchronizer_GetL1InfoRootPerIndex_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *Synchronizer_GetL1InfoRootPerIndex_Call) Return(_a0 common.Hash, _a1 error) *Synchronizer_GetL1InfoRootPerIndex_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Synchronizer_GetL1InfoRootPerIndex_Call) RunAndReturn(run func(context.Context, uint32) (common.Hash, error)) *Synchronizer_GetL1InfoRootPerIndex_Call {
	_c.Call.Return(run)
	return _c
}

// GetL1InfoTreeLeaves provides a mock function with given fields: ctx, indexLeaves
func (_m *Synchronizer) GetL1InfoTreeLeaves(ctx context.Context, indexLeaves []uint32) (map[uint32]synchronizer.L1InfoTreeLeaf, error) {
	ret := _m.Called(ctx, indexLeaves)

	if len(ret) == 0 {
		panic("no return value specified for GetL1InfoTreeLeaves")
	}

	var r0 map[uint32]synchronizer.L1InfoTreeLeaf
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []uint32) (map[uint32]synchronizer.L1InfoTreeLeaf, error)); ok {
		return rf(ctx, indexLeaves)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []uint32) map[uint32]synchronizer.L1InfoTreeLeaf); ok {
		r0 = rf(ctx, indexLeaves)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[uint32]synchronizer.L1InfoTreeLeaf)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []uint32) error); ok {
		r1 = rf(ctx, indexLeaves)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Synchronizer_GetL1InfoTreeLeaves_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetL1InfoTreeLeaves'
type Synchronizer_GetL1InfoTreeLeaves_Call struct {
	*mock.Call
}

// GetL1InfoTreeLeaves is a helper method to define mock.On call
//   - ctx context.Context
//   - indexLeaves []uint32
func (_e *Synchronizer_Expecter) GetL1InfoTreeLeaves(ctx interface{}, indexLeaves interface{}) *Synchronizer_GetL1InfoTreeLeaves_Call {
	return &Synchronizer_GetL1InfoTreeLeaves_Call{Call: _e.mock.On("GetL1InfoTreeLeaves", ctx, indexLeaves)}
}

func (_c *Synchronizer_GetL1InfoTreeLeaves_Call) Run(run func(ctx context.Context, indexLeaves []uint32)) *Synchronizer_GetL1InfoTreeLeaves_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]uint32))
	})
	return _c
}

func (_c *Synchronizer_GetL1InfoTreeLeaves_Call) Return(_a0 map[uint32]synchronizer.L1InfoTreeLeaf, _a1 error) *Synchronizer_GetL1InfoTreeLeaves_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Synchronizer_GetL1InfoTreeLeaves_Call) RunAndReturn(run func(context.Context, []uint32) (map[uint32]synchronizer.L1InfoTreeLeaf, error)) *Synchronizer_GetL1InfoTreeLeaves_Call {
	_c.Call.Return(run)
	return _c
}

// GetLastL1Block provides a mock function with given fields: ctx
func (_m *Synchronizer) GetLastL1Block(ctx context.Context) (*synchronizer.L1Block, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetLastL1Block")
	}

	var r0 *synchronizer.L1Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*synchronizer.L1Block, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *synchronizer.L1Block); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synchronizer.L1Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Synchronizer_GetLastL1Block_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLastL1Block'
type Synchronizer_GetLastL1Block_Call struct {
	*mock.Call
}

// GetLastL1Block is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Synchronizer_Expecter) GetLastL1Block(ctx interface{}) *Synchronizer_GetLastL1Block_Call {
	return &Synchronizer_GetLastL1Block_Call{Call: _e.mock.On("GetLastL1Block", ctx)}
}

func (_c *Synchronizer_GetLastL1Block_Call) Run(run func(ctx context.Context)) *Synchronizer_GetLastL1Block_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Synchronizer_GetLastL1Block_Call) Return(_a0 *synchronizer.L1Block, _a1 error) *Synchronizer_GetLastL1Block_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Synchronizer_GetLastL1Block_Call) RunAndReturn(run func(context.Context) (*synchronizer.L1Block, error)) *Synchronizer_GetLastL1Block_Call {
	_c.Call.Return(run)
	return _c
}

// GetLastestVirtualBatchNumber provides a mock function with given fields: ctx
func (_m *Synchronizer) GetLastestVirtualBatchNumber(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetLastestVirtualBatchNumber")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Synchronizer_GetLastestVirtualBatchNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLastestVirtualBatchNumber'
type Synchronizer_GetLastestVirtualBatchNumber_Call struct {
	*mock.Call
}

// GetLastestVirtualBatchNumber is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Synchronizer_Expecter) GetLastestVirtualBatchNumber(ctx interface{}) *Synchronizer_GetLastestVirtualBatchNumber_Call {
	return &Synchronizer_GetLastestVirtualBatchNumber_Call{Call: _e.mock.On("GetLastestVirtualBatchNumber", ctx)}
}

func (_c *Synchronizer_GetLastestVirtualBatchNumber_Call) Run(run func(ctx context.Context)) *Synchronizer_GetLastestVirtualBatchNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Synchronizer_GetLastestVirtualBatchNumber_Call) Return(_a0 uint64, _a1 error) *Synchronizer_GetLastestVirtualBatchNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Synchronizer_GetLastestVirtualBatchNumber_Call) RunAndReturn(run func(context.Context) (uint64, error)) *Synchronizer_GetLastestVirtualBatchNumber_Call {
	_c.Call.Return(run)
	return _c
}

// GetLeafsByL1InfoRoot provides a mock function with given fields: ctx, l1InfoRoot
func (_m *Synchronizer) GetLeafsByL1InfoRoot(ctx context.Context, l1InfoRoot common.Hash) ([]synchronizer.L1InfoTreeLeaf, error) {
	ret := _m.Called(ctx, l1InfoRoot)

	if len(ret) == 0 {
		panic("no return value specified for GetLeafsByL1InfoRoot")
	}

	var r0 []synchronizer.L1InfoTreeLeaf
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) ([]synchronizer.L1InfoTreeLeaf, error)); ok {
		return rf(ctx, l1InfoRoot)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) []synchronizer.L1InfoTreeLeaf); ok {
		r0 = rf(ctx, l1InfoRoot)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]synchronizer.L1InfoTreeLeaf)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, l1InfoRoot)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Synchronizer_GetLeafsByL1InfoRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLeafsByL1InfoRoot'
type Synchronizer_GetLeafsByL1InfoRoot_Call struct {
	*mock.Call
}

// GetLeafsByL1InfoRoot is a helper method to define mock.On call
//   - ctx context.Context
//   - l1InfoRoot common.Hash
func (_e *Synchronizer_Expecter) GetLeafsByL1InfoRoot(ctx interface{}, l1InfoRoot interface{}) *Synchronizer_GetLeafsByL1InfoRoot_Call {
	return &Synchronizer_GetLeafsByL1InfoRoot_Call{Call: _e.mock.On("GetLeafsByL1InfoRoot", ctx, l1InfoRoot)}
}

func (_c *Synchronizer_GetLeafsByL1InfoRoot_Call) Run(run func(ctx context.Context, l1InfoRoot common.Hash)) *Synchronizer_GetLeafsByL1InfoRoot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Hash))
	})
	return _c
}

func (_c *Synchronizer_GetLeafsByL1InfoRoot_Call) Return(_a0 []synchronizer.L1InfoTreeLeaf, _a1 error) *Synchronizer_GetLeafsByL1InfoRoot_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Synchronizer_GetLeafsByL1InfoRoot_Call) RunAndReturn(run func(context.Context, common.Hash) ([]synchronizer.L1InfoTreeLeaf, error)) *Synchronizer_GetLeafsByL1InfoRoot_Call {
	_c.Call.Return(run)
	return _c
}

// GetSequenceByBatchNumber provides a mock function with given fields: ctx, batchNumber
func (_m *Synchronizer) GetSequenceByBatchNumber(ctx context.Context, batchNumber uint64) (*synchronizer.SequencedBatches, error) {
	ret := _m.Called(ctx, batchNumber)

	if len(ret) == 0 {
		panic("no return value specified for GetSequenceByBatchNumber")
	}

	var r0 *synchronizer.SequencedBatches
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*synchronizer.SequencedBatches, error)); ok {
		return rf(ctx, batchNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *synchronizer.SequencedBatches); ok {
		r0 = rf(ctx, batchNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synchronizer.SequencedBatches)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, batchNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Synchronizer_GetSequenceByBatchNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSequenceByBatchNumber'
type Synchronizer_GetSequenceByBatchNumber_Call struct {
	*mock.Call
}

// GetSequenceByBatchNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - batchNumber uint64
func (_e *Synchronizer_Expecter) GetSequenceByBatchNumber(ctx interface{}, batchNumber interface{}) *Synchronizer_GetSequenceByBatchNumber_Call {
	return &Synchronizer_GetSequenceByBatchNumber_Call{Call: _e.mock.On("GetSequenceByBatchNumber", ctx, batchNumber)}
}

func (_c *Synchronizer_GetSequenceByBatchNumber_Call) Run(run func(ctx context.Context, batchNumber uint64)) *Synchronizer_GetSequenceByBatchNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *Synchronizer_GetSequenceByBatchNumber_Call) Return(_a0 *synchronizer.SequencedBatches, _a1 error) *Synchronizer_GetSequenceByBatchNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Synchronizer_GetSequenceByBatchNumber_Call) RunAndReturn(run func(context.Context, uint64) (*synchronizer.SequencedBatches, error)) *Synchronizer_GetSequenceByBatchNumber_Call {
	_c.Call.Return(run)
	return _c
}

// GetVirtualBatchByBatchNumber provides a mock function with given fields: ctx, batchNumber
func (_m *Synchronizer) GetVirtualBatchByBatchNumber(ctx context.Context, batchNumber uint64) (*synchronizer.VirtualBatch, error) {
	ret := _m.Called(ctx, batchNumber)

	if len(ret) == 0 {
		panic("no return value specified for GetVirtualBatchByBatchNumber")
	}

	var r0 *synchronizer.VirtualBatch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*synchronizer.VirtualBatch, error)); ok {
		return rf(ctx, batchNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *synchronizer.VirtualBatch); ok {
		r0 = rf(ctx, batchNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*synchronizer.VirtualBatch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, batchNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Synchronizer_GetVirtualBatchByBatchNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetVirtualBatchByBatchNumber'
type Synchronizer_GetVirtualBatchByBatchNumber_Call struct {
	*mock.Call
}

// GetVirtualBatchByBatchNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - batchNumber uint64
func (_e *Synchronizer_Expecter) GetVirtualBatchByBatchNumber(ctx interface{}, batchNumber interface{}) *Synchronizer_GetVirtualBatchByBatchNumber_Call {
	return &Synchronizer_GetVirtualBatchByBatchNumber_Call{Call: _e.mock.On("GetVirtualBatchByBatchNumber", ctx, batchNumber)}
}

func (_c *Synchronizer_GetVirtualBatchByBatchNumber_Call) Run(run func(ctx context.Context, batchNumber uint64)) *Synchronizer_GetVirtualBatchByBatchNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64))
	})
	return _c
}

func (_c *Synchronizer_GetVirtualBatchByBatchNumber_Call) Return(_a0 *synchronizer.VirtualBatch, _a1 error) *Synchronizer_GetVirtualBatchByBatchNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Synchronizer_GetVirtualBatchByBatchNumber_Call) RunAndReturn(run func(context.Context, uint64) (*synchronizer.VirtualBatch, error)) *Synchronizer_GetVirtualBatchByBatchNumber_Call {
	_c.Call.Return(run)
	return _c
}

// IsSynced provides a mock function with given fields:
func (_m *Synchronizer) IsSynced() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsSynced")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Synchronizer_IsSynced_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsSynced'
type Synchronizer_IsSynced_Call struct {
	*mock.Call
}

// IsSynced is a helper method to define mock.On call
func (_e *Synchronizer_Expecter) IsSynced() *Synchronizer_IsSynced_Call {
	return &Synchronizer_IsSynced_Call{Call: _e.mock.On("IsSynced")}
}

func (_c *Synchronizer_IsSynced_Call) Run(run func()) *Synchronizer_IsSynced_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Synchronizer_IsSynced_Call) Return(_a0 bool) *Synchronizer_IsSynced_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Synchronizer_IsSynced_Call) RunAndReturn(run func() bool) *Synchronizer_IsSynced_Call {
	_c.Call.Return(run)
	return _c
}

// SetCallbackOnReorgDone provides a mock function with given fields: callback
func (_m *Synchronizer) SetCallbackOnReorgDone(callback func(synchronizer.ReorgExecutionResult)) {
	_m.Called(callback)
}

// Synchronizer_SetCallbackOnReorgDone_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetCallbackOnReorgDone'
type Synchronizer_SetCallbackOnReorgDone_Call struct {
	*mock.Call
}

// SetCallbackOnReorgDone is a helper method to define mock.On call
//   - callback func(synchronizer.ReorgExecutionResult)
func (_e *Synchronizer_Expecter) SetCallbackOnReorgDone(callback interface{}) *Synchronizer_SetCallbackOnReorgDone_Call {
	return &Synchronizer_SetCallbackOnReorgDone_Call{Call: _e.mock.On("SetCallbackOnReorgDone", callback)}
}

func (_c *Synchronizer_SetCallbackOnReorgDone_Call) Run(run func(callback func(synchronizer.ReorgExecutionResult))) *Synchronizer_SetCallbackOnReorgDone_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(synchronizer.ReorgExecutionResult)))
	})
	return _c
}

func (_c *Synchronizer_SetCallbackOnReorgDone_Call) Return() *Synchronizer_SetCallbackOnReorgDone_Call {
	_c.Call.Return()
	return _c
}

func (_c *Synchronizer_SetCallbackOnReorgDone_Call) RunAndReturn(run func(func(synchronizer.ReorgExecutionResult))) *Synchronizer_SetCallbackOnReorgDone_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *Synchronizer) Stop() {
	_m.Called()
}

// Synchronizer_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type Synchronizer_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *Synchronizer_Expecter) Stop() *Synchronizer_Stop_Call {
	return &Synchronizer_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *Synchronizer_Stop_Call) Run(run func()) *Synchronizer_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Synchronizer_Stop_Call) Return() *Synchronizer_Stop_Call {
	_c.Call.Return()
	return _c
}

func (_c *Synchronizer_Stop_Call) RunAndReturn(run func()) *Synchronizer_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// Sync provides a mock function with given fields: returnOnSync
func (_m *Synchronizer) Sync(returnOnSync bool) error {
	ret := _m.Called(returnOnSync)

	if len(ret) == 0 {
		panic("no return value specified for Sync")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(returnOnSync)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Synchronizer_Sync_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Sync'
type Synchronizer_Sync_Call struct {
	*mock.Call
}

// Sync is a helper method to define mock.On call
//   - returnOnSync bool
func (_e *Synchronizer_Expecter) Sync(returnOnSync interface{}) *Synchronizer_Sync_Call {
	return &Synchronizer_Sync_Call{Call: _e.mock.On("Sync", returnOnSync)}
}

func (_c *Synchronizer_Sync_Call) Run(run func(returnOnSync bool)) *Synchronizer_Sync_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *Synchronizer_Sync_Call) Return(_a0 error) *Synchronizer_Sync_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Synchronizer_Sync_Call) RunAndReturn(run func(bool) error) *Synchronizer_Sync_Call {
	_c.Call.Return(run)
	return _c
}

// NewSynchronizer creates a new instance of Synchronizer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSynchronizer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Synchronizer {
	mock := &Synchronizer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
