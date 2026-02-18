package mvkt

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/hashicorp/go-version"
	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/constants"
	"github.com/mavryk-network/mvgo/mavryk"

	"github.com/samber/lo"
)

const (
	DELEGATOR_FETCH_LIMIT = 10000
)

type splitDelegator struct {
	Address string `json:"address"`

	DelegatedBalance int64 `json:"delegatedBalance"`
	StakedBalance    int64 `json:"stakedBalance"`

	Emptied bool `json:"emptied,omitempty"`
}

type mvktBakersCycleData struct {
	BakingPower              int64 `json:"bakingPower"`
	OwnDelegatedBalance      int64 `json:"ownDelegatedBalance"`
	ExternalDelegatedBalance int64 `json:"externalDelegatedBalance"`
	OwnStakedBalance         int64 `json:"ownStakedBalance"`      // OwnDelegatedBalance + ExternalDelegatedBalance
	ExternalStakedBalance    int64 `json:"externalStakedBalance"` // ExternalDelegatedBalance

	BlockRewardsDelegated    int64 `json:"blockRewardsDelegated"`
	BlockRewardsStakedOwn    int64 `json:"blockRewardsStakedOwn"`
	BlockRewardsStakedEdge   int64 `json:"blockRewardsStakedEdge"`
	BlockRewardsStakedShared int64 `json:"blockRewardsStakedShared"`
	// BlockRewards             int64            `json:"blockRewards"` // BlockRewardsLiquid + BlockRewardsStakedOwn
	MissedBlockRewards int64 `json:"missedBlockRewards"`

	AttestationRewardsDelegated    int64 `json:"attestationRewardsDelegated"`
	AttestationRewardsStakedOwn    int64 `json:"attestationRewardsStakedOwn"`
	AttestationRewardsStakedEdge   int64 `json:"attestationRewardsStakedEdge"`
	AttestationRewardsStakedShared int64 `json:"attestationRewardsStakedShared"`
	// AttestationRewards       int64            `json:"attestationRewards"` // AttestationRewardsLiquid + AttestationRewardsStakedOwn
	MissedAttestationRewards int64 `json:"missedAttestationRewards"`

	DalRewardsDelegated    int64 `json:"dalAttestationRewardsDelegated"`
	DalRewardsStakedOwn    int64 `json:"dalAttestationRewardsStakedOwn"`
	DalRewardsStakedEdge   int64 `json:"dalAttestationRewardsStakedEdge"`
	DalRewardsStakedShared int64 `json:"dalAttestationRewardsStakedShared"`
	MissedDalRewards       int64 `json:"missedDalAttestationRewards"`

	DelegatorsCount int32 `json:"delegatorsCount"`
	StakersCount    int32 `json:"stakersCount"`
	// NumDelegators            int32            `json:"numDelegators"` // DelegatorsCount
	BlockFees  int64            `json:"blockFees"`
	Delegators []splitDelegator `json:"delegators"`
}

type bakerData struct {
	FrozenDepositLimit       int64 `json:"frozenDepositLimit"`
	LimitOfStakingOverBaking int64 `json:"limitOfStakingOverBaking"`
}

type setDelegateParamtersOps []struct {
	ActivationCycle          int64 `json:"activationCycle"`
	LimitOfStakingOverBaking int64 `json:"limitOfStakingOverBaking"`
}

type Client struct {
	*http.Client

	rootUrl *url.URL
}

type MvktClientOptions struct {
	HttpClient *http.Client
}

func InitClient(rootUrl string, options *MvktClientOptions) (*Client, error) {
	if options == nil {
		options = &MvktClientOptions{}
	}

	root, err := url.Parse(rootUrl)
	if err != nil {
		return nil, err
	}
	if options.HttpClient == nil {
		options.HttpClient = &http.Client{
			Timeout: 10 * time.Second,
		}
	}

	client := &Client{
		Client:  options.HttpClient,
		rootUrl: root,
	}

	return client, nil
}

func (client *Client) Get(ctx context.Context, path string) (*http.Response, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	request, _ := http.NewRequestWithContext(ctx, "GET", client.rootUrl.ResolveReference(rel).String(), nil)

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (client *Client) Options(ctx context.Context, path string) (*http.Response, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	request, _ := http.NewRequestWithContext(ctx, "OPTIONS", client.rootUrl.ResolveReference(rel).String(), nil)

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (client *Client) IsMvktVersionHigherOrEqual(ctx context.Context, rawDesiredVersion string) (bool, error) {
	u := "/v1/accounts/count"
	resp, err := client.Options(ctx, u)
	if err != nil {
		return false, fmt.Errorf("failed to fetch mvkt version: %w", err)
	}
	defer resp.Body.Close()

	rawVersion := resp.Header.Get("mvkt-version")
	if rawVersion == "" {
		return false, fmt.Errorf("mvkt-version header is missing")
	}
	// mvkt-version: 1.14.9.0
	if strings.Count(rawVersion, ".") < 2 || strings.Count(rawDesiredVersion, ".") < 2 {
		return false, fmt.Errorf("invalid version format: version=%s, desiredVersion=%s", rawVersion, rawDesiredVersion)
	}

	ver, err := version.NewVersion(rawVersion)
	if err != nil {
		return false, fmt.Errorf("failed to parse mvkt version: %w", err)
	}
	desiredVer, err := version.NewVersion(rawDesiredVersion)
	if err != nil {
		return false, fmt.Errorf("failed to parse desired version: %w", err)
	}

	return ver.Compare(desiredVer) >= 0, nil
}

func unmarshallMvktResponse[T any](resp *http.Response, result *T) error {
	defer resp.Body.Close()
	err := json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return errors.Join(constants.ErrCycleDataUnmarshalFailed, err)
	}
	return nil
}

// https://api.mavryk.network/v1/rewards/split/${baker}/${cycle}?limit=${limit}&offset=${offset}
func (client *Client) getDelegatorsCycleData(ctx context.Context, baker []byte, cycle int64, limit int32, offset int) ([]splitDelegator, error) {
	u := fmt.Sprintf("v1/rewards/split/%s/%d?limit=%d&offset=%d", baker, cycle, limit, offset)
	slog.Debug("getting delegators data", "baker", baker, "cycle", cycle, "url", u)
	resp, err := client.Get(ctx, u)
	if err != nil {
		return nil, errors.Join(constants.ErrCycleDataFetchFailed, err)
	}
	data := &mvktBakersCycleData{}
	err = unmarshallMvktResponse(resp, data)
	if err != nil {
		return nil, err
	}
	return data.Delegators, nil
}

func (client *Client) getBakerData(ctx context.Context, baker []byte, cycle int64) (*bakerData, error) {
	u := fmt.Sprintf("v1/delegates/%s", baker)
	slog.Debug("getting baker data", "baker", baker, "url", u)
	resp, err := client.Get(ctx, u)
	if err != nil {
		return nil, errors.Join(constants.ErrCycleDataFetchFailed, err)
	}
	data := &bakerData{}
	err = unmarshallMvktResponse(resp, data)
	if err != nil {
		return nil, err
	}

	u = fmt.Sprintf("v1/operations/set_delegate_parameters?sender=%s&status=applied&select=activationCycle%%2ClimitOfStakingOverBaking&sort.desc=id&limit=10000", baker)
	slog.Debug("getting baker limit of staking over baking", "baker", baker, "url", u)
	resp, err = client.Get(ctx, u)
	if err != nil {
		return nil, errors.Join(constants.ErrCycleDataFetchFailed, err)
	}
	var setDelegateParamtersOps setDelegateParamtersOps
	err = unmarshallMvktResponse(resp, &setDelegateParamtersOps)
	if err != nil {
		return nil, err
	}
	for _, op := range setDelegateParamtersOps {
		if op.ActivationCycle+1 <= cycle { // +1 because it activates after cycle is generated
			data.LimitOfStakingOverBaking = op.LimitOfStakingOverBaking
			break
		}
	}

	return data, nil
}

func (client *Client) getCycleData(ctx context.Context, baker []byte, cycle int64) (*mvktBakersCycleData, error) {
	u := fmt.Sprintf("v1/rewards/split/%s/%d?limit=0", baker, cycle)
	slog.Debug("getting cycle data", "baker", baker, "cycle", cycle, "url", u)
	resp, err := client.Get(ctx, u)
	if err != nil {
		return nil, errors.Join(constants.ErrCycleDataFetchFailed, err)
	}
	if resp.StatusCode == 204 {
		return nil, errors.Join(constants.ErrNoCycleDataAvailable, fmt.Errorf("baker: %s", baker))
	}
	mvktBakerCycleData := &mvktBakersCycleData{}
	err = unmarshallMvktResponse(resp, mvktBakerCycleData)
	if err != nil {
		return nil, err
	}
	return mvktBakerCycleData, nil
}

func (client *Client) getFirstBlockCycleAfterTimestamp(ctx context.Context, timestamp time.Time) (int64, error) {
	u := fmt.Sprintf("v1/blocks?select=cycle&limit=1&timestamp.gt=%s", timestamp.Format(time.RFC3339))
	slog.Debug("getting first block cycle after timestamp", "timestamp", timestamp, "url", u)
	resp, err := client.Get(ctx, u)
	if err != nil {
		return 0, errors.Join(constants.ErrCycleDataFetchFailed, err)
	}
	defer resp.Body.Close()
	var cycles []int64
	err = json.NewDecoder(resp.Body).Decode(&cycles)
	if err != nil {
		return 0, errors.Join(constants.ErrCycleDataUnmarshalFailed, err)
	}
	if len(cycles) == 0 {
		return 0, errors.Join(constants.ErrCycleDataFetchFailed, fmt.Errorf("no cycles found"))
	}
	return cycles[0], nil
}

// https://api.mavryk.network/v1/blocks?select=cycle,level&limit=1&timestamp.lt=2020-02-20T02:40:57Z
func (client *Client) GetCyclesInDateRange(ctx context.Context, startDate time.Time, endDate time.Time) ([]int64, error) {
	firstCycle, err := client.getFirstBlockCycleAfterTimestamp(ctx, startDate)
	if err != nil {
		return nil, err
	}
	firstCycleAfterTheRange, err := client.getFirstBlockCycleAfterTimestamp(ctx, endDate)
	if err != nil {
		return nil, err
	}

	cycles := make([]int64, 0, 20)
	for cycle := firstCycle; cycle < firstCycleAfterTheRange; cycle++ {
		cycles = append(cycles, cycle)
	}
	return cycles, nil
}

// https://api.mavryk.network/v1/rewards/split/${baker}/${cycle}?limit=0
func (client *Client) GetCycleData(ctx context.Context, chainId mavryk.ChainIdHash, baker mavryk.Address, cycle int64) (bakersCycleData *common.BakersCycleData, err error) {
	bakerAddr, _ := baker.MarshalText()

	mvktBakerData, err := client.getBakerData(ctx, bakerAddr, cycle-2 /* we check against cycle which was used to calculate rewards - c - 2 */)
	if err != nil {
		return nil, err
	}
	mvktBakerCycleData, err := client.getCycleData(ctx, bakerAddr, cycle)
	if err != nil {
		return nil, err
	}

	collectedDelegators := make([]splitDelegator, 0)
	fetched := DELEGATOR_FETCH_LIMIT
	for fetched == DELEGATOR_FETCH_LIMIT {
		newDelegators, err := client.getDelegatorsCycleData(ctx, bakerAddr, cycle, DELEGATOR_FETCH_LIMIT, len(collectedDelegators))
		if err != nil {
			return nil, err
		}
		collectedDelegators = append(collectedDelegators, newDelegators...)
		fetched = len(newDelegators)
	}

	// handle delegator parsing errors
	defer (func() {
		panicError := recover()
		if panicError != nil {
			err = panicError.(error)
			return
		}
	})()
	slog.Debug("fetched baker data", "delegators_count", len(collectedDelegators))

	precision := int64(10000)

	var blockDelegatedRewards, endorsingDelegatedRewards, delegationShare mavryk.Z

	blockDelegatedRewards = mavryk.NewZ(mvktBakerCycleData.BlockRewardsDelegated)
	endorsingDelegatedRewards = mavryk.NewZ(mvktBakerCycleData.AttestationRewardsDelegated)
	dalDelegatedRewards := mavryk.NewZ(mvktBakerCycleData.DalRewardsDelegated)
	delegationShare = mavryk.NewZ(mvktBakerCycleData.BakingPower - mvktBakerCycleData.OwnStakedBalance - mvktBakerCycleData.ExternalStakedBalance).Mul64(precision).Div64(mvktBakerCycleData.BakingPower)

	// all block fees are distributed as liquid balance only
	blockDelegatedFees := mavryk.NewZ(mvktBakerCycleData.BlockFees)

	return &common.BakersCycleData{
		DelegatorsCount:                   mvktBakerCycleData.DelegatorsCount,
		OwnDelegatedBalance:               mavryk.NewZ(mvktBakerCycleData.OwnDelegatedBalance),
		ExternalDelegatedBalance:          mavryk.NewZ(mvktBakerCycleData.ExternalDelegatedBalance),
		BlockDelegatedRewards:             blockDelegatedRewards,
		IdealBlockDelegatedRewards:        blockDelegatedRewards.Add(delegationShare.Mul64(mvktBakerCycleData.MissedBlockRewards).Div64(precision)),
		AttestationsDelegatedRewards:      endorsingDelegatedRewards,
		IdealAttestationsDelegatedRewards: endorsingDelegatedRewards.Add(delegationShare.Mul64(mvktBakerCycleData.MissedAttestationRewards).Div64(precision)),
		DalDelegatedRewards:               dalDelegatedRewards,
		IdealDalDelegatedRewards:          dalDelegatedRewards.Add(delegationShare.Mul64(mvktBakerCycleData.MissedDalRewards).Div64(precision)),
		BlockDelegatedFees:                blockDelegatedFees,

		StakersCount:                  mvktBakerCycleData.StakersCount,
		OwnStakedBalance:              mavryk.NewZ(mvktBakerCycleData.OwnStakedBalance),
		ExternalStakedBalance:         mavryk.NewZ(mvktBakerCycleData.ExternalStakedBalance),
		BlockStakingRewardsEdge:       mavryk.NewZ(mvktBakerCycleData.BlockRewardsStakedEdge),
		AttestationStakingRewardsEdge: mavryk.NewZ(mvktBakerCycleData.AttestationRewardsStakedEdge),
		BlockStakingFees:              mavryk.Zero, // block fees are distributed as liquid balance only

		FrozenDepositLimit: mavryk.NewZ(mvktBakerData.FrozenDepositLimit),
		Delegators: lo.Map(collectedDelegators, func(delegator splitDelegator, _ int) common.Delegator {
			addr, err := mavryk.ParseAddress(delegator.Address)
			if err != nil {
				panic(err)
			}
			return common.Delegator{
				Address:          addr,
				DelegatedBalance: mavryk.NewZ(delegator.DelegatedBalance),
				StakedBalance:    mavryk.NewZ(delegator.StakedBalance),
				Emptied:          delegator.Emptied,
			}
		}),
	}, nil

}

// https://api.mavryk.network/v1/operations/transactions/onyUK7ZnQHzeNYbWSLL4zVATBtvLLk5GpPDv3VfoQPLtsBCjPX1/status
func (client *Client) WasOperationApplied(ctx context.Context, opHash mavryk.OpHash) (common.OperationStatus, error) {
	op, _ := opHash.MarshalText()

	path := fmt.Sprintf("v1/operations/transactions/%s/status", op)
	resp, err := client.Get(ctx, path)
	if err != nil {
		return common.OPERATION_STATUS_UNKNOWN, errors.Join(constants.ErrOperationStatusCheckFailed, err)
	}
	if resp.StatusCode == 204 {
		return common.OPERATION_STATUS_NOT_EXISTS, nil
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return common.OPERATION_STATUS_UNKNOWN, errors.Join(constants.ErrOperationStatusCheckFailed, err)
	}
	if bytes.Equal(body, []byte("true")) {
		return common.OPERATION_STATUS_APPLIED, nil
	}
	if bytes.Equal(body, []byte("false")) {
		return common.OPERATION_STATUS_FAILED, nil
	}
	return common.OPERATION_STATUS_NOT_EXISTS, nil
}
