// Copyright 2020 Tony Mackay.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

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
