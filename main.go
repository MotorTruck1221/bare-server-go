package main

import (
    "fmt"
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/cors"

    //imports from this project
    "github.com/Ruby-Network/bare-go/internal/utils"
    "github.com/Ruby-Network/bare-go/internal/v3"
)

func main() {
    router := chi.NewRouter()
    router.Use(middleware.Logger)
    router.Use(cors.Handler(cors.Options{
         AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"*"},
        ExposedHeaders: []string{"*"},
        MaxAge: 300,
    }))

    router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        e := utils.GetJson()
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(e))
    }))
    router.Handle("/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        e := `{"status": "ok"}`
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(e))
    }))

    router.Handle("/v3/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        method := string(r.Method)
        userAgent := string(r.Header.Get("User-Agent"))
        headers := r.Header
        v3.Handler(method, userAgent, w, r, headers)
    }))

    fmt.Println("Server is running on port 8081")
    http.ListenAndServe(":8081", router)
}
