package app

import (
	"github.com/zenazn/goji/web"
	"log"
	"net/http"
	"strings"
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
	log.Println("fn http - view client")
}

func (b *BetwixtWebApp) fnHttpApiGetClients(c web.C, w http.ResponseWriter, r *http.Request) {
	log.Println("fn http api - get clients")
}

func (b *BetwixtWebApp) fnHttpApiGetServerStats(c web.C, w http.ResponseWriter, r *http.Request) {
	log.Println("fn http api - get server stats")
}

func (b *BetwixtWebApp) fnHttpApiGetClientMessages(c web.C, w http.ResponseWriter, r *http.Request) {
	log.Println("fn http api - get client messages")
}

func (b *BetwixtWebApp) fnHttpApiGetClient(c web.C, w http.ResponseWriter, r *http.Request) {
	log.Println("fn http api - get client")
}

func (b *BetwixtWebApp) fnHttpApiGetClientResource(c web.C, w http.ResponseWriter, r *http.Request)           {}
func (b *BetwixtWebApp) fnHttpApiGetClientInstance(c web.C, w http.ResponseWriter, r *http.Request)           {}
func (b *BetwixtWebApp) fnHttpApiPutClientResource(c web.C, w http.ResponseWriter, r *http.Request)           {}
func (b *BetwixtWebApp) fnHttpApiPutClientInstance(c web.C, w http.ResponseWriter, r *http.Request)           {}
func (b *BetwixtWebApp) fnHttpApiDeleteClientInstance(c web.C, w http.ResponseWriter, r *http.Request)        {}
func (b *BetwixtWebApp) fnHttpApiObserveClientResource(c web.C, w http.ResponseWriter, r *http.Request)       {}
func (b *BetwixtWebApp) fnHttpApiDeleteObserveClientResource(c web.C, w http.ResponseWriter, r *http.Request) {}
func (b *BetwixtWebApp) fnHttpApiPostClientObservation(c web.C, w http.ResponseWriter, r *http.Request)       {}
func (b *BetwixtWebApp) fnHttpApiPostClientInstance(c web.C, w http.ResponseWriter, r *http.Request)          {}

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
