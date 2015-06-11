package basic

import (
	. "github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/betwixt/objectdefs/oma"
	"github.com/zubairhamed/betwixt/core/response"
)

type ConnectivityMonitoringObject struct {
	Model ObjectDefinition
}

func (o *ConnectivityMonitoringObject) OnExecute(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func (o *ConnectivityMonitoringObject) OnCreate(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func (o *ConnectivityMonitoringObject) OnDelete(instanceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func (o *ConnectivityMonitoringObject) OnRead(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func (o *ConnectivityMonitoringObject) OnWrite(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func NewExampleConnectivityMonitoringObject(reg Registry) *ConnectivityMonitoringObject {
	return &ConnectivityMonitoringObject{
		Model: reg.GetDefinition(oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING),
	}
}
