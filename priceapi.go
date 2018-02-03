package main

import (
	"time"
)

// CoinAPI - A service that allows you to retrieve the price of a coin
type CoinAPI interface {
	getPrice(coin string, priceIn string) (PriceInfo, error)
}

// PriceInfo - A struct containing price info for a coin
type PriceInfo struct {
	Coin    string
	Price   float32
	PriceIn string
	Time    time.Time
}
