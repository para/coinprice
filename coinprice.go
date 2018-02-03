package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Flag parsing
	coin := flag.String("coin", "BTC", "The symbol of the coin you want to get (e.g., BTC for Bitcoin).")
	in := flag.String("in", "USD", "The currency the coin price will be expressed in.")

	flag.Parse()

	// API invocation
	var api CoinAPI
	api = CryptoCompareAPI{}
	priceInfo, err := api.getPrice(*coin, *in)

	// Output
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error retrieving price info for coin %s in %s\n", *coin, *in)
		os.Exit(1)
	} else {
		var formatter PriceFormatter
		formatter = DefaultFormatter{}
		output := formatter.format(priceInfo)
		fmt.Println(output)
	}
}
