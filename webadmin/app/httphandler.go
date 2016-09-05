package app

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"strings"

	"github.com/zenazn/goji/web"
	"github.com/zubairhamed/betwixt"
)

func (b *BetwixtWebApp) fnHttpIndexPage(c web.C, w http.ResponseWriter, r *http.Request) {
	b.tpl.ExecuteTemplate(w, "page_index", nil)
}

func (b *BetwixtWebApp) fnHttpLogsPage(c web.C, w http.ResponseWriter, r *http.Request) {
	b.tpl.ExecuteTemplate(w, "page_logs", nil)
}

func (b *BetwixtWebApp) fnHttpSettingsPage(c web.C, w http.ResponseWriter, r *http.Request) {
	b.tpl.ExecuteTemplate(w, "page_settings", nil)
}

func (b *BetwixtWebApp) fnHttpClientView(c web.C, w http.ResponseWriter, r *http.Request) {
	b.tpl.ExecuteTemplate(w, "page_client", nil)
}

func (b *BetwixtWebApp) fnHttpApiGetClients(c web.C, w http.ResponseWriter, r *http.Request) {
	cl := []ClientModel{}

	for _, v := range b.server.GetClients() {

		objs := make(map[string]ObjectModel)
		for key, val := range v.GetObjects() {
			objectModel := ObjectModel{
				Instances:  val.GetInstances(),
				Definition: val.GetDefinition(),
			}
			typeKey := strconv.Itoa(int(key))
			objs[typeKey] = objectModel
		}

		c := ClientModel{
			Endpoint:         v.GetName(),
			RegistrationID:   v.GetId(),
			RegistrationDate: v.GetRegistrationDate().Format("Jan 2, 2006, 3:04pm (SGT)"),
			LastUpdate:       v.LastUpdate().Format("Jan 2, 2006, 3:04pm (SGT)"),
			Objects:          objs,
		}
		cl = append(cl, c)
	}

	if jsonBytes, err := json.Marshal(cl); err == nil {
		w.Write(jsonBytes)
	} else {
		w.WriteHeader(500)
	}
}

func (b *BetwixtWebApp) fnHttpApiGetServerStats(c web.C, w http.ResponseWriter, r *http.Request) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	clientsCount := len(b.server.GetClients())

	model := &StatsModel{
		ClientsCount: clientsCount,
		MemUsage:     strconv.Itoa(int(mem.Alloc / 1000)),
		Requests:     b.server.GetServerStats().GetRequestsCount(),
		Errors:       0,
	}

	if jsonBytes, err := json.Marshal(model); err == nil {
		w.Write(jsonBytes)
	} else {
		w.WriteHeader(500)
	}
}

func (b *BetwixtWebApp) fnHttpApiGetClientMessages(c web.C, w http.ResponseWriter, r *http.Request) {
	log.Println("fn http api - get client messages")
	w.WriteHeader(501)
}

func (b *BetwixtWebApp) fnHttpApiGetClient(c web.C, w http.ResponseWriter, r *http.Request) {

	clientId := c.URLParams["client"]

	v := b.server.GetClient(clientId)
	log.Println(clientId, v)
	if v == nil {
		w.WriteHeader(500)
	}

	objs := make(map[string]ObjectModel)
	for key, val := range v.GetObjects() {
		objectModel := ObjectModel{
			Instances:  val.GetInstances(),
			Definition: val.GetDefinition(),
		}
		typeKey := strconv.Itoa(int(key))
		objs[typeKey] = objectModel
	}

	cl := ClientModel{
		Endpoint:         v.GetName(),
		RegistrationID:   v.GetId(),
		RegistrationDate: v.GetRegistrationDate().Format("Jan 2, 2006, 3:04pm (SGT)"),
		LastUpdate:       v.LastUpdate().Format("Jan 2, 2006, 3:04pm (SGT)"),
		Objects:          objs,
	}

	if jsonBytes, err := json.Marshal(cl); err == nil {
		w.Write(jsonBytes)
	} else {
		w.WriteHeader(500)
	}
}

