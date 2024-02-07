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
    bare "github.com/tomphttp/bare-server-go/pkg/chi"
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

## Having CORS Issues?
- Update your main.go to the following:
```go
package main

import (
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/cors"
    "net/http"
    "fmt"
    bare "github.com/tomphttp/bare-server-go/pkg/chi"
)

func main() {
    chiRouter = chi.NewRouter()
    chiRouter.Use(middleware.Logger)
    chiRouter.Use(cors.Handler(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"*"},
        ExposedHeaders: []string{"*"},
        MaxAge: 300,
    }))
    router := bare.HandleBare("/bare/", chiRouter)
    router.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))
    })
    fmt.Println("Listening on http://localhost:8080")
    http.ListenAndServe(":8080", router)
}
```
