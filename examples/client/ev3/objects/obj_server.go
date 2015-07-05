package ev3

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
	return &ServerObject{
		Model: reg.GetDefinition(OMA_OBJECT_LWM2M_SERVER),
	}
}
