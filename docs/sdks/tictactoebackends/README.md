# TicTacToeBackends SDK


## Overview

Game Engine API for Tic Tac Toe: Game Engine API for Tic Tac Toe

### Available Operations

* [Get](#get) - Root endpoint.
* [GetVersion](#getversion) - Root endpoint.
* [PutGames](#putgames) - Games endpoint. Creates the next game state from the previous game state.

## Get

<br/>Returns the package name and version.<br/><br/>

### Example Usage

```go
package main

import(
	"context"
	"log"
	tictactoebackends "tic-tac-toe-backends/v2"
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetResponse](../../pkg/models/operations/getresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetVersion

<br/>Returns the package name and version.<br/><br/>

### Example Usage

```go
package main

import(
	"context"
	"log"
	tictactoebackends "tic-tac-toe-backends/v2"
)

func main() {
    s := tictactoebackends.New()

    ctx := context.Background()
    res, err := s.GetVersion(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetVersionResponse](../../pkg/models/operations/getversionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PutGames

<br/>Accepts a GameState and Move.<br/><br/>Returns a Move including the before and after GameStates.<br/>

### Example Usage

```go
package main

import(
	"context"
	"log"
	tictactoebackends "tic-tac-toe-backends/v2"
	"tic-tac-toe-backends/v2/pkg/models/shared"
)

func main() {
    s := tictactoebackends.New()

    ctx := context.Background()
    res, err := s.PutGames(ctx, []byte("0x8BCDbF9B8f"))
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]byte](../../.md)                                   | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PutGamesResponse](../../pkg/models/operations/putgamesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
