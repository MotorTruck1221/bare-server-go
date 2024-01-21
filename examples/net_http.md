# Usage with basic net/http

## Setup
```bash
go mod init your-module-name
```
---

> main.go
```go
package main

import (
    "net/http"
    "fmt"
    bare "github.com/ruby-network/bare-go/pkg/http"
)

func main() {
    router := bare.HandleBare("/bare/", http.NewServeMux())
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World"))
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

## Having CORS issues?
- Update your `main.go` to include the following:
```go
import (
    "net/http"
    "fmt"
    bare "github.com/ruby-network/bare-go/pkg/http"
    "github.com/rs/cors"
)

func main() {
    cors := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"*"},
        ExposedHeaders: []string{"*"},
    })
    router := bare.HandleBare("/bare/", http.NewServeMux())
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World"))
    })
    fmt.Println("Listening on http://localhost:8080")
    http.ListenAndServe(":8080", cors.Handler(router))
}
```

