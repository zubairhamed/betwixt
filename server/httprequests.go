package server
import (
    . "github.com/zubairhamed/go-commons/network"
    "github.com/zubairhamed/go-lwm2m/server/pages"
)

func SetupHttpRoutes(server *DefaultServer) {
    http := server.httpServer

    http.NewRoute("/", METHOD_GET, handleHttpHome(server))
}

func handleHttpHome(server *DefaultServer) RouteHandler {
    return func (r Request) (Response) {

        page := &pages.HomePage{}

        model := struct {
            Title   string
            Content string
        }{
            "Page Title",
            "Page Content",
        }

        return &HttpResponse{
            TemplateModel: model,
            Payload: NewBytesPayload(page.GetContent()),
        }
    }
}
