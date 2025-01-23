# Goingecko

Coingecko API client for golang.

<p align="center">
    <img src="docs/images/goin.png" alt="goingecko" height="200" />
</p> 



## Endpoints

### Simple Endpoints
| Endpoint                                                   |  Status | Function                  |
|------------------------------------------------------------|--|---------------------------|
| /ping                                                      | ✓ | Ping                      |
| /simple/price                                              | ✓ | SimplePrice               |
| /simple/token_price/{id}                                   | ✓ | SimpleTokenPrice          |
| /simple/supported_vs_currencies                            | ✓ | SimpleSupportedVsCurrency |

### Coins Endpoints
| Endpoint                                                   |  Status | Function                  |
|------------------------------------------------------------|--|---------------------------|
| /coins/list                                                | ✓ | CoinsList                 |
| /coins/markets                                             | ✓ | CoinsMarket               |
| /coins/{id}                                                | ✓ | CoinsId                   |
| /coins/{id}/tickers                                        | ✓ | CoinsIdTickers            |
| /coins/{id}/history                                        | ✓ | CoinsIdHistory            |
| /coins/{id}/market_chart                                   | ✓ | CoinsIdMarketChart        |
| /coins/{id}/market_chart/range                             | ✓ | CoinsIdMarketChartRange   |
| /coins/{id}/ohlc                                           | ✓ | CoinsOhlc                 |
| /coins/{id}/contract/{contract_address}                    | ✓ | ContractInfo              |
| /coins/{id}/contract/{contract_address}/market_chart/      | ✓ | ContractMarketChart       |
| /coins/{id}/contract/{contract_address}/market_chart/range | ✓ | ContractMarketChartRange  |
| /coins/categories/list                                     | ✓ | CategoriesList            |
| /coins/categories/                                         | ✓ | Categories                |

### Exchange Endpoints
| Endpoint                                                   |  Status | Function                  |
|------------------------------------------------------------|--|---------------------------|
| /exchanges                                                 | ✓ | Exchanges                 |
| /exchanges/list                                            | ✓ | ExchangesList             |
| /exchanges/{id}                                            | ✓ | ExchangesId               |
| /exchanges/{id}/tickers                                    | ✓ | ExchangesIdTickers        |
| /exchanges/{id}/volume_chart                               | ✓ | ExchangesIdVolumeChart    |

### Derivatives Endpoints
| Endpoint                                                   |  Status | Function                  |
|------------------------------------------------------------|--|---------------------------|
| /derivaties                                                | ✓ | Derivatives               |
| /derivaties/exchanges                                      | ✓ | DerivativesExchanges      |
| /derivaties/exchanges/{id}                                 | ✓ | DerivativesExchangesId    |
| /derivaties/exchanges/list                                 | ✓ | DerivativesExchangesList  |

### NFT Endpoints
| Endpoint                                                   |  Status | Function                  |
|------------------------------------------------------------|--|---------------------------|
| /nfts/list                                                 | ✓ | NftsList                  |
| /nfts/{id}                                                 | ✓ | NftsId                    |
| /nfts/{asset_platform_id}/contract/{contract_address}      | ✓ | NftsContract              |

### Other Endpoints
| Endpoint                                                   |  Status | Function                  |
|------------------------------------------------------------|--|---------------------------|
| /asset_platforms                                           | ✓ | AssetPlatforms            |
| /exchange_rates                                            | ✓ | ExchangeRates             |
| /search                                                    | ✓ | Search                    |
| /search/trending                                           | ✓ | Trending                  |
| /global                                                    | ✓ | Global                    |
| /global/decentralized_finance_defi                         | ✓ | DecentrilizedFinanceDEFI  |
| /companies/public_treasury/{coin_id}                       | ✓ | PublicTreasuryCoinId      |

## Usage

```golang
package main

import (
	"context"
	"fmt"

	"github.com/JulianToledano/goingecko/v3/api"
	"github.com/JulianToledano/goingecko/v3/api/coins"
)

func main() {
	cgClient := api.NewDefaultClient()

	data, err := cgClient.CoinsId(context.Background(), "bitcoin", coins.WithTickers(false))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Bitcoin price is: %f$", data.MarketData.CurrentPrice.Usd)
}
```

Check dir [examples](docs/examples) for more.

## Todo

 - [ ] Implement On Chain Dex Api

## Thanks

Image was created with [Gophers](https://github.com/egonelbre/gophers)