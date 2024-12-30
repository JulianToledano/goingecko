package types

type Nft struct {
	Id              string `json:"id"`
	ContractAddress string `json:"contract_address"`
	Name            string `json:"name"`
	AssetPlatformId string `json:"asset_platform_id"`
	Symbol          string `json:"symbol"`
}

type NftId struct {
	Id                                         string      `json:"id"`
	ContractAddress                            string      `json:"contract_address"`
	AssetPlatformId                            string      `json:"asset_platform_id"`
	Name                                       string      `json:"name"`
	Symbol                                     string      `json:"symbol"`
	Image                                      Image       `json:"image"`
	Description                                string      `json:"description"`
	NativeCurrency                             string      `json:"native_currency"`
	NativeCurrencySymbol                       string      `json:"native_currency_symbol"`
	FloorPrice                                 Price       `json:"floor_price"`
	MarketCap                                  Price       `json:"market_cap"`
	Volume24H                                  Price       `json:"volume_24h"`
	FloorPriceInUsd24HPercentageChange         float64     `json:"floor_price_in_usd_24h_percentage_change"`
	FloorPrice24HPercentageChange              Price       `json:"floor_price_24h_percentage_change"`
	MarketCap24HPercentageChange               Price       `json:"market_cap_24h_percentage_change"`
	Volume24HPercentageChange                  Price       `json:"volume_24h_percentage_change"`
	NumberOfUniqueAddresses                    float64     `json:"number_of_unique_addresses"`
	NumberOfUniqueAddresses24HPercentageChange float64     `json:"number_of_unique_addresses_24h_percentage_change"`
	VolumeInUsd24HPercentageChange             float64     `json:"volume_in_usd_24h_percentage_change"`
	TotalSupply                                float64     `json:"total_supply"`
	OneDaySales                                float64     `json:"one_day_sales"`
	OneDaySales24HPercentageChange             float64     `json:"one_day_sales_24h_percentage_change"`
	OneDayAverageSalePrice                     float64     `json:"one_day_average_sale_price"`
	OneDayAverageSalePrice24HPercentageChange  float64     `json:"one_day_average_sale_price_24h_percentage_change"`
	Links                                      Links       `json:"links"`
	FloorPrice7DPercentageChange               Price       `json:"floor_price_7d_percentage_change"`
	FloorPrice14DPercentageChange              Price       `json:"floor_price_14d_percentage_change"`
	FloorPrice30DPercentageChange              Price       `json:"floor_price_30d_percentage_change"`
	FloorPrice60DPercentageChange              Price       `json:"floor_price_60d_percentage_change"`
	FloorPrice1YPercentageChange               Price       `json:"floor_price_1y_percentage_change"`
	Explorers                                  []Explorers `json:"explorers"`
}

type Price struct {
	Usd            float64 `json:"usd"`
	NativeCurrency float64 `json:"native_currency"`
}

type Image struct {
	Small string `json:"small"`
}

type Links struct {
	Homepage string `json:"homepage"`
	Twitter  string `json:"twitter"`
	Discord  string `json:"discord"`
}

type Explorers struct {
	Name string `json:"name"`
	Link string `json:"link"`
}
