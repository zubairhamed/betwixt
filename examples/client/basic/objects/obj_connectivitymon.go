package basic

import (
	. "github.com/zubairhamed/betwixt"
)

type ConnectivityMonitoringObject struct {
	Model ObjectDefinition
}

func (o *ConnectivityMonitoringObject) OnExecute(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *ConnectivityMonitoringObject) OnCreate(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *ConnectivityMonitoringObject) OnDelete(instanceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *ConnectivityMonitoringObject) OnRead(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *ConnectivityMonitoringObject) OnWrite(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func NewExampleConnectivityMonitoringObject(reg Registry) *ConnectivityMonitoringObject {
	return &ConnectivityMonitoringObject{
		Model: reg.GetDefinition(OMA_OBJECT_LWM2M_CONNECTIVITY_MONITORING),
	}
}
