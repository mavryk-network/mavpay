package common

import (
	"github.com/mavryk-network/mavpay/constants"
	"github.com/mavryk-network/mavpay/constants/enums"
	"github.com/mavryk-network/mvgo/codec"
	"github.com/mavryk-network/mvgo/contract"
	"github.com/mavryk-network/mvgo/mavryk"
	"github.com/mavryk-network/mvgo/rpc"
)

type OpExecutionContext struct {
	Op         *codec.Op
	Transactor TransactorEngine
	result     OpResult
}

func InitOpExecutionContext(op *codec.Op, transactor TransactorEngine) *OpExecutionContext {
	return &OpExecutionContext{
		Op:         op,
		result:     nil,
		Transactor: transactor,
	}
}

func (ctx *OpExecutionContext) GetOpHash() mavryk.OpHash {
	if ctx.result == nil {
		return mavryk.ZeroOpHash
	}
	return ctx.result.GetOpHash()
}

func (ctx *OpExecutionContext) Dispatch(opts *rpc.CallOptions) error {
	result, err := ctx.Transactor.Dispatch(ctx.Op, opts)
	if err != nil {
		return err
	}
	ctx.result = result
	return err
}

func (ctx *OpExecutionContext) WaitForApply() error {
	if ctx.result == nil {
		return constants.ErrOperationNotDispatched
	}
	return ctx.result.WaitForApply()
}

type TransferArgs interface {
	GetTxKind() enums.EPayoutTransactionKind
	GetFAContract() mavryk.Address
	GetFATokenId() mavryk.Z
	GetDestination() mavryk.Address
	GetAmount() mavryk.Z
}

func InjectTransferContents(op *codec.Op, source mavryk.Address, p TransferArgs) error {
	switch p.GetTxKind() {
	case enums.PAYOUT_TX_KIND_FA1_2:
		if p.GetFAContract().Equal(mavryk.ZeroAddress) || p.GetFAContract().Equal(mavryk.InvalidAddress) {
			return constants.ErrOperationInvalidContractAddress
		}
		args := contract.NewFA1TransferArgs().WithTransfer(source, p.GetDestination(), p.GetAmount()).
			WithSource(source).
			WithDestination(p.GetFAContract())
		op.WithContents(args.Encode())
	case enums.PAYOUT_TX_KIND_FA2:
		if p.GetFAContract().Equal(mavryk.ZeroAddress) || p.GetFAContract().Equal(mavryk.InvalidAddress) {
			return constants.ErrOperationInvalidContractAddress
		}
		args := contract.NewFA2TransferArgs().WithTransfer(source, p.GetDestination(), p.GetFATokenId(), p.GetAmount()).
			WithSource(source).
			WithDestination(p.GetFAContract())
		op.WithContents(args.Encode())
	default:
		op.WithTransfer(p.GetDestination(), p.GetAmount().Int64())
	}
	return nil
}

func InjectTransferContentsWithLimits(op *codec.Op, source mavryk.Address, p TransferArgs, limits mavryk.Limits) error {
	err := InjectTransferContents(op, source, p)
	if err != nil {
		return err
	}
	op.Contents[len(op.Contents)-1].WithLimits(limits)
	return nil
}

func InjectLimits(op *codec.Op, limits []mavryk.Limits) error {
	if len(limits) == 0 {
		return nil
	}
	if len(limits) != len(op.Contents) {
		return constants.ErrOperationInvalidLimits
	}
	for i := range op.Contents {
		op.Contents[i].WithLimits(limits[i])
	}
	return nil
}
