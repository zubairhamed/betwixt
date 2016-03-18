package app

import (
	"bytes"
	"flag"
	"github.com/zenazn/goji"
	"github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/canopus"
	"log"
	"strconv"
	"strings"
	"html/template"
)

type ServerConfig map[string]string

func NewWebApp(store Store, cfg ServerConfig) *BetwixtWebApp {

	name := cfg["name"]
	httpPort := cfg["http-port"]

	w := &BetwixtWebApp{
		name:             name,
		store:            store,
		httpPort:         httpPort,
		coapServer:       canopus.NewServer("5683", ""),
		config:           cfg,
		connectedClients: make(map[string]betwixt.RegisteredClient),
		stats:            &BetwixtServerStatistics{},
		events:           make(map[betwixt.EventType]betwixt.FnEvent),
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
	tpl 						 *template.Template
	connectedClients map[string]betwixt.RegisteredClient
	stats            betwixt.ServerStatistics
	events           map[betwixt.EventType]betwixt.FnEvent
	registry         betwixt.Registry
}

func (b *BetwixtWebApp) cacheWebTemplates() {
	tplBuf := bytes.NewBuffer([]byte{})

	var tpls = []string{"index", "head", "logs", "settings", "stats", "client"}
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
		b.stats.IncrementCoapRequestsCount()
	})

	b.coapServer.Post("/rd", FnCoapRegisterClient(b))
	b.coapServer.Put("/rd/:id", FnCoapUpdateClient(b))
	b.coapServer.Delete("/rd/:id", FnCoapDeleteClient(b))
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

func (b *BetwixtWebApp) getClient(id string) betwixt.RegisteredClient {
	log.Println(b.connectedClients)
	return b.connectedClients[id]
}

func (b *BetwixtWebApp) getServerStats() betwixt.ServerStatistics {
	return b.stats
}

func (b *BetwixtWebApp) register(ep string, addr string, resources []*canopus.CoreResource) (string, error) {
	clientId := canopus.GenerateToken(8)
	cli := NewRegisteredClient(ep, clientId, addr)

	objs := make(map[betwixt.LWM2MObjectType]betwixt.Object)

	for _, o := range resources {
		t := o.Target[1:len(o.Target)]
		sp := strings.Split(t, "/")

		objectId, _ := strconv.Atoi(sp[0])
		lwId := betwixt.LWM2MObjectType(objectId)

		obj, ok := objs[lwId]
		if !ok {
			obj = betwixt.NewObject(lwId, nil, b.registry)
		}

		if len(sp) > 1 {
			// Has Object Instance
			instanceId, _ := strconv.Atoi(sp[1])
			obj.AddInstance(instanceId)
		}
		objs[lwId] = obj
	}
	cli.SetObjects(objs)
	b.connectedClients[ep] = cli

	return clientId, nil
}

func (b *BetwixtWebApp) delete(id string) {
	for k, v := range b.connectedClients {
		if v.GetId() == id {

			delete(b.connectedClients, k)
			return
		}
	}
}

func (b *BetwixtWebApp) update(id string) {
	for k, v := range b.connectedClients {
		if v.GetId() == id {
			v.Update()
			b.connectedClients[k] = v
		}
	}
}

func (b *BetwixtWebApp) UseRegistry(reg betwixt.Registry) {
	b.registry = reg
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
