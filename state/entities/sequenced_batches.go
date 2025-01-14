package entities

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type SequencedBatches struct {
	FromBatchNumber uint64
	ToBatchNumber   uint64
	L1BlockNumber   uint64
	ForkID          uint64
	Timestamp       time.Time
	ReceivedAt      time.Time
	L1InfoRoot      common.Hash
	Source          string
}

func (s *SequencedBatches) IsEqual(o interface{}) bool {
	other, ok := o.(*SequencedBatches)
	if !ok {
		return false
	}
	if s == other {
		return true
	}
	return *s == *other
}

func (s *SequencedBatches) Key() uint64 {
	return s.FromBatchNumber
}

func NewSequencedBatches(fromBatchNumber, toBatchNumber, l1BlockNumber, forkID uint64,
	timestamp time.Time, receivedAt time.Time, l1InfoRoot common.Hash, source string) *SequencedBatches {
	return &SequencedBatches{
		FromBatchNumber: fromBatchNumber,
		ToBatchNumber:   toBatchNumber,
		L1BlockNumber:   l1BlockNumber,
		ForkID:          forkID,
		Timestamp:       timestamp,
		ReceivedAt:      receivedAt,
		L1InfoRoot:      l1InfoRoot,
		Source:          source,
	}
}
