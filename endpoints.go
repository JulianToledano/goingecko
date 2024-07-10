package goingecko

import "fmt"

var (
	version    = "v3"
	BaseURL    = fmt.Sprintf("https://api.coingecko.com/api/%s", version)
	ProBaseURL = fmt.Sprintf("https://pro-api.coingecko.com/api/%s", version)
)

func (c *Client) getPingURL() string {
	return fmt.Sprintf("%s/ping", c.baseUrl)
}

func (c *Client) getSimpleURL() string {
	return fmt.Sprintf("%s/simple", c.baseUrl)
}

func (c *Client) getCoinsURL() string {
	return fmt.Sprintf("%s/coins", c.baseUrl)
}

func (c *Client) getAssetPlatformsURL() string {
	return fmt.Sprintf("%s/asset_platforms", c.baseUrl)
}

func (c *Client) getCategoriesURL() string {
	return fmt.Sprintf("%s/categories", c.getCoinsURL())
}

func (c *Client) getExchangesURL() string {
	return fmt.Sprintf("%s/exchanges", c.baseUrl)
}

func (c *Client) getDerivativesURL() string {
	return fmt.Sprintf("%s/derivatives", c.baseUrl)
}

func (c *Client) getNftsURL() string {
	return fmt.Sprintf("%s/nfts", c.baseUrl)
}

func (c *Client) getSearchURL() string {
	return fmt.Sprintf("%s/search", c.baseUrl)
}

func (c *Client) getContractURL() string {
	return fmt.Sprintf("%s/coins", c.baseUrl)
}

func (c *Client) getExchangeRatesURL() string {
	return fmt.Sprintf("%s/exchange_rates", c.baseUrl)
}

func (c *Client) getTrendingURL() string {
	return fmt.Sprintf("%s/search/trending", c.baseUrl)
}

func (c *Client) getGlobalURL() string {
	return fmt.Sprintf("%s/global", c.baseUrl)
}

func (c *Client) getCompaniesURL() string {
	return fmt.Sprintf("%s/companies", c.baseUrl)
}
