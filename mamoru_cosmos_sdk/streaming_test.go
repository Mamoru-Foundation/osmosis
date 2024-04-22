package mamoru_cosmos_sdk

import (
	types2 "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cometbft/cometbft/proto/tendermint/version"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestListenBeginBlock(t *testing.T) {
	t.Run("TestListenBeginBlock", func(t *testing.T) {

		logger := log.TestingLogger()
		header := types.Header{}
		ischeck := true
		ctx := sdk.NewContext(nil, header, ischeck, logger)

		ss := NewStreamingService(logger, nil)

		req := types2.RequestBeginBlock{Header: types.Header{
			Version:            version.Consensus{},
			ChainID:            "",
			Height:             1234,
			Time:               time.Time{},
			LastBlockId:        types.BlockID{},
			LastCommitHash:     []byte{'a', 'b', 'c'},
			DataHash:           []byte{'a', 'b', 'c'},
			ValidatorsHash:     []byte{'a', 'b', 'c'},
			NextValidatorsHash: []byte{'a', 'b', 'c'},
			ConsensusHash:      []byte{'a', 'b', 'c'},
			AppHash:            []byte{'a', 'b', 'c'},
			LastResultsHash:    []byte{'a', 'b', 'c'}, EvidenceHash: []byte{'a', 'b', 'c'},
			ProposerAddress: []byte{'a', 'b', 'c'},
		}}
		res := types2.ResponseBeginBlock{}

		err := ss.ListenBeginBlock(ctx, req, res)
		assert.NilError(t, err)
		assert.Equal(t, ss.blockMetadata.RequestBeginBlock.Header.Height, req.Header.Height)
		assert.Equal(t, ss.currentBlockNumber, req.Header.Height)
		assert.DeepEqual(t, ss.blockMetadata.RequestBeginBlock.Header.LastCommitHash, req.Header.LastCommitHash)
		assert.DeepEqual(t, ss.blockMetadata.RequestBeginBlock.Header.DataHash, req.Header.DataHash)
		assert.DeepEqual(t, ss.blockMetadata.RequestBeginBlock.Header.ValidatorsHash, req.Header.ValidatorsHash)
		assert.DeepEqual(t, ss.blockMetadata.RequestBeginBlock.Header.NextValidatorsHash, req.Header.NextValidatorsHash)
		assert.DeepEqual(t, ss.blockMetadata.RequestBeginBlock.Header.ConsensusHash, req.Header.ConsensusHash)
		assert.DeepEqual(t, ss.blockMetadata.RequestBeginBlock.Header.AppHash, req.Header.AppHash)
		assert.DeepEqual(t, ss.blockMetadata.RequestBeginBlock.Header.LastResultsHash, req.Header.LastResultsHash)
		assert.DeepEqual(t, ss.blockMetadata.RequestBeginBlock.Header.EvidenceHash, req.Header.EvidenceHash)
		assert.DeepEqual(t, ss.blockMetadata.RequestBeginBlock.Header.ProposerAddress, req.Header.ProposerAddress)
	})
}
