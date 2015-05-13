package basic

import "github.com/zubairhamed/lwm2m/core"

type Firmware struct {

}

func (o *Firmware) OnRead(r *core.ResourceModel, resourceId int) core.ResourceValue {
    return core.NewEmptyValue()
}
