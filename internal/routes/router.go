package routes

import (
    "fmt"
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/cors"

    //imports from project
    "github.com/Ruby-Network/bare-go/internal/utils"
    "github.com/Ruby-Network/bare-go/internal/v3"
)

func Init(host string, port string, directory string) {
    router := chi.NewRouter()
    router.Use(middleware.Logger)
    router.Use(cors.Handler(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"*"},
        ExposedHeaders: []string{"*"},
        MaxAge: 300,
    }))

    router.Handle(directory, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        e := utils.GetJson()
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(e))
    }))
    router.Handle(directory + "health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        e := `{"status": "ok"}`
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(e))
    }))

    router.Handle(directory + "v3/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        method := string(r.Method)
        userAgent := string(r.Header.Get("User-Agent"))
        headers := r.Header
        v3.Handler(method, userAgent, w, r, headers)
    }))

    fmt.Println("Server listening on http://" + host + ":" + port + directory)
    if (host == "0.0.0.0") { fmt.Println("Server also listening on http://localhost:" + port + directory) }
    http.ListenAndServe(host + ":" + port, router)
}
