package marketdata

import (
	"context"
	"encoding/json"
)

type TimeResponse struct {
	ServerTime int64 `json:"serverTime"`
}

type ExchangeInfoRequest struct {
	Symbols     []string
	Permissions []string
}

type Symbol struct {
	Symbol                          string   `json:"symbol"`
	Status                          string   `json:"status"`
	BaseAssetPrecision              int      `json:"baseAssetPrecision"`
	QuoteAsset                      string   `json:"quoteAsset"`
	QuotePrecision                  int      `json:"quotePrecision"`
	QuoteAssetPrecision             int      `json:"quoteAssetPrecision"`
	OrderTypes                      []string `json:"orderTypes"`
	IcebergAllowed                  bool     `json:"icebergAllowed"`
	OcoAllowed                      bool     `json:"ocoAllowed"`
	QuoteOrderQtyMarketAllowed      bool     `json:"quoteOrderQtyMarketAllowed"`
	AllowTrailingStop               bool     `json:"allowTrailingStop"`
	CancelReplaceAllowed            bool     `json:"cancelReplaceAllowed"`
	IsSpotTradingAllowed            bool     `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed          bool     `json:"isMarginTradingAllowed"`
	Permissions                     []string `json:"permissions"`
	DefaultSelfTradePreventionMode  string   `json:"defaultSelfTradePreventionMode"`
	AllowedSelfTradePreventionModes []string `json:"allowedSelfTradePreventionModes"`
}

type ExchangeInfoResponse struct {
	Timezone   string   `json:"timezone"`
	ServerTime int64    `json:"serverTime"`
	Symbols    []Symbol `json:"symbols"`
}

type OrderBookRequest struct {
	Symbol string `json:"symbol" validate:"required"`
	Limit  int    `json:"limit"`
}

type OrderBookResponse struct {
	LastUpdateId int        `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

type RecentTradesRequest struct {
	Symbol string `json:"symbol" validate:"required"`
	Limit  int    `json:"limit"`
}

type Trade struct {
	Id           int64  `json:"id"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	QuoteQty     string `json:"quoteQty"`
	Time         int64  `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}

type OldTradesLookupRequest struct {
	Symbol string `json:"symbol" validate:"required"`
	Limit  int    `json:"limit"`
	FromId int64  `json:"fromId"`
}

type KlinesRequest struct {
	Symbol    string `json:"symbol" validate:"required"`
	Interval  string `json:"interval" validate:"required"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
	Limit     int    `json:"limit"`
}

type Kline struct {
	OpenTime                 float64
	OpenPrice                string
	HighPrice                string
	LowPrice                 string
	ClosePrice               string
	Volume                   string
	KlineCloseTime           int64
	QuoteAssetVolume         string
	NumberOfTrades           int64
	TakerBuyBaseAssetVolume  string
	TakerBuyQuoteAssetVolume string
	UnusedField              string
}

func (k *Kline) UnmarshalJSON(bs []byte) error {
	arr := []interface{}{}

	err := json.Unmarshal(bs, &arr)
	if err != nil {
		return err
	}

	opTime := arr[0].(float64)
	k.OpenTime = opTime

	opPrice := arr[1].(string)
	k.OpenPrice = opPrice

	highPrice := arr[2].(string)
	k.HighPrice = highPrice

	lowPrice := arr[3].(string)
	k.LowPrice = lowPrice

	closePrice := arr[4].(string)
	k.ClosePrice = closePrice

	volume := arr[5].(string)
	k.Volume = volume

	klineCloseTime := arr[6].(float64)
	k.KlineCloseTime = int64(klineCloseTime)

	quoteAssetVolume := arr[7].(string)
	k.QuoteAssetVolume = quoteAssetVolume

	numOfTrades := arr[8].(float64)
	k.NumberOfTrades = int64(numOfTrades)

	takerBuyBaseAssetVolume := arr[9].(string)
	k.TakerBuyBaseAssetVolume = takerBuyBaseAssetVolume

	takerBuyQuoteAssetVolume := arr[10].(string)
	k.TakerBuyQuoteAssetVolume = takerBuyQuoteAssetVolume

	unusedField := arr[11].(string)
	k.UnusedField = unusedField

	return nil
}

type AveragePriceRequest struct {
	Symbol string `json:"symbol" validate:"required"`
}

type AveragePriceResponse struct {
	Mins  int    `json:"mins"`
	Price string `json:"price"`
}

type PriceChangeRequest struct {
	Symbols []string
	Type    string `json:"type"` // FULL or MINI
}

type PriceChangeResponse struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstId            int64  `json:"firstId"`
	LastId             int64  `json:"lastId"`
	Count              int64  `json:"count"`
}

type SymbolPriceTickerRequest struct {
	Symbols []string
}

type SymbolPriceTicker struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type SymbolOrderBookTickerRequest struct {
	Symbols []string
}

type SymbolOrderBookTicker struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
}

type RollingWindowPriceChangeRequest struct {
	Symbols    []string `json:"symbol" validate:"required"`
	WindowSize string   `json:"windowSize"`
	Type       string   `json:"type"`
}

type RollingWindowPriceChange struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	LastPrice          string `json:"lastPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstId            int64  `json:"firstId"`
	LastId             int64  `json:"lastId"`
	Count              int64  `json:"count"`
}

type MarketData interface {
	Time(ctx context.Context) (*TimeResponse, error)
	ExchangeInfo(ctx context.Context, req *ExchangeInfoRequest) (*ExchangeInfoResponse, error)
	OrderBook(ctx context.Context, req *OrderBookRequest) (*OrderBookResponse, error)
	RecentTrades(ctx context.Context, req *RecentTradesRequest) ([]Trade, error)
	OldTradesLookup(ctx context.Context, req *OldTradesLookupRequest) ([]Trade, error)
	Klines(ctx context.Context, req *KlinesRequest) ([]Kline, error)
	UIKlines(ctx context.Context, req *KlinesRequest) ([]Kline, error)
	AveragePrice(ctx context.Context, req *AveragePriceRequest) (*AveragePriceResponse, error)
	PriceChange24H(ctx context.Context, req *PriceChangeRequest) ([]PriceChangeResponse, error)
	SymbolPriceTicker(ctx context.Context, req *SymbolPriceTickerRequest) ([]SymbolPriceTicker, error)
	SymbolOrderBookTicker(ctx context.Context, req *SymbolOrderBookTickerRequest) ([]SymbolOrderBookTicker, error)
	RollingWindowPriceChange(ctx context.Context, req *RollingWindowPriceChangeRequest) ([]RollingWindowPriceChange, error)
}
