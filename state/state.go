package state

import (
	"github.com/0xPolygonHermez/zkevm-synchronizer-l1/state/model"
	"github.com/0xPolygonHermez/zkevm-synchronizer-l1/state/storage"
)

type State struct {
	*model.TxManager
	*model.ForkIdState
	*model.L1InfoTreeState
	*model.BatchState
	storage.BlockStorer
}

func NewState(storageImpl storage.Storer) *State {
	res := &State{
		model.NewTxManager(storageImpl),
		model.NewForkIdState(storageImpl),
		model.NewL1InfoTreeManager(storageImpl),
		model.NewBatchState(storageImpl),
		storageImpl,
	}
	return res
}
