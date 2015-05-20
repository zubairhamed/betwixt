package basic

import (
    "github.com/zubairhamed/lwm2m/core"
    . "github.com/zubairhamed/lwm2m/api"
)

type ConnectivityMonitoring struct {

}

func (o *ConnectivityMonitoring) OnRead(r ResourceModel, resourceId int) ResourceValue {
    return core.NewEmptyValue()
}
