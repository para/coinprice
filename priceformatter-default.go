package main

import "fmt"

// DefaultFormatter - The default formatter for PriceInfo
type DefaultFormatter struct{}

func (f DefaultFormatter) format(info PriceInfo) string {
	time := info.Time.Format("2006-01-02 15:04 MST")
	output := fmt.Sprintf("%s: %0.2f %s (%s)", info.Coin, info.Price, info.PriceIn, time)
	return output
}
