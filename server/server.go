package server

import (
	. "github.com/zubairhamed/go-commons/network"
	"github.com/zubairhamed/go-lwm2m/api"
	"github.com/zubairhamed/goap"
	"log"
	"net"
)

func NewDefaultCoapServer() *goap.CoapServer {
	localAddr, err := net.ResolveUDPAddr("udp", ":5683")
	if err != nil {
		log.Println("Error starting CoAP Server: ", err)
	}
	return goap.NewServer(localAddr, nil)
}

func NewDefaultServer() api.Server {
	return &DefaultServer{
		coapServer: NewDefaultCoapServer(),
		httpServer: NewDefaultHttpServer(),
		clients:    make(map[string]api.RegisteredClient),
	}
}

type DefaultServer struct {
	coapServer *goap.CoapServer
	httpServer *HttpServer
	registry   api.Registry
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

}

func (server *DefaultServer) register(ep string) (string, error) {
	clientId := goap.GenerateToken(8)
	newClient := NewRegisteredClient(ep, clientId)
	server.clients[clientId] = newClient

	return clientId, nil
}
