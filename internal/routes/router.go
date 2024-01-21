package routes

import (
    "fmt"
    "net/http"
    "github.com/go-chi/chi/v5"
)

func Init(host string, port string, directory string) {
    chiRouter := chi.NewRouter()
    router := ChiRouter(directory, chiRouter)
    fmt.Println("Listening on http://" + host + ":" + port + directory)
    if host == "0.0.0.0" { fmt.Println("Also listening on http://localhost:" + port + directory) }
    http.ListenAndServe(host + ":" + port, router)
}
