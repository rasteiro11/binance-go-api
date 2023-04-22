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
	_, _ = binanceService.OldTradesLookup(ctx, &marketdata.OldTradesLookupRequest{Symbol: "BTCUSDT", FromId: 28457})
}

func TestKlines(t *testing.T) {
	ctx := context.TODO()
	binanceService := binance.NewService()
	res, err := binanceService.Klines(ctx, &marketdata.KlinesRequest{Symbol: "BTCUSDT", Interval: "1h"})
	assert.Nil(t, err, "binanceService.Klines() returned error")
	assert.NotNil(t, res, "response can not be nil")
}

func TestUIKlines(t *testing.T) {
	ctx := context.TODO()
	binanceService := binance.NewService()
	res, err := binanceService.UIKlines(ctx, &marketdata.KlinesRequest{Symbol: "BTCUSDT", Interval: "1h"})
	assert.Nil(t, err, "binanceService.UIKlines() returned error")
	assert.NotNil(t, res, "response can not be nil")
}

func TestAveragePrice(t *testing.T) {
	ctx := context.TODO()
	binanceService := binance.NewService()
	res, err := binanceService.AveragePrice(ctx, &marketdata.AveragePriceRequest{Symbol: "BTCUSDT"})
	assert.Nil(t, err, "binanceService.AveragePrice() returned error")
	assert.NotNil(t, res, "response can not be nil")
	fmt.Println(res)
}

func TestPriceChange24H(t *testing.T) {
	ctx := context.TODO()
	binanceService := binance.NewService()
	res, err := binanceService.PriceChange24H(ctx, &marketdata.PriceChangeRequest{
		Symbols: []string{"BTCUSDT"},
		Type:    "MINI",
	})
	assert.Nil(t, err, "binanceService.PriceChange24H]() returned error")
	assert.NotNil(t, res, "response can not be nil")
	fmt.Println(res)
}

func TestPriceTicker(t *testing.T) {
	ctx := context.TODO()
	binanceService := binance.NewService()
	res, err := binanceService.SymbolPriceTicker(ctx, &marketdata.SymbolPriceTickerRequest{
		Symbols: []string{"BTCUSDT", "LTCBTC"},
	})
	assert.Nil(t, err, "binanceService.SymbolPriceTicker() returned error")
	assert.NotNil(t, res, "response can not be nil")
	fmt.Println(res)
}

func TestSymbolOrderBookTicker(t *testing.T) {
	ctx := context.TODO()
	binanceService := binance.NewService()
	res, err := binanceService.SymbolOrderBookTicker(ctx, &marketdata.SymbolOrderBookTickerRequest{
		Symbols: []string{"BTCUSDT", "LTCBTC"},
	})
	assert.Nil(t, err, "binanceService.SymbolOrderBookTicker() returned error")
	assert.NotNil(t, res, "response can not be nil")
	fmt.Println(res)
}
func TestRollingWindowPriceChange(t *testing.T) {
	ctx := context.TODO()
	binanceService := binance.NewService()
	res, err := binanceService.RollingWindowPriceChange(ctx, &marketdata.RollingWindowPriceChangeRequest{
		Symbols: []string{"BTCUSDT", "LTCBTC"},
		Type:    "MINI",
	})
	assert.Nil(t, err, "binanceService.SymbolOrderBookTicker() returned error")
	assert.NotNil(t, res, "response can not be nil")
	fmt.Println(res)
	t.Fail()
}
