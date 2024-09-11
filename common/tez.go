package common

import (
	"fmt"

	"github.com/mavryk-network/mavpay/constants"
	"github.com/mavryk-network/mavpay/constants/enums"
	"github.com/mavryk-network/mvgo/mavryk"
)

func FormatAmount(kind enums.EPayoutTransactionKind, amount int64) string {
	if amount == 0 {
		return ""
	}
	switch kind {
	case enums.PAYOUT_TX_KIND_FA1_2:
		return fmt.Sprintf("%d FA1", amount)
	case enums.PAYOUT_TX_KIND_FA2:
		return fmt.Sprintf("%d FA2", amount)
	default:
		return MutezToTezS(amount)
	}
}

func MutezToTezS(amount int64) string {
	if amount == 0 {
		return ""
	}
	tez := float64(amount) / constants.MUTEZ_FACTOR
	return fmt.Sprintf("%f TEZ", tez)
}

func FloatToPercentage(f float64) string {
	if f == 0 {
		return ""
	}
	return fmt.Sprintf("%.2f%%", f*100)
}

func ShortenAddress(taddr mavryk.Address) string {
	if taddr.Equal(mavryk.ZeroAddress) || taddr.Equal(mavryk.InvalidAddress) {
		return ""
	}
	addr := taddr.String()
	total := len(addr)
	if total <= 13 {
		return addr
	}
	return fmt.Sprintf("%s...%s", addr[:5], addr[total-5:])
}

func ToStringEmptyIfZero[T comparable](value T) string {
	var zero T
	if value == zero {
		return ""
	}
	return fmt.Sprintf("%v", value)
}
