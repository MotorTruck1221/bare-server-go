package main

import (
    "fmt"
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/cors"

    //imports from this project
    "github.com/Ruby-Network/bare-go/utils"
    "github.com/Ruby-Network/bare-go/v3"
)

func main() {
    router := chi.NewRouter()
    router.Use(middleware.Logger)
    router.Use(cors.Handler(cors.Options{
        //ALLOW ALL ORIGINS
        AllowedOrigins: []string{"*"},
    }))

    router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        //content-type: application/json
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
        v3.Handler()
    }))

    fmt.Println("Server is running on port 8080")
    http.ListenAndServe(":8080", router)
}
