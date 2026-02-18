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
		OwnStakedBalance:             mavryk.NewZ(500_000),
		OwnDelegatedBalance:          mavryk.NewZ(500_000),
		ExternalDelegatedBalance:     mavryk.NewZ(19_000_000),
		BlockDelegatedRewards:        mavryk.NewZ(1000),
		AttestationsDelegatedRewards: mavryk.NewZ(10000),
		DalDelegatedRewards:          mavryk.NewZ(100),
	}

	bakerBondsAmount := getBakerBondsAmount(&cycleData, mavryk.NewZ(19_000_000), &configWithOverdelegationProtectionEnabled)
	assert.Equal(bakerBondsAmount.Int64(), mavryk.NewZ(1233).Int64())

	bakerBondsAmount = getBakerBondsAmount(&cycleData, mavryk.NewZ(19_000_000), &configWithOverdelegationProtectionDisabled)
	assert.Equal(bakerBondsAmount.Int64(), mavryk.NewZ(284).Int64())

	bakerBondsAmount = getBakerBondsAmount(&cycleData, mavryk.NewZ(9_000_000), &configWithOverdelegationProtectionEnabled)
	assert.Equal(bakerBondsAmount.Int64(), mavryk.NewZ(1233).Int64())

	bakerBondsAmount = getBakerBondsAmount(&cycleData, mavryk.NewZ(9_000_000), &configWithOverdelegationProtectionDisabled)
	assert.Equal(bakerBondsAmount.Int64(), mavryk.NewZ(584).Int64())

	cycleData = common.BakersCycleData{
		OwnStakedBalance:             mavryk.NewZ(600_000),
		OwnDelegatedBalance:          mavryk.NewZ(400_000),
		ExternalDelegatedBalance:     mavryk.NewZ(9_000_000),
		BlockDelegatedRewards:        mavryk.NewZ(1000),
		AttestationsDelegatedRewards: mavryk.NewZ(10000),
		DalDelegatedRewards:          mavryk.NewZ(100),
	}

	bakerBondsAmount = getBakerBondsAmount(&cycleData, mavryk.NewZ(9_000_000), &configWithOverdelegationProtectionEnabled)
	assert.Equal(bakerBondsAmount.Int64(), mavryk.NewZ(822).Int64())

	bakerBondsAmount = getBakerBondsAmount(&cycleData, mavryk.NewZ(9_000_000), &configWithOverdelegationProtectionDisabled)
	assert.Equal(bakerBondsAmount.Int64(), mavryk.NewZ(472).Int64())
}
