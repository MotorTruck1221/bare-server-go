package bare_go_gorilla

import (
   "github.com/gorilla/mux" 
   "github.com/tomphttp/bare-server-go/internal/routes"
)

func HandleBare(directory string, gorillaRouter *mux.Router) *mux.Router {
    router := routes.GorillaRouter(directory, gorillaRouter)
    return router
}
