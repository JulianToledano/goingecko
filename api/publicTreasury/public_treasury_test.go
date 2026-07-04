package publicTreasury

import (
	"context"
	"net/url"
	"testing"
)

type testClient struct {
	requestURL string
	response   []byte
}

func (c *testClient) MakeReq(_ context.Context, requestURL string) ([]byte, error) {
	c.requestURL = requestURL
	return c.response, nil
}

func TestPublicTreasuryClient_EntitiesList(t *testing.T) {
	httpClient := &testClient{
		response: []byte(`[{"id":"texas","symbol":"","name":"Texas","country":"US"}]`),
	}
	c := NewClient(httpClient, "https://api.coingecko.com/api/v3")

	got, err := c.EntitiesList(
		context.Background(),
		WithEntitiesEntityType("government"),
		WithEntitiesPerPage(25),
		WithEntitiesPage(2),
	)
	if err != nil {
		t.Fatalf("EntitiesList() error = %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("EntitiesList() returned %d entities, want 1", len(got))
	}
	if got[0].ID != "texas" {
		t.Fatalf("EntitiesList()[0].ID = %q, want texas", got[0].ID)
	}

	requestURL, err := url.Parse(httpClient.requestURL)
	if err != nil {
		t.Fatalf("url.Parse() error = %v", err)
	}
	if requestURL.Path != "/api/v3/entities/list" {
		t.Fatalf("request path = %q, want /api/v3/entities/list", requestURL.Path)
	}

	query := requestURL.Query()
	if query.Get("entity_type") != "government" {
		t.Errorf("entity_type = %q, want government", query.Get("entity_type"))
	}
	if query.Get("per_page") != "25" {
		t.Errorf("per_page = %q, want 25", query.Get("per_page"))
	}
	if query.Get("page") != "2" {
		t.Errorf("page = %q, want 2", query.Get("page"))
	}
}

func TestPublicTreasuryClient_PublicTreasuryCoinIdByEntity(t *testing.T) {
	httpClient := &testClient{
		response: []byte(`{"total_holdings":1.5,"total_value_usd":2.5,"market_cap_dominance":3.5,"governments":[{"name":"Texas","symbol":"","country":"US","total_holdings":1.5,"total_entry_value_usd":2.5,"total_current_value_usd":3.5,"percentage_of_total_supply":4.5}]}`),
	}
	c := NewClient(httpClient, "https://api.coingecko.com/api/v3")

	got, err := c.PublicTreasuryCoinIdByEntity(
		context.Background(),
		"governments",
		"bitcoin",
		WithPublicTreasuryPerPage(25),
		WithPublicTreasuryPage(2),
		WithPublicTreasuryOrder("total_holdings_usd_asc"),
	)
	if err != nil {
		t.Fatalf("PublicTreasuryCoinIdByEntity() error = %v", err)
	}
	if got == nil {
		t.Fatal("PublicTreasuryCoinIdByEntity() got nil")
	}
	if len(got.Governments) != 1 {
		t.Fatalf("PublicTreasuryCoinIdByEntity() returned %d governments, want 1", len(got.Governments))
	}
	if got.Governments[0].TotalCurrentValueUsd != 3.5 {
		t.Fatalf("TotalCurrentValueUsd = %v, want 3.5", got.Governments[0].TotalCurrentValueUsd)
	}

	requestURL, err := url.Parse(httpClient.requestURL)
	if err != nil {
		t.Fatalf("url.Parse() error = %v", err)
	}
	if requestURL.Path != "/api/v3/governments/public_treasury/bitcoin" {
		t.Fatalf("request path = %q, want /api/v3/governments/public_treasury/bitcoin", requestURL.Path)
	}

	query := requestURL.Query()
	if query.Get("per_page") != "25" {
		t.Errorf("per_page = %q, want 25", query.Get("per_page"))
	}
	if query.Get("page") != "2" {
		t.Errorf("page = %q, want 2", query.Get("page"))
	}
	if query.Get("order") != "total_holdings_usd_asc" {
		t.Errorf("order = %q, want total_holdings_usd_asc", query.Get("order"))
	}
}

func TestPublicTreasuryClient_PublicTreasuryCoinIdUsesCompaniesEntity(t *testing.T) {
	httpClient := &testClient{
		response: []byte(`{"total_holdings":1,"total_value_usd":2,"market_cap_dominance":3,"companies":[]}`),
	}
	c := NewClient(httpClient, "https://api.coingecko.com/api/v3")

	_, err := c.PublicTreasuryCoinId(context.Background(), "bitcoin")
	if err != nil {
		t.Fatalf("PublicTreasuryCoinId() error = %v", err)
	}

	requestURL, err := url.Parse(httpClient.requestURL)
	if err != nil {
		t.Fatalf("url.Parse() error = %v", err)
	}
	if requestURL.Path != "/api/v3/companies/public_treasury/bitcoin" {
		t.Fatalf("request path = %q, want /api/v3/companies/public_treasury/bitcoin", requestURL.Path)
	}
}

func TestPublicTreasuryClient_PublicTreasuryEntity(t *testing.T) {
	symbol := "MSTR.US"
	httpClient := &testClient{
		response: []byte(`{"name":"Strategy","id":"strategy","type":"company","symbol":"MSTR.US","country":"US","website_url":"https://www.strategy.com/","twitter_screen_name":"Strategy","total_treasury_value_usd":64383151578.86817,"unrealized_pnl":513095879.8681717,"m_nav":1.03,"total_asset_value_per_share_usd":179.74079167746558,"holdings":[{"coin_id":"bitcoin","amount":843738,"percentage_of_total_supply":4.018,"amount_per_share":0.002355494137353434,"entity_value_usd_percentage":100,"current_value_usd":64383151578.86817,"total_entry_value_usd":63870055699,"average_entry_value_usd":75698.9203982753,"unrealized_pnl":513095879.8681717,"holding_amount_change":{"7d":0,"14d":24869},"holding_change_percentage":{"7d":0,"14d":3.037}}]}`),
	}
	c := NewClient(httpClient, "https://api.coingecko.com/api/v3")

	got, err := c.PublicTreasuryEntity(
		context.Background(),
		"strategy",
		WithHoldingAmountChange("7d,14d"),
		WithHoldingChangePercentage("7d,14d"),
	)
	if err != nil {
		t.Fatalf("PublicTreasuryEntity() error = %v", err)
	}
	if got == nil {
		t.Fatal("PublicTreasuryEntity() got nil")
	}
	if got.Symbol == nil || *got.Symbol != symbol {
		t.Fatalf("Symbol = %v, want %q", got.Symbol, symbol)
	}
	if len(got.Holdings) != 1 {
		t.Fatalf("PublicTreasuryEntity() returned %d holdings, want 1", len(got.Holdings))
	}
	if got.Holdings[0].HoldingAmountChange["14d"] != 24869 {
		t.Fatalf("HoldingAmountChange[14d] = %v, want 24869", got.Holdings[0].HoldingAmountChange["14d"])
	}

	requestURL, err := url.Parse(httpClient.requestURL)
	if err != nil {
		t.Fatalf("url.Parse() error = %v", err)
	}
	if requestURL.Path != "/api/v3/public_treasury/strategy" {
		t.Fatalf("request path = %q, want /api/v3/public_treasury/strategy", requestURL.Path)
	}

	query := requestURL.Query()
	if query.Get("holding_amount_change") != "7d,14d" {
		t.Errorf("holding_amount_change = %q, want 7d,14d", query.Get("holding_amount_change"))
	}
	if query.Get("holding_change_percentage") != "7d,14d" {
		t.Errorf("holding_change_percentage = %q, want 7d,14d", query.Get("holding_change_percentage"))
	}
}

func TestPublicTreasuryClient_PublicTreasuryHoldingChart(t *testing.T) {
	httpClient := &testClient{
		response: []byte(`{"holdings":[[1748736000000,580955],[1749340800000,582000]],"holding_value_in_usd":[[1748736000000,60818730878.617355],[1749340800000,61506606585.45032]]}`),
	}
	c := NewClient(httpClient, "https://api.coingecko.com/api/v3")

	got, err := c.PublicTreasuryHoldingChart(
		context.Background(),
		"strategy",
		"bitcoin",
		WithHoldingChartDays("365"),
		WithHoldingChartIncludeEmptyIntervals(true),
	)
	if err != nil {
		t.Fatalf("PublicTreasuryHoldingChart() error = %v", err)
	}
	if got == nil {
		t.Fatal("PublicTreasuryHoldingChart() got nil")
	}
	if len(got.Holdings) != 2 {
		t.Fatalf("PublicTreasuryHoldingChart() returned %d holdings, want 2", len(got.Holdings))
	}
	if got.HoldingValueInUsd[1][1] != 61506606585.45032 {
		t.Fatalf("HoldingValueInUsd[1][1] = %v, want 61506606585.45032", got.HoldingValueInUsd[1][1])
	}

	requestURL, err := url.Parse(httpClient.requestURL)
	if err != nil {
		t.Fatalf("url.Parse() error = %v", err)
	}
	if requestURL.Path != "/api/v3/public_treasury/strategy/bitcoin/holding_chart" {
		t.Fatalf("request path = %q, want /api/v3/public_treasury/strategy/bitcoin/holding_chart", requestURL.Path)
	}

	query := requestURL.Query()
	if query.Get("days") != "365" {
		t.Errorf("days = %q, want 365", query.Get("days"))
	}
	if query.Get("include_empty_intervals") != "true" {
		t.Errorf("include_empty_intervals = %q, want true", query.Get("include_empty_intervals"))
	}
}

func TestPublicTreasuryClient_PublicTreasuryTransactionHistory(t *testing.T) {
	httpClient := &testClient{
		response: []byte(`{"transactions":[{"date":1779062400000,"source_url":"https://example.com/form-8-k.pdf","coin_id":"bitcoin","type":"buy","holding_net_change":24869,"transaction_value_usd":2014015965,"holding_balance":843738,"average_entry_value_usd":80985}]}`),
	}
	c := NewClient(httpClient, "https://api.coingecko.com/api/v3")

	got, err := c.PublicTreasuryTransactionHistory(
		context.Background(),
		"strategy",
		WithTransactionHistoryPerPage(25),
		WithTransactionHistoryPage(2),
		WithTransactionHistoryOrder("transaction_value_usd_desc"),
		WithTransactionHistoryCoinIDs("bitcoin,ethereum"),
	)
	if err != nil {
		t.Fatalf("PublicTreasuryTransactionHistory() error = %v", err)
	}
	if got == nil {
		t.Fatal("PublicTreasuryTransactionHistory() got nil")
	}
	if len(got.Transactions) != 1 {
		t.Fatalf("PublicTreasuryTransactionHistory() returned %d transactions, want 1", len(got.Transactions))
	}
	if got.Transactions[0].TransactionValueUsd != 2014015965 {
		t.Fatalf("TransactionValueUsd = %v, want 2014015965", got.Transactions[0].TransactionValueUsd)
	}

	requestURL, err := url.Parse(httpClient.requestURL)
	if err != nil {
		t.Fatalf("url.Parse() error = %v", err)
	}
	if requestURL.Path != "/api/v3/public_treasury/strategy/transaction_history" {
		t.Fatalf("request path = %q, want /api/v3/public_treasury/strategy/transaction_history", requestURL.Path)
	}

	query := requestURL.Query()
	if query.Get("per_page") != "25" {
		t.Errorf("per_page = %q, want 25", query.Get("per_page"))
	}
	if query.Get("page") != "2" {
		t.Errorf("page = %q, want 2", query.Get("page"))
	}
	if query.Get("order") != "transaction_value_usd_desc" {
		t.Errorf("order = %q, want transaction_value_usd_desc", query.Get("order"))
	}
	if query.Get("coin_ids") != "bitcoin,ethereum" {
		t.Errorf("coin_ids = %q, want bitcoin,ethereum", query.Get("coin_ids"))
	}
}
