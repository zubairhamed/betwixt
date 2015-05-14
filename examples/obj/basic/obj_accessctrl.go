package basic

import (
    "github.com/zubairhamed/lwm2m/core"
)

type AccessControl struct {

}


func (o *AccessControl) OnRead(instanceId int, resourceId int) (core.ResourceValue) {
    return core.NewEmptyValue()
}
