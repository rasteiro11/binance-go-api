package main

import (
	"binance-go-api/binance"
	"binance-go-api/httpclient"
	"context"
	"log"
	"net/http"
)

func main() {
	_, _ = httpclient.NewRequest(http.MethodGet, "http://gamer.com",
		httpclient.WithQueryParam("GAMER", "42069"),
		httpclient.WithHeader(map[string][]string{"Authorization": {"GAMER"}}))

	ctx := context.TODO()
	svcBinance := binance.NewService()

	_, err := svcBinance.Time(ctx)
	if err != nil {
		log.Fatalf("ERROR: %+v", err)
	}

}
