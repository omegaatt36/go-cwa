# go-cwa

An idiomatic, zero-dependency Go SDK for the [Taiwan Central Weather Administration (CWA)](https://opendata.cwa.gov.tw/) Open Data API.

## Features

- **Zero External Dependencies**: Built entirely with the Go standard library.
- **Idiomatic Go**: Uses the Functional Options pattern for configuration.
- **Type-Safe**: Provides constants for Taiwan counties and cities.
- **Simplified API**: Descriptive method names for common datasets like 36-hour forecasts and township forecasts.

## Installation

```bash
go get github.com/omegaatt36/go-cwa
```

## Usage

### 1. Initialize Client

```go
package main

import (
    "github.com/omegaatt36/go-cwa"
)

func main() {
    client := cwa.NewClient("YOUR_CWA_API_KEY")
    // ...
}
```

### 2. Get 36-Hour Forecast

```go
forecast, err := client.Get36hForecast(context.Background(), &cwa.Forecast36hParams{
    LocationNames: []cwa.County{cwa.TaipeiCity, cwa.NewTaipeiCity},
})
if err != nil {
    panic(err)
}

if forecast.IsSuccess() {
    for _, loc := range forecast.Records.Location {
        fmt.Printf("Location: %s\n", loc.LocationName)
    }
}
```

### 3. Get Township Forecast

```go
resp, err := client.GetTownshipForecast(context.Background(), &cwa.TownshipForecastParams{
	County: cwa.TaipeiCity,
	Period: cwa.ThreeDays,
})

if err != nil {
    panic(err)
}
```

## Supported Datasets

- `F-C0032-001`: 36-hour weather forecast.
- `F-D0047-*`: Township forecasts (3-day and 1-week).

## License

MIT (See [LICENSE](LICENSE) file).
