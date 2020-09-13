// Copyright 2020 Tony Mackay.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package yahoo

import (
	"reflect"
	"testing"
)

func TestQuote(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "TSLA", want: "Tesla, Inc."},
		{input: "GOOG", want: "Alphabet Inc."},
	}

	for _, tc := range tests {
		quote, _ := Quote(tc.input)
		got := quote.QuoteSummary.Result[0].Price.ShortName
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
