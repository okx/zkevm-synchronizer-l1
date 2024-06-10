package l1_check_block_test

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/0xPolygonHermez/zkevm-synchronizer-l1/synchronizer/l1_check_block"
	mock_l1_check_block "github.com/0xPolygonHermez/zkevm-synchronizer-l1/synchronizer/l1_check_block/mocks"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBlockNumber(t *testing.T) {
	ctx := context.Background()
	mockRequester := mock_l1_check_block.NewL1Requester(t)
	//safeBlockPoint := big.NewInt(50)
	offset := 10
	syncPoint := l1_check_block.L1BlockPointWithOffset{
		BlockPoint: l1_check_block.SafeBlockNumber,
		Offset:     offset,
	}
	safeL1Block := l1_check_block.NewL1BlockNumberByNameFetch(syncPoint)

	mockRequester.EXPECT().HeaderByNumber(ctx, mock.Anything).Return(&types.Header{
		Number: big.NewInt(100),
	}, nil)
	blockNumber, err := safeL1Block.BlockNumber(ctx, mockRequester)
	assert.NoError(t, err)
	expectedBlockNumber := uint64(100 + offset)
	assert.Equal(t, expectedBlockNumber, blockNumber)
}

func TestBlockNumberMutliplesCases(t *testing.T) {
	tests := []struct {
		name                string
		blockPoint          string
		offset              int
		l1ReturnBlockNumber uint64
		expectedCallToGeth  *big.Int
		expectedBlockNumber uint64
	}{
		{
			name:                "SafeBlockNumber+10",
			blockPoint:          "safe",
			offset:              10,
			l1ReturnBlockNumber: 100,
			expectedCallToGeth:  big.NewInt(int64(rpc.SafeBlockNumber)),
			expectedBlockNumber: 110,
		},
		{
			name:                "FinalizedBlockNumber+10",
			blockPoint:          "finalized",
			offset:              10,
			l1ReturnBlockNumber: 100,
			expectedCallToGeth:  big.NewInt(int64(rpc.FinalizedBlockNumber)),
			expectedBlockNumber: 110,
		},
		{
			name:                "PendingBlockNumber+10",
			blockPoint:          "pending",
			offset:              10,
			l1ReturnBlockNumber: 100,
			expectedCallToGeth:  big.NewInt(int64(rpc.PendingBlockNumber)),
			expectedBlockNumber: 110,
		},
		{
			name:                "LastBlockNumber+10, can't add 10 to latest block number. So must return latest block number and ignore positive offset",
			blockPoint:          "latest",
			offset:              10,
			l1ReturnBlockNumber: 100,
			expectedCallToGeth:  nil,
			expectedBlockNumber: 100,
		},
		{
			name:                "FinalizedBlockNumber-1000. negative blockNumbers are not welcome. So must return 0",
			blockPoint:          "finalized",
			offset:              -1000,
			l1ReturnBlockNumber: 100,
			expectedCallToGeth:  big.NewInt(int64(rpc.FinalizedBlockNumber)),
			expectedBlockNumber: 0,
		},
		{
			name:                "FinalizedBlockNumber(1000)-1000. is 0 ",
			blockPoint:          "finalized",
			offset:              -1000,
			l1ReturnBlockNumber: 1000,
			expectedCallToGeth:  big.NewInt(int64(rpc.FinalizedBlockNumber)),
			expectedBlockNumber: 0,
		},
		{
			name:                "FinalizedBlockNumber(1001)-1000. is 1 ",
			blockPoint:          "finalized",
			offset:              -1000,
			l1ReturnBlockNumber: 1001,
			expectedCallToGeth:  big.NewInt(int64(rpc.FinalizedBlockNumber)),
			expectedBlockNumber: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockRequester := mock_l1_check_block.NewL1Requester(t)
			syncPointStr := fmt.Sprintf("%s/%d", tt.blockPoint, tt.offset)
			syncPoint, err := l1_check_block.StringToL1BlockPointWithOffset(syncPointStr)
			assert.NoError(t, err)
			safeL1Block := l1_check_block.NewL1BlockNumberByNameFetch(syncPoint)

			mockRequester.EXPECT().HeaderByNumber(ctx, tt.expectedCallToGeth).Return(&types.Header{
				Number: big.NewInt(int64(tt.l1ReturnBlockNumber)),
			}, nil)
			blockNumber, err := safeL1Block.BlockNumber(ctx, mockRequester)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedBlockNumber, blockNumber)
		})
	}
}

func TestBlockNumberWithOffsetMutliplesCases(t *testing.T) {
	tests := []struct {
		name       string
		blockPoint string
		expected   l1_check_block.L1BlockPointWithOffset
		err        bool
	}{
		{
			name:       "SafeBlockNumber+10",
			blockPoint: "safe/10",
			expected:   l1_check_block.L1BlockPointWithOffset{BlockPoint: l1_check_block.SafeBlockNumber, Offset: 10},
		},
		{
			name:       "SafeBlockNumber+10",
			blockPoint: "safe/-10",
			expected:   l1_check_block.L1BlockPointWithOffset{BlockPoint: l1_check_block.SafeBlockNumber, Offset: -10},
		},
		{
			name:       "SafeBlockNumber+10",
			blockPoint: "safe",
			expected:   l1_check_block.L1BlockPointWithOffset{BlockPoint: l1_check_block.SafeBlockNumber, Offset: 0},
		},
		{
			name:       "bad1",
			blockPoint: "safe/10/20",
			expected:   l1_check_block.L1BlockPointWithOffset{BlockPoint: l1_check_block.SafeBlockNumber, Offset: 0},
			err:        true,
		},
		{
			name:       "bad2",
			blockPoint: "unknown_value/10",
			expected:   l1_check_block.L1BlockPointWithOffset{BlockPoint: l1_check_block.SafeBlockNumber, Offset: 0},
			err:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := l1_check_block.StringToL1BlockPointWithOffset(tt.blockPoint)
			if tt.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestBlockNumberNotFound(t *testing.T) {
	ctx := context.Background()
	mockRequester := mock_l1_check_block.NewL1Requester(t)
	//safeBlockPoint := big.NewInt(50)
	offset := 10
	syncPoint := l1_check_block.L1BlockPointWithOffset{
		BlockPoint: l1_check_block.SafeBlockNumber,
		Offset:     offset,
	}
	mockRequester.EXPECT().HeaderByNumber(ctx, mock.Anything).Return(nil, fmt.Errorf("block not found"))
	safeL1Block := l1_check_block.NewL1BlockNumberByNameFetch(syncPoint)
	_, err := safeL1Block.BlockNumber(ctx, mockRequester)
	assert.Error(t, err)

	mockRequester.EXPECT().HeaderByNumber(ctx, mock.Anything).Return(nil, fmt.Errorf("block not found"))
	safeL1Block = l1_check_block.NewL1BlockNumberByNameFetch(syncPoint).SetIfNotFoundReturnsZero()
	block, err := safeL1Block.BlockNumber(ctx, mockRequester)
	assert.NoError(t, err)
	assert.Equal(t, uint64(0), block)

}

func TestGreaterThan(t *testing.T) {
	finalized := &l1_check_block.L1BlockPointWithOffset{
		BlockPoint: l1_check_block.FinalizedBlockNumber,
		Offset:     0,
	}
	safe := &l1_check_block.L1BlockPointWithOffset{
		BlockPoint: l1_check_block.SafeBlockNumber,
		Offset:     0,
	}
	lastest := &l1_check_block.L1BlockPointWithOffset{
		BlockPoint: l1_check_block.LastBlockNumber,
		Offset:     0,
	}
	assert.False(t, finalized.GreaterThan(safe))
	assert.False(t, safe.GreaterThan(lastest))
	assert.False(t, finalized.GreaterThan(lastest))
	assert.True(t, lastest.GreaterThan(finalized))
	assert.True(t, lastest.GreaterThan(safe))
	assert.True(t, safe.GreaterThan(finalized))

	assert.False(t, safe.GreaterThan(safe))
}

func TestWithOffsetGreaterThan(t *testing.T) {
	safe10 := &l1_check_block.L1BlockPointWithOffset{
		BlockPoint: l1_check_block.SafeBlockNumber,
		Offset:     10,
	}
	safem10 := &l1_check_block.L1BlockPointWithOffset{
		BlockPoint: l1_check_block.SafeBlockNumber,
		Offset:     5,
	}

	assert.True(t, safe10.GreaterThan(safem10))
	assert.False(t, safe10.GreaterThan(safe10))
}
