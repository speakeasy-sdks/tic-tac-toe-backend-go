<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"log"
	tictactoebackends "tic-tac-toe-backends/v3"
)

func main() {
	s := tictactoebackends.New()

	ctx := context.Background()
	res, err := s.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Body != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->