package websocket

import (
	"bytes"
	"encoding/json"
	"strconv"
)

// Message is a typed websocket payload returned by Read or ParseMessage.
type Message interface {
	isMessage()
}

// EventMessage represents websocket lifecycle events such as welcome or confirm_subscription.
type EventMessage struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier,omitempty"`
}

// StatusMessage represents a command acknowledgement or server-side error payload.
type StatusMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// UnknownMessage captures payloads the client does not recognize yet.
type UnknownMessage struct {
	Raw json.RawMessage
}

// Number decodes CoinGecko websocket numbers that may be emitted as JSON numbers or quoted strings.
type Number float64

// Float64 returns the underlying float64 value.
func (n Number) Float64() float64 {
	return float64(n)
}

// CGSimplePriceUpdate represents a CGSimplePrice update payload.
type CGSimplePriceUpdate struct {
	ChannelType              string
	CoinID                   string
	VSCurrency               string
	MarketCapUSD             *Number
	Price                    *Number
	Price24hChangePercentage *Number
	LastUpdatedAt            *Number
	Volume24h                *Number
}

// OnchainSimpleTokenPriceUpdate represents an OnchainSimpleTokenPrice update payload.
type OnchainSimpleTokenPriceUpdate struct {
	ChannelType              string
	NetworkID                string
	TokenAddress             string
	PriceUSD                 *Number
	Price24hChangePercentage *Number
	MarketCapUSD             *Number
	Volume24h                *Number
	LastUpdatedAt            *Number
}

// OnchainTradeUpdate represents an OnchainTrade update payload.
type OnchainTradeUpdate struct {
	ChannelType           string
	NetworkID             string
	PoolAddress           string
	TransactionHash       string
	TradeType             string
	TokenAmount           *Number
	QuoteTokenAmount      *Number
	VolumeUSD             *Number
	PriceInNativeCurrency *Number
	PriceUSD              *Number
	LastUpdatedAt         *Number
}

// OnchainOHLCVUpdate represents an OnchainOHLCV update payload.
type OnchainOHLCVUpdate struct {
	ChannelType string
	NetworkID   string
	PoolAddress string
	Token       OHLCVToken
	Interval    OHLCVInterval
	Open        *Number
	High        *Number
	Low         *Number
	Close       *Number
	Volume      *Number
	Timestamp   *Number
}

func (EventMessage) isMessage()                  {}
func (StatusMessage) isMessage()                 {}
func (UnknownMessage) isMessage()                {}
func (CGSimplePriceUpdate) isMessage()           {}
func (OnchainSimpleTokenPriceUpdate) isMessage() {}
func (OnchainTradeUpdate) isMessage()            {}
func (OnchainOHLCVUpdate) isMessage()            {}

// ParseMessage decodes a raw websocket payload into a typed message.
func ParseMessage(data []byte) (Message, error) {
	var probe struct {
		Type string `json:"type"`
		Code *int   `json:"code"`
		I    string `json:"i"`
		VS   string `json:"vs"`
		TA   string `json:"ta"`
		TX   string `json:"tx"`
		PA   string `json:"pa"`
	}

	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, err
	}

	switch {
	case probe.Type != "":
		var msg EventMessage
		if err := json.Unmarshal(data, &msg); err != nil {
			return nil, err
		}
		return msg, nil
	case probe.Code != nil:
		var msg StatusMessage
		if err := json.Unmarshal(data, &msg); err != nil {
			return nil, err
		}
		return msg, nil
	case probe.TA != "":
		var msg OnchainSimpleTokenPriceUpdate
		if err := json.Unmarshal(data, &msg); err != nil {
			return nil, err
		}
		return msg, nil
	case probe.TX != "":
		var msg OnchainTradeUpdate
		if err := json.Unmarshal(data, &msg); err != nil {
			return nil, err
		}
		return msg, nil
	case probe.PA != "":
		var msg OnchainOHLCVUpdate
		if err := json.Unmarshal(data, &msg); err != nil {
			return nil, err
		}
		return msg, nil
	case probe.VS != "" || probe.I != "":
		var msg CGSimplePriceUpdate
		if err := json.Unmarshal(data, &msg); err != nil {
			return nil, err
		}
		return msg, nil
	default:
		return UnknownMessage{Raw: append([]byte(nil), data...)}, nil
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (n *Number) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if len(data) == 0 || bytes.Equal(data, []byte("null")) {
		return nil
	}

	if data[0] == '"' {
		var value string
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		parsed, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		*n = Number(parsed)
		return nil
	}

	parsed, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return err
	}
	*n = Number(parsed)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (u *CGSimplePriceUpdate) UnmarshalJSON(data []byte) error {
	var payload struct {
		C  string  `json:"c"`
		CH string  `json:"ch"`
		I  string  `json:"i"`
		VS string  `json:"vs"`
		M  *Number `json:"m"`
		P  *Number `json:"p"`
		PP *Number `json:"pp"`
		T  *Number `json:"t"`
		V  *Number `json:"v"`
	}

	if err := json.Unmarshal(data, &payload); err != nil {
		return err
	}

	u.ChannelType = pickChannelType(payload.C, payload.CH)
	u.CoinID = payload.I
	u.VSCurrency = payload.VS
	u.MarketCapUSD = payload.M
	u.Price = payload.P
	u.Price24hChangePercentage = payload.PP
	u.LastUpdatedAt = payload.T
	u.Volume24h = payload.V

	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (u *OnchainSimpleTokenPriceUpdate) UnmarshalJSON(data []byte) error {
	var payload struct {
		C  string  `json:"c"`
		CH string  `json:"ch"`
		N  string  `json:"n"`
		TA string  `json:"ta"`
		P  *Number `json:"p"`
		PP *Number `json:"pp"`
		M  *Number `json:"m"`
		V  *Number `json:"v"`
		T  *Number `json:"t"`
	}

	if err := json.Unmarshal(data, &payload); err != nil {
		return err
	}

	u.ChannelType = pickChannelType(payload.C, payload.CH)
	u.NetworkID = payload.N
	u.TokenAddress = payload.TA
	u.PriceUSD = payload.P
	u.Price24hChangePercentage = payload.PP
	u.MarketCapUSD = payload.M
	u.Volume24h = payload.V
	u.LastUpdatedAt = payload.T

	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (u *OnchainTradeUpdate) UnmarshalJSON(data []byte) error {
	var payload struct {
		C   string  `json:"c"`
		CH  string  `json:"ch"`
		N   string  `json:"n"`
		PA  string  `json:"pa"`
		TX  string  `json:"tx"`
		TY  string  `json:"ty"`
		TO  *Number `json:"to"`
		TOQ *Number `json:"toq"`
		VO  *Number `json:"vo"`
		PC  *Number `json:"pc"`
		PU  *Number `json:"pu"`
		T   *Number `json:"t"`
	}

	if err := json.Unmarshal(data, &payload); err != nil {
		return err
	}

	u.ChannelType = pickChannelType(payload.C, payload.CH)
	u.NetworkID = payload.N
	u.PoolAddress = payload.PA
	u.TransactionHash = payload.TX
	u.TradeType = payload.TY
	u.TokenAmount = payload.TO
	u.QuoteTokenAmount = payload.TOQ
	u.VolumeUSD = payload.VO
	u.PriceInNativeCurrency = payload.PC
	u.PriceUSD = payload.PU
	u.LastUpdatedAt = payload.T

	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (u *OnchainOHLCVUpdate) UnmarshalJSON(data []byte) error {
	var payload map[string]json.RawMessage
	if err := json.Unmarshal(data, &payload); err != nil {
		return err
	}

	if err := unmarshalField(payload, "ch", &u.ChannelType); err != nil {
		return err
	}
	if err := unmarshalField(payload, "n", &u.NetworkID); err != nil {
		return err
	}
	if err := unmarshalField(payload, "pa", &u.PoolAddress); err != nil {
		return err
	}
	if err := unmarshalField(payload, "to", &u.Token); err != nil {
		return err
	}
	if err := unmarshalField(payload, "i", &u.Interval); err != nil {
		return err
	}
	if err := unmarshalField(payload, "o", &u.Open); err != nil {
		return err
	}
	if err := unmarshalField(payload, "h", &u.High); err != nil {
		return err
	}
	if err := unmarshalField(payload, "l", &u.Low); err != nil {
		return err
	}
	if err := unmarshalField(payload, "c", &u.Close); err != nil {
		return err
	}
	if err := unmarshalField(payload, "v", &u.Volume); err != nil {
		return err
	}
	if err := unmarshalField(payload, "t", &u.Timestamp); err != nil {
		return err
	}

	return nil
}

func pickChannelType(primary, secondary string) string {
	if primary != "" {
		return primary
	}
	return secondary
}

func unmarshalField[T any](payload map[string]json.RawMessage, key string, target *T) error {
	raw, ok := payload[key]
	if !ok {
		return nil
	}
	return json.Unmarshal(raw, target)
}
