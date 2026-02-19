package mock

import (
	"time"

	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/gomavryk/codec"
	"github.com/mavryk-network/gomavryk/rpc"
	"github.com/mavryk-network/gomavryk/mavryk"
)

type EmptyCollector struct {
}

func (engine *EmptyCollector) GetId() string {
	panic("not implemented")
}

func (engine *EmptyCollector) RefreshParams() error {
	panic("not implemented")
}

func (engine *EmptyCollector) IsRevealed(address mavryk.Address) (bool, error) {
	panic("not implemented")
}

func (engine *EmptyCollector) GetCurrentCycleNumber() (int64, error) {
	panic("not implemented")
}

func (engine *EmptyCollector) GetLastCompletedCycle() (int64, error) {
	panic("not implemented")
}

func (engine *EmptyCollector) GetCycleStakingData(baker mavryk.Address, cycle int64) (*common.BakersCycleData, error) {
	panic("not implemented")
}

func (engine *EmptyCollector) GetCyclesInDateRange(startDate time.Time, endDate time.Time) ([]int64, error) {
	panic("not implemented")
}

func (engine *EmptyCollector) WasOperationApplied(op mavryk.OpHash) (common.OperationStatus, error) {
	panic("not implemented")
}

func (engine *EmptyCollector) CreateCycleMonitor(options common.CycleMonitorOptions) (common.CycleMonitor, error) {
	panic("not implemented")
}

func (engine *EmptyCollector) GetBranch(offset int64) (hash mavryk.BlockHash, err error) {
	panic("not implemented")
}

func (engine *EmptyCollector) GetExpectedTxCosts() int64 {
	panic("not implemented")
}

func (engine *EmptyCollector) Simulate(o *codec.Op, publicKey mavryk.Key) (*rpc.Receipt, error) {
	panic("not implemented")
}

func (engine *EmptyCollector) GetBalance(addr mavryk.Address) (mavryk.Z, error) {
	panic("not implemented")
}

func (engine *EmptyCollector) SendAnalytics(bakerId string, version string) {}

func (engine *EmptyCollector) GetCurrentProtocol() (mavryk.ProtocolHash, error) {
	panic("not implemented")
}
