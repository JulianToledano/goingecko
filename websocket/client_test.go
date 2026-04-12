package websocket

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"
)

func TestDialConfigUsesQueryAuthByDefault(t *testing.T) {
	t.Parallel()

	client := NewClient("secret", WithURL("wss://stream.coingecko.com/v1?foo=bar"))

	dialURL, header, err := client.dialConfig()
	if err != nil {
		t.Fatal(err)
	}

	if dialURL != "wss://stream.coingecko.com/v1?foo=bar&x_cg_pro_api_key=secret" {
		t.Fatalf("unexpected dial url: %s", dialURL)
	}

	if got := header.Get(proAPIKeyHeaderKey); got != "" {
		t.Fatalf("expected no auth header, got %q", got)
	}
}

func TestDialConfigSupportsHeaderAuth(t *testing.T) {
	t.Parallel()

	header := make(http.Header)
	header.Set("X-Test", "1")

	client := NewClient(
		"secret",
		WithAuthMode(AuthModeHeader),
		WithHeader(header),
	)

	dialURL, gotHeader, err := client.dialConfig()
	if err != nil {
		t.Fatal(err)
	}

	if dialURL != BaseURL {
		t.Fatalf("unexpected dial url: %s", dialURL)
	}

	if got := gotHeader.Get(proAPIKeyHeaderKey); got != "secret" {
		t.Fatalf("unexpected auth header: %q", got)
	}

	if got := gotHeader.Get("X-Test"); got != "1" {
		t.Fatalf("missing custom header: %q", got)
	}
}

