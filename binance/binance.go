package binance

import (
	binance "binance-go-api/binance/market_data"
	"binance-go-api/httpclient"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const BASE_ENDPOINT = "https://api.binance.com/api/v3"

type service struct {
	client httpclient.HttpClient
}

type ServiceOpts func(*service)

var _ BinanceApiService = (*service)(nil)

func (s *service) Time(ctx context.Context) (*binance.TimeResponse, error) {
	return call[binance.TimeResponse](ctx, http.MethodGet, "/time", s)
}

func (s *service) ExchangeInfo(ctx context.Context, req *binance.ExchangeInfoRequest) (*binance.ExchangeInfoResponse, error) {
	opts := []httpclient.RequestOption{}

	if len(req.Permissions) == 1 {
		opts = append(opts, httpclient.WithQueryParam("permissions", req.Permissions[0]))
	} else {
		opts = append(opts, httpclient.WithQueryListParam("permissions", req.Permissions))
	}

	queryName := "symbol"
	if len(req.Symbols) == 1 {
		opts = append(opts, httpclient.WithQueryParam(queryName, req.Symbols[0]))
		return call[binance.ExchangeInfoResponse](ctx, http.MethodGet, "/exchangeInfo", s, opts...)
	}

	opts = append(opts, httpclient.WithQueryListParam("symbols", req.Symbols))
	res, err := call[binance.ExchangeInfoResponse](ctx, http.MethodGet, "/exchangeInfo", s, opts...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *service) OrderBook(ctx context.Context, req *binance.OrderBookRequest) (*binance.OrderBookResponse, error) {
	opts := []httpclient.RequestOption{}
	opts = append(opts, httpclient.WithQueryParam("symbol", req.Symbol))

	if req.Limit != 0 {
		limit := fmt.Sprintf("%d", req.Limit)
		opts = append(opts, httpclient.WithQueryParam("limit", limit))
	}

	return call[binance.OrderBookResponse](ctx, http.MethodGet, "/depth", s, opts...)
}

func (s *service) RecentTrades(ctx context.Context, req *binance.RecentTradesRequest) ([]binance.Trade, error) {
	opts := []httpclient.RequestOption{}
	opts = append(opts, httpclient.WithQueryParam("symbol", req.Symbol))

	if req.Limit != 0 {
		limit := fmt.Sprintf("%d", req.Limit)
		opts = append(opts, httpclient.WithQueryParam("limit", limit))
	}

	res, err := call[[]binance.Trade](ctx, http.MethodGet, "/trades", s, opts...)
	if err != nil {
		return nil, err
	}

	return *res, err
}

func (s *service) OldTradesLookup(ctx context.Context, req *binance.OldTradesLookup) ([]binance.Trade, error) {
	opts := []httpclient.RequestOption{}
	opts = append(opts, httpclient.WithQueryParam("symbol", req.Symbol))

	if req.Limit != 0 {
		limit := fmt.Sprintf("%d", req.Limit)
		opts = append(opts, httpclient.WithQueryParam("limit", limit))
	}

	if req.FromId != 0 {
		fromId := fmt.Sprintf("%d", req.FromId)
		opts = append(opts, httpclient.WithQueryParam("fromId", fromId))
	}

	res, err := call[[]binance.Trade](ctx, http.MethodGet, "/historicalTrades", s, opts...)
	if err != nil {
		return nil, err
	}

	return *res, err
}

func call[T any](ctx context.Context, method, endpoint string, s *service, opts ...httpclient.RequestOption) (*T, error) {
	var entity T

	req, err := httpclient.NewRequest(method, fmt.Sprintf("%s%s", BASE_ENDPOINT, endpoint), opts...)
	if err != nil {
		return nil, err
	}

	res, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	fmt.Println("Body: ", string(res.Body()))
	if res.Body() != nil {
		err = json.Unmarshal(res.Body(), &entity)
		if err != nil {
			return nil, err
		}
	}

	return &entity, nil
}

func NewService(opts ...ServiceOpts) BinanceApiService {
	s := &service{}

	for _, opt := range opts {
		opt(s)
	}

	if s.client == nil {
		s.client = httpclient.NewClient()
	}

	return s
}
