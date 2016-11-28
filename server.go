package betwixt

import (
	"log"
	"strconv"
	"strings"

	"github.com/zubairhamed/canopus"
)

type ServerConfig map[string]string

func NewLwm2mServer(name string, store Store, cfg ServerConfig) *LWM2MServer {
	coapServer := canopus.NewServer()

	return &LWM2MServer{
		Name:       name,
		Store:      store,
		Config:     cfg,
		Stats:      &BetwixtServerStatistics{},
		Events:     make(map[EventType]FnEvent),
		CoapServer: coapServer,
	}
}

type LWM2MServer struct {
	Name       string
	Store      Store
	CoapServer canopus.CoapServer
	Config     ServerConfig
	Stats      ServerStatistics
	Events     map[EventType]FnEvent
	Registry   Registry

	EvtOnRegistered   FnOnRegistered
	EvtOnDeregistered FnOnDeregistered
}

func (c *LWM2MServer) OnRegistered(fn FnOnRegistered) {
	c.EvtOnRegistered = fn
}

func (c *LWM2MServer) OnDeregistered(fn FnOnDeregistered) {
	c.EvtOnDeregistered = fn
}

func (b *LWM2MServer) Serve() error {
	b.CoapServer.OnMessage(func(msg *canopus.Message, inbound bool) {
		b.Stats.IncrementCoapRequestsCount()
	})

	b.CoapServer.Post("/rd", FnCoapRegisterClient(b))
	b.CoapServer.Put("/rd/:id", FnCoapUpdateClient(b))
	b.CoapServer.Delete("/rd/:id", FnCoapDeleteClient(b))

	go b.CoapServer.Start()

	return nil
}

func (b *LWM2MServer) Register(ep string, addr string, resources []*canopus.CoreResource) (string, error) {
	clientId := canopus.GenerateToken(8)
	cli := NewRegisteredClient(ep, clientId, addr, b.CoapServer)

	objs := make(map[LWM2MObjectType]Object)

	for _, o := range resources {
		t := o.Target[1:len(o.Target)]
		sp := strings.Split(t, "/")

		objectId, _ := strconv.Atoi(sp[0])
		lwId := LWM2MObjectType(objectId)

		obj, ok := objs[lwId]
		if !ok {
			obj = NewObject(lwId, nil, b.Registry)
		}

		if len(sp) > 1 {
			// Has Object Instance
			instanceId, _ := strconv.Atoi(sp[1])
			obj.AddInstance(instanceId)
		}
		objs[lwId] = obj
	}
	cli.SetObjects(objs)
	b.Store.PutClient(ep, cli)

	if b.EvtOnRegistered != nil {
		b.EvtOnRegistered(cli)
	}

	return clientId, nil
}

func (b *LWM2MServer) Delete(id string) {
	b.Store.DeleteClient(id)
}

func (b *LWM2MServer) Update(id string) {
	b.Store.UpdateTS(id)
}

func (b *LWM2MServer) UseRegistry(reg Registry) {
	b.Registry = reg
}

func (b *LWM2MServer) GetClients() map[string]RegisteredClient {
	return b.Store.GetClients()
}

func (b *LWM2MServer) GetClient(id string) RegisteredClient {
	return b.Store.GetClient(id)
}

func (b *LWM2MServer) GetServerStats() ServerStatistics {
	return b.Stats
}

func FnCoapRegisterClient(b *LWM2MServer) canopus.RouteHandler {
	return func(req canopus.Request) canopus.Response {
		ep := req.GetURIQuery("ep")

		// lt := req.GetUriQuery("lt")
		// sms := req.GetUriQuery("sms")
		// binding := req.GetUriQuery("b")

		resources := canopus.CoreResourcesFromString(req.GetMessage().GetPayload().String())
		clientId, err := b.Register(ep, req.GetAddress().String(), resources)
		if err != nil {
			log.Println("Error registering client ", ep)
		}

		msg := canopus.NewMessageOfType(canopus.MessageAcknowledgment, req.GetMessage().GetMessageId(), canopus.NewEmptyPayload()).(*canopus.CoapMessage)
		msg.Token = req.GetMessage().GetToken()
		msg.AddOption(canopus.OptionLocationPath, "rd/"+clientId)
		msg.Code = canopus.CoapCodeCreated

		return canopus.NewResponseWithMessage(msg)
	}
}

func FnCoapUpdateClient(b *LWM2MServer) canopus.RouteHandler {
	return func(req canopus.Request) canopus.Response {
		id := req.GetAttribute("id")

		b.Update(id)

		msg := canopus.NewMessageOfType(canopus.MessageAcknowledgment, req.GetMessage().GetMessageId(), canopus.NewEmptyPayload()).(*canopus.CoapMessage)
		msg.Token = req.GetMessage().GetToken()
		msg.Code = canopus.CoapCodeChanged

		return canopus.NewResponseWithMessage(msg)
	}
}

func FnCoapDeleteClient(b *LWM2MServer) canopus.RouteHandler {
	return func(req canopus.Request) canopus.Response {
		id := req.GetAttribute("id")

		b.Delete(id)

		msg := canopus.NewMessageOfType(canopus.MessageAcknowledgment, req.GetMessage().GetMessageId(), canopus.NewEmptyPayload()).(*canopus.CoapMessage)
		msg.Token = req.GetMessage().GetToken()
		msg.Code = canopus.CoapCodeDeleted

		return canopus.NewResponseWithMessage(msg)
	}
}

type BetwixtServerStatistics struct {
	requestCount int
}

func (s *BetwixtServerStatistics) IncrementCoapRequestsCount() {
	s.requestCount++
}

func (s *BetwixtServerStatistics) GetRequestsCount() int {
	return s.requestCount
}
