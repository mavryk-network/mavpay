package collector_engines

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/configuration"
	"github.com/mavryk-network/mavpay/engines/mvkt"
	"github.com/mavryk-network/mavpay/utils"
	"github.com/mavryk-network/gomavryk/codec"
	"github.com/mavryk-network/gomavryk/mavryk"
	"github.com/mavryk-network/gomavryk/rpc"
)

type DefaultRpcAndMvktColletor struct {
	rpcs []*rpc.Client
	mvkt *mvkt.Client
}

var (
	defaultCtx context.Context = context.Background()
)

func InitDefaultRpcAndMvktColletor(config *configuration.RuntimeConfiguration) (*DefaultRpcAndMvktColletor, error) {
	http_client := &http.Client{
		Timeout: 10 * time.Second,
	}

	rpc_clients, err := utils.InitializeRpcClients(context.Background(), config.Network.RpcPool, http_client)
	if err != nil {
		return nil, err
	}

	mvkt_client, err := mvkt.InitClient(config.Network.MvktUrl, &mvkt.MvktClientOptions{
		HttpClient: http_client,
	})
	if err != nil {
		return nil, err
	}

	result := &DefaultRpcAndMvktColletor{
		rpcs: rpc_clients,
		mvkt: mvkt_client,
	}

	return result, result.RefreshParams()
}

func (engine *DefaultRpcAndMvktColletor) GetId() string {
	return "DefaultRpcAndMvktColletor"
}

func (engine *DefaultRpcAndMvktColletor) RefreshParams() error {
	failures := 0
	for _, rpc := range engine.rpcs {
		err := rpc.Init(context.Background())
		if err != nil {
			slog.Debug("failed to refresh rpc params", "error", err.Error(), "rpc_url", rpc.BaseURL.String())
			failures++
		}
	}
	if failures == len(engine.rpcs) {
		return fmt.Errorf("failed to refresh rpc params for all clients, all %d failed", failures)
	}
	return nil
}

func (engine *DefaultRpcAndMvktColletor) GetCurrentProtocol() (mavryk.ProtocolHash, error) {
	params, err := utils.AttemptWithRpcClients(defaultCtx, engine.rpcs, func(client *rpc.Client) (*mavryk.Params, error) {
		return client.GetParams(context.Background(), rpc.Head)
	})
	if err != nil {
		return mavryk.ZeroProtocolHash, err
	}
	return params.Protocol, nil
}

func (engine *DefaultRpcAndMvktColletor) IsRevealed(addr mavryk.Address) (bool, error) {
	state, err := utils.AttemptWithRpcClients(defaultCtx, engine.rpcs, func(client *rpc.Client) (*rpc.ContractInfo, error) {
		return client.GetContractExt(defaultCtx, addr, rpc.Head)
	})
	if err != nil {
		return false, err
	}
	return state.IsRevealed(), nil
}

func (engine *DefaultRpcAndMvktColletor) GetCurrentCycleNumber() (int64, error) {
	head, err := utils.AttemptWithRpcClients(defaultCtx, engine.rpcs, func(client *rpc.Client) (*rpc.BlockMetadata, error) {
		return client.GetBlockMetadata(defaultCtx, rpc.Head)
	})
	if err != nil {
		return 0, err
	}

	return head.LevelInfo.Cycle, nil
}

func (engine *DefaultRpcAndMvktColletor) GetLastCompletedCycle() (int64, error) {
	cycle, err := engine.GetCurrentCycleNumber()
	return cycle - 1, err
}

func (engine *DefaultRpcAndMvktColletor) GetChainId() (mavryk.ChainIdHash, error) {
	chainId, err := utils.AttemptWithRpcClients(defaultCtx, engine.rpcs, func(client *rpc.Client) (mavryk.ChainIdHash, error) {
		return client.GetChainId(defaultCtx)
	})
	return chainId, err
}

func (engine *DefaultRpcAndMvktColletor) GetCycleStakingData(baker mavryk.Address, cycle int64) (*common.BakersCycleData, error) {
	chainId, err := engine.GetChainId()
	if err != nil {
		return nil, err
	}

	return engine.mvkt.GetCycleData(context.Background(), chainId, baker, cycle)
}

func (engine *DefaultRpcAndMvktColletor) GetCyclesInDateRange(startDate time.Time, endDate time.Time) ([]int64, error) {
	return engine.mvkt.GetCyclesInDateRange(context.Background(), startDate, endDate)
}

func (engine *DefaultRpcAndMvktColletor) WasOperationApplied(op mavryk.OpHash) (common.OperationStatus, error) {
	return engine.mvkt.WasOperationApplied(context.Background(), op)
}

func (engine *DefaultRpcAndMvktColletor) GetBranch(offset int64) (hash mavryk.BlockHash, err error) {
	hash, err = utils.AttemptWithRpcClients(defaultCtx, engine.rpcs, func(client *rpc.Client) (mavryk.BlockHash, error) {
		return client.GetBlockHash(context.Background(), rpc.NewBlockOffset(rpc.Head, offset))
	})
	return
}

func (engine *DefaultRpcAndMvktColletor) Simulate(o *codec.Op, publicKey mavryk.Key) (rcpt *rpc.Receipt, err error) {
	params, err := utils.AttemptWithRpcClients(defaultCtx, engine.rpcs, func(client *rpc.Client) (*mavryk.Params, error) {
		return client.GetParams(context.Background(), rpc.Head)
	})

	if err != nil {
		return nil, err
	}

	o = o.WithParams(params)
	for i := 0; i < 5; i++ {
		_, err = utils.AttemptWithRpcClients(defaultCtx, engine.rpcs, func(client *rpc.Client) (bool, error) {
			err := client.Complete(context.Background(), o, publicKey)
			if err != nil {
				return false, err
			}

			rcpt, err = client.Simulate(context.Background(), o, nil)
			if err != nil && rcpt == nil { // we do not retry on receipt errors
				slog.Debug("Internal simulate error - likely networking, retrying", "error", err.Error())
				// sleep 5s * i
				time.Sleep(time.Duration(i*5) * time.Second)
				return false, err
			}
			return true, nil
		})
		if err == nil {
			break
		}
	}
	return rcpt, err
}

func (engine *DefaultRpcAndMvktColletor) GetBalance(addr mavryk.Address) (mavryk.Z, error) {
	return utils.AttemptWithRpcClients(defaultCtx, engine.rpcs, func(client *rpc.Client) (mavryk.Z, error) {
		return client.GetContractBalance(context.Background(), addr, rpc.Head)
	})
}

func (engine *DefaultRpcAndMvktColletor) CreateCycleMonitor(options common.CycleMonitorOptions) (common.CycleMonitor, error) {
	ctx := context.Background()
	monitor, err := utils.AttemptWithRpcClients(ctx, engine.rpcs, func(client *rpc.Client) (common.CycleMonitor, error) {
		return common.NewCycleMonitor(ctx, client, options)
	})
	if err != nil {
		return nil, err
	}
	utils.CallbackOnInterrupt(ctx, monitor.Cancel)
	slog.Info("tracking cycles... (cancel with Ctrl-C/SIGINT)\n\n")
	return monitor, nil
}

func (engine *DefaultRpcAndMvktColletor) SendAnalytics(bakerId string, version string) {
	go func() {
		// body := fmt.Sprintf(`{"bakerId": "%s", "version": "%s"}`, bakerId, version)
		// resp, err := http.Post("https://analytics.tez.capital/pay", "application/json", strings.NewReader(body))
		// if err != nil {
		// 	return
		// }
		// defer resp.Body.Close()
	}()
}