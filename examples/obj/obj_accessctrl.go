package obj

import (
    "github.com/zubairhamed/lwm2m/core"
)

type AccessControl struct {

}

func (o *AccessControl) OnRead(t core.LWM2MObjectType, m *core.ObjectModel, i *core.ObjectInstance, r *core.ResourceModel) core.ResourceValue {
    return nil
}
