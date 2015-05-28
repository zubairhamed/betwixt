package basic

import (
	. "github.com/zubairhamed/go-lwm2m/api"
	"github.com/zubairhamed/go-lwm2m/core"
	"github.com/zubairhamed/go-lwm2m/objects/oma"
	"github.com/zubairhamed/goap"
)

type ConnectivityStatistics struct {
	Model ObjectModel
	Data  *core.ObjectsData
}

func (o *ConnectivityStatistics) OnExecute(instanceId int, resourceId int, req Request) goap.CoapCode {
	return goap.COAPCODE_401_UNAUTHORIZED
}

func (o *ConnectivityStatistics) OnCreate(instanceId int, resourceId int, req Request) goap.CoapCode {
	return goap.COAPCODE_401_UNAUTHORIZED
}

func (o *ConnectivityStatistics) OnDelete(instanceId int, req Request) goap.CoapCode {
	return goap.COAPCODE_401_UNAUTHORIZED
}

func (o *ConnectivityStatistics) OnRead(instanceId int, resourceId int, req Request) (ResponseValue, goap.CoapCode) {
	return core.NewEmptyValue(), goap.COAPCODE_401_UNAUTHORIZED
}

func (o *ConnectivityStatistics) OnWrite(instanceId int, resourceId int, req Request) goap.CoapCode {
	return goap.COAPCODE_401_UNAUTHORIZED
}

func NewExampleConnectivityStatisticsObject(reg Registry) *ConnectivityStatistics {
	data := &core.ObjectsData{
		Data: make(map[string]interface{}),
	}

	return &ConnectivityStatistics{
		Model: reg.GetModel(oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS),
		Data:  data,
	}
}
