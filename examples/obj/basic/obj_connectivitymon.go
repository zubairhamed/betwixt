package basic

import "github.com/zubairhamed/lwm2m/core"

type ConnectivityMonitoring struct {

}

func (o *ConnectivityMonitoring) OnRead(r *core.ResourceModel, resourceId int) core.ResourceValue {
    return core.NewEmptyValue()
}
