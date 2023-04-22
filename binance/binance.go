package binance

import (
	binance "binance-go-api/binance/market_data"
	"binance-go-api/httpclient"
	"binance-go-api/validator"
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

	if len(req.Symbols) == 1 {
		opts = append(opts, httpclient.WithQueryParam("symbol", req.Symbols[0]))
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

	if err := validator.IsValid(req); err != nil {
		return nil, err
	}

	opts = append(opts, httpclient.WithQueryParam("symbol", req.Symbol))

	if req.Limit != 0 {
		limit := fmt.Sprintf("%d", req.Limit)
		opts = append(opts, httpclient.WithQueryParam("limit", limit))
	}

	return call[binance.OrderBookResponse](ctx, http.MethodGet, "/depth", s, opts...)
}

func (s *service) RecentTrades(ctx context.Context, req *binance.RecentTradesRequest) ([]binance.Trade, error) {
	opts := []httpclient.RequestOption{}

	if err := validator.IsValid(req); err != nil {
		return nil, err
	}

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

func (s *service) OldTradesLookup(ctx context.Context, req *binance.OldTradesLookupRequest) ([]binance.Trade, error) {
	opts := []httpclient.RequestOption{}

	if err := validator.IsValid(req); err != nil {
		return nil, err
	}

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

func (s *service) Klines(ctx context.Context, req *binance.KlinesRequest) ([]binance.Kline, error) {
	opts := []httpclient.RequestOption{}

	if err := validator.IsValid(req); err != nil {
		return nil, err
	}

	opts = append(opts, httpclient.WithQueryParam("symbol", req.Symbol), httpclient.WithQueryParam("interval", req.Interval))

	if req.StartTime != 0 {
		startTime := fmt.Sprintf("%d", req.StartTime)
		opts = append(opts, httpclient.WithQueryParam("startTime", startTime))
	}

	if req.EndTime != 0 {
		endTime := fmt.Sprintf("%d", req.EndTime)
		opts = append(opts, httpclient.WithQueryParam("endTime", endTime))
	}

	if req.Limit != 0 {
		limit := fmt.Sprintf("%d", req.Limit)
		opts = append(opts, httpclient.WithQueryParam("limit", limit))
	}

	res, err := call[[]binance.Kline](ctx, http.MethodGet, "/klines", s, opts...)
	if err != nil {
		return nil, err
	}

	return *res, err
}

func (s *service) UIKlines(ctx context.Context, req *binance.KlinesRequest) ([]binance.Kline, error) {
	opts := []httpclient.RequestOption{}

	if err := validator.IsValid(req); err != nil {
		return nil, err
	}

	opts = append(opts, httpclient.WithQueryParam("symbol", req.Symbol), httpclient.WithQueryParam("interval", req.Interval))

	if req.StartTime != 0 {
		startTime := fmt.Sprintf("%d", req.StartTime)
		opts = append(opts, httpclient.WithQueryParam("startTime", startTime))
	}

	if req.EndTime != 0 {
		endTime := fmt.Sprintf("%d", req.EndTime)
		opts = append(opts, httpclient.WithQueryParam("endTime", endTime))
	}

	if req.Limit != 0 {
		limit := fmt.Sprintf("%d", req.Limit)
		opts = append(opts, httpclient.WithQueryParam("limit", limit))
	}

	res, err := call[[]binance.Kline](ctx, http.MethodGet, "/uiKlines", s, opts...)
	if err != nil {
		return nil, err
	}

	return *res, err
}

func (s *service) AveragePrice(ctx context.Context, req *binance.AveragePriceRequest) (*binance.AveragePriceResponse, error) {
	if err := validator.IsValid(req); err != nil {
		return nil, err
	}
	return call[binance.AveragePriceResponse](ctx, http.MethodGet, "/avgPrice", s, httpclient.WithQueryParam("symbol", req.Symbol))
}

func (s *service) PriceChange24H(ctx context.Context, req *binance.PriceChangeRequest) ([]binance.PriceChangeResponse, error) {
	opts := []httpclient.RequestOption{}

	if req.Type != "" {
		opts = append(opts, httpclient.WithQueryParam("type", req.Type))
	}

	if len(req.Symbols) == 1 {
		opts = append(opts, httpclient.WithQueryParam("symbol", req.Symbols[0]))
		res, err := call[binance.PriceChangeResponse](ctx, http.MethodGet, "/ticker/24hr", s, opts...)
		if err != nil {
			return nil, err
		}

		return []binance.PriceChangeResponse{*res}, nil
	}

	opts = append(opts, httpclient.WithQueryListParam("symbols", req.Symbols))
	res, err := call[[]binance.PriceChangeResponse](ctx, http.MethodGet, "/ticker/24hr", s, opts...)
	if err != nil {
		return nil, err
	}

	return *res, nil
}

func (s *service) SymbolPriceTicker(ctx context.Context, req *binance.SymbolPriceTickerRequest) ([]binance.SymbolPriceTicker, error) {
	if len(req.Symbols) == 1 {
		res, err := call[binance.SymbolPriceTicker](ctx, http.MethodGet, "/ticker/price", s, httpclient.WithQueryParam("symbol", req.Symbols[0]))
		if err != nil {
			return nil, err
		}

		return []binance.SymbolPriceTicker{*res}, nil
	}

	res, err := call[[]binance.SymbolPriceTicker](ctx, http.MethodGet, "/ticker/price", s, httpclient.WithQueryListParam("symbols", req.Symbols))
	if err != nil {
		return nil, err
	}

	return *res, nil
}

func (s *service) SymbolOrderBookTicker(ctx context.Context, req *binance.SymbolOrderBookTickerRequest) ([]binance.SymbolOrderBookTicker, error) {
	if len(req.Symbols) == 1 {
		res, err := call[binance.SymbolOrderBookTicker](ctx, http.MethodGet, "/ticker/bookTicker", s, httpclient.WithQueryParam("symbol", req.Symbols[0]))
		if err != nil {
			return nil, err
		}

		return []binance.SymbolOrderBookTicker{*res}, nil
	}

	res, err := call[[]binance.SymbolOrderBookTicker](ctx, http.MethodGet, "/ticker/bookTicker", s, httpclient.WithQueryListParam("symbols", req.Symbols))
	if err != nil {
		return nil, err
	}

	return *res, nil
}

func (s *service) RollingWindowPriceChange(ctx context.Context, req *binance.RollingWindowPriceChangeRequest) ([]binance.RollingWindowPriceChange, error) {
	opts := []httpclient.RequestOption{}

	if err := validator.IsValid(req); err != nil {
		return nil, err
	}

	if req.Type != "" {
		opts = append(opts, httpclient.WithQueryParam("type", req.Type))
	}

	if req.WindowSize != "" {
		opts = append(opts, httpclient.WithQueryParam("windowSize", req.WindowSize))
	}

	if len(req.Symbols) == 1 {
		opts = append(opts, httpclient.WithQueryParam("symbol", req.Symbols[0]))
		res, err := call[binance.RollingWindowPriceChange](ctx, http.MethodGet, "/ticker", s, opts...)
		if err != nil {
			return nil, err
		}

		return []binance.RollingWindowPriceChange{*res}, nil
	}

	opts = append(opts, httpclient.WithQueryListParam("symbols", req.Symbols))
	res, err := call[[]binance.RollingWindowPriceChange](ctx, http.MethodGet, "/ticker", s, opts...)
	if err != nil {
		return nil, err
	}

	return *res, nil
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
