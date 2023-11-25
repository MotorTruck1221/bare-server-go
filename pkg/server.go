package bare_go

import (
    "github.com/go-chi/chi/v5"
    "github.com/Ruby-Network/bare-go/internal/routes"
)

func CreateBare(directory string) *chi.Mux {
    router := routes.Router(directory)
    return router
}
