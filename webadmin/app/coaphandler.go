package app

import (
	"github.com/zubairhamed/canopus"
	"log"
)

func FnCoapRegisterClient(b *BetwixtWebApp) canopus.RouteHandler {
	return func(req canopus.CoapRequest) canopus.CoapResponse {
		ep := req.GetURIQuery("ep")

		// lt := req.GetUriQuery("lt")
		// sms := req.GetUriQuery("sms")
		// binding := req.GetUriQuery("b")

		resources := canopus.CoreResourcesFromString(req.GetMessage().Payload.String())
		clientId, err := b.register(ep, req.GetAddress().String(), resources)
		if err != nil {
			log.Println("Error registering client ", ep)
		}

		msg := canopus.NewMessageOfType(canopus.MessageAcknowledgment, req.GetMessage().MessageID)
		msg.Token = req.GetMessage().Token
		msg.AddOption(canopus.OptionLocationPath, "rd/" + clientId)
		msg.Code = canopus.CoapCodeCreated

		return canopus.NewResponseWithMessage(msg)
	}
}

func FnCoapUpdateClient(b *BetwixtWebApp) canopus.RouteHandler {
	return func(req canopus.CoapRequest) canopus.CoapResponse {
		id := req.GetAttribute("id")

		b.update(id)

		msg := canopus.NewMessageOfType(canopus.MessageAcknowledgment, req.GetMessage().MessageID)
		msg.Token = req.GetMessage().Token
		msg.Code = canopus.CoapCodeChanged

		return canopus.NewResponseWithMessage(msg)
	}
}

func FnCoapDeleteClient(b *BetwixtWebApp) canopus.RouteHandler {
	return func(req canopus.CoapRequest) canopus.CoapResponse {
		id := req.GetAttribute("id")

		b.delete(id)

		msg := canopus.NewMessageOfType(canopus.MessageAcknowledgment, req.GetMessage().MessageID)
		msg.Token = req.GetMessage().Token
		msg.Code = canopus.CoapCodeDeleted

		return canopus.NewResponseWithMessage(msg)
	}
}

