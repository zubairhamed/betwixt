package basic

import (
    . "github.com/zubairhamed/lwm2m/api"
    "github.com/zubairhamed/lwm2m/core"
)

type Firmware struct {

}

func (o *Firmware) OnRead(r ResourceModel, resourceId int) ResourceValue {
    return core.NewEmptyValue()
}
