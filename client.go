// Copyright 2020 Tony Mackay.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package yahoo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const baseURL = "https://query2.finance.yahoo.com/v10/finance/quoteSummary/"

// Quote returns stock price
func Quote(symbol string) (QuoteResult, error) {
	result := QuoteResult{}
	quoteEndpoint := baseURL + strings.ToUpper(symbol) + "?modules=price"

	resp, err := http.Get(quoteEndpoint)
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return result, readErr
	}

	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		return result, jsonErr
	}

	return result, nil
}
