package app

import (
	"bytes"
	"flag"
	"github.com/alecthomas/template"
	"github.com/zenazn/goji"
	"github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/canopus"
	"log"
	"strings"
)

type ServerConfig map[string]string

func NewWebApp(store Store, cfg ServerConfig) *BetwixtWebApp {

	name := cfg["name"]
	httpPort := cfg["http-port"]

	w := &BetwixtWebApp{
		name:       name,
		httpPort:   httpPort,
		store:      store,
		coapServer: canopus.NewServer("5683", ""),
		config:     cfg,
		stats:      &BetwixtServerStatistics{},
	}

	return w
}

type BetwixtWebApp struct {
	name             string
	store            Store
	httpPort         string
	coapServer       canopus.CoapServer
	config           ServerConfig
	wait             chan struct{}
	tpl              *template.Template
	connectedClients map[string]betwixt.RegisteredClient
	stats            betwixt.ServerStatistics
	events           map[betwixt.EventType]betwixt.FnEvent
}

func (b *BetwixtWebApp) cacheWebTemplates() {
	tplBuf := bytes.NewBuffer([]byte{})

	var tpls = []string{
		"index", "nav", "head", "logs", "settings", "clients_list", "client_content",
	}

	var tpl []byte

	for _, v := range tpls {
		tpl, _ = AssetContent("tpl/" + v + ".html")
		tplBuf.Write(tpl)
	}

	b.tpl, _ = template.New("tpls").Delims("#{", "}#").Parse(tplBuf.String())
}

func (b *BetwixtWebApp) Serve() error {
	b.cacheWebTemplates()

	b.setupCoap()
	b.setupHttp()

	go func() {
		b.coapServer.Start()
	}()

	go func() {
		httpPort := b.config["http-port"]
		flag.Set("bind", ":"+httpPort)

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
	goji.Get("/", b.fnHttpIndexPage)
	goji.Get("/logs", b.fnHttpLogsPage)
	goji.Get("/settings", b.fnHttpSettingsPage)
	goji.Get("/client/:client/view", b.fnHttpClientView)

	// Static Resources
	goji.Get("/css/*", b.fnWebUiResources)
	goji.Get("/img/*", b.fnWebUiResources)
	goji.Get("/js/*", b.fnWebUiResources)

	// REST API
	goji.Get("/api/clients", b.fnHttpApiGetClients)
	goji.Get("/api/server/stats", b.fnHttpApiGetServerStats)
	goji.Get("/api/server/:client/messages", b.fnHttpApiGetClientMessages)
	goji.Get("/api/clients/:client", b.fnHttpApiGetClient)
	goji.Get("/api/clients/:client/:object/:instance/:resource", b.fnHttpApiGetClientResource)
	goji.Get("/api/clients/:client/:object/:instance", b.fnHttpApiGetClientInstance)
	goji.Put("/api/clients/:client/:object/:instance/:resource", b.fnHttpApiPutClientResource)
	goji.Put("/api/clients/:client/:object/:instance", b.fnHttpApiPutClientInstance)
	goji.Delete("/api/clients/:client/:object/:instance", b.fnHttpApiDeleteClientInstance)
	goji.Delete("/api/clients/:client/:object/:instance/:resource/observe", b.fnHttpApiDeleteObserveClientResource)
	goji.Post("/api/clients/:client/:object/:instance/:resource/observe", b.fnHttpApiObserveClientResource)
	goji.Post("/api/clients/:client/:object/:instance/:resource", b.fnHttpApiPostClientObservation)
	goji.Post("/api/clients/:client/:object/:instance", b.fnHttpApiPostClientInstance)
}

func (b *BetwixtWebApp) getClients() map[string]betwixt.RegisteredClient {
	return b.connectedClients
}

func (b *BetwixtWebApp) getServerStats() betwixt.ServerStatistics {
	return b.stats
}

func AssetContent(path string) ([]byte, error) {
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}

	data, err := Asset("resources/" + path)

	if err != nil {
		log.Println(err)
	}

	return data, err
}

// Instantiate CoAP Server
// Register all CoAP Endpoints
/*
	POST		/rd
	PUT			/rd/:id
	DELETE	/rd/:id
*/
