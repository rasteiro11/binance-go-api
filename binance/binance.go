package binance

import (
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

type TimeResponse struct {
	ServerTime int64 `json:"serverTime"`
}

type BinanceApiService interface {
	Time(ctx context.Context) (*TimeResponse, error)
}

type ServiceOpts func(*service)

func (s *service) Time(ctx context.Context) (*TimeResponse, error) {
	req, err := httpclient.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", BASE_ENDPOINT, "/time"))
	if err != nil {
		return nil, err
	}

	res, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(res.Body()))

	timeResponse := &TimeResponse{}

	err = json.Unmarshal(res.Body(), timeResponse)
	if err != nil {
		return nil, err
	}

	return timeResponse, nil
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
