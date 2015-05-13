package basic

import "github.com/zubairhamed/lwm2m/core"

type ConnectivityStatistics struct {

}

func (o *ConnectivityStatistics) OnRead(r *core.ResourceModel, resourceId int) core.ResourceValue {
    return core.NewEmptyValue()
}
