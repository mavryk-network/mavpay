package constants

const (
	MAVPAY_REPOSITORY = "mavryk-network/mavpay"

	MUMAV_FACTOR = 1000000

	DELEGATION_CAPACITY_FACTOR             = 9
	LIMIT_OF_STAKING_OVER_BAKING_PRECISION = 1_000_000

	DEFAULT_BAKER_FEE                     = float64(.05)
	DEFAULT_DELEGATOR_MINIMUM_BALANCE     = float64(0)
	DEFAULT_PAYOUT_MINIMUM_AMOUNT         = float64(0)
	DEFAULT_MVKT_URL                      = "https://api.mavryk.network/"
	DEFAULT_EXPLORER_URL                  = "https://nexus.mavryk.org/"
	DEFAULT_REQUIRED_CONFIRMATIONS        = int64(2)
	DEFAULT_TX_GAS_LIMIT_BUFFER           = int64(100)
	DEFAULT_KT_TX_GAS_LIMIT_BUFFER        = int64(200)
	DEFAULT_TX_DESERIALIZATION_GAS_BUFFER = int64(2) // just because of integer division
	DEFAULT_TX_FEE_BUFFER                 = int64(0)
	DEFAULT_KT_TX_FEE_BUFFER              = int64(0)
	DEFAULT_SIMULATION_TX_BATCH_SIZE      = 50

	// buffer for signature, branch etc.
	DEFAULT_BATCHING_OPERATION_DATA_BUFFER = 3000

	PAYOUT_FEE_BUFFER  = 1000 // buffer per payout to check baker balance is sufficient
	MAX_OPERATION_TTL  = 12   // 12 blocks
	ALLOCATION_STORAGE = 257

	DEFAULT_CYCLE_MONITOR_MAXIMUM_DELAY = int64(1500)
	DEFAULT_CYCLE_MONITOR_MINIMUM_DELAY = int64(500)

	CONFIG_FILE_BACKUP_SUFFIX = ".backup"
	PAYOUT_REPORT_FILE_NAME   = "payouts.csv"
	INVALID_REPORT_FILE_NAME  = "invalid.csv"
	REPORT_SUMMARY_FILE_NAME  = "summary.json"
	REPORTS_DIRECTORY         = "reports"

	DEFAULT_DONATION_ADDRESS    = "mv1V4h45W3p4e1sjSBvRkK2uYbvkTnSuHg8g"
	DEFAULT_DONATION_PERCENTAGE = 0.05

	// periodic payouts
	MINIMUM_PAYOUT_INTERVAL_CYCLES = int64(1)
	MAXIMUM_PAYOUT_INTERVAL_CYCLES = int64(365)

	DRY_RUN_NOTE = "[DRY RUN]"
)

var (
	DEFAULT_RPC_POOL = []string{
		"https://rpc.mavryk.network/",
	}
)
