package yahoo

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestQuote(t *testing.T) {
	actual, err := Quote("TSLA")
	assert.Nil(t, err)
	assert.Equal(t, "Tesla, Inc.", actual.QuoteSummary.Result[0].Price.ShortName)
}
