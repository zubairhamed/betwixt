package basic

import (
	. "github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/betwixt/core/response"
	"github.com/zubairhamed/betwixt/objects/oma"
)

type SecurityObject struct {
	Model ObjectDefinition
}

func (o *SecurityObject) OnExecute(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func (o *SecurityObject) OnCreate(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func (o *SecurityObject) OnDelete(instanceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func (o *SecurityObject) OnRead(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func (o *SecurityObject) OnWrite(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func NewExampleSecurityObject(reg Registry) *SecurityObject {
	/*
	data.Put("/0/0", "coap://bootstrap.example.com")
	data.Put("/0/1", true)
	data.Put("/0/2", 0)
	data.Put("/0/3", "[identity string]")
	data.Put("/0/4", "[secret key data]")
	data.Put("/0/10", 0)
	data.Put("/0/11", 3600)

	data.Put("/1/0", "coap://server1.example.com")
	data.Put("/1/1", false)
	data.Put("/1/2", 0)
	data.Put("/1/3", "[identity string]")
	data.Put("/1/4", "[secret key data]")
	data.Put("/1/10", 101)
	data.Put("/1/11", 0)

	data.Put("/2/0", "coap://server2.example.com")
	data.Put("/2/1", false)
	data.Put("/2/2", 0)
	data.Put("/2/3", "[identity string]")
	data.Put("/2/4", "[secret key data]")
	data.Put("/2/10", 102)
	data.Put("/2/11", 0)
	*/

	return &SecurityObject{
		Model: reg.GetModel(oma.OBJECT_LWM2M_SECURITY),
	}
}
