package basic


import (
    . "github.com/zubairhamed/lwm2m/api"
    "github.com/zubairhamed/lwm2m/core"
)


type Security struct {

}
func (o *Security) OnRead(instanceId int, resourceId int) (ResourceValue) {
    return core.NewEmptyValue()
}
