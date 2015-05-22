package basic

import (
	. "github.com/zubairhamed/go-lwm2m/api"
	"github.com/zubairhamed/go-lwm2m/core"
	"github.com/zubairhamed/go-lwm2m/objects/oma"
	"github.com/zubairhamed/goap"
)

type Server struct {
	Model ObjectModel
	Data  *core.ObjectsData
}

func (o *Server) OnExecute(instanceId int, resourceId int) goap.CoapCode {
	return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *Server) OnCreate(instanceId int, resourceId int) goap.CoapCode {
	return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *Server) OnDelete(instanceId int) goap.CoapCode {
	return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *Server) OnRead(instanceId int, resourceId int) (ResponseValue, goap.CoapCode) {
	return core.NewEmptyValue(), goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *Server) OnWrite(instanceId int, resourceId int) goap.CoapCode {
	return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func NewExampleServerObject(reg Registry) *Server {
	data := &core.ObjectsData{
		Data: make(map[string]interface{}),
	}

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

	return &Server{
		Model: reg.GetModel(oma.OBJECT_LWM2M_SERVER),
		Data:  data,
	}
}
