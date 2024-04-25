package resolvers

import (
	"context"

	"github.com/omni-network/omni/explorer/graphql/utils"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type XMsgRangeArgs struct {
	From hexutil.Big
	To   hexutil.Big
}

type XMsgArgs struct {
	SourceChainID hexutil.Big
	DestChainID   hexutil.Big
	StreamOffset  hexutil.Big
}

type XMsgsArgs struct {
	Limit         *hexutil.Big
	Cursor        *hexutil.Big
	SourceChainID *hexutil.Big
	DestChainID   *hexutil.Big
	Address       *common.Address
}

const MsgsLimit uint64 = 25

func (b *BlocksResolver) XMsgCount(ctx context.Context) (*hexutil.Big, error) {
	res, found, err := b.BlocksProvider.XMsgCount(ctx)
	if err != nil {
		return nil, errors.New("failed to fetch message count")
	}
	if !found {
		return nil, errors.New("message count not found")
	}

	return res, nil
}

func (b *BlocksResolver) XMsgRange(ctx context.Context, args XMsgRangeArgs) ([]*XMsg, error) {
	if args.From.ToInt().Cmp(args.To.ToInt()) >= 0 {
		return nil, errors.New("invalid range")
	}

	res, found, err := b.BlocksProvider.XMsgRange(ctx, args.From.ToInt().Uint64(), args.To.ToInt().Uint64())
	if err != nil {
		return nil, errors.New("failed to fetch messages")
	}
	if !found {
		return nil, errors.New("messages not found")
	}

	return res, nil
}

func (b *BlocksResolver) XMsg(ctx context.Context, args XMsgArgs) (*XMsg, error) {
	res, found, err := b.BlocksProvider.XMsg(ctx, args.SourceChainID.ToInt().Uint64(), args.DestChainID.ToInt().Uint64(), args.StreamOffset.ToInt().Uint64())
	if err != nil {
		return nil, errors.New("failed to fetch message")
	}
	if !found {
		return nil, errors.New("message not found")
	}

	return res, nil
}

func (b *BlocksResolver) XMsgs(ctx context.Context, args XMsgsArgs) (*XMsgResult, error) {
	limit := uint64(1)
	var cursor, sourceChainID, destChainID *uint64

	if args.Limit != nil {
		limit = args.Limit.ToInt().Uint64()
	}

	if limit > MsgsLimit || args.Limit == nil {
		limit = MsgsLimit
	}

	cursor = utils.TryParseHexUtilBigToUint64(args.Cursor)
	sourceChainID = utils.TryParseHexUtilBigToUint64(args.SourceChainID)
	destChainID = utils.TryParseHexUtilBigToUint64(args.DestChainID)

	res, found, err := b.BlocksProvider.XMsgs(ctx, limit, cursor, sourceChainID, destChainID, args.Address)
	if err != nil {
		return nil, errors.New("failed to fetch messages")
	}
	if !found {
		return nil, errors.New("messages not found")
	}

	return res, nil
}
