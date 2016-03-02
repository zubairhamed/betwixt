package app

import (
	"github.com/zenazn/goji/web"
	"net/http"
	"log"
)

func FnHttpIndexPage(c web.C, w http.ResponseWriter, r *http.Request) {
	log.Println("fn http - index")
}

func FnHttpClientView (c web.C, w http.ResponseWriter, r *http.Request) {}
func FnHttpApiGetClients (c web.C, w http.ResponseWriter, r *http.Request) {}
func FnHttpApiGetServerStats (c web.C, w http.ResponseWriter, r *http.Request) {}
func FnHttpApiGetClientMessages (c web.C, w http.ResponseWriter, r *http.Request) {}
func FnHttpApiGetClient (c web.C, w http.ResponseWriter, r *http.Request) {}
func FnHttpApiGetClientResource (c web.C, w http.ResponseWriter, r *http.Request) {}
func FnHttpApiGetClientInstance (c web.C, w http.ResponseWriter, r *http.Request) {}
func FnHttpApiPutClientResource (c web.C, w http.ResponseWriter, r *http.Request) {}
func FnHttpApiPutClientInstance (c web.C, w http.ResponseWriter, r *http.Request) {}
func FnHttpApiDeleteClientInstance (c web.C, w http.ResponseWriter, r *http.Request) { }
func FnHttpApiObserveClientResource (c web.C, w http.ResponseWriter, r *http.Request) {}
func FnHttpApiDeleteObserveClientResource (c web.C, w http.ResponseWriter, r *http.Request) { }
func FnHttpApiPostClientObservation (c web.C, w http.ResponseWriter, r *http.Request) { }
func FnHttpApiPostClientInstance (c web.C, w http.ResponseWriter, r *http.Request) { }

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
