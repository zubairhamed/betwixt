package webadmin

import (
	. "github.com/zubairhamed/canopus"
	"log"
)

func SetupCoapRoutes(server *DefaultServer) {
	coap := server.coapServer

	coap.OnMessage(func(msg *Message, inbound bool) {
		server.stats.IncrementCoapRequestsCount()
	})

	coap.Post("/rd", handleRegister(server))
	coap.Put("/rd/:id", handleUpdate(server))
	coap.Delete("/rd/:id", handleDelete(server))
}

func handleRegister(server *DefaultServer) RouteHandler {
	return func(req CoapRequest) CoapResponse {
		ep := req.GetURIQuery("ep")
		// lt := req.GetUriQuery("lt")
		// sms := req.GetUriQuery("sms")
		// binding := req.GetUriQuery("b")

		resources := CoreResourcesFromString(req.GetMessage().Payload.String())
		clientId, err := server.register(ep, req.GetAddress().String(), resources)
		if err != nil {
			log.Println("Error registering client ", ep)
		}

		msg := NewMessageOfType(MessageAcknowledgment, req.GetMessage().MessageID)
		msg.Token = req.GetMessage().Token
		msg.AddOption(OptionLocationPath, "rd/"+clientId)
		msg.Code = CoapCodeCreated

		return NewResponseWithMessage(msg)
	}
}

func handleUpdate(server *DefaultServer) RouteHandler {
	return func(req CoapRequest) CoapResponse {
		id := req.GetAttribute("id")

		server.update(id)

		msg := NewMessageOfType(MessageAcknowledgment, req.GetMessage().MessageID)
		msg.Token = req.GetMessage().Token
		msg.Code = CoapCodeChanged

		return NewResponseWithMessage(msg)
	}
}

func handleDelete(server *DefaultServer) RouteHandler {
	return func(req CoapRequest) CoapResponse {
		id := req.GetAttribute("id")

		server.delete(id)

		msg := NewMessageOfType(MessageAcknowledgment, req.GetMessage().MessageID)
		msg.Token = req.GetMessage().Token
		msg.Code = CoapCodeDeleted

		return NewResponseWithMessage(msg)
	}
}
