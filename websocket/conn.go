package websocket

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	coderws "github.com/coder/websocket"
)

var (
	// ErrMissingTargets is returned when a channel action is sent without any target ids.
	ErrMissingTargets = errors.New("at least one subscription target is required")
	// ErrInvalidOHLCVInterval is returned when the requested OHLCV interval is unsupported by CoinGecko.
	ErrInvalidOHLCVInterval = errors.New("invalid ohlcv interval")
	// ErrInvalidOHLCVToken is returned when the requested OHLCV token side is unsupported by CoinGecko.
	ErrInvalidOHLCVToken = errors.New("invalid ohlcv token side")
)

const (
	commandMessage     = "message"
	commandSubscribe   = "subscribe"
	commandUnsubscribe = "unsubscribe"

	actionSetPools    = "set_pools"
	actionSetTokens   = "set_tokens"
	actionUnsetPools  = "unset_pools"
	actionUnsetTokens = "unset_tokens"
)

// Channel identifies a CoinGecko websocket channel.
type Channel string

const (
	// ChannelCGSimplePrice streams CoinGecko simple price updates.
	ChannelCGSimplePrice Channel = "CGSimplePrice"
	// ChannelOnchainSimpleTokenPrice streams GeckoTerminal token price updates.
	ChannelOnchainSimpleTokenPrice Channel = "OnchainSimpleTokenPrice"
	// ChannelOnchainTrade streams GeckoTerminal trade updates.
	ChannelOnchainTrade Channel = "OnchainTrade"
	// ChannelOnchainOHLCV streams GeckoTerminal OHLCV updates.
	ChannelOnchainOHLCV Channel = "OnchainOHLCV"
)

// OHLCVInterval identifies a supported OHLCV candle interval.
type OHLCVInterval string

const (
	OHLCVInterval1s  OHLCVInterval = "1s"
	OHLCVInterval1m  OHLCVInterval = "1m"
	OHLCVInterval5m  OHLCVInterval = "5m"
	OHLCVInterval15m OHLCVInterval = "15m"
	OHLCVInterval1h  OHLCVInterval = "1h"
	OHLCVInterval2h  OHLCVInterval = "2h"
	OHLCVInterval4h  OHLCVInterval = "4h"
	OHLCVInterval8h  OHLCVInterval = "8h"
	OHLCVInterval12h OHLCVInterval = "12h"
	OHLCVInterval1d  OHLCVInterval = "1d"
)

// OHLCVToken identifies whether OHLCV updates are streamed for the base or quote token of a pool.
type OHLCVToken string

const (
	OHLCVTokenBase  OHLCVToken = "base"
	OHLCVTokenQuote OHLCVToken = "quote"
)

// Conn wraps an active CoinGecko websocket connection.
type Conn struct {
	conn *coderws.Conn
}

// Raw returns the underlying websocket connection.
func (c *Conn) Raw() *coderws.Conn {
	return c.conn
}

// Close closes the websocket connection with a normal closure status.
func (c *Conn) Close() error {
	return c.conn.Close(coderws.StatusNormalClosure, "")
}

// CloseWithStatus closes the websocket connection with a caller-provided status and reason.
func (c *Conn) CloseWithStatus(code coderws.StatusCode, reason string) error {
	return c.conn.Close(code, reason)
}

// ReadRaw reads the next websocket message as raw JSON bytes.
func (c *Conn) ReadRaw(ctx context.Context) ([]byte, error) {
	_, data, err := c.conn.Read(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Read reads and decodes the next websocket JSON message into one of the typed message variants.
func (c *Conn) Read(ctx context.Context) (Message, error) {
	data, err := c.ReadRaw(ctx)
	if err != nil {
		return nil, err
	}
	return ParseMessage(data)
}

// Subscribe subscribes the connection to a websocket channel.
func (c *Conn) Subscribe(ctx context.Context, channel Channel) error {
	data, err := buildChannelCommand(commandSubscribe, channel)
	if err != nil {
		return err
	}
	return c.write(ctx, data)
}

// Unsubscribe unsubscribes the connection from a websocket channel.
func (c *Conn) Unsubscribe(ctx context.Context, channel Channel) error {
	data, err := buildChannelCommand(commandUnsubscribe, channel)
	if err != nil {
		return err
	}
	return c.write(ctx, data)
}

// SetCGSimplePrice configures the CGSimplePrice channel for the provided coin ids and quote currencies.
func (c *Conn) SetCGSimplePrice(ctx context.Context, coinIDs []string, vsCurrencies ...string) error {
	data, err := buildActionCommand(
		ChannelCGSimplePrice,
		channelAction{
			Action:       actionSetTokens,
			CoinIDs:      cleanValues(coinIDs),
			VSCurrencies: cleanValues(vsCurrencies),
		},
	)
	if err != nil {
		return err
	}
	return c.write(ctx, data)
}

// UnsetCGSimplePrice removes specific CGSimplePrice subscriptions for the provided coin ids and quote currencies.
func (c *Conn) UnsetCGSimplePrice(ctx context.Context, coinIDs []string, vsCurrencies ...string) error {
	data, err := buildActionCommand(
		ChannelCGSimplePrice,
		channelAction{
			Action:       actionUnsetTokens,
			CoinIDs:      cleanValues(coinIDs),
			VSCurrencies: cleanValues(vsCurrencies),
		},
	)
	if err != nil {
		return err
	}
	return c.write(ctx, data)
}

// SetOnchainSimpleTokenPrice configures the OnchainSimpleTokenPrice channel for the provided token targets.
func (c *Conn) SetOnchainSimpleTokenPrice(ctx context.Context, targets []string) error {
	data, err := buildActionCommand(
		ChannelOnchainSimpleTokenPrice,
		channelAction{
			Action:                actionSetTokens,
			NetworkTokenAddresses: cleanValues(targets),
		},
	)
	if err != nil {
		return err
	}
	return c.write(ctx, data)
}

// UnsetOnchainSimpleTokenPrice removes specific OnchainSimpleTokenPrice subscriptions for the provided token targets.
func (c *Conn) UnsetOnchainSimpleTokenPrice(ctx context.Context, targets []string) error {
	data, err := buildActionCommand(
		ChannelOnchainSimpleTokenPrice,
		channelAction{
			Action:                actionUnsetTokens,
			NetworkTokenAddresses: cleanValues(targets),
		},
	)
	if err != nil {
		return err
	}
	return c.write(ctx, data)
}

// SetOnchainTrade configures the OnchainTrade channel for the provided pool targets.
func (c *Conn) SetOnchainTrade(ctx context.Context, targets []string) error {
	data, err := buildActionCommand(
		ChannelOnchainTrade,
		channelAction{
			Action:               actionSetPools,
			NetworkPoolAddresses: cleanValues(targets),
		},
	)
	if err != nil {
		return err
	}
	return c.write(ctx, data)
}

// UnsetOnchainTrade removes specific OnchainTrade subscriptions for the provided pool targets.
func (c *Conn) UnsetOnchainTrade(ctx context.Context, targets []string) error {
	data, err := buildActionCommand(
		ChannelOnchainTrade,
		channelAction{
			Action:               actionUnsetPools,
			NetworkPoolAddresses: cleanValues(targets),
		},
	)
	if err != nil {
		return err
	}
	return c.write(ctx, data)
}

// SetOnchainOHLCV configures the OnchainOHLCV channel for the provided pool targets.
func (c *Conn) SetOnchainOHLCV(ctx context.Context, targets []string, interval OHLCVInterval, token OHLCVToken) error {
	data, err := buildOHLCVActionCommand(
		ChannelOnchainOHLCV,
		channelAction{
			Action:               actionSetPools,
			NetworkPoolAddresses: cleanValues(targets),
			Interval:             interval,
			Token:                token,
		},
	)
	if err != nil {
		return err
	}
	return c.write(ctx, data)
}

// UnsetOnchainOHLCV removes specific OnchainOHLCV subscriptions for the provided pool targets.
func (c *Conn) UnsetOnchainOHLCV(ctx context.Context, targets []string, interval OHLCVInterval, token OHLCVToken) error {
	data, err := buildOHLCVActionCommand(
		ChannelOnchainOHLCV,
		channelAction{
			Action:               actionUnsetPools,
			NetworkPoolAddresses: cleanValues(targets),
			Interval:             interval,
			Token:                token,
		},
	)
	if err != nil {
		return err
	}
	return c.write(ctx, data)
}

type channelIdentifier struct {
	Channel Channel `json:"channel"`
}

type wireCommand struct {
	Command    string `json:"command"`
	Identifier string `json:"identifier"`
	Data       string `json:"data,omitempty"`
}

type channelAction struct {
	Action                string        `json:"action"`
	CoinIDs               []string      `json:"coin_id,omitempty"`
	VSCurrencies          []string      `json:"vs_currencies,omitempty"`
	NetworkTokenAddresses []string      `json:"network_id:token_addresses,omitempty"`
	NetworkPoolAddresses  []string      `json:"network_id:pool_addresses,omitempty"`
	Interval              OHLCVInterval `json:"interval,omitempty"`
	Token                 OHLCVToken    `json:"token,omitempty"`
}

func (c *Conn) write(ctx context.Context, data []byte) error {
	return c.conn.Write(ctx, coderws.MessageText, data)
}

func buildChannelCommand(command string, channel Channel) ([]byte, error) {
	identifier, err := marshalEmbeddedJSON(channelIdentifier{Channel: channel})
	if err != nil {
		return nil, err
	}
	return json.Marshal(wireCommand{
		Command:    command,
		Identifier: identifier,
	})
}

func buildActionCommand(channel Channel, payload channelAction) ([]byte, error) {
	if len(payload.CoinIDs) == 0 && len(payload.NetworkTokenAddresses) == 0 && len(payload.NetworkPoolAddresses) == 0 {
		return nil, ErrMissingTargets
	}

	identifier, err := marshalEmbeddedJSON(channelIdentifier{Channel: channel})
	if err != nil {
		return nil, err
	}

	data, err := marshalEmbeddedJSON(payload)
	if err != nil {
		return nil, err
	}

	return json.Marshal(wireCommand{
		Command:    commandMessage,
		Identifier: identifier,
		Data:       data,
	})
}

func buildOHLCVActionCommand(channel Channel, payload channelAction) ([]byte, error) {
	if !isValidOHLCVInterval(payload.Interval) {
		return nil, ErrInvalidOHLCVInterval
	}
	if !isValidOHLCVToken(payload.Token) {
		return nil, ErrInvalidOHLCVToken
	}
	return buildActionCommand(channel, payload)
}

func marshalEmbeddedJSON(v any) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func cleanValues(values []string) []string {
	if len(values) == 0 {
		return nil
	}

	cleaned := make([]string, 0, len(values))
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed == "" {
			continue
		}
		cleaned = append(cleaned, trimmed)
	}

	if len(cleaned) == 0 {
		return nil
	}

	return cleaned
}

func isValidOHLCVInterval(interval OHLCVInterval) bool {
	switch interval {
	case OHLCVInterval1s, OHLCVInterval1m, OHLCVInterval5m, OHLCVInterval15m,
		OHLCVInterval1h, OHLCVInterval2h, OHLCVInterval4h, OHLCVInterval8h,
		OHLCVInterval12h, OHLCVInterval1d:
		return true
	default:
		return false
	}
}

func isValidOHLCVToken(token OHLCVToken) bool {
	switch token {
	case OHLCVTokenBase, OHLCVTokenQuote:
		return true
	default:
		return false
	}
}
