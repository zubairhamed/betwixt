package basic

import "github.com/zubairhamed/lwm2m/core"

type Security struct {

}

func (o *Security) OnRead(r *core.ResourceModel, resourceId int) core.ResourceValue {
    return core.NewEmptyValue()
}
