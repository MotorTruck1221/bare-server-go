package routes

import (
    "fmt"
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/cors"
)

//router that is used by the CLI

func Init(host string, port string, directory string) {
    chiRouter := chi.NewRouter()
    chiRouter.Use(middleware.Logger)
    chiRouter.Use(cors.Handler(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"*"},
        ExposedHeaders: []string{"*"},
        MaxAge: 300,
    }))

    router := ChiRouter(directory, chiRouter)
    fmt.Println("Listening on http://" + host + ":" + port + directory)
    if host == "0.0.0.0" { fmt.Println("Also listening on http://localhost:" + port + directory) }
    http.ListenAndServe(host + ":" + port, router)
}
