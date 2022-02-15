# Goingecko

Coingecko API client for golang.

<p align="center">
    <img src="images/goin.png" alt="goingecko" height="200" />
</p> 



## Endpoints
|  Endpoint |  Status | Function  |
|---|---|---|
|  ping | ✓  | Ping  |
| simple/price  |  ✓ | SimplePrice  |
| simple/token_price/{id}  | ✓  | SimpleTokenPrice  |
| simple/supported_vs_currencies | ✓  | SimpleSupportedVsCurrency  |
| coins/list | ✓  |  CoinsList |
| coins/markets | ✓  | CoinsMarket  |
| coins/{id} | ✓  |  CoinsId |
| coins/{id}/tickers | ✓  |  CoinsIdTickers |
| coins/{id}/history | ✓  | CoinsIdHistory  |
|  coins/{id}/market_chart | ✓  | CoinsIdMarketChart  |
|  coins/{id}/market_chart/range | ✓  | CoinsIdMarketChartRange  |
|  coins/{id}/status_updates | ✓  |   |
|  coins/{id}/ohlc | ✓  | CoinsOhlc  |
|  /coins/{id}/contract/{contract_address} | ✓  |  ContractInfo |
|  /coins/{id}/contract/{contract_address}/market_chart/ | ✓  | ContractMarketChart  |
|  /coins/{id}/contract/{contract_address}/market_chart/range | ✓  | ContractMarketChartRange  |
| /events | ✓  |  Events |
| /events/countries | ✓  |  EventsCountries |
| /events/types | ✓  | EventsTypes  |
| /exchange_rates  | ✓  | ExchangeRates  |
| /search/trending  | ✓  |  Trending |
| /global  | ✓  |  Global |
| /global/decentralized_finance_defi  | ✓  |  DecentrilizedFinanceDEFI |

### Beta endpoints
Beta endpoins are to be implemented in a matter of time :smile:

## Usage

```golang
package main

import (
	"fmt"

	"github.com/JulianToledano/goingecko"
)

func main() {
	cgClient := goingecko.NewClient(nil)
	defer cgClient.Close()

	data, err := cgClient.CoinsId("bitcoin", true, true, true, false, false, false)
	if err != nil {
		fmt.Print("Somethig went wrong...")
		return
	}
	fmt.Printf("Bitcoin price is: %f$", data.MarketData.CurrentPrice.Usd)
}

```
Check dir [examples](examples) for more.

## Thanks
This repo is based some how in [superoo7/go-gecko](https://github.com/superoo7/go-gecko) work.

Image was created with [Gophers](https://github.com/egonelbre/gophers)