package bare_go_chi

import (
    "github.com/go-chi/chi/v5"
    "github.com/ruby-network/bare-server-go/internal/routes"
)

func HandleBare(directory string, chiRouter *chi.Mux) *chi.Mux {
    router := routes.ChiRouter(directory, chiRouter)
    return router
}
