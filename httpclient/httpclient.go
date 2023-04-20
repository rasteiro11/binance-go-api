package httpclient

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

type httpclient struct {
	client *http.Client
}

type HttpClient interface {
	Do(ctx context.Context, req Request) (Response, error)
}

type HttpClientOpts func(*httpclient)

var _ HttpClient = (*httpclient)(nil)

func WithHttpClient(client *http.Client) HttpClientOpts {
	return func(h *httpclient) {
		h.client = client
	}
}

func (c *httpclient) Do(ctx context.Context, req Request) (Response, error) {
	httpReq, err := http.NewRequestWithContext(ctx, req.Method(), req.URL().String(), req.Payload())
	if err != nil {
		return nil, err
	}

	for k, val := range req.Header() {
		for _, v := range val {
			httpReq.Header.Add(k, v)
		}
	}

	res, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err := res.Body.Close(); err != nil {
		return nil, err
	}

	return &response{
		contentLength: int(res.ContentLength),
		statusCode:    res.StatusCode,
		header:        res.Header.Clone(),
		body:          body,
	}, nil
}

func newDefaultHttpClient() *http.Client {
	return &http.Client{
		Timeout:   time.Second * 420,
		Transport: http.DefaultClient.Transport,
	}
}

func NewClient(opts ...HttpClientOpts) HttpClient {
	c := &httpclient{}

	for _, opt := range opts {
		opt(c)
	}

	if c.client == nil {
		c.client = newDefaultHttpClient()
	}

	return c
}
