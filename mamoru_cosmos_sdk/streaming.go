package mamoru_cosmos_sdk

import (
	"context"
	"strconv"
	"strings"
	"sync"

	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer/cosmos"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/store/types"
	store "github.com/cosmos/cosmos-sdk/store/types"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
)

var _ baseapp.StreamingService = (*StreamingService)(nil)

type StreamingService struct {
	logger             log.Logger
	blockMetadata      types.BlockMetadata
	currentBlockNumber int64
	storeListeners     []*types.MemoryListener

	sniffer SnifferI
}

func NewStreamingService(logger log.Logger, sniffer SnifferI) *StreamingService {
	logger.Info("Mamoru MockStreamingService start")

	return &StreamingService{
		sniffer: sniffer,
		logger:  logger,
	}
}

func (ss *StreamingService) ListenBeginBlock(ctx context.Context, req abci.RequestBeginBlock, res abci.ResponseBeginBlock) error {
	ss.blockMetadata = types.BlockMetadata{}
	ss.blockMetadata.RequestBeginBlock = &req
	ss.blockMetadata.ResponseBeginBlock = &res
	ss.currentBlockNumber = req.Header.Height
	ss.logger.Info("Mamoru ListenBeginBlock", "height", ss.currentBlockNumber)

	return nil
}

func (ss *StreamingService) ListenDeliverTx(ctx context.Context, req abci.RequestDeliverTx, res abci.ResponseDeliverTx) error {
	ss.blockMetadata.DeliverTxs = append(ss.blockMetadata.DeliverTxs, &types.BlockMetadata_DeliverTx{
		Request:  &req,
		Response: &res,
	})
	ss.logger.Info("Mamoru ListenDeliverTx", "height", ss.currentBlockNumber)

	return nil
}

func (ss *StreamingService) ListenEndBlock(ctx context.Context, req abci.RequestEndBlock, res abci.ResponseEndBlock) error {
	ss.blockMetadata.RequestEndBlock = &req
	ss.blockMetadata.ResponseEndBlock = &res
	ss.logger.Info("Mamoru ListenEndBlock", "height", ss.currentBlockNumber)

	return nil
}

