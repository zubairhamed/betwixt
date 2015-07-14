package server

import (
	"github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/canopus"
	"net"
	"strconv"
	"strings"
	"github.com/zubairhamed/sugoi"
	"log"
)

func NewDefaultCoapServer() *canopus.CoapServer {
	localAddr, err := net.ResolveUDPAddr("udp", ":5683")
	if err != nil {
		log.Fatal("Error starting CoAP Server: ", err)
	}
	return canopus.NewServer(localAddr, nil)
}

func NewDefaultServer(port string) betwixt.Server {
	return &DefaultServer{
		coapServer: NewDefaultCoapServer(),
		httpServer: sugoi.NewSugoi("8081"),
		clients:    make(map[string]betwixt.RegisteredClient),
		stats:      &DefaultServerStatistics{},
		events:     make(map[betwixt.EventType]betwixt.FnEvent),
	}
}

type DefaultServer struct {
	coapServer *canopus.CoapServer
	httpServer *sugoi.SugoiServer
	registry   betwixt.Registry
	stats      betwixt.ServerStatistics
	clients    map[string]betwixt.RegisteredClient
	events     map[betwixt.EventType]betwixt.FnEvent
}

func (server *DefaultServer) GetCoapServer() *canopus.CoapServer {
	return server.coapServer
}

func (server *DefaultServer) GetHttpServer() *sugoi.SugoiServer {
	return server.httpServer
}

func (server *DefaultServer) Start() {
	coap := server.coapServer

	// Setup CoAP Routes
	SetupCoapRoutes(server)

	// Start CoAP Server
	go func() {
		coap.Start()
	}()

	// Setup HTTP Routes
	http := server.httpServer
	SetupHttpRoutes(server)

	betwixt.CallLwm2mEvent(betwixt.EVENT_START, server.events[betwixt.EVENT_START])

	// Start HTTP Server
	http.Serve()
}

func (server *DefaultServer) GetStats() betwixt.ServerStatistics {
	return server.stats
}

func (server *DefaultServer) GetClients() map[string]betwixt.RegisteredClient {
	return server.clients
}

func (server *DefaultServer) UseRegistry(reg betwixt.Registry) {
	server.registry = reg
}

func (server *DefaultServer) GetClient(id string) betwixt.RegisteredClient {
	return server.clients[id]
}

func (server *DefaultServer) update(id string) {
	for k, v := range server.clients {
		if v.GetId() == id {
			v.Update()
			server.clients[k] = v
		}
	}
}

func (server *DefaultServer) register(ep string, addr string, resources []*canopus.CoreResource) (string, error) {
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
			obj = betwixt.NewObject(lwId, nil, server.registry)
		}

		if len(sp) > 1 {
			// Has Object Instance
			instanceId, _ := strconv.Atoi(sp[1])
			obj.AddInstance(instanceId)
		}
		objs[lwId] = obj
	}
	cli.SetObjects(objs)
	server.clients[ep] = cli

	return clientId, nil
}

func (server *DefaultServer) delete(id string) {
	for k, v := range server.clients {
		if v.GetId() == id {

			delete(server.clients, k)
			return
		}
	}
}

func (server *DefaultServer) On(e betwixt.EventType, fn betwixt.FnEvent) {
	server.events[e] = fn
}

func (server *DefaultServer) callEvent(e betwixt.EventType) {
	fn, ok := server.events[e]
	if ok {
		go fn()
	}
}
