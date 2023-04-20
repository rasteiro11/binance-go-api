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

type ExchangeInfoResponse struct {
	Timezone   string `json:"timezone"`
	ServerTime int64  `json:"serverTime"`
}

type MarketData interface {
	Time(ctx context.Context) (*TimeResponse, error)
	ExchangeInfo(ctx context.Context, req *ExchangeInfoRequest) (*ExchangeInfoResponse, error)
}
