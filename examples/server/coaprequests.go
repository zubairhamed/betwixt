package server

import (
	. "github.com/zubairhamed/canopus"
	"log"
)

func SetupCoapRoutes(server *DefaultServer) {
	coap := server.coapServer

	coap.OnMessage(func (msg *Message, inbound bool){
		server.stats.IncrementCoapRequestsCount()
	})

	coap.Post("/rd", handleRegister(server))
	coap.Put("/rd/:id", handleUpdate(server))
	coap.Delete("/rd/:id", handleDelete(server))
}

func handleRegister(server *DefaultServer) RouteHandler {
	return func(req *Request) *Response {
		ep := req.GetUriQuery("ep")
		// lt := req.GetUriQuery("lt")
		// sms := req.GetUriQuery("sms")
		// binding := req.GetUriQuery("b")

		resources := CoreResourcesFromString(req.GetMessage().Payload.String())
		clientId, err := server.register(ep, req.GetAddress().String(), resources)
		if err != nil {
			log.Println("Error registering client ", ep)
		}

		msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
		msg.Token = req.GetMessage().Token
		msg.AddOption(OPTION_LOCATION_PATH, "rd/"+clientId)
		msg.Code = COAPCODE_201_CREATED

		return NewResponseWithMessage(msg)
	}
}

func handleUpdate(server *DefaultServer) RouteHandler {
	return func(req *Request) *Response {
		id := req.GetAttribute("id")

		server.update(id)

		msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
		msg.Token = req.GetMessage().Token
		msg.Code = COAPCODE_204_CHANGED

		return NewResponseWithMessage(msg)
	}
}

func handleDelete(server *DefaultServer) RouteHandler {
	return func(req *Request) *Response {
		id := req.GetAttribute("id")

		server.delete(id)

		msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
		msg.Token = req.GetMessage().Token
		msg.Code = COAPCODE_202_DELETED

		return NewResponseWithMessage(msg)
	}
}
