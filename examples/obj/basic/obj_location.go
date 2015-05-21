package basic

import (
    . "github.com/zubairhamed/lwm2m/api"
    "github.com/zubairhamed/lwm2m/core"
    "github.com/zubairhamed/goap"
)

type Location struct {

}

func (o *Location) OnExecute(instanceId int, resourceId int) (goap.CoapCode) {
    return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *Location) OnCreate(instanceId int, resourceId int) (goap.CoapCode) {
    return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *Location) OnDelete(instanceId int) (goap.CoapCode) {
    return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *Location) OnRead(instanceId int, resourceId int) (ResponseValue, goap.CoapCode) {
    return core.NewEmptyValue(),  goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *Location) OnWrite(instanceId int, resourceId int) (goap.CoapCode) {
    return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

