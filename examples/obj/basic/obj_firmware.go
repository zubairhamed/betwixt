package basic

import (
	. "github.com/zubairhamed/go-lwm2m/api"
	"github.com/zubairhamed/go-lwm2m/core"
	"github.com/zubairhamed/go-lwm2m/objects/oma"
	"github.com/zubairhamed/goap"
)

type Firmware struct {
	Model ObjectModel
	Data  *core.ObjectsData
}

func (o *Firmware) OnExecute(instanceId int, resourceId int) goap.CoapCode {
	return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *Firmware) OnCreate(instanceId int, resourceId int) goap.CoapCode {
	return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *Firmware) OnDelete(instanceId int) goap.CoapCode {
	return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *Firmware) OnRead(instanceId int, resourceId int) (ResponseValue, goap.CoapCode) {
	return core.NewEmptyValue(), goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *Firmware) OnWrite(instanceId int, resourceId int) goap.CoapCode {
	return 0
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
