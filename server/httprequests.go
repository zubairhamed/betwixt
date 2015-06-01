package server
import (
    . "github.com/zubairhamed/go-commons/network"
    "log"
)

func SetupHttpRoutes(server *DefaultServer) {
    http := server.httpServer

    http.NewRoute("/", METHOD_GET, handleHttpHome(server))
}

func handleHttpHome(server *DefaultServer) RouteHandler {
    return func (r Request) (Response) {
        log.Println("Handle Home")

        return nil
    }
}
