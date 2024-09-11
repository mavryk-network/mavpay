package generate

import (
	"testing"

	"blockwatch.cc/tzgo/tezos"
	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/configuration"
	"github.com/stretchr/testify/assert"
)

func TestDelegatorToPayoutCandidate(t *testing.T) {
	assert := assert.New(t)

	config := configuration.GetDefaultRuntimeConfiguration()

	maximumBalance := tezos.NewZ(100000000)
	config.Delegators.Overrides = map[string]configuration.RuntimeDelegatorOverride{
		"tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM": {
			MaximumBalance: &maximumBalance,
		},
		"tz1hZvgjekGo7DmQjWh7XnY5eLQD8wNYPczE": {
			MaximumBalance: &maximumBalance,
		},
	}

	delegators := []common.Delegator{
		{
			Address:          tezos.MustParseAddress("tz1P6WKJu2rcbxKiKRZHKQKmKrpC9TfW1AwM"),
			DelegatedBalance: tezos.NewZ(100000000),
		},
		{
			Address:          tezos.MustParseAddress("tz1hZvgjekGo7DmQjWh7XnY5eLQD8wNYPczE"),
			DelegatedBalance: tezos.NewZ(200000000),
		},
	}

	delegator := delegators[0]
	candidate := DelegatorToPayoutCandidate(delegator, &config)
	assert.True(candidate.GetDelegatedBalance().Equal(tezos.MinZ(delegator.DelegatedBalance, maximumBalance)))

	delegator = delegators[1]
	candidate = DelegatorToPayoutCandidate(delegator, &config)
	assert.True(candidate.GetDelegatedBalance().Equal(tezos.MinZ(delegator.DelegatedBalance, maximumBalance)))

	config.Delegators.Overrides = map[string]configuration.RuntimeDelegatorOverride{}

	delegator = delegators[0]
	candidate = DelegatorToPayoutCandidate(delegator, &config)
	assert.True(candidate.GetDelegatedBalance().Equal(delegator.DelegatedBalance))

	delegator = delegators[1]
	candidate = DelegatorToPayoutCandidate(delegator, &config)
	assert.True(candidate.GetDelegatedBalance().Equal(delegator.DelegatedBalance))
}
