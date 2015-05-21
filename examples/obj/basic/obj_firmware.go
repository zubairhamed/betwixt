package basic

import (
    . "github.com/zubairhamed/lwm2m/api"
    "github.com/zubairhamed/lwm2m/core"
)

type Firmware struct {

}

func (o *Firmware) OnExecute(instanceId int, resourceId int) (bool, int) {
    return true, 0
}

func (o *Firmware) OnCreate(instanceId int, resourceId int) (bool, int) {
    return true, 0
}

func (o *Firmware) OnDelete(instanceId int) (bool) {
    return true
}

func (o *Firmware) OnRead(instanceId int, resourceId int) (ResponseValue) {
    return core.NewEmptyValue()
}

func (o *Firmware) OnWrite(instanceId int, resourceId int) (bool) {
    return true
}
