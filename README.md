# Syniac Go SDK

A simple Go SDK for interacting with the Syniac Cloud API.

## 📦 Installation

```bash
go get github.com/yourorg/syniac-sdk-go
```

## 🧪 Example Usage

```go
package main

import (
    "fmt"
    "log"
    "github.com/yourorg/syniac-sdk-go/syniac"
)

func main() {
    client := syniac.NewClient("your_api_key")

    instances, err := client.ListInstances()
    if err != nil {
        log.Fatal(err)
    }

    for _, inst := range instances {
        fmt.Println(inst.ID, inst.Name, inst.Status)
    }
}
```

## 🔐 Authentication

API Key is passed via the `Authorization: Bearer <API_KEY>` header.

## 📘 Endpoints Supported

- `GET /v1/instances`
- `GET /v1/instances/{id}`

## ✅ License

MIT
