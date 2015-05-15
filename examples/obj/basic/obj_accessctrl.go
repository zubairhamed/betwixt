package basic

import (
    "github.com/zubairhamed/lwm2m/core"
)

type AccessControl struct {
    Model       *core.ObjectModel
}


func (o *AccessControl) OnRead(instanceId int, resourceId int) (core.ResourceValue) {
    return core.NewEmptyValue()
}
