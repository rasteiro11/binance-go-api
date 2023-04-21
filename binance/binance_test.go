package binance_test

import (
	"binance-go-api/binance"
	marketdata "binance-go-api/binance/market_data"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	ctx := context.TODO()
	binanceService := binance.NewService()

	timeResponse, err := binanceService.Time(ctx)
	assert.Nil(t, err, "binanceService.Time() returned error")
	assert.NotNil(t, timeResponse, "response can not be nil")
	assert.NotZero(t, timeResponse.ServerTime, "server time can not be zero")
}

func TestExchangeInfo(t *testing.T) {
	ctx := context.TODO()
	binanceService := binance.NewService()
	res, err := binanceService.ExchangeInfo(ctx, &marketdata.ExchangeInfoRequest{
		Symbols: []string{"BTCUSDT", "BNBBTC"},
	})

	assert.Nil(t, err, "binanceService.ExchangeInfo() returned error")
	assert.NotNil(t, res, "response can not be nil")
}

func TestOrderBook(t *testing.T) {
	ctx := context.TODO()
	binanceService := binance.NewService()
	res, err := binanceService.OrderBook(ctx, &marketdata.OrderBookRequest{Symbol: "BTCUSDT"})
	assert.Nil(t, err, "binanceService.OrderBook() returned error")
	assert.NotNil(t, res, "response can not be nil")
}

func TestRecentTrades(t *testing.T) {
	ctx := context.TODO()
	binanceService := binance.NewService()
	res, err := binanceService.RecentTrades(ctx, &marketdata.RecentTradesRequest{Symbol: "BTCUSDT"})
	fmt.Println(res)
	assert.Nil(t, err, "binanceService.RecentTrades() returned error")
	assert.NotNil(t, res, "response can not be nil")
}

func TestOldTadeLookup(t *testing.T) {
	ctx := context.TODO()
	binanceService := binance.NewService()
	_, _ = binanceService.OldTradesLookup(ctx, &marketdata.OldTradesLookup{Symbol: "BTCUSDT", FromId: 28457})
}
