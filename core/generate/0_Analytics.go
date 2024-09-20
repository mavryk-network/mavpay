package generate

import (
	"os"

	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/constants"
)

func SendAnalytics(ctx *PayoutGenerationContext, options *common.GeneratePayoutsOptions) (*PayoutGenerationContext, error) {
	configuration := ctx.GetConfiguration()

	if os.Getenv("DISABLE_MAVPAY_ANALYTICS") == "true" {
		return ctx, nil
	}

	if configuration.DisableAnalytics {
		return ctx, nil
	}

	ctx.GetCollector().SendAnalytics(configuration.BakerPKH.String(), constants.VERSION)

	return ctx, nil
}
