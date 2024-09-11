//go:build js && wasm

package main

import (
	"encoding/json"
	"errors"
	"log/slog"
	"syscall/js"

	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/configuration"
	"github.com/mavryk-network/mavpay/constants"
	"github.com/mavryk-network/mavpay/core"
	"github.com/mavryk-network/mvgo/mavryk"
)

func main() {
	slog.Info("mavpay wasm loaded", "version", constants.VERSION)
}

//export generate_payouts
func generate_payouts(key js.Value, cycle int64, configurationJs js.Value) (js.Value, error) {
	configurationBytes := []byte(configurationJs.String())
	config, err := configuration.LoadFromString(configurationBytes)
	if err != nil {
		return js.Null(), err
	}

	bakerKey, err := mavryk.ParseKey(key.String())
	if err != nil {
		return js.Null(), err
	}

	payoutBlueprint, err := core.GeneratePayoutsWithPayoutAddress(bakerKey, config, common.GeneratePayoutsOptions{
		Cycle:            cycle,
		SkipBalanceCheck: true,
		Engines: common.GeneratePayoutsEngines{
			//FIXME possible JSCollector/JSSigner interfaced from JS
			Collector: nil,
		},
	})
	if err != nil {
		return js.Null(), err
	}

	result, err := json.Marshal(payoutBlueprint)

	return js.ValueOf(string(result)), err
}

//export test
func test(data js.Value) (js.Value, error) {
	x := data.String()
	slog.Info(x)
	return js.ValueOf(x), errors.New("test")
}
