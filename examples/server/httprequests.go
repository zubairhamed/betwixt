package server

import (
	"github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/betwixt/examples/server/models"
	"github.com/zubairhamed/go-commons/logging"
	"github.com/zubairhamed/go-commons/typeval"
	"runtime"
	"strconv"
	"github.com/zubairhamed/sugoi"
)

func SetupHttpRoutes(server betwixt.Server) {
	httpServer := server.GetHttpServer()

	// Pages
	httpServer.Get("/", handleHttpHome(server))
	httpServer.Get("/client/:client/view", handleHttpViewClient(server))

	httpServer.Get("/api/clients", func(r *sugoi.Request) sugoi.Content {
		cl := []models.ClientModel{}
		for _, v := range server.GetClients() {

			objs := make(map[string]models.ObjectModel)
			for key, val := range v.GetObjects() {
				objectModel := models.ObjectModel{
					Instances:  val.GetInstances(),
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
				Objects:          objs,
			}
			cl = append(cl, c)
		}

		return cl
	})

	httpServer.Get("/api/server/stats", func(r *sugoi.Request) sugoi.Content {
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)

		clientsCount := len(server.GetClients())

		model := &models.StatsModel{
			ClientsCount: clientsCount,
			MemUsage:     strconv.Itoa(int(mem.Alloc / 1000)),
			Requests:     server.GetStats().GetRequestsCount(),
			Errors:       0,
		}

		return model
	})

	// Get Message, Logs
	httpServer.Get("/api/server/:client/messages", func(r *sugoi.Request) sugoi.Content {
		return sugoi.OK()
	})

	// Read
	httpServer.Get("/api/clients/:client", func(req *sugoi.Request) sugoi.Content {
		clientId := req.GetAttribute("client")

		v := server.GetClient(clientId)
		if v == nil {

		}

		objs := make(map[string]models.ObjectModel)
		for key, val := range v.GetObjects() {
			objectModel := models.ObjectModel{
				Instances:  val.GetInstances(),
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
			Objects:          objs,
		}

		return c
	})

	httpServer.Get("/api/clients/:client/:object/:instance/:resource", func(req *sugoi.Request) sugoi.Content {
		clientId := req.GetAttribute("client")
		object := req.GetAttributeAsInt("object")
		instance := req.GetAttributeAsInt("instance")
		resource := req.GetAttributeAsInt("resource")
		cli := server.GetClient(clientId)

		val, _ := cli.ReadResource(uint16(object), uint16(instance), uint16(resource))

		if val == nil {
			logging.LogError("Value returned by ReadResource is nil")
		}
		contentModels := []*models.ContentValueModel{}
		if val.GetType() == typeval.VALUETYPE_MULTIRESOURCE {
			resources := val.(*betwixt.MultipleResourceValue).GetValue().([]*betwixt.ResourceValue)

			for _, resource := range resources {
				contentModels = append(contentModels, &models.ContentValueModel{
					Id:    resource.GetId(),
					Value: resource.GetValue(),
				})
			}
		} else {
			resource := val.(*betwixt.ResourceValue)
			contentModels = append(contentModels, &models.ContentValueModel{
				Id:    resource.GetId(),
				Value: resource.GetValue(),
			})
		}

		payload := &models.ExecuteResponseModel{
			Content: contentModels,
		}

		return payload
	})

	httpServer.Get("/api/clients/:client/:object/:instance", func(r *sugoi.Request) sugoi.Content {
		return sugoi.OK()
	})

	// Write
	httpServer.Put("/api/clients/:client/:object/:instance/:resource", func(r *sugoi.Request) sugoi.Content {
		return sugoi.OK()
	})

	httpServer.Put("/api/clients/:client/:object/:instance", func(r *sugoi.Request) sugoi.Content {
		return sugoi.OK()
	})

	// Delete
	httpServer.Delete("/api/clients/:client/:object/:instance", func(r *sugoi.Request) sugoi.Content {
		return sugoi.OK()
	})

	// Observe
	httpServer.Post("/api/clients/:client/:object/:instance/:resource/observe", func(r *sugoi.Request) sugoi.Content {
		return sugoi.OK()
	})

	// Cancel Observe
	httpServer.Delete("/api/clients/:client/:object/:instance/:resource/observe", func(r *sugoi.Request) sugoi.Content {
		return sugoi.OK()
	})

	// Execute
	httpServer.Post("/api/clients/:client/:object/:instance/:resource", func(r *sugoi.Request) sugoi.Content {
		return sugoi.OK()
	})

	// Create
	httpServer.Post("/api/clients/:client/:object/:instance", func(r *sugoi.Request) sugoi.Content {
		return sugoi.OK()
	})
}

func handleHttpViewClient(server betwixt.Server) sugoi.RouteHandler {
	return func(r *sugoi.Request) sugoi.Content {
		return sugoi.StaticHtml("details.html")
	}
}

func handleHttpHome(server betwixt.Server) sugoi.RouteHandler {
	return func(r *sugoi.Request) sugoi.Content {
		return sugoi.StaticHtml("index.html")
	}
}
