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
