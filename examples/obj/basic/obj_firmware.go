package basic

import (
    . "github.com/zubairhamed/lwm2m/api"
    "github.com/zubairhamed/lwm2m/core"
)

type Firmware struct {

}

func (o *Firmware) OnDelete(instanceId int) (bool) {
    return true
}

func (o *Firmware) OnRead(instanceId int, resourceId int) (ResourceValue) {
    return core.NewEmptyValue()
}

func (o *Firmware) OnWrite(instanceId int, resourceId int) (bool) {
    return true
}
