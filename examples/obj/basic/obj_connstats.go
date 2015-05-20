package basic

import (
    "github.com/zubairhamed/lwm2m/core"
    . "github.com/zubairhamed/lwm2m/api"
)


type ConnectivityStatistics struct {

}

func (o *ConnectivityStatistics) OnRead(instanceId int, resourceId int) (ResourceValue) {
    return core.NewEmptyValue()
}
