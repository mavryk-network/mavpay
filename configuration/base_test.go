package configuration

import (
	"strings"
	"testing"

	mavpay_configuration "github.com/mavryk-network/mavpay/configuration/v"
	"github.com/mavryk-network/mvgo/mavryk"
	test_assert "github.com/stretchr/testify/assert"
)

func TestConfigurationToRuntimeConfiguration(t *testing.T) {
	assert := test_assert.New(t)
	runtime, _ := ConfigurationToRuntimeConfiguration(&LatestConfigurationType{
		Delegators: mavpay_configuration.DelegatorsConfigurationV0{
			FeeOverrides: map[string][]mavryk.Address{
				".5": {mavryk.InvalidAddress, mavryk.BurnAddress},
				"1":  {mavryk.ZeroAddress},
			},
		},
	})
	val, ok := runtime.Delegators.Overrides[mavryk.InvalidAddress.String()]
	assert.True(ok)
	assert.Equal(*val.Fee, 0.5)

	val, ok = runtime.Delegators.Overrides[mavryk.BurnAddress.String()]
	assert.True(ok)
	assert.Equal(*val.Fee, 0.5)

	val, ok = runtime.Delegators.Overrides[mavryk.ZeroAddress.String()]
	assert.True(ok)
	assert.Equal(*val.Fee, float64(1))

	runtime, _ = ConfigurationToRuntimeConfiguration(&LatestConfigurationType{
		Delegators: mavpay_configuration.DelegatorsConfigurationV0{
			FeeOverrides: map[string][]mavryk.Address{
				"0": {mavryk.InvalidAddress, mavryk.BurnAddress},
			},
		},
	})

	val, ok = runtime.Delegators.Overrides[mavryk.InvalidAddress.String()]
	assert.True(ok)
	assert.Equal(*val.Fee, 0.)

	val, ok = runtime.Delegators.Overrides[mavryk.BurnAddress.String()]
	assert.True(ok)
	assert.Equal(*val.Fee, 0.)

	fee := 1.0
	runtime, _ = ConfigurationToRuntimeConfiguration(&LatestConfigurationType{
		Delegators: mavpay_configuration.DelegatorsConfigurationV0{
			FeeOverrides: map[string][]mavryk.Address{
				"0": {mavryk.InvalidAddress, mavryk.BurnAddress},
			},
			Overrides: map[string]mavpay_configuration.DelegatorOverrideV0{
				mavryk.InvalidAddress.String(): {
					Fee: &fee,
				},
			},
		},
	})

	val, ok = runtime.Delegators.Overrides[mavryk.InvalidAddress.String()]
	assert.True(ok)
	assert.Equal(*val.Fee, float64(1))

	val, ok = runtime.Delegators.Overrides[mavryk.BurnAddress.String()]
	assert.True(ok)
	assert.Equal(*val.Fee, 0.)

	runtime, _ = ConfigurationToRuntimeConfiguration(&LatestConfigurationType{
		Delegators: mavpay_configuration.DelegatorsConfigurationV0{
			FeeOverrides: map[string][]mavryk.Address{
				"1.1": {mavryk.InvalidAddress, mavryk.BurnAddress},
			},
		},
	})

	err := runtime.Validate()
	assert.NotNil(err)
	assert.True(strings.Contains(err.Error(), "fee must be between 0 and 1"))
}
