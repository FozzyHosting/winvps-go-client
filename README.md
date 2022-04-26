# winvps-go-client

A Winvps API client to interact with Winvps service  

[![Test Status](https://github.com/FozzyHosting/winvps-go-client/actions/workflows/test.yml/badge.svg)](https://github.com/FozzyHosting/winvps-go-client/actions/workflows/test.yml)

## Developer Documentation

The actual API Documentation available on this [link](https://winvps.fozzy.com/api/v2_docs).

## Usage

```go
import "github.com/fozzyhosting/winvps-go-client"

winClient, err := winvps.NewClient("token")
if err != nil {
  log.Fatalf("Failed to create client: %v", err)
}

machines, _, err := winClient.GetMachines()
```

### Examples

The [examples](examples) directory contains serveral examples of using this library.
