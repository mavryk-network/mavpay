package generate

import (
	"testing"

	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/configuration"
	"github.com/mavryk-network/mvgo/mavryk"
	"github.com/stretchr/testify/assert"
)

func TestGetBakerBondsAmount(t *testing.T) {
	assert := assert.New(t)

	configWithOverdelegationProtectionEnabled := configuration.GetDefaultRuntimeConfiguration()
	configWithOverdelegationProtectionDisabled := configuration.GetDefaultRuntimeConfiguration()
	configWithOverdelegationProtectionDisabled.Overdelegation.IsProtectionEnabled = false

	cycleData := common.BakersCycleData{
		OwnStakedBalance:            mavryk.NewZ(500_000),
		OwnDelegatedBalance:         mavryk.NewZ(500_000),
		ExternalDelegatedBalance:    mavryk.NewZ(19_000_000),
		BlockDelegatedRewards:       mavryk.NewZ(1000),
		EndorsementDelegatedRewards: mavryk.NewZ(10000),
	}

	bakerBondsAmount := getBakerBondsAmount(&cycleData, mavryk.NewZ(19_000_000), &configWithOverdelegationProtectionEnabled)
	assert.Equal(bakerBondsAmount.Int64(), mavryk.NewZ(1222).Int64())

	bakerBondsAmount = getBakerBondsAmount(&cycleData, mavryk.NewZ(19_000_000), &configWithOverdelegationProtectionDisabled)
	assert.Equal(bakerBondsAmount.Int64(), mavryk.NewZ(282).Int64())

	bakerBondsAmount = getBakerBondsAmount(&cycleData, mavryk.NewZ(9_000_000), &configWithOverdelegationProtectionEnabled)
	assert.Equal(bakerBondsAmount.Int64(), mavryk.NewZ(1222).Int64())

	bakerBondsAmount = getBakerBondsAmount(&cycleData, mavryk.NewZ(9_000_000), &configWithOverdelegationProtectionDisabled)
	assert.Equal(bakerBondsAmount.Int64(), mavryk.NewZ(578).Int64())

	cycleData = common.BakersCycleData{
		OwnStakedBalance:            mavryk.NewZ(600_000),
		OwnDelegatedBalance:         mavryk.NewZ(400_000),
		ExternalDelegatedBalance:    mavryk.NewZ(9_000_000),
		BlockDelegatedRewards:       mavryk.NewZ(1000),
		EndorsementDelegatedRewards: mavryk.NewZ(10000),
	}

	bakerBondsAmount = getBakerBondsAmount(&cycleData, mavryk.NewZ(9_000_000), &configWithOverdelegationProtectionEnabled)
	assert.Equal(bakerBondsAmount.Int64(), mavryk.NewZ(814).Int64())

	bakerBondsAmount = getBakerBondsAmount(&cycleData, mavryk.NewZ(9_000_000), &configWithOverdelegationProtectionDisabled)
	assert.Equal(bakerBondsAmount.Int64(), mavryk.NewZ(468).Int64())
}
