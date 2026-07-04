package types

type RecentlyAddedCoin struct {
	ID          string `json:"id"`
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	ActivatedAt int64  `json:"activated_at"`
}
