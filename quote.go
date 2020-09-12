// Copyright 2020 Tony Mackay.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package yahoo

// QuoteResult type contains data returned from the Yahoo Finance Quote Endpoint
type QuoteResult struct {
	QuoteSummary struct {
		Result []struct {
			Price struct {
				MaxAge          int `json:"maxAge"`
				PreMarketChange struct {
				} `json:"preMarketChange"`
				PreMarketPrice struct {
				} `json:"preMarketPrice"`
				PreMarketSource         string `json:"preMarketSource"`
				PostMarketChangePercent struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"postMarketChangePercent"`
				PostMarketChange struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"postMarketChange"`
				PostMarketTime  int `json:"postMarketTime"`
				PostMarketPrice struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"postMarketPrice"`
				PostMarketSource           string `json:"postMarketSource"`
				RegularMarketChangePercent struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketChangePercent"`
				RegularMarketChange struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketChange"`
				RegularMarketTime int `json:"regularMarketTime"`
				PriceHint         struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"priceHint"`
				RegularMarketPrice struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketPrice"`
				RegularMarketDayHigh struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketDayHigh"`
				RegularMarketDayLow struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketDayLow"`
				RegularMarketVolume struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"regularMarketVolume"`
				AverageDailyVolume10Day struct {
				} `json:"averageDailyVolume10Day"`
				AverageDailyVolume3Month struct {
				} `json:"averageDailyVolume3Month"`
				RegularMarketPreviousClose struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketPreviousClose"`
				RegularMarketSource string `json:"regularMarketSource"`
				RegularMarketOpen   struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketOpen"`
				StrikePrice struct {
				} `json:"strikePrice"`
				OpenInterest struct {
				} `json:"openInterest"`
				Exchange              string      `json:"exchange"`
				ExchangeName          string      `json:"exchangeName"`
				ExchangeDataDelayedBy int         `json:"exchangeDataDelayedBy"`
				MarketState           string      `json:"marketState"`
				QuoteType             string      `json:"quoteType"`
				Symbol                string      `json:"symbol"`
				UnderlyingSymbol      interface{} `json:"underlyingSymbol"`
				ShortName             string      `json:"shortName"`
				LongName              string      `json:"longName"`
				Currency              string      `json:"currency"`
				QuoteSourceName       string      `json:"quoteSourceName"`
				CurrencySymbol        string      `json:"currencySymbol"`
				FromCurrency          interface{} `json:"fromCurrency"`
				ToCurrency            interface{} `json:"toCurrency"`
				LastMarket            interface{} `json:"lastMarket"`
				Volume24Hr            struct {
				} `json:"volume24Hr"`
				VolumeAllCurrencies struct {
				} `json:"volumeAllCurrencies"`
				CirculatingSupply struct {
				} `json:"circulatingSupply"`
				MarketCap struct {
					Raw     int64  `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"marketCap"`
			} `json:"price"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"quoteSummary"`
}
