package types

type SupplyBreakdown struct {
	ID                    string                        `json:"id"`
	Symbol                string                        `json:"symbol"`
	Name                  string                        `json:"name"`
	SupplyData            SupplyData                    `json:"supply_data"`
	NonCirculatingWallets []*NonCirculatingSupplyWallet `json:"non_circulating_wallets"`
}

type SupplyData struct {
	TotalSupply          float64 `json:"total_supply"`
	CirculatingSupply    float64 `json:"circulating_supply"`
	OutstandingSupply    float64 `json:"outstanding_supply"`
	NonCirculatingSupply float64 `json:"non_circulating_supply"`
	LastUpdated          string  `json:"last_updated"`
}

type NonCirculatingSupplyWallet struct {
	Address                 string  `json:"address"`
	Label                   string  `json:"label"`
	Balance                 float64 `json:"balance"`
	PercentageOfTotalSupply float64 `json:"percentage_of_total_supply"`
	Anomaly                 bool    `json:"anomaly"`
	LastUpdated             string  `json:"last_updated"`
}
