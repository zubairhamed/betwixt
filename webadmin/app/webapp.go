package app

import (
	"github.com/zubairhamed/canopus"
	"github.com/zenazn/goji"
	"flag"
	"log"
)

type ServerConfig map[string]string

func NewWebApp(store Store, cfg ServerConfig) *BetwixtWebApp {

	name := cfg["name"]
	httpPort := cfg["http-port"]

	return &BetwixtWebApp{
		name: name,
		httpPort: httpPort,
		store: store,
		coapServer: canopus.NewServer("5683", ""),
		config: cfg,
	}
}

type BetwixtWebApp struct {
	name 				string
	store 			Store
	httpPort 		string
	coapServer 	canopus.CoapServer
	config 			ServerConfig
	wait 				chan struct{}
}

func (b *BetwixtWebApp) Serve() error {
	b.setupCoap()
	b.setupHttp()

	go func() {
		b.coapServer.Start()
	}()

	go func() {
		httpPort := b.config["http-port"]
		flag.Set("bind", ":" + httpPort)

		log.Println("Start Server on port " + httpPort)
		goji.Serve()
	}()

	b.wait = make(chan struct{})

	<-b.wait

	return nil
}

func (b *BetwixtWebApp) setupCoap() {
	b.coapServer.OnMessage(func(msg *canopus.Message, inbound bool) {
		// server.stats.IncrementCoapRequestsCount()
	})

	b.coapServer.Post("/rd", FnCoapRegisterClient)
	b.coapServer.Put("/rd/:id", FnCoapUpdateClient)
	b.coapServer.Delete("/rd/:id", FnCoapDeleteClient)

}

func (b *BetwixtWebApp) setupHttp() {
	goji.Get("/", FnHttpIndexPage)
	goji.Get("/client/:client/view", FnHttpClientView)
	goji.Get("/api/clients", FnHttpApiGetClients)
	goji.Get("/api/server/stats", FnHttpApiGetServerStats)
	goji.Get("/api/server/:client/messages", FnHttpApiGetClientMessages)
	goji.Get("/api/clients/:client", FnHttpApiGetClient)
	goji.Get("/api/clients/:client/:object/:instance/:resource", FnHttpApiGetClientResource)
	goji.Get("/api/clients/:client/:object/:instance", FnHttpApiGetClientInstance)

	goji.Put("/api/clients/:client/:object/:instance/:resource", FnHttpApiPutClientResource)
	goji.Put("/api/clients/:client/:object/:instance", FnHttpApiPutClientInstance)

	goji.Delete("/api/clients/:client/:object/:instance", FnHttpApiDeleteClientInstance)
	goji.Delete("/api/clients/:client/:object/:instance/:resource/observe", FnHttpApiDeleteObserveClientResource)

	goji.Post("/api/clients/:client/:object/:instance/:resource/observe", FnHttpApiObserveClientResource)
	goji.Post("/api/clients/:client/:object/:instance/:resource", FnHttpApiPostClientObservation)
	goji.Post("/api/clients/:client/:object/:instance", FnHttpApiPostClientInstance)
}

// Instantiate CoAP Server
// Register all CoAP Endpoints
/*
		POST		/rd
		PUT			/rd/:id
		DELETE	/rd/:id
 */