package server

import (
	"github.com/zubairhamed/betwixt/core/server/pages"
	. "github.com/zubairhamed/go-commons/network"
	"log"
	"runtime"
	"strconv"
	"github.com/zubairhamed/betwixt/core/server/pages/models"
)

func SetupHttpRoutes(server *DefaultServer) {
	http := server.httpServer

	// Pages
	http.NewRoute("/", METHOD_GET, handleHttpHome(server))
	http.NewRoute("/client/{client}/view", METHOD_GET, handleHttpViewClient(server))
	http.NewRoute("/client/{client}/delete", METHOD_GET, handleHttpDeleteClient(server))

	// APIs
	http.NewRoute("/api/clients", METHOD_GET, func(r Request) Response {
		cl := []models.ClientModel{}
		for _, v := range server.clients {

			objs := make(map[string]models.ObjectModel)
			for key, val := range v.GetObjects() {
				objectModel := models.ObjectModel{
					Instances: val.GetInstances(),
					Definition: val.GetDefinition(),
				}
				typeKey := strconv.Itoa(int(key))
				objs[typeKey] = objectModel
			}

			c := models.ClientModel{
				Endpoint:         v.GetName(),
				RegistrationID:   v.GetId(),
				RegistrationDate: v.GetRegistrationDate().Format("Jan 2, 2006, 3:04pm (SGT)"),
				LastUpdate:       v.LastUpdate().Format("Jan 2, 2006, 3:04pm (SGT)"),
				Objects: 		  objs,
			}
			cl = append(cl, c)
		}

		return &HttpResponse{
			Payload: NewJsonPayload(cl),
		}
	})

	http.NewRoute("/api/server/stats", METHOD_GET, func(r Request) Response {
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)

		clientsCount := len(server.clients)

		model := &models.StatsModel{
			ClientsCount: clientsCount,
			MemUsage: strconv.Itoa(int(mem.Alloc / 1000)),
			Requests: server.stats.GetRequestsCount(),
			Errors: 0,
		}

		return &HttpResponse{
			Payload: NewJsonPayload(model),
		}
	})

	// Get Message, Logs
	http.NewRoute("/api/server/{client}/messages", METHOD_GET, func(r Request) Response {
		return &HttpResponse{
			Payload: NewJsonPayload(""),
		}
	})

	// Read
	http.NewRoute("/api/clients/{client}", METHOD_GET, func(r Request) Response {
		req := r.(*HttpRequest)
		clientId := req.GetAttribute("client")

		v := server.GetRegisteredClient(clientId)
		if v == nil {

		}

		objs := make(map[string]models.ObjectModel)
		for key, val := range v.GetObjects() {
			objectModel := models.ObjectModel{
				Instances: val.GetInstances(),
				Definition: val.GetDefinition(),
			}
			typeKey := strconv.Itoa(int(key))
			objs[typeKey] = objectModel
		}

		c := models.ClientModel{
			Endpoint:         v.GetName(),
			RegistrationID:   v.GetId(),
			RegistrationDate: v.GetRegistrationDate().Format("Jan 2, 2006, 3:04pm (SGT)"),
			LastUpdate:       v.LastUpdate().Format("Jan 2, 2006, 3:04pm (SGT)"),
			Objects: 		  objs,
		}

		return &HttpResponse{
			Payload: NewJsonPayload(c),
		}
	})

	http.NewRoute("/api/clients/{client}/{object}/{instance}/{resource}", METHOD_GET, func(r Request) Response {
		req := r.(*HttpRequest)
		clientId := req.GetAttribute("client")
		object := req.GetAttributeAsInt("object")
		instance := req.GetAttributeAsInt("instance")
		resource := req.GetAttributeAsInt("resource")
		cli := server.GetRegisteredClient(clientId)

		cli.Read(object, instance, resource)
		return &HttpResponse{
			Payload: NewJsonPayload(cli),
		}
	})

	http.NewRoute("/api/clients/{client}/{object}/{instance}", METHOD_GET, func(r Request) Response {
		return &HttpResponse{
			Payload: NewJsonPayload(""),
		}
	})

	// Write
	http.NewRoute("/api/clients/{client}/{object}/{instance}/{resource}", METHOD_PUT, func(r Request) Response {
		return &HttpResponse{
			Payload: NewJsonPayload(""),
		}
	})

	http.NewRoute("/api/clients/{client}/{object}/{instance}", METHOD_PUT, func(r Request) Response {
		return &HttpResponse{
			Payload: NewJsonPayload(""),
		}
	})

	// Delete
	http.NewRoute("/api/clients/{client}/{object}/{instance}", METHOD_DELETE, func(r Request) Response {
		return &HttpResponse{
			Payload: NewJsonPayload(""),
		}
	})

	// Observe
	http.NewRoute("/api/clients/{client}/{object}/{instance}/{resource}/observe", METHOD_POST, func(r Request) Response {
		return &HttpResponse{
			Payload: NewJsonPayload(""),
		}
	})

	// Cancel Observe
	http.NewRoute("/api/clients/{client}/{object}/{instance}/{resource}/observe", METHOD_DELETE, func(r Request) Response {
		return &HttpResponse{
			Payload: NewJsonPayload(""),
		}
	})

	// Execute
	http.NewRoute("/api/clients/{client}/{object}/{instance}/{resource}", METHOD_POST, func(r Request) Response {
		return &HttpResponse{
			Payload: NewJsonPayload(""),
		}
	})

	// Create
	http.NewRoute("/api/clients/{client}/{object}/{instance}", METHOD_POST, func(r Request) Response {
		return &HttpResponse{
			Payload: NewJsonPayload(""),
		}
	})
}

func handleHttpViewClient(server *DefaultServer) RouteHandler {
	return func(r Request) Response {
		page := &pages.ClientDetailPage{}
		/*
		req := r.(*HttpRequest)

clientId := req.GetAttribute("client")
cli := server.GetRegisteredClient(clientId)

type model struct {
	ClientId string
	Objects  map[betwixt.LWM2MObjectType]betwixt.Object
}

m := &model{
	Objects:  cli.GetObjects(),
	ClientId: clientId,
}
*/

		return &HttpResponse{
			// TemplateModel: m,
			Payload: NewBytesPayload(page.GetContent()),
		}
	}
}

func handleHttpDeleteClient(server *DefaultServer) RouteHandler {
	return func(r Request) Response {
		page := &pages.HomePage{}

		log.Println("Handle Deleting of Client")

		return &HttpResponse{
			Payload: NewBytesPayload(page.GetContent()),
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
			Clients      []*client
			ClientsCount int
			MemUsage     string
			RequestCount int
			ErrorsCount  int
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
			Clients:      cl,
			MemUsage:     strconv.Itoa(int(mem.Alloc / 1000)),
			RequestCount: server.stats.GetRequestsCount(),
			ErrorsCount:  0,
		}

		return &HttpResponse{
			TemplateModel: m,
			Payload:       NewBytesPayload(page.GetContent()),
		}
	}
}
