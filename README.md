# go-yahoo-finance
A basic go package that you can use to query the Yahoo Finance API for stock pricing data. 

## Using
Add the package to your project using the go get command:

```
go get github.com/tonymackay/go-yahoo-finance
```

## Example
The code below shows a basic example of how to query the Yahoo Finance API and get the stock price for Tesla.

```
package main

import (
	"fmt"
	"log"

	"github.com/tonymackay/go-yahoo-finance"
)

func main() {
	result, err := yahoo.Quote("TSLA")
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Tesla Price: %v %s\n",
		result.QuoteSummary.Result[0].Price.RegularMarketPrice.Raw,
		result.QuoteSummary.Result[0].Price.Currency,
	)
}
```



## License
[MIT License](LICENSE)
