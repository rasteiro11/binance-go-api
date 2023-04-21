package binance

import (
	md "binance-go-api/binance/market_data"
)

type BinanceApiService interface {
	md.MarketData
}
