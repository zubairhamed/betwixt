package server

import (
	"github.com/zubairhamed/betwixt/server/pages"
	. "github.com/zubairhamed/go-commons/network"
	"log"
	"runtime"
	"strconv"
)

func SetupHttpRoutes(server *DefaultServer) {
	http := server.httpServer

	http.NewRoute("/", METHOD_GET, handleHttpHome(server))
	http.NewRoute("/client/{client}/view", METHOD_GET, handleHttpViewClient(server))
	http.NewRoute("/client/{client}/delete", METHOD_GET, handleHttpDeleteClient(server))
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

func handleHttpDeleteClient(server *DefaultServer) RouteHandler {
	return func(r Request) Response {
		page := &pages.HomePage{}

		log.Println("Handle Deleting of Client")

		return &HttpResponse{
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

		type model struct {
			Clients 		 []*client
			ClientsCount 	 int
			MemUsage 		 string
			RequestCount 	 int
			ErrorsCount	 	 int
		}

		cl := []*client{}
		for _, v := range server.clients {
			c := &client{
				Endpoint:         v.GetName(),
				RegistrationID:   v.GetId(),
				RegistrationDate: v.GetRegistrationDate().Format("Jan 2, 2006, 3:04pm (SGT)"),
				LastUpdate:       v.LastUpdate().Format("Jan 2, 2006, 3:04pm (SGT)"),
			}
			cl = append(cl, c)
		}

		// Memory Usage
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)

		m := &model{
			ClientsCount: len(cl),
			Clients: cl,
			MemUsage: strconv.Itoa(int(mem.Alloc/1000)),
			RequestCount: server.stats.GetRequestsCount(),
			ErrorsCount: 0,

		}


		return &HttpResponse{
			TemplateModel: m,
			Payload:       NewBytesPayload(page.GetContent()),
		}
	}
}
