package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// CryptoCompareAPI - Price API offered by cryptocompare.com
type CryptoCompareAPI struct{}

func (c CryptoCompareAPI) getPrice(coin string, priceIn string) (PriceInfo, error) {
	url := fmt.Sprintf("https://min-api.cryptocompare.com/data/price?fsym=%s&tsyms=%s", coin, priceIn)
	response, err := http.Get(url)

	if err != nil {
		return PriceInfo{}, err
	}

	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return PriceInfo{}, err
	}

	var r map[string]float32
	err = json.Unmarshal(responseData, &r)

	if err != nil {
		return PriceInfo{}, err
	}

	if price, exists := r[priceIn]; exists {
		var priceInfo PriceInfo
		priceInfo.Coin = coin
		priceInfo.Price = price
		priceInfo.PriceIn = priceIn
		priceInfo.Time = time.Now()
		return priceInfo, nil
	}

	return PriceInfo{}, fmt.Errorf("No price info for currency: %s", priceIn)
}
