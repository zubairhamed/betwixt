package server
import (
    "github.com/zubairhamed/goap"
    "github.com/zubairhamed/go-lwm2m/api"
    "log"
    "net"
)

func NewDefaultServer() (Server) {
    localAddr, err := net.ResolveUDPAddr("udp", ":5683")
    if err != nil {
        log.Println("Error starting CoAP Server: ", err)
    }
    coapServer := goap.NewServer(localAddr, nil)

    return &DefaultServer{
        coapServer:  coapServer,
        clients:     make(map[string]RegisteredClient),
    }
}

type Server interface {
    Start()
}

type DefaultServer struct {
    coapServer     *goap.CoapServer
    registry       api.Registry
    clients        map[string]RegisteredClient
}

func (server *DefaultServer) Start() {
    s := server.coapServer

    s.NewRoute("rd", goap.POST, server.handleRegister)
    s.NewRoute("rd/{id}", goap.PUT, server.handleUpdate)

    s.Start()
}

func (server *DefaultServer) handleRegister(req *goap.CoapRequest) *(goap.CoapResponse) {
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

func (server *DefaultServer) handleUpdate(req *goap.CoapRequest) *(goap.CoapResponse) {
    id := req.GetAttribute("id")

    log.Println("Updating id", id)
    // 204 changed

    msg := goap.NewMessageOfType(goap.TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
    msg.Token = req.GetMessage().Token
    msg.Code = goap.COAPCODE_204_CHANGED

    return goap.NewResponseWithMessage(msg)

}

func (server *DefaultServer) GetRegisteredClient(id string) (RegisteredClient){
    return server.clients[id]
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

}

type DefaultRegisteredClient struct {
    id          string
    name        string
    lifetime    int
    version     string
    bindingMode api.BindingMode
    smsNumber   string
}