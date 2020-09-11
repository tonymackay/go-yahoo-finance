package yahoo

import (
	"errors"
)

// QuoteResult contains
type QuoteResult struct {
	Name  string
	Price float64
}

// Quote returns stock price
func Quote(symbol string) (QuoteResult, error) {
	return QuoteResult{}, errors.New("not implemented")
}
