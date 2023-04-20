package binance_test

import (
	"binance-go-api/binance"
	marketdata "binance-go-api/binance/market_data"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	ctx := context.TODO()
	binanceService := binance.NewService()

	timeResponse, err := binanceService.Time(ctx)
	assert.Nil(t, err, "binanceService.Time() returned error")
	assert.NotZero(t, timeResponse.ServerTime, "server time can not be zero")
}

func TestExchangeInfo(t *testing.T) {
	ctx := context.TODO()
	binanceService := binance.NewService()
	_, err := binanceService.ExchangeInfo(ctx, &marketdata.ExchangeInfoRequest{
		Symbols: []string{"BTCUSDT", "BNBBTC"},
	})

	assert.Nil(t, err, "binanceService.ExchangeInfo() returned error")
	t.Fail()
}
