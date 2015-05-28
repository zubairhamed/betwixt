package basic

import (
	. "github.com/zubairhamed/go-lwm2m/api"
	"github.com/zubairhamed/go-lwm2m/core"
	"github.com/zubairhamed/go-lwm2m/objects/oma"
)

type ConnectivityMonitoring struct {
	Model ObjectModel
	Data  *core.ObjectsData
}

func (o *ConnectivityMonitoring) OnExecute(instanceId int, resourceId int, req Request) (Response) {
	return core.NewUnauthorizedResponse()
}

func (o *ConnectivityMonitoring) OnCreate(instanceId int, resourceId int, req Request) (Response) {
	return core.NewUnauthorizedResponse()
}

func (o *ConnectivityMonitoring) OnDelete(instanceId int, req Request) (Response) {
	return core.NewUnauthorizedResponse()
}

func (o *ConnectivityMonitoring) OnRead(instanceId int, resourceId int, req Request) (Response) {
	return core.NewUnauthorizedResponse()
}

func (o *ConnectivityMonitoring) OnWrite(instanceId int, resourceId int, req Request) (Response) {
	return core.NewUnauthorizedResponse()
}

func NewExampleConnectivityMonitoringObject(reg Registry) *ConnectivityMonitoring {
	data := &core.ObjectsData{
		Data: make(map[string]interface{}),
	}

	data.Put("/0/0", 0)
	data.Put("/0/1", 0)
	data.Put("/0/2", 92)
	data.Put("/0/3", 2)
	data.Put("/0/4/0", "192.168.0.100")
	data.Put("/0/5/0", "192.168.1.1")
	data.Put("/0/6", 5)
	data.Put("/0/7/0", "internet")

	return &ConnectivityMonitoring{
		Model: reg.GetModel(oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING),
		Data:  data,
	}
}

/*

*/
