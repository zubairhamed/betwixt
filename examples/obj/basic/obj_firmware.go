package basic

import (
    . "github.com/zubairhamed/lwm2m/api"
    "github.com/zubairhamed/lwm2m/core"
    "github.com/zubairhamed/goap"
)

type Firmware struct {

}

func (o *Firmware) OnExecute(instanceId int, resourceId int) (goap.CoapCode) {
    return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *Firmware) OnCreate(instanceId int, resourceId int) (goap.CoapCode) {
    return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *Firmware) OnDelete(instanceId int) (goap.CoapCode) {
    return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *Firmware) OnRead(instanceId int, resourceId int) (ResponseValue, goap.CoapCode) {
    return core.NewEmptyValue(), goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *Firmware) OnWrite(instanceId int, resourceId int) (goap.CoapCode) {
    return 0
}
