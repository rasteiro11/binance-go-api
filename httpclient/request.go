package httpclient

import (
	"fmt"
	"io"
	"net/url"
)

type request struct {
	mehtod  string
	payload io.Reader
	url     *url.URL
	header  map[string][]string
}

type Request interface {
	Method() string
	Payload() io.Reader
	URL() *url.URL
	Header() map[string][]string
}

type RequestOption func(*request)

var _ Request = (*request)(nil)

func (r *request) Method() string {
	return r.mehtod
}

func (r *request) Header() map[string][]string {
	return r.header
}

func (r *request) Payload() io.Reader {
	return r.payload
}

func (r *request) URL() *url.URL {
	return r.url
}

func WithQueryParam(key, value string) RequestOption {
	return func(r *request) {
		query := r.url.Query()
		query.Add(key, value)
		r.url.RawQuery = query.Encode()
	}
}

func WithHeader(header map[string][]string) RequestOption {
	return func(r *request) {
		for k, val := range header {
			r.header[k] = val
		}
	}
}

func WithPayload(payload io.Reader) RequestOption {
	return func(r *request) {
		r.payload = payload
	}
}

func NewRequest(method, endpoint string, opts ...RequestOption) (Request, error) {
	parsedUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	req := &request{
		mehtod: method,
		url:    parsedUrl,
		header: make(map[string][]string),
	}

	for _, opt := range opts {
		opt(req)
	}

	fmt.Println("HEADER")
	for k, val := range req.header {
		fmt.Println("KEY: ", k)
		fmt.Println("VAL: ", val)
	}

	fmt.Println("QUERY PARAMS")
	fmt.Println("QUERY: ", req.URL().Scheme, req.URL().Host, req.URL().RawQuery)
	for k, val := range req.url.Query() {
		fmt.Println("KEY: ", k)
		fmt.Println("VAL: ", val)
	}

	return req, nil
}
