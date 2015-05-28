package basic

import (
	. "github.com/zubairhamed/go-lwm2m/api"
	"github.com/zubairhamed/go-lwm2m/core"
	"github.com/zubairhamed/go-lwm2m/objects/oma"
)

type Firmware struct {
	Model ObjectModel
	Data  *core.ObjectsData
}

func (o *Firmware) OnExecute(instanceId int, resourceId int, req Request) (Response) {
	return core.NewUnauthorizedResponse()
}

func (o *Firmware) OnCreate(instanceId int, resourceId int, req Request) (Response) {
	return core.NewUnauthorizedResponse()
}

func (o *Firmware) OnDelete(instanceId int, req Request) (Response) {
	return core.NewUnauthorizedResponse()
}

func (o *Firmware) OnRead(instanceId int, resourceId int, req Request) (Response) {
	return core.NewUnauthorizedResponse()
}

func (o *Firmware) OnWrite(instanceId int, resourceId int, req Request) (Response) {
	return core.NewUnauthorizedResponse()
}

func NewExampleFirmwareUpdateObject(reg Registry) *Firmware {
	data := &core.ObjectsData{
		Data: make(map[string]interface{}),
	}

	return &Firmware{
		Model: reg.GetModel(oma.OBJECT_LWM2M_FIRMWARE_UPDATE),
		Data:  data,
	}
}
