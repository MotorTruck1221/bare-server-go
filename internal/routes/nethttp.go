package routes


import (
    "net/http"
    //imports from project
    "github.com/tomphttp/bare-server-go/internal/utils"
    "github.com/tomphttp/bare-server-go/internal/v3"
)

func NetHttp(directory string, router *http.ServeMux) *http.ServeMux { 
    router.HandleFunc(directory, func(w http.ResponseWriter, r *http.Request) {
        e := utils.GetJson()
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(e))
    })
    router.HandleFunc(directory + "health", func(w http.ResponseWriter, r *http.Request) {
        e := `{"status": "ok"}`
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(e))
    })

    router.HandleFunc(directory + "v3/", func(w http.ResponseWriter, r *http.Request) {
        method := string(r.Method)
        userAgent := string(r.Header.Get("User-Agent"))
        headers := r.Header
        v3.Handler(method, userAgent, w, r, headers)
    })

    return router
}
