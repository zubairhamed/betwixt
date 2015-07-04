package basic

import (
	. "github.com/zubairhamed/betwixt"
)

type LocationObject struct {
	Model ObjectDefinition
}

func (o *LocationObject) OnExecute(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *LocationObject) OnCreate(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *LocationObject) OnDelete(instanceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *LocationObject) OnRead(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *LocationObject) OnWrite(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func NewExampleLocationObject(reg Registry) *LocationObject {
	return &LocationObject{
		Model: reg.GetDefinition(OMA_OBJECT_LWM2M_LOCATION),
	}
}
