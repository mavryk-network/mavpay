package common

import (
	"github.com/mavryk-network/mavpay/constants/enums"
	"github.com/mavryk-network/mvgo/mavryk"
)

type Delegator struct {
	Address          mavryk.Address
	DelegatedBalance mavryk.Z
	StakedBalance    mavryk.Z
	Emptied          bool
}

type BakersCycleData struct {
	OwnDelegatedBalance              mavryk.Z
	ExternalDelegatedBalance         mavryk.Z
	BlockDelegatedRewards            mavryk.Z
	IdealBlockDelegatedRewards       mavryk.Z
	EndorsementDelegatedRewards      mavryk.Z
	IdealEndorsementDelegatedRewards mavryk.Z
	BlockDelegatedFees               mavryk.Z
	DelegatorsCount                  int32

	OwnStakedBalance              mavryk.Z
	ExternalStakedBalance         mavryk.Z
	BlockStakingRewardsEdge       mavryk.Z
	EndorsementStakingRewardsEdge mavryk.Z
	BlockStakingFees              mavryk.Z
	StakersCount                  int32

	FrozenDepositLimit mavryk.Z
	Delegators         []Delegator
}

type ShareInfo struct {
	Baker      mavryk.Z
	Delegators map[string]mavryk.Z
}

func (cycleData *BakersCycleData) getActualDelegatedRewards() mavryk.Z {
	return cycleData.BlockDelegatedFees.Add(cycleData.BlockDelegatedRewards).Add(cycleData.EndorsementDelegatedRewards)
}

func (cycleData *BakersCycleData) getIdealDelegatedRewards() mavryk.Z {
	return cycleData.IdealBlockDelegatedRewards.Add(cycleData.IdealEndorsementDelegatedRewards).Add(cycleData.BlockDelegatedFees)
}

// GetTotalDelegatedRewards returns the total rewards for the cycle based on payout mode
func (cycleData *BakersCycleData) GetTotalDelegatedRewards(payoutMode enums.EPayoutMode) mavryk.Z {
	switch payoutMode {
	case enums.PAYOUT_MODE_IDEAL:
		return cycleData.getIdealDelegatedRewards()
	default:
		return cycleData.getActualDelegatedRewards()
	}
}

func (cycleData *BakersCycleData) GetBakerDelegatedBalance() mavryk.Z {
	return cycleData.OwnDelegatedBalance
}

func (cycleData *BakersCycleData) GetBakerStakedBalance() mavryk.Z {
	return cycleData.OwnStakedBalance
}

type OperationLimits struct {
	HardGasLimitPerOperation     int64
	HardStorageLimitPerOperation int64
	MaxOperationDataLength       int
}
