package server

import (
	"github.com/zubairhamed/betwixt/api"
	"github.com/zubairhamed/canopus"
	. "github.com/zubairhamed/go-commons/network"
	"log"
	"net"
)

func NewDefaultCoapServer() *canopus.CoapServer {
	localAddr, err := net.ResolveUDPAddr("udp", ":5683")
	if err != nil {
		log.Println("Error starting CoAP Server: ", err)
	}
	return canopus.NewServer(localAddr, nil)
}

func NewDefaultServer() api.Server {
	return &DefaultServer{
		coapServer: NewDefaultCoapServer(),
		httpServer: NewDefaultHttpServer(),
		clients:    make(map[string]api.RegisteredClient),
		stats:      &ServerStatistics{},
	}
}

type DefaultServer struct {
	coapServer *canopus.CoapServer
	httpServer *HttpServer
	registry   api.Registry
	stats      *ServerStatistics
	clients    map[string]api.RegisteredClient
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

func (server *DefaultServer) UseRegistry(reg api.Registry) {
	server.registry = reg
}

func (server *DefaultServer) GetRegisteredClient(id string) api.RegisteredClient {
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

func (server *DefaultServer) register(ep string, addr string) (string, error) {
	clientId := canopus.GenerateToken(8)
	newClient := NewRegisteredClient(ep, clientId, addr)

	server.clients[ep] = newClient

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
