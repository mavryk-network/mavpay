package main

import (
	"log/slog"
	"math"

	"github.com/mavryk-network/gomavryk/mavryk"
)

type Exchanger interface {
	RefreshExchangeRate() error
	ExchangeMavToToken(mumav mavryk.Z) mavryk.Z
}

type FixedAmountExchanger struct {
	amount float64
	token  TokenConfiguration
}

func NewFixedAmountExchanger(amount float64, token TokenConfiguration) FixedAmountExchanger {
	return FixedAmountExchanger{
		amount: amount,
		token:  token,
	}
}

func (e FixedAmountExchanger) RefreshExchangeRate() error {
	return nil
}

func (e FixedAmountExchanger) ExchangeMavToToken(_ mavryk.Z) mavryk.Z {
	slog.Debug("Exchanging fixed amount", "amount", e.amount, "token", e.token)
	// we need to multiply by 1_000_000 because other functions assume we calculated with mumav
	return mavryk.NewZ(int64(e.amount * math.Pow10(e.token.Decimals)))
}

type FixedRateExchanger struct {
	rate  mavryk.Z
	fee   mavryk.Z
	token TokenConfiguration
}

func NewFixedRateExchanger(rate, fee float64, token TokenConfiguration) FixedRateExchanger {
	return FixedRateExchanger{
		rate:  mavryk.NewZ(int64(rate * float64(PRECISION))),
		fee:   mavryk.NewZ(int64(fee * float64(PRECISION))),
		token: token,
	}
}

func (e FixedRateExchanger) RefreshExchangeRate() error {
	return nil
}

func (e FixedRateExchanger) ExchangeMavToToken(mumav mavryk.Z) mavryk.Z {
	decimalsMultiplier := int64(math.Pow10(e.token.Decimals))

	token_amount := mumav.Mul(e.rate).Mul64(decimalsMultiplier).Div(mavryk.NewZ(PRECISION).Sub(e.fee)).Div64(MUMAV_FACTOR)
	slog.Debug("Exchanging amount", "amount", mumav, "token_amount", token_amount, "rate", e.rate, "fee", e.fee)
	return token_amount
}
