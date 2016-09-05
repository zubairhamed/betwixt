package objects

import (
	"log"

	. "github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/canopus"
)

type FirmwareObject struct {
	Model ObjectDefinition
}

func (o *FirmwareObject) OnExecute(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	log.Println("Executing: ", instanceId, resourceId)
	canopus.PrintMessage(req.GetMessage())

	return Unauthorized()
}

func (o *FirmwareObject) OnCreate(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	log.Println("Creating: ", instanceId, resourceId)
	canopus.PrintMessage(req.GetMessage())

	return Unauthorized()
}

func (o *FirmwareObject) OnDelete(instanceId int, req Lwm2mRequest) Lwm2mResponse {
	log.Println("Deleting: ", instanceId)
	canopus.PrintMessage(req.GetMessage())

	return Unauthorized()
}

func (o *FirmwareObject) OnRead(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	log.Println("Reading: ", instanceId, resourceId)
	canopus.PrintMessage(req.GetMessage())
	return Unauthorized()
}

func (o *FirmwareObject) OnWrite(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	log.Println("Writing: ", instanceId, resourceId)
	canopus.PrintMessage(req.GetMessage())
	return Unauthorized()
}

func NewExampleFirmwareUpdateObject(reg Registry) *FirmwareObject {
	return &FirmwareObject{
		Model: reg.GetDefinition(OMA_OBJECT_LWM2M_FIRMWARE_UPDATE),
	}
}
