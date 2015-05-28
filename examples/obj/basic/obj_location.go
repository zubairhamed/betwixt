package basic

import (
	. "github.com/zubairhamed/go-lwm2m/api"
	"github.com/zubairhamed/go-lwm2m/core"
	"github.com/zubairhamed/go-lwm2m/objects/oma"
	"github.com/zubairhamed/goap"
)

type Location struct {
	Model ObjectModel
	Data  *core.ObjectsData
}

func (o *Location) OnExecute(instanceId int, resourceId int, req Request) goap.CoapCode {
	return goap.COAPCODE_401_UNAUTHORIZED
}

func (o *Location) OnCreate(instanceId int, resourceId int, req Request) goap.CoapCode {
	return goap.COAPCODE_401_UNAUTHORIZED
}

func (o *Location) OnDelete(instanceId int, req Request) goap.CoapCode {
	return goap.COAPCODE_401_UNAUTHORIZED
}

func (o *Location) OnRead(instanceId int, resourceId int, req Request) (ResponseValue, goap.CoapCode) {
	return core.NewEmptyValue(), goap.COAPCODE_401_UNAUTHORIZED
}

func (o *Location) OnWrite(instanceId int, resourceId int, req Request) goap.CoapCode {
	return goap.COAPCODE_401_UNAUTHORIZED
}

func NewExampleLocationObject(reg Registry) *Location {
	data := &core.ObjectsData{
		Data: make(map[string]interface{}),
	}

	return &Location{
		Model: reg.GetModel(oma.OBJECT_LWM2M_LOCATION),
		Data:  data,
	}
}
