package asocks

type BalanceModel struct {
	Success                bool    `json:"success"`
	Balance                float64 `json:"balance"`
	BalanceTraffic         float64 `json:"balance_traffic"`
	AllAvailableTraffic    float64 `json:"all_available_traffic"`
	PreparedTrafficBalance float64 `json:"prepared_traffic_balance"`
	BalanceHold            float64 `json:"balance_hold"`
}
