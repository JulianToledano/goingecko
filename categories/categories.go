package categories

type CategoriesList []Category

type Category struct {
	CategoryId string `json:"category_id"`
	Name       string `json:"name"`
}

type CategoriesWithMarketDataList []CategoryWithMarketData

type CategoryWithMarketData struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	MarketCap          float64  `json:"market_cap"`
	MarketCapChange24h float64  `json:"market_cap_change_24h"`
	Content            string   `json:"content"`
	Top3Coins          []string `json:"top_3_coins"`
	Volume24h          float64  `json:"volume_24h"`
	UpdatedAt          string   `json:"updated_at"`
}
