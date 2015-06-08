package server

import (
	"github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/betwixt/objects"
	"github.com/zubairhamed/canopus"
	. "github.com/zubairhamed/go-commons/network"
	"log"
	"net"
	"strconv"
	"strings"
)

func NewDefaultCoapServer() *canopus.CoapServer {
	localAddr, err := net.ResolveUDPAddr("udp", ":5683")
	if err != nil {
		log.Println("Error starting CoAP Server: ", err)
	}
	return canopus.NewServer(localAddr, nil)
}

func NewDefaultServer(port string) betwixt.Server {
	return &DefaultServer{
		coapServer: NewDefaultCoapServer(),
		httpServer: NewDefaultHttpServer(),
		clients:    make(map[string]betwixt.RegisteredClient),
		stats:      &ServerStatistics{},
	}
}

type DefaultServer struct {
	coapServer *canopus.CoapServer
	httpServer *HttpServer
	registry   betwixt.Registry
	stats      *ServerStatistics
	clients    map[string]betwixt.RegisteredClient
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

	// Start HTTP Server
	http.Start()
}

func (server *DefaultServer) UseRegistry(reg betwixt.Registry) {
	server.registry = reg
}

func (server *DefaultServer) GetRegisteredClient(id string) betwixt.RegisteredClient {
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
			obj = objects.NewObject(lwId, nil, server.registry)
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
