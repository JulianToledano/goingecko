package types

type SupplyChart struct {
	CirculatingSupply []SupplyChartPoint `json:"circulating_supply,omitempty"`
	TotalSupply       []SupplyChartPoint `json:"total_supply,omitempty"`
}

type SupplyChartPoint []interface{}
