package ev3

import (
	. "github.com/zubairhamed/betwixt"
)

type SecurityObject struct {
	Model ObjectDefinition
}

func (o *SecurityObject) OnExecute(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *SecurityObject) OnCreate(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *SecurityObject) OnDelete(instanceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *SecurityObject) OnRead(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *SecurityObject) OnWrite(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func NewExampleSecurityObject(reg Registry) *SecurityObject {
	return &SecurityObject{
		Model: reg.GetDefinition(OMA_OBJECT_LWM2M_SECURITY),
	}
}
