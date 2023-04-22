package marketdata

import (
	"context"
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
	Symbol string `json:"symbol"`
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

type OldTradesLookup struct {
	Symbol string `json:"symbol"`
	Limit  int    `json:"limit"`
	FromId int64  `json:"fromId"`
}

type MarketData interface {
	Time(ctx context.Context) (*TimeResponse, error)
	ExchangeInfo(ctx context.Context, req *ExchangeInfoRequest) (*ExchangeInfoResponse, error)
	OrderBook(ctx context.Context, req *OrderBookRequest) (*OrderBookResponse, error)
	RecentTrades(ctx context.Context, req *RecentTradesRequest) ([]Trade, error)
	OldTradesLookup(ctx context.Context, req *OldTradesLookup) ([]Trade, error)
}
