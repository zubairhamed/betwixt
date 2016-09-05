package objects

import (
	. "github.com/zubairhamed/betwixt"
)

type ConnectivityStatisticsObject struct {
	Model ObjectDefinition
}

func (o *ConnectivityStatisticsObject) OnExecute(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *ConnectivityStatisticsObject) OnCreate(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *ConnectivityStatisticsObject) OnDelete(instanceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *ConnectivityStatisticsObject) OnRead(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func (o *ConnectivityStatisticsObject) OnWrite(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Unauthorized()
}

func NewExampleConnectivityStatisticsObject(reg Registry) *ConnectivityStatisticsObject {
	return &ConnectivityStatisticsObject{
		Model: reg.GetDefinition(OMA_OBJECT_LWM2M_CONNECTIVITY_STATISTICS),
	}
}
