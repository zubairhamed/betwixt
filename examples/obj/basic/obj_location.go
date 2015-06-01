package basic

import (
	. "github.com/zubairhamed/go-lwm2m/api"
	"github.com/zubairhamed/go-lwm2m/core"
	"github.com/zubairhamed/go-lwm2m/core/response"
	"github.com/zubairhamed/go-lwm2m/objects/oma"
)

type LocationObject struct {
	Model ObjectModel
	Data  *core.ObjectsData
}

func (o *LocationObject) OnExecute(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func (o *LocationObject) OnCreate(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func (o *LocationObject) OnDelete(instanceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func (o *LocationObject) OnRead(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func (o *LocationObject) OnWrite(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return response.Unauthorized()
}

func NewExampleLocationObject(reg Registry) *LocationObject {
	data := &core.ObjectsData{
		Data: make(map[string]interface{}),
	}

	return &LocationObject{
		Model: reg.GetModel(oma.OBJECT_LWM2M_LOCATION),
		Data:  data,
	}
}
