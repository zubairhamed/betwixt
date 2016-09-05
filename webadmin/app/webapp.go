package app

import (
	"bytes"
	"flag"
	"html/template"
	"log"
	"strings"

	"github.com/zenazn/goji"
	"github.com/zubairhamed/betwixt"
)

func NewWebApp(store betwixt.Store, cfg betwixt.ServerConfig) *BetwixtWebApp {

	name := cfg["name"]
	httpPort := cfg["http-port"]

	server := betwixt.NewLwm2mServer(name, store, cfg)

	w := &BetwixtWebApp{
		server:   server,
		httpPort: httpPort,
	}

	return w
}

type BetwixtWebApp struct {
	server   *betwixt.LWM2MServer
	httpPort string
	wait     chan struct{}
	tpl      *template.Template
}

func (b *BetwixtWebApp) UseRegistry(reg betwixt.Registry) {
	b.server.UseRegistry(reg)
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
	b.setupHttp()

	b.server.Serve()

	go func() {
		httpPort := b.server.Config["http-port"]
		flag.Set("bind", ":"+httpPort)

		log.Println("Start Server on port " + httpPort)
		goji.Serve()
	}()

	b.wait = make(chan struct{})

	<-b.wait

	return nil
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
