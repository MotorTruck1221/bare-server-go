# Usage with Chi

## Setup
```bash
go mod init your-module-name
```
---

> main.go
```go
package main

import (
    "github.com/go-chi/chi/v5"
    "net/http"
    "fmt"
    bare "github.com/ruby-network/bare-server-go/pkg/chi"
)

func main() {
    router := bare.HandleBare("/bare/", chi.NewRouter())
    router.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))
    })
    fmt.Println("Listening on http://localhost:8080")
    http.ListenAndServe(":8080", router)
}
```
---
## Run
```bash
go get && go run main.go
```
And navigate to http://localhost:8080/bare/