func (ss *StreamingService) ListenCommit(ctx context.Context, res abci.ResponseCommit) error {
	ss.blockMetadata.ResponseCommit = &res
	ss.logger.Info("Mamoru ListenCommit", "height", ss.currentBlockNumber, "res", res.String())

	if ss.sniffer == nil || !ss.sniffer.CheckRequirements() {
		return nil
	}

	builder := cosmos.NewCosmosCtxBuilder()

	blockHeight := uint64(ss.blockMetadata.RequestEndBlock.Height)
	block := cosmos.Block{
		Seq:                           blockHeight,
		Height:                        ss.blockMetadata.RequestEndBlock.Height,
		Hash:                          ss.blockMetadata.RequestBeginBlock.Hash,
		VersionBlock:                  ss.blockMetadata.RequestBeginBlock.Header.Version.Block,
		VersionApp:                    ss.blockMetadata.RequestBeginBlock.Header.Version.App,
		ChainId:                       ss.blockMetadata.RequestBeginBlock.Header.ChainID,
		Time:                          ss.blockMetadata.RequestBeginBlock.Header.Time.Unix(),
		LastBlockIdHash:               ss.blockMetadata.RequestBeginBlock.Header.LastBlockId.Hash,
		LastBlockIdPartSetHeaderTotal: ss.blockMetadata.RequestBeginBlock.Header.LastBlockId.PartSetHeader.Total,
		LastBlockIdPartSetHeaderHash:  ss.blockMetadata.RequestBeginBlock.Header.LastBlockId.PartSetHeader.Hash,
		LastCommitHash:                ss.blockMetadata.RequestBeginBlock.Header.LastCommitHash,
		DataHash:                      ss.blockMetadata.RequestBeginBlock.Header.DataHash,
		ValidatorsHash:                ss.blockMetadata.RequestBeginBlock.Header.ValidatorsHash,
		NextValidatorsHash:            ss.blockMetadata.RequestBeginBlock.Header.NextValidatorsHash,
		ConsensusHash:                 ss.blockMetadata.RequestBeginBlock.Header.ConsensusHash,
		AppHash:                       ss.blockMetadata.RequestBeginBlock.Header.AppHash,
		LastResultsHash:               ss.blockMetadata.RequestBeginBlock.Header.LastResultsHash,
		EvidenceHash:                  ss.blockMetadata.RequestBeginBlock.Header.EvidenceHash,
		ProposerAddress:               ss.blockMetadata.RequestBeginBlock.Header.ProposerAddress,
		LastCommitInfoRound:           ss.blockMetadata.RequestBeginBlock.LastCommitInfo.Round,
	}

	if ss.blockMetadata.ResponseEndBlock.ConsensusParamUpdates != nil {
		block.ConsensusParamUpdatesBlockMaxBytes = ss.blockMetadata.ResponseEndBlock.ConsensusParamUpdates.Block.MaxBytes
		block.ConsensusParamUpdatesBlockMaxGas = ss.blockMetadata.ResponseEndBlock.ConsensusParamUpdates.Block.MaxGas
		block.ConsensusParamUpdatesEvidenceMaxAgeNumBlocks = ss.blockMetadata.ResponseEndBlock.ConsensusParamUpdates.Evidence.MaxAgeNumBlocks
		block.ConsensusParamUpdatesEvidenceMaxAgeDuration = ss.blockMetadata.ResponseEndBlock.ConsensusParamUpdates.Evidence.MaxAgeDuration.Milliseconds()
		block.ConsensusParamUpdatesEvidenceMaxBytes = ss.blockMetadata.ResponseEndBlock.ConsensusParamUpdates.Evidence.MaxBytes
		block.ConsensusParamUpdatesValidatorPubKeyTypes = strings.Join(ss.blockMetadata.ResponseEndBlock.ConsensusParamUpdates.Validator.PubKeyTypes[:], ",") //todo  []string to string
		block.ConsensusParamUpdatesVersionApp = ss.blockMetadata.ResponseEndBlock.ConsensusParamUpdates.Version.App
	}

	builder.SetBlock(block)

	for _, validatorUpdate := range ss.blockMetadata.ResponseEndBlock.ValidatorUpdates {
		builder.AppendValidatorUpdates([]cosmos.ValidatorUpdate{
			{
				Seq:    blockHeight,
				PubKey: validatorUpdate.PubKey.GetSecp256K1(),
				Power:  validatorUpdate.Power,
			},
		})
	}

	for _, misbehavior := range ss.blockMetadata.RequestBeginBlock.ByzantineValidators {
		builder.AppendMisbehaviors([]cosmos.Misbehavior{
			{
				Seq:              blockHeight,
				BlockSeq:         blockHeight,
				Typ:              misbehavior.Type.String(),
				ValidatorPower:   misbehavior.Validator.Power,
				ValidatorAddress: sdktypes.ValAddress(misbehavior.Validator.Address).String(),
				Height:           misbehavior.Height,
				Time:             misbehavior.Time.Unix(),
				TotalVotingPower: misbehavior.TotalVotingPower,
			},
		})
	}

	for _, tx := range ss.blockMetadata.DeliverTxs {
		builder.AppendTxs([]cosmos.Transaction{
			{
				Seq:  blockHeight,
				Tx:   tx.Request.Tx,
				Code: tx.Response.Code,
				Data: tx.Response.Data,
			},
		})
	}

	for _, event := range ss.blockMetadata.ResponseEndBlock.Events {
		builder.AppendEvents([]cosmos.Event{
			{
				Seq:       blockHeight,
				EventType: event.Type,
			},
		})
		for _, attribute := range event.Attributes {
			builder.AppendEventAttributes([]cosmos.EventAttribute{
				{
					Seq:      blockHeight,
					EventSeq: blockHeight,
					Key:      attribute.Key,
					Value:    attribute.Value,
					Index:    attribute.Index,
				},
			})
		}
	}

	builder.SetBlockData(strconv.FormatUint(blockHeight, 10), string(ss.blockMetadata.RequestBeginBlock.Hash))
	statTxs := uint64(len(ss.blockMetadata.DeliverTxs))
	statEvn := uint64(len(ss.blockMetadata.ResponseEndBlock.Events))
	builder.SetStatistics(uint64(1), statTxs, statEvn, 0)

	cosmosCtx := builder.Finish()

	ss.logger.Info("Mamoru Send", "height", ss.currentBlockNumber)

	if client := ss.sniffer.Client(); client != nil {
		client.ObserveCosmosData(cosmosCtx)
	}

	return nil
}

func (ss *StreamingService) Stream(wg *sync.WaitGroup) error {
	return nil
}

func (ss *StreamingService) Listeners() map[store.StoreKey][]store.WriteListener {
	listeners := make(map[types.StoreKey][]types.WriteListener, len(ss.storeListeners))
	//for _, listener := range ss.storeListeners {
	//	listeners[listener.StoreKey()] = []types.WriteListener{listener}
	//}
	return listeners
}

func (ss StreamingService) Close() error {
	return nil
}
