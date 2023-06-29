<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"tic-tac-toe-backends"
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
<!-- End SDK Example Usage -->