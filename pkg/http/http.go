package bare_go_http

import (
    "net/http"
    "github.com/tomphttp/bare-server-go/internal/routes"
)

func HandleBare(directory string, router *http.ServeMux) *http.ServeMux {
    router = routes.NetHttp(directory, router)
    return router
}
