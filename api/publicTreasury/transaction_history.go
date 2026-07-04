package publicTreasury

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/JulianToledano/goingecko/v3/api/internal"
	"github.com/JulianToledano/goingecko/v3/api/publicTreasury/types"
)

// transactionHistoryOption is an interface that extends internal.Option to provide options specific to the public treasury transaction history endpoint.
type transactionHistoryOption interface {
	internal.Option

	IsTransactionHistoryOption()
}

// WithTransactionHistoryPerPage returns a transactionHistoryOption that sets the number of transactions returned per page.
func WithTransactionHistoryPerPage(perPage int) transactionHistoryOption {
	return transactionHistoryPerPageOption{perPage}
}

// WithTransactionHistoryPage returns a transactionHistoryOption that sets the page of transactions to return.
func WithTransactionHistoryPage(page int) transactionHistoryOption {
	return transactionHistoryPageOption{page}
}

// WithTransactionHistoryOrder returns a transactionHistoryOption that sets the sort order for transactions.
func WithTransactionHistoryOrder(order string) transactionHistoryOption {
	return transactionHistoryOrderOption{order}
}

// WithTransactionHistoryCoinIDs returns a transactionHistoryOption that filters transactions by comma-separated coin IDs.
func WithTransactionHistoryCoinIDs(coinIDs string) transactionHistoryOption {
	return transactionHistoryCoinIDsOption{coinIDs}
}

// PublicTreasuryTransactionHistory allows you to query public companies' or governments' cryptocurrency transaction history by entity ID.
func (c *PublicTreasuryClient) PublicTreasuryTransactionHistory(ctx context.Context, entityID string, options ...transactionHistoryOption) (*types.TransactionHistory, error) {
	params := url.Values{}
	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/public_treasury/%s/transaction_history?%s", c.URL, entityID, params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data *types.TransactionHistory
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type transactionHistoryPerPageOption struct{ perPage int }

func (o transactionHistoryPerPageOption) Apply(v *url.Values) {
	v.Set("per_page", strconv.Itoa(o.perPage))
}

func (o transactionHistoryPerPageOption) IsTransactionHistoryOption() {}

type transactionHistoryPageOption struct{ page int }

func (o transactionHistoryPageOption) Apply(v *url.Values) {
	v.Set("page", strconv.Itoa(o.page))
}

func (o transactionHistoryPageOption) IsTransactionHistoryOption() {}

type transactionHistoryOrderOption struct{ order string }

func (o transactionHistoryOrderOption) Apply(v *url.Values) {
	v.Set("order", o.order)
}

func (o transactionHistoryOrderOption) IsTransactionHistoryOption() {}

type transactionHistoryCoinIDsOption struct{ coinIDs string }

func (o transactionHistoryCoinIDsOption) Apply(v *url.Values) {
	v.Set("coin_ids", o.coinIDs)
}

func (o transactionHistoryCoinIDsOption) IsTransactionHistoryOption() {}
