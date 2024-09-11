package generate

import (
	"testing"

	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/configuration"
	"github.com/mavryk-network/mvgo/mavryk"
	"github.com/stretchr/testify/assert"
)

func TestDelegatorToPayoutCandidate(t *testing.T) {
	assert := assert.New(t)

	config := configuration.GetDefaultRuntimeConfiguration()

	maximumBalance := mavryk.NewZ(100000000)
	config.Delegators.Overrides = map[string]configuration.RuntimeDelegatorOverride{
		"mv1HCXRedE7zVSwmSqxDe3XZcMPLeF7xYqP3": {
			MaximumBalance: &maximumBalance,
		},
		"mv1Qe2hoRHRHYxYCHzD8vUX2We8uEJrEdWAb": {
			MaximumBalance: &maximumBalance,
		},
	}

	delegators := []common.Delegator{
		{
			Address:          mavryk.MustParseAddress("mv1HCXRedE7zVSwmSqxDe3XZcMPLeF7xYqP3"),
			DelegatedBalance: mavryk.NewZ(100000000),
		},
		{
			Address:          mavryk.MustParseAddress("mv1Qe2hoRHRHYxYCHzD8vUX2We8uEJrEdWAb"),
			DelegatedBalance: mavryk.NewZ(200000000),
		},
	}

	delegator := delegators[0]
	candidate := DelegatorToPayoutCandidate(delegator, &config)
	assert.True(candidate.GetDelegatedBalance().Equal(mavryk.MinZ(delegator.DelegatedBalance, maximumBalance)))

	delegator = delegators[1]
	candidate = DelegatorToPayoutCandidate(delegator, &config)
	assert.True(candidate.GetDelegatedBalance().Equal(mavryk.MinZ(delegator.DelegatedBalance, maximumBalance)))

	config.Delegators.Overrides = map[string]configuration.RuntimeDelegatorOverride{}

	delegator = delegators[0]
	candidate = DelegatorToPayoutCandidate(delegator, &config)
	assert.True(candidate.GetDelegatedBalance().Equal(delegator.DelegatedBalance))

	delegator = delegators[1]
	candidate = DelegatorToPayoutCandidate(delegator, &config)
	assert.True(candidate.GetDelegatedBalance().Equal(delegator.DelegatedBalance))
}
