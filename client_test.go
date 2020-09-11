package yahoo

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestQuote(t *testing.T) {
	actual, err := Quote("TSLA")
	assert.Nil(t, err)
	assert.Equal(t, "Tesla", actual.Name)
}
