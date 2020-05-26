package mlit

import (
	"strings"

	"github.com/go-numb/go-mlit-estate/areas"
	"github.com/go-numb/go-mlit-estate/prices"
)

const (
	API = "https://www.land.mlit.go.jp/webland/api"
)

type Client struct {
	Prices prices.Client
	Areas  areas.Client
}

func New(isEnglish bool) *Client {
	api := API
	if isEnglish {
		api = strings.Replace(api, "webland", "webland_english", 1)
	}

	return &Client{
		Prices: prices.Client{
			Endpoint: api + prices.Endpoint,
		},
		Areas: areas.Client{
			Endpoint: api + areas.Endpoint,
		},
	}
}
