package basic

import (
	. "github.com/zubairhamed/betwixt/api"
	"github.com/zubairhamed/betwixt/core"
	"github.com/zubairhamed/betwixt/core/response"
	"github.com/zubairhamed/betwixt/objects/oma"
)

type ConnectivityStatisticsObject struct {
	Model ObjectModel
	Data  *core.ObjectsData
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
	data := &core.ObjectsData{
		Data: make(map[string]interface{}),
	}

	return &ConnectivityStatisticsObject{
		Model: reg.GetModel(oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS),
		Data:  data,
	}
}
