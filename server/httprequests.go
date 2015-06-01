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
	return func(r Request) Response {

		page := &pages.HomePage{}

		type client struct {
			Endpoint         string
			RegistrationID   string
			RegistrationDate string
			LastUpdate       string
		}

		model := []client{
			{
				Endpoint: "abc",
			}, {
				Endpoint: "def",
			},
		}

		return &HttpResponse{
			TemplateModel: model,
			Payload:       NewBytesPayload(page.GetContent()),
		}
	}
}
