package generate

import (
	"blockwatch.cc/tzgo/tezos"
	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/constants"
	"github.com/mavryk-network/mvgo/codec"
)

func buildOpForEstimation[T common.TransferArgs](ctx *PayoutGenerationContext, batch []T, injectBurnTransactions bool) (*codec.Op, error) {
	var err error
	op := codec.NewOp().WithSource(ctx.PayoutKey.Address())
	op.WithTTL(constants.MAX_OPERATION_TTL)
	if injectBurnTransactions {
		op.WithTransfer(tezos.BurnAddress, 1)
	}
	for _, p := range batch {
		if err = common.InjectTransferContents(op, ctx.PayoutKey.Address(), p); err != nil {
			break
		}
	}
	if injectBurnTransactions {
		op.WithTransfer(tezos.BurnAddress, 1)
	}
	return op, err
}
