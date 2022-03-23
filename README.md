# go-winvps

A Winvps API client to interact with Winvps service

## Developer Documentation

The actual API Documentation available on this [link](https://winvps.fozzy.com/api/v2_docs).

## Usage

```go
import "github.com/FozzyHosting/go-winvps"

winClient, err := winvps.NewClient("token")
if err != nil {
  log.Fatalf("Failed to create client: %v", err)
}

machines, _, err := winClient.GetMachines()
```

### Examples

The [examples](examples) directory contains serveral examples of using this library.
