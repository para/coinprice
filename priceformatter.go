package main

// PriceFormatter - A component that can format a PriceInfo instance
// in a way that can be displayed on the screen
type PriceFormatter interface {
	format(info PriceInfo) string
}
