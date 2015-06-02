package server

import (
	. "github.com/zubairhamed/go-commons/network"
	"github.com/zubairhamed/go-lwm2m/server/pages"
)

func SetupHttpRoutes(server *DefaultServer) {
	http := server.httpServer

	http.NewRoute("/", METHOD_GET, handleHttpHome(server))
	http.NewRoute("/reg/{client}", METHOD_GET, handleHttpViewClient(server))
}

func handleHttpViewClient(server *DefaultServer) RouteHandler {
	return func(r Request) Response {
		page := &pages.ClientDetailPage{}

		type clientdetails struct {
		}

		model := clientdetails{}

		return &HttpResponse{
			TemplateModel: model,
			Payload:       NewBytesPayload(page.GetContent()),
		}
	}
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

		model := []client{}

		for _, v := range server.clients {
			c := client{
				Endpoint:         v.GetName(),
				RegistrationID:   v.GetId(),
				RegistrationDate: v.GetRegistrationDate().String(),
				LastUpdate: 	  v.LastUpdate().String(),
			}
			model = append(model, c)
		}

		return &HttpResponse{
			TemplateModel: model,
			Payload:       NewBytesPayload(page.GetContent()),
		}
	}
}
