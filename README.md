## Coinprice 

Coinprice is a simple CLI tool written in Go that allows you to retrieve the price of a cryptocurrency.

It was created to test the [CryptoCompare API](https://www.cryptocompare.com/api/) and designed to be easily extended in order to support other APIs and output formatting (to be enabled via additional flags).

## Usage 

**Flags**

```
-coin   [Default: BTC] The symbol of the coin you want to get 
-in     [Default: USD] The currency the coin price will be expressed in 
```

**Examples**

Get the price of Ethereum in euros

```
coinprice -coin ETH -in EUR
```

Get the price of Bitcoin Cash in US dollars

```
coinprice -coin BCH -in USD
```
