package collector_engines

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/configuration"
	"github.com/mavryk-network/mavpay/engines/mvkt"
	"github.com/mavryk-network/mavpay/utils"
	"github.com/mavryk-network/mvgo/codec"
	"github.com/mavryk-network/mvgo/mavryk"
	"github.com/mavryk-network/mvgo/rpc"
)

type DefaultRpcAndMvktColletor struct {
	rpc  *rpc.Client
	mvkt *mvkt.Client
}

var (
	defaultCtx context.Context = context.Background()
)

func InitDefaultRpcAndMvktColletor(config *configuration.RuntimeConfiguration) (*DefaultRpcAndMvktColletor, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	rpcClient, err := rpc.NewClient(config.Network.RpcUrl, client)
	if err != nil {
		return nil, err
	}

	mvktClient, err := mvkt.InitClient(config.Network.MvktUrl, config.Network.ProtocolRewardsUrl, &mvkt.MvktClientOptions{
		HttpClient:       client,
		BalanceCheckMode: config.PayoutConfiguration.BalanceCheckMode,
	})
	if err != nil {
		return nil, err
	}

	result := &DefaultRpcAndMvktColletor{
		rpc:  rpcClient,
		mvkt: mvktClient,
	}

	return result, result.RefreshParams()
}

func (engine *DefaultRpcAndMvktColletor) GetId() string {
	return "DefaultRpcAndMvktColletor"
}

func (engine *DefaultRpcAndMvktColletor) RefreshParams() error {
	return engine.rpc.Init(context.Background())
}

func (engine *DefaultRpcAndMvktColletor) GetCurrentProtocol() (mavryk.ProtocolHash, error) {
	params, err := engine.rpc.GetParams(context.Background(), rpc.Head)

	if err != nil {
		return mavryk.ZeroProtocolHash, err
	}
	return params.Protocol, nil
}

func (engine *DefaultRpcAndMvktColletor) IsRevealed(addr mavryk.Address) (bool, error) {
	state, err := engine.rpc.GetContractExt(defaultCtx, addr, rpc.Head)
	if err != nil {
		return false, err
	}
	return state.IsRevealed(), nil
}

func (engine *DefaultRpcAndMvktColletor) GetCurrentCycleNumber() (int64, error) {
	head, err := engine.rpc.GetHeadBlock(defaultCtx)
	if err != nil {
		return 0, err
	}

	return head.GetLevelInfo().Cycle, err
}

func (engine *DefaultRpcAndMvktColletor) GetLastCompletedCycle() (int64, error) {
	cycle, err := engine.GetCurrentCycleNumber()
	return cycle - 1, err
}

func (engine *DefaultRpcAndMvktColletor) GetCycleStakingData(baker mavryk.Address, cycle int64) (*common.BakersCycleData, error) {
	return engine.mvkt.GetCycleData(context.Background(), baker, cycle)
}

func (engine *DefaultRpcAndMvktColletor) GetCyclesInDateRange(startDate time.Time, endDate time.Time) ([]int64, error) {
	return engine.mvkt.GetCyclesInDateRange(context.Background(), startDate, endDate)
}

func (engine *DefaultRpcAndMvktColletor) WasOperationApplied(op mavryk.OpHash) (common.OperationStatus, error) {
	return engine.mvkt.WasOperationApplied(context.Background(), op)
}

func (engine *DefaultRpcAndMvktColletor) GetBranch(offset int64) (hash mavryk.BlockHash, err error) {
	hash, err = engine.rpc.GetBlockHash(context.Background(), rpc.NewBlockOffset(rpc.Head, offset))
	return
}

func (engine *DefaultRpcAndMvktColletor) Simulate(o *codec.Op, publicKey mavryk.Key) (rcpt *rpc.Receipt, err error) {
	o = o.WithParams(engine.rpc.Params)
	for i := 0; i < 5; i++ {
		err = engine.rpc.Complete(context.Background(), o, publicKey)
		if err != nil {
			continue
		}

		rcpt, err = engine.rpc.Simulate(context.Background(), o, nil)
		if err != nil && rcpt == nil { // we do not retry on receipt errors
			slog.Debug("Internal simulate error - likely networking, retrying", "error", err.Error())
			// sleep 5s * i
			time.Sleep(time.Duration(i*5) * time.Second)
			continue
		}
		break
	}
	return rcpt, err
}

func (engine *DefaultRpcAndMvktColletor) GetBalance(addr mavryk.Address) (mavryk.Z, error) {
	return engine.rpc.GetContractBalance(context.Background(), addr, rpc.Head)
}

func (engine *DefaultRpcAndMvktColletor) CreateCycleMonitor(options common.CycleMonitorOptions) (common.CycleMonitor, error) {
	ctx := context.Background()
	monitor, err := common.NewCycleMonitor(ctx, engine.rpc, options)
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