func TestBuildActionCommand(t *testing.T) {
	t.Parallel()

	data, err := buildActionCommand(
		ChannelOnchainTrade,
		channelAction{
			Action:               actionSetPools,
			NetworkPoolAddresses: []string{"bsc:0xpool"},
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	var command wireCommand
	if err := json.Unmarshal(data, &command); err != nil {
		t.Fatal(err)
	}

	if command.Command != commandMessage {
		t.Fatalf("unexpected command: %s", command.Command)
	}

	var identifier channelIdentifier
	if err := json.Unmarshal([]byte(command.Identifier), &identifier); err != nil {
		t.Fatal(err)
	}
	if identifier.Channel != ChannelOnchainTrade {
		t.Fatalf("unexpected channel: %s", identifier.Channel)
	}

	var payload channelAction
	if err := json.Unmarshal([]byte(command.Data), &payload); err != nil {
		t.Fatal(err)
	}
	if payload.Action != actionSetPools {
		t.Fatalf("unexpected action: %s", payload.Action)
	}
	if len(payload.NetworkPoolAddresses) != 1 || payload.NetworkPoolAddresses[0] != "bsc:0xpool" {
		t.Fatalf("unexpected targets: %#v", payload.NetworkPoolAddresses)
	}
}

func TestBuildOHLCVActionCommandValidatesIntervalAndToken(t *testing.T) {
	t.Parallel()

	_, err := buildOHLCVActionCommand(
		ChannelOnchainOHLCV,
		channelAction{
			Action:               actionSetPools,
			NetworkPoolAddresses: []string{"bsc:0xpool"},
			Interval:             "30m",
			Token:                OHLCVTokenBase,
		},
	)
	if !errors.Is(err, ErrInvalidOHLCVInterval) {
		t.Fatalf("expected ErrInvalidOHLCVInterval, got %v", err)
	}

	_, err = buildOHLCVActionCommand(
		ChannelOnchainOHLCV,
		channelAction{
			Action:               actionSetPools,
			NetworkPoolAddresses: []string{"bsc:0xpool"},
			Interval:             OHLCVInterval1m,
			Token:                "mid",
		},
	)
	if !errors.Is(err, ErrInvalidOHLCVToken) {
		t.Fatalf("expected ErrInvalidOHLCVToken, got %v", err)
	}
}

func TestParseMessage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		payload string
		assert  func(t *testing.T, message Message)
	}{
		{
			name:    "event",
			payload: `{"type":"confirm_subscription","identifier":"{\"channel\":\"OnchainTrade\"}"}`,
			assert: func(t *testing.T, message Message) {
				t.Helper()

				event, ok := message.(EventMessage)
				if !ok {
					t.Fatalf("expected EventMessage, got %T", message)
				}
				if event.Type != "confirm_subscription" {
					t.Fatalf("unexpected event type: %s", event.Type)
				}
			},
		},
		{
			name:    "status",
			payload: `{"code":2000,"message":"Subscription successful"}`,
			assert: func(t *testing.T, message Message) {
				t.Helper()

				status, ok := message.(StatusMessage)
				if !ok {
					t.Fatalf("expected StatusMessage, got %T", message)
				}
				if status.Code != 2000 {
					t.Fatalf("unexpected status code: %d", status.Code)
				}
			},
		},
		{
			name:    "cg simple price",
			payload: `{"c":"C1","i":"bitcoin","vs":"usd","p":100000.12,"pp":1.2,"m":2000000,"v":42.5,"t":1747808150.269067}`,
			assert: func(t *testing.T, message Message) {
				t.Helper()

				update, ok := message.(CGSimplePriceUpdate)
				if !ok {
					t.Fatalf("expected CGSimplePriceUpdate, got %T", message)
				}
				if update.ChannelType != "C1" || update.CoinID != "bitcoin" || update.VSCurrency != "usd" {
					t.Fatalf("unexpected update: %#v", update)
				}
				if update.Price == nil || update.Price.Float64() != 100000.12 {
					t.Fatalf("unexpected price: %#v", update.Price)
				}
			},
		},
		{
			name:    "onchain trade",
			payload: `{"ch":"G2","n":"bsc","pa":"0xpool","tx":"0xtx","ty":"b","to":11.08,"toq":0.01,"vo":"2.75","pc":"0.0002","pu":"3656.89","t":1724927796000}`,
			assert: func(t *testing.T, message Message) {
				t.Helper()

				update, ok := message.(OnchainTradeUpdate)
				if !ok {
					t.Fatalf("expected OnchainTradeUpdate, got %T", message)
				}
				if update.ChannelType != "G2" || update.NetworkID != "bsc" || update.TradeType != "b" {
					t.Fatalf("unexpected update: %#v", update)
				}
				if update.VolumeUSD == nil || update.VolumeUSD.Float64() != 2.75 {
					t.Fatalf("unexpected volume: %#v", update.VolumeUSD)
				}
			},
		},
		{
			name:    "onchain ohlcv",
			payload: `{"ch":"G3","n":"eth","pa":"0xpool","to":"base","i":"1m","o":1.1,"h":2.2,"l":0.9,"c":1.8,"v":10.5,"t":1753803600}`,
			assert: func(t *testing.T, message Message) {
				t.Helper()

				update, ok := message.(OnchainOHLCVUpdate)
				if !ok {
					t.Fatalf("expected OnchainOHLCVUpdate, got %T", message)
				}
				if update.ChannelType != "G3" || update.Interval != OHLCVInterval1m || update.Token != OHLCVTokenBase {
					t.Fatalf("unexpected update: %#v", update)
				}
				if update.Close == nil || update.Close.Float64() != 1.8 {
					t.Fatalf("unexpected close: %#v", update.Close)
				}
			},
		},
		{
			name:    "unknown",
			payload: `{"hello":"world"}`,
			assert: func(t *testing.T, message Message) {
				t.Helper()

				if _, ok := message.(UnknownMessage); !ok {
					t.Fatalf("expected UnknownMessage, got %T", message)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			message, err := ParseMessage([]byte(tt.payload))
			if err != nil {
				t.Fatal(err)
			}

			tt.assert(t, message)
		})
	}
}
