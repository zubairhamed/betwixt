package basic

import (
    . "github.com/zubairhamed/lwm2m/api"
    "github.com/zubairhamed/lwm2m/core"
)

type Location struct {

}

func (o *Location) OnExecute(instanceId int, resourceId int) (bool, int) {
    return true, 0
}

func (o *Location) OnCreate(instanceId int, resourceId int) (bool, int) {
    return true, 0
}

func (o *Location) OnDelete(instanceId int) (bool) {
    return true
}

func (o *Location) OnRead(instanceId int, resourceId int) (ResponseValue) {
    return core.NewEmptyValue()
}

func (o *Location) OnWrite(instanceId int, resourceId int) (bool) {
    return true
}

