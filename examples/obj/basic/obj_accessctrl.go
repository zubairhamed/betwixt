package basic

import (
    "github.com/zubairhamed/lwm2m/core"
    . "github.com/zubairhamed/lwm2m/api"
)

type AccessControl struct {
    Model       ObjectModel
}


func (o *AccessControl) OnRead(instanceId int, resourceId int) (ResourceValue) {
    return core.NewEmptyValue()
}
