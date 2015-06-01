package server
import (
    "github.com/zubairhamed/goap"
    "github.com/zubairhamed/go-lwm2m/api"
    "log"
    "net"
    . "github.com/zubairhamed/go-commons/network"
)

func NewDefaultServer() (api.Server) {
    localAddr, err := net.ResolveUDPAddr("udp", ":5683")
    if err != nil {
        log.Println("Error starting CoAP Server: ", err)
    }
    coapServer := goap.NewServer(localAddr, nil)
    httpServer := NewDefaultHttpServer()

    return &DefaultServer{
        coapServer:  coapServer,
        httpServer:  httpServer,
        clients:     make(map[string]api.RegisteredClient),
    }
}

type DefaultServer struct {
    coapServer     *goap.CoapServer
    httpServer     *HttpServer
    registry       api.Registry
    clients        map[string]api.RegisteredClient
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

func (server *DefaultServer) GetRegisteredClient(id string) (api.RegisteredClient){
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

func NewRegisteredClient (ep string, id string) (api.RegisteredClient) {
    return &DefaultRegisteredClient{
        name: ep,
        id: id,
    }
}

type DefaultRegisteredClient struct {
    id          string
    name        string
    lifetime    int
    version     string
    bindingMode api.BindingMode
    smsNumber   string
}

func (c *DefaultRegisteredClient) GetId() string {
    return c.id
}

func (c *DefaultRegisteredClient) GetName() string {
    return c.name
}

func (c *DefaultRegisteredClient) GetLifetime() int {
    return c.lifetime
}

func (c *DefaultRegisteredClient) GetVersion() string {
    return c.version
}

func (c *DefaultRegisteredClient) GetBindingMode() api.BindingMode {
    return c.bindingMode
}

func (c *DefaultRegisteredClient) GetSmsNumber() string {
    return c.smsNumber
}
