package types

import "github.com/JulianToledano/goingecko/v3/api/types"

type StatusUpdates struct {
	StatusUpdates []Status `json:"status_updates"`
}

type Status struct {
	Description string  `json:"description"`
	Category    string  `json:"category"`
	CreatedAt   string  `json:"created_at"`
	User        string  `json:"user"`
	UserTitle   string  `json:"user_title"`
	Pin         bool    `json:"pin"`
	Project     Project `json:"project"`
}

type Project struct {
	Type   string      `json:"type"`
	ID     string      `json:"id"`
	Name   string      `json:"name"`
	Symbol string      `json:"symbol"`
	Image  types.Image `json:"image"`
}