func (b *BetwixtWebApp) fnHttpApiGetClientResource(c web.C, w http.ResponseWriter, r *http.Request) {
	clientId := c.URLParams["client"]
	object, err := strconv.Atoi(c.URLParams["object"])
	if err != nil {
		w.WriteHeader(500)
	}

	instance, err := strconv.Atoi(c.URLParams["instance"])
	if err != nil {
		w.WriteHeader(500)
	}

	resource, err := strconv.Atoi(c.URLParams["resource"])
	if err != nil {
		w.WriteHeader(500)
	}

	cli := b.server.GetClient(clientId)
	val, _ := cli.ReadResource(uint16(object), uint16(instance), uint16(resource))

	if val == nil {
		log.Println("Value returned by ReadResource is nil")
		w.WriteHeader(500)
	}

	contentModels := []*ContentValueModel{}
	if val.GetType() == betwixt.VALUETYPE_MULTIRESOURCE {
		resources := val.(*betwixt.MultipleResourceValue).GetValue().([]*betwixt.ResourceValue)

		for _, resource := range resources {
			contentModels = append(contentModels, &ContentValueModel{
				Id:    uint16(resource.GetId()),
				Value: resource.GetValue(),
			})
		}
	} else {
		resource := val.(*betwixt.ResourceValue)
		contentModels = append(contentModels, &ContentValueModel{
			Id:    uint16(resource.GetId()),
			Value: resource.GetValue(),
		})
	}

	payload := &ExecuteResponseModel{
		Content: contentModels,
	}

	if jsonBytes, err := json.Marshal(payload); err == nil {
		w.Write(jsonBytes)
	} else {
		w.WriteHeader(500)
	}
}

func (b *BetwixtWebApp) fnHttpApiGetClientInstance(c web.C, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
}

func (b *BetwixtWebApp) fnHttpApiPutClientResource(c web.C, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
}

func (b *BetwixtWebApp) fnHttpApiPutClientInstance(c web.C, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
}

func (b *BetwixtWebApp) fnHttpApiDeleteClientInstance(c web.C, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
}

func (b *BetwixtWebApp) fnHttpApiObserveClientResource(c web.C, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
}

func (b *BetwixtWebApp) fnHttpApiDeleteObserveClientResource(c web.C, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
}

func (b *BetwixtWebApp) fnHttpApiPostClientObservation(c web.C, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
}

func (b *BetwixtWebApp) fnHttpApiPostClientInstance(c web.C, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
}

func (b *BetwixtWebApp) fnWebUiResources(c web.C, w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch {
	case strings.HasPrefix(path, "/css"):
		w.Header().Set("Content-Type", "text/css")

	case strings.HasPrefix(path, "/js"):
		w.Header().Set("Content-Type", "text/javascript")
	}

	data, _ := AssetContent(path)

	w.Write(data)
}

/*
			GET 		/
			GET			/client/:client/view
			GET			/api/clients
			GET			/api/server/stats
			GET			/api/server/:client/messages
			GET			/api/clients/:client
			GET			/api/clients/:client/:object/:instance/:resource
			GET			/api/clients/:client/:object/:instance
			PUT			/api/clients/:client/:object/:instance/:resource
			PUT			/api/clients/:client/:object/:instance
			DELETE	/api/clients/:client/:object/:instance
			POST		/api/clients/:client/:object/:instance/:resource/observe
			DELETE	/api/clients/:client/:object/:instance/:resource/observe
			POST		/api/clients/:client/:object/:instance/:resource
			POST		/api/clients/:client/:object/:instance

gs)
	goji.Get("/sensors.html", s.HandleWebUiSensors)
	goji.Get("/observations.html", s.HandleWebUiObservations)
	goji.Get("/observedproperties.html", s.HandleWebUiObservedProperties)
	goji.Get("/locations.html", s.HandleWebUiLocations)
	goji.Get("/datastreams.html", s.HandleWebUiDatastreams)
	goji.Get("/featuresofinterest.html", s.HandleWebUiFeaturesOfInterest)
	goji.Get("/historiclocations.html", s.HandleWebUiHistoricLocations)
	goji.Get("/css/*", s.HandleWebUiResources)
	goji.Get("/img/*", s.HandleWebUiResources)
	goji.Get("/js/*", s.HandleWebUiResources)

	goji.Get("/v1.0", s.handleRootResource)
	goji.Get("/v1.0/", s.handleRootResource)

	goji.Get("/v1.0/*", s.HandleGet)
	goji.Post("/v1.0/*", s.HandlePost)
	goji.Put("/v1.0/*", s.HandlePut)
	goji.Delete("/v1.0/*", s.HandleDelete)
	goji.Patch("/v1.0/*", s.HandlePatch)
*/
