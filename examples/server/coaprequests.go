package server

import (
	"github.com/zubairhamed/canopus"
	"github.com/zubairhamed/go-commons/logging"
	. "github.com/zubairhamed/go-commons/network"
)

func SetupCoapRoutes(server *DefaultServer) {
	coap := server.coapServer

	coap.On(canopus.EVT_MESSAGE, func() {
		server.stats.IncrementCoapRequestsCount()
	})

	coap.NewRoute("rd", canopus.POST, handleRegister(server))
	coap.NewRoute("rd/{id}", canopus.PUT, handleUpdate(server))
	coap.NewRoute("rd/{id}", canopus.DELETE, handleDelete(server))
}

func handleRegister(server *DefaultServer) RouteHandler {
	return func(r Request) Response {
		req := r.(*canopus.CoapRequest)

		ep := req.GetUriQuery("ep")
		// lt := req.GetUriQuery("lt")
		// sms := req.GetUriQuery("sms")
		// binding := req.GetUriQuery("b")

		resources := canopus.CoreResourcesFromString(req.GetMessage().Payload.String())
		clientId, err := server.register(ep, req.GetAddress().String(), resources)
		if err != nil {
			logging.LogWarn("Error registering client ", ep)
		}

		msg := canopus.NewMessageOfType(canopus.TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
		msg.Token = req.GetMessage().Token
		msg.AddOption(canopus.OPTION_LOCATION_PATH, "rd/"+clientId)
		msg.Code = canopus.COAPCODE_201_CREATED

		return canopus.NewResponseWithMessage(msg)
	}
}

func handleUpdate(server *DefaultServer) RouteHandler {
	return func(r Request) Response {
		req := r.(*canopus.CoapRequest)
		id := req.GetAttribute("id")

		server.update(id)

		msg := canopus.NewMessageOfType(canopus.TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
		msg.Token = req.GetMessage().Token
		msg.Code = canopus.COAPCODE_204_CHANGED

		return canopus.NewResponseWithMessage(msg)
	}
}

func handleDelete(server *DefaultServer) RouteHandler {
	return func(r Request) Response {
		req := r.(*canopus.CoapRequest)
		id := req.GetAttribute("id")

		server.delete(id)

		msg := canopus.NewMessageOfType(canopus.TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
		msg.Token = req.GetMessage().Token
		msg.Code = canopus.COAPCODE_202_DELETED

		return canopus.NewResponseWithMessage(msg)
	}
}
