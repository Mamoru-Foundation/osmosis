package mamoru_cosmos_sdk

import (
	"context"
	"github.com/cometbft/cometbft/abci/types"
	log2 "github.com/cometbft/cometbft/libs/log"
	tmtypes "github.com/cometbft/cometbft/proto/tendermint/types"
	tmversion "github.com/cometbft/cometbft/proto/tendermint/version"
	"gotest.tools/v3/assert"
	"testing"
	"time"
)

func TestListenBeginBlock(t *testing.T) {
	t.Run("TestListenBeginBlock", func(t *testing.T) {
		//buf := &bytes.Buffer{}
		logger := log2.TestingLogger()
		ss := NewStreamingService(logger, nil)

		ctx := context.Background()
		req := types.RequestBeginBlock{Header: tmtypes.Header{
			Version:            tmversion.Consensus{},
			ChainID:            "",
			Height:             1234,
			Time:               time.Time{},
			LastBlockId:        tmtypes.BlockID{},
			LastCommitHash:     nil,
			DataHash:           nil,
			ValidatorsHash:     nil,
			NextValidatorsHash: nil,
			ConsensusHash:      nil,
			AppHash:            nil,
			LastResultsHash:    nil,
			EvidenceHash:       nil,
			ProposerAddress:    nil,
		}}
		res := types.ResponseBeginBlock{}

		err := ss.ListenBeginBlock(ctx, req, res)
		assert.NilError(t, err)
		assert.Equal(t, ss.blockMetadata.RequestBeginBlock.Header.Height, req.Header.Height)
		assert.Equal(t, ss.currentBlockNumber, req.Header.Height)
	})
}
