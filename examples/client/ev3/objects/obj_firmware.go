package ev3

import (
	. "github.com/zubairhamed/betwixt"
)

type FirmwareObject struct {
	Model ObjectDefinition
}

func (o *FirmwareObject) OnExecute(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *FirmwareObject) OnCreate(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *FirmwareObject) OnDelete(instanceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *FirmwareObject) OnRead(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *FirmwareObject) OnWrite(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func NewExampleFirmwareUpdateObject(reg Registry) *FirmwareObject {
	return &FirmwareObject{
		Model: reg.GetDefinition(OMA_OBJECT_LWM2M_FIRMWARE_UPDATE),
	}
}
