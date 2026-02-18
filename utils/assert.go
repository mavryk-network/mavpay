package utils

import "github.com/mavryk-network/gomavryk/mavryk"

func AssertZAmountPositiveOrZero(amount mavryk.Z) {
	if amount.IsNeg() {
		panic("amount is negative, this should never happen")
	}
}
