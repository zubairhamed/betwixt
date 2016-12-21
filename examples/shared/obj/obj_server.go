package obj

import (
	. "github.com/zubairhamed/betwixt"
)

type ServerObject struct {
	Model ObjectDefinition
}

func (o *ServerObject) OnExecute(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *ServerObject) OnCreate(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *ServerObject) OnDelete(instanceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *ServerObject) OnRead(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *ServerObject) OnWrite(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func NewExampleServerObject(reg Registry) *ServerObject {
	/*
		data.Put("/1/0", 101)
		data.Put("/1/1", 86400)
		data.Put("/1/2", 300)
		data.Put("/1/3", 6000)
		data.Put("/1/5", 86400)
		data.Put("/1/6", true)
		data.Put("/1/7", "U")

		data.Put("/2/0", 102)
		data.Put("/2/1", 86400)
		data.Put("/2/2", 60)
		data.Put("/2/3", 6000)
		data.Put("/2/5", 86400)
		data.Put("/2/6", false)
		data.Put("/2/7", "UQ")
	*/

	return &ServerObject{
		Model: reg.GetDefinition(OMA_OBJECT_LWM2M_SERVER),
	}
}
