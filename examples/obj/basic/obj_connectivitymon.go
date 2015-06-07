package basic

import (
	. "github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/betwixt/core/response"
	"github.com/zubairhamed/betwixt/core/values"
	"github.com/zubairhamed/betwixt/objects/oma"
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
	if resourceId == -1 {
		// Read Object Instance
	} else {
		// Read Resource Instance
		var val ResponseValue

		switch resourceId {
		case 0:
			val = values.Integer(0)
			break

		case 1:
			val = values.Integer(0)
			break

		case 2:
			val = values.Integer(92)
			break

		case 3:
			val = values.Integer(2)
			break

		case 4:
			val = values.String("192.168.0.100")
			break

		case 5:
			val = values.String("192.168.1.1")
			break

		case 6:
			val = values.Integer(5)
			break

		case 7:
			val = values.String("internet")
			break

		case 8:
			val = values.String("")
			break

		case 9:
			val = values.String("")
			break

		case 10:
			val = values.String("")
			break

		default:
			break

		}
		return response.Content(val)
	}
	return response.NotFound()
}

func (o *ConnectivityMonitoringObject) OnWrite(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func NewExampleConnectivityMonitoringObject(reg Registry) *ConnectivityMonitoringObject {
	return &ConnectivityMonitoringObject{
		Model: reg.GetModel(oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING),
	}
}
