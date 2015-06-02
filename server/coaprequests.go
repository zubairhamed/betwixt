package server

import (
	"github.com/zubairhamed/canopus"
	. "github.com/zubairhamed/go-commons/network"
	"log"
)

func SetupCoapRoutes(server *DefaultServer) {
	coap := server.coapServer

	coap.NewRoute("rd", canopus.POST, handleRegister(server))
	coap.NewRoute("rd/{id}", canopus.PUT, handleUpdate(server))
	coap.NewRoute("rd/{id}", canopus.DELETE, handleDelete(server))
}

func handleRegister(server *DefaultServer) RouteHandler {
	return func(r Request) Response {
		req := r.(*canopus.CoapRequest)
		ep := req.GetUriQuery("ep")

		clientId, err := server.register(ep, req.GetAddress().String())
		if err != nil {
			log.Println("Error registering client ", ep)
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
