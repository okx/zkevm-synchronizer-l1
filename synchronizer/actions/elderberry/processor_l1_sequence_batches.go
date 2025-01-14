package elderberry

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/0xPolygonHermez/zkevm-synchronizer-l1/etherman"
	"github.com/0xPolygonHermez/zkevm-synchronizer-l1/log"
	"github.com/0xPolygonHermez/zkevm-synchronizer-l1/state/entities"
	"github.com/0xPolygonHermez/zkevm-synchronizer-l1/synchronizer/actions"
)

var (
	// ErrInvalidInitialBatchNumber is returned when the initial batch number is not the expected one
	ErrInvalidInitialBatchNumber = errors.New("invalid initial batch number")
)

// PreviousProcessor is the interface that the previous processor (Etrog)
type PreviousProcessor interface {
	Process(ctx context.Context, forkId actions.ForkIdType, order etherman.Order, l1Block *etherman.Block, dbTx entities.Tx) error
	ProcessSequenceBatches(ctx context.Context, forkId actions.ForkIdType, sequencedBatches []etherman.SequencedBatch, blockNumber uint64, l1BlockTimestamp time.Time, dbTx entities.Tx) error
}

// ProcessorL1SequenceBatchesElderberry is the processor for SequenceBatches for Elderberry
type ProcessorL1SequenceBatchesElderberry struct {
	actions.ProcessorBase[ProcessorL1SequenceBatchesElderberry]
	previousProcessor PreviousProcessor
}

// NewProcessorL1SequenceBatchesElderberry returns instance of a processor for SequenceBatchesOrder
func NewProcessorL1SequenceBatchesElderberry(previousProcessor PreviousProcessor) *ProcessorL1SequenceBatchesElderberry {
	return &ProcessorL1SequenceBatchesElderberry{
		ProcessorBase: actions.ProcessorBase[ProcessorL1SequenceBatchesElderberry]{
			SupportedEvent:    []etherman.EventOrder{etherman.SequenceBatchesOrder},
			SupportedForkdIds: &actions.ForksIdOnlyElderberry},
		previousProcessor: previousProcessor,
	}
}

// Process process event
func (g *ProcessorL1SequenceBatchesElderberry) Process(ctx context.Context, forkId actions.ForkIdType, order etherman.Order, l1Block *etherman.Block, dbTx entities.Tx) error {
	if l1Block == nil || len(l1Block.SequencedBatches) <= order.Pos {
		return actions.ErrInvalidParams
	}
	if len(l1Block.SequencedBatches[order.Pos]) == 0 {
		log.Warnf("No sequenced batches for position")
		return nil
	}

	sbatch := l1Block.SequencedBatches[order.Pos][0]

	if sbatch.SequencedBatchElderberryData == nil {
		log.Errorf("No elderberry sequenced batch data for batch %d", sbatch.BatchNumber)
		return fmt.Errorf("no elderberry sequenced batch data for batch %d", sbatch.BatchNumber)
	}

	// We known that the MaxSequenceTimestamp is the same for all the batches so we can use the first one
	err := g.previousProcessor.ProcessSequenceBatches(ctx, forkId, l1Block.SequencedBatches[order.Pos], l1Block.BlockNumber, time.Unix(int64(sbatch.SequencedBatchElderberryData.MaxSequenceTimestamp), 0), dbTx)
	// The last L2block timestamp must match MaxSequenceTimestamp
	if err != nil {
		return err
	}
	return nil
}
