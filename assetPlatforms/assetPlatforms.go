package assetPlatforms

type AssetPlatforms []Asset

type Asset struct {
	ID              string `json:"id"`
	ChainIdentifier int64  `json:"chain_identifier"`
	Name            string `json:"name"`
	ShortName       string `json:"shortName"`
	NativeCoinId    string `json:"native_coin_id"`
}
