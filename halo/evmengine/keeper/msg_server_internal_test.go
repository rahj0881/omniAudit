package keeper

import (
	"encoding/json"
	"math/big"
	"testing"
	"time"

	"github.com/omni-network/omni/halo/evmengine/types"
	"github.com/omni-network/omni/lib/k1util"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"
)

func Test_msgServer_ExecutionPayload(t *testing.T) {
	t.Parallel()
	cdc := getCodec(t)
	txConfig := authtx.NewTxConfig(cdc, nil)

	mockEngine, err := newMockEngineAPI()
	require.NoError(t, err)
	cmtAPI := newMockCometAPI(t, nil)
	// set the header and proposer so we have the correct next proposer
	header := cmtproto.Header{Height: 1}
	header.ProposerAddress = cmtAPI.validatorSet.Validators[0].Address
	nxtAddr, err := k1util.PubKeyToAddress(cmtAPI.validatorSet.Validators[1].PubKey)
	require.NoError(t, err)

	ctx, storeService := setupCtxStore(t, &header)
	ctx = ctx.WithExecMode(sdk.ExecModeFinalize)

	keeper := NewKeeper(cdc, storeService, &mockEngine, txConfig)
	ap := mockAddressProvider{
		address: nxtAddr,
	}
	keeper.SetAddressProvider(ap)
	keeper.SetCometAPI(&cmtAPI)
	msgSrv := NewMsgServerImpl(keeper)

	var payloadData []byte
	newPayload := func() {
		// get latest block to build on top of
		latestHeight, err := mockEngine.BlockNumber(ctx)
		require.NoError(t, err)
		latestBlock, err := mockEngine.BlockByNumber(ctx, big.NewInt(int64(latestHeight)))
		require.NoError(t, err)

		payloadID := mockEngine.pushPayload(t, ctx, ap.LocalAddress(), latestBlock.Hash(), time.Now())
		payloadResp, err := mockEngine.GetPayloadV2(ctx, *payloadID)
		require.NoError(t, err)

		// Create execution payload message
		payloadData, err = json.Marshal(payloadResp.ExecutionPayload)
		require.NoError(t, err)
	}

	assertExecutionPayload := func() {
		resp, err := msgSrv.ExecutionPayload(ctx, &types.MsgExecutionPayload{
			Authority: authtypes.NewModuleAddress(types.ModuleName).String(),
			Data:      payloadData,
		})
		require.NoError(t, err)
		require.NotNil(t, resp)
	}

	newPayload()
	assertExecutionPayload()

	// not lets run optimistic flow
	newPayload()
	keeper.SetBuildOptimistic(true)
	assertExecutionPayload()
}
