package basic

import (
	. "github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/betwixt/core/response"
	"github.com/zubairhamed/betwixt/objectdefs/oma"
)

type ConnectivityStatisticsObject struct {
	Model ObjectDefinition
}

func (o *ConnectivityStatisticsObject) OnExecute(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func (o *ConnectivityStatisticsObject) OnCreate(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func (o *ConnectivityStatisticsObject) OnDelete(instanceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func (o *ConnectivityStatisticsObject) OnRead(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func (o *ConnectivityStatisticsObject) OnWrite(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func NewExampleConnectivityStatisticsObject(reg Registry) *ConnectivityStatisticsObject {
	return &ConnectivityStatisticsObject{
		Model: reg.GetDefinition(oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS),
	}
}
