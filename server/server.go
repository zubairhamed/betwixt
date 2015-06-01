package server
import (
    "github.com/zubairhamed/goap"
    "github.com/zubairhamed/go-lwm2m/api"
    "log"
    "net"
    . "github.com/zubairhamed/go-commons/network"
)

func NewDefaultServer() (Server) {
    localAddr, err := net.ResolveUDPAddr("udp", ":5683")
    if err != nil {
        log.Println("Error starting CoAP Server: ", err)
    }
    coapServer := goap.NewServer(localAddr, nil)
    httpServer := NewDefaultHttpServer()

    return &DefaultServer{
        coapServer:  coapServer,
        httpServer:  httpServer,
        clients:     make(map[string]RegisteredClient),
    }
}

type Server interface {
    UseRegistry(api.Registry)
    Start()
}

type DefaultServer struct {
    coapServer     *goap.CoapServer
    httpServer     *HttpServer
    registry       api.Registry
    clients        map[string]RegisteredClient
}

func (server *DefaultServer) Start() {
    s := server.coapServer

    // Setup Routes
    s.NewRoute("rd", goap.POST, server.handleRegister)
    s.NewRoute("rd/{id}", goap.PUT, server.handleUpdate)

    // Start CoAP Server
    go func() {
        s.Start()
    }()

    // Start HTTP Server
    server.httpServer.Start()
}

func (server *DefaultServer) UseRegistry(reg api.Registry) {
    server.registry = reg
}

func (server *DefaultServer) handleRegister(r Request) (Response) {
    req := r.(*goap.CoapRequest)
    ep := req.GetUriQuery("ep")

    clientId, err := server.register(ep)
    if err != nil {
        log.Println("Error registering client ", ep)
    }

    msg := goap.NewMessageOfType(goap.TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
    msg.Token = req.GetMessage().Token
    msg.AddOption(goap.OPTION_LOCATION_PATH, "rd/" + clientId)
    msg.Code = goap.COAPCODE_201_CREATED

    return goap.NewResponseWithMessage(msg)
}

func (server *DefaultServer) handleUpdate(r Request) (Response) {
    req := r.(*goap.CoapRequest)
    id := req.GetAttribute("id")

    server.update(id)

    msg := goap.NewMessageOfType(goap.TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
    msg.Token = req.GetMessage().Token
    msg.Code = goap.COAPCODE_204_CHANGED

    return goap.NewResponseWithMessage(msg)
}

func (server *DefaultServer) GetRegisteredClient(id string) (RegisteredClient){
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

func NewRegisteredClient (ep string, id string) (RegisteredClient) {
    return &DefaultRegisteredClient{
        name: ep,
        id: id,
    }
}

type RegisteredClient interface {
    GetId() string
    GetName() string
    GetLifetime() int
    GetVersion() string
    GetBindingMode() api.BindingMode
    GetSmsNumber() string
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
