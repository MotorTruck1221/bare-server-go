package routes 

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    //imports from project
    "github.com/tomphttp/bare-server-go/internal/utils"
    "github.com/tomphttp/bare-server-go/internal/v3"
)

func ChiRouter(directory string, router *chi.Mux) *chi.Mux { 
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

    return router
}
