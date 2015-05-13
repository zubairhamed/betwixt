package basic

import "github.com/zubairhamed/lwm2m/core"

type Location struct {

}

func (o *Location) OnRead(r *core.ResourceModel, resourceId int) core.ResourceValue {
    return core.NewEmptyValue()
}
