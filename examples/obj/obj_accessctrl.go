package obj

import (
    "github.com/zubairhamed/lwm2m/core"
)

type AccessControl struct {

}

func (o *AccessControl) OnRead(r *core.ResourceModel, resourceId int) core.ResourceValue {
    return nil
}
