package basic

import (
	. "github.com/zubairhamed/go-lwm2m/api"
	"github.com/zubairhamed/go-lwm2m/core"
	"github.com/zubairhamed/go-lwm2m/objects/oma"
	"github.com/zubairhamed/goap"
)

type AccessControl struct {
	Model ObjectModel
	Data  *core.ObjectsData
}

func (o *AccessControl) OnExecute(instanceId int, resourceId int) goap.CoapCode {
	return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *AccessControl) OnCreate(instanceId int, resourceId int) goap.CoapCode {
	return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *AccessControl) OnDelete(instanceId int) goap.CoapCode {
	return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *AccessControl) OnRead(instanceId int, resourceId int) (ResponseValue, goap.CoapCode) {
	return core.NewEmptyValue(), goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *AccessControl) OnWrite(instanceId int, resourceId int) goap.CoapCode {
	return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func NewExampleAccessControlObject(reg Registry) *AccessControl {
	data := &core.ObjectsData{
		Data: make(map[string]interface{}),
	}

	data.Put("/0/0", 1)
	data.Put("/0/1", 0)
	data.Put("/0/2/101", []byte{0, 15})
	data.Put("/0/3", 101)

	data.Put("1/0", 1)
	data.Put("1/1", 1)
	data.Put("1/2/102", []byte{0, 15})
	data.Put("1/3", 102)

	data.Put("2/0", 3)
	data.Put("2/1", 0)
	data.Put("2/2/101", []byte{0, 15})
	data.Put("2/2/102", []byte{0, 1})
	data.Put("2/3", 101)

	data.Put("3/0", 4)
	data.Put("3/1", 0)
	data.Put("3/2/101", []byte{0, 1})
	data.Put("3/2/0", []byte{0, 1})
	data.Put("3/3", 101)

	data.Put("4/0", 5)
	data.Put("4/1", 65535)
	data.Put("4/2/101", []byte{0, 16})
	data.Put("4/3", 65535)

	return &AccessControl{
		Model: reg.GetModel(oma.OBJECT_LWM2M_SECURITY),
		Data:  data,
	}
}
