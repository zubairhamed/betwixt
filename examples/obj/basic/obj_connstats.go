package basic

import (
    "github.com/zubairhamed/lwm2m/core"
    . "github.com/zubairhamed/lwm2m/api"
    "github.com/zubairhamed/goap"
)


type ConnectivityStatistics struct {

}

func (o *ConnectivityStatistics) OnExecute(instanceId int, resourceId int) (goap.CoapCode) {
    return 0
}

func (o *ConnectivityStatistics) OnCreate(instanceId int, resourceId int) (goap.CoapCode) {
    return 0
}

func (o *ConnectivityStatistics) OnDelete(instanceId int) (goap.CoapCode) {
    return 0
}

func (o *ConnectivityStatistics) OnRead(instanceId int, resourceId int) (ResponseValue, goap.CoapCode) {
    return core.NewEmptyValue(), 0
}

func (o *ConnectivityStatistics) OnWrite(instanceId int, resourceId int) (goap.CoapCode) {
    return 0
}
