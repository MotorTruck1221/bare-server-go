package bare_go

import (
    "github.com/Ruby-Network/bare-go/internal/routes"
)

func Start(host string, port string, directory string) {
    routes.Init(host, port, directory)
}
