package coins

import "github.com/JulianToledano/goingecko/types"

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

/*
{
description: "Another big milestone for Fuse 🎉\r\n\r\nMajor crypto exchange MEXC Global has integrated the Fuse RPC and enabled FUSE deposits and withdrawals natively on the Fuse Network blockchain 😎\r\n\r\nRead more 👉 https://bit.ly/3rOopN4",
category: "general",
created_at: "2022-02-15T11:07:39.610Z",
user: "Robert Miller",
user_title: "Marcom Director ",
pin: false,
project: {
type: "Coin",
id: "fuse-network-token",
name: "Fuse",
symbol: "fuse",
image: {
thumb: "https://assets.coingecko.com/coins/images/10347/thumb/vUXKHEe.png?1601523640",
small: "https://assets.coingecko.com/coins/images/10347/small/vUXKHEe.png?1601523640",
large: "https://assets.coingecko.com/coins/images/10347/large/vUXKHEe.png?1601523640"
      }
}
    }
*/
