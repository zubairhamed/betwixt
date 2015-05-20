package basic

import (
    "github.com/zubairhamed/lwm2m/core"
    . "github.com/zubairhamed/lwm2m/api"
)

type ConnectivityMonitoring struct {

}

func (o *ConnectivityMonitoring) OnDelete(instanceId int) (bool) {
    return true
}

func (o *ConnectivityMonitoring) OnRead(instanceId int, resourceId int) (ResourceValue) {
    return core.NewEmptyValue()
}

func (o *ConnectivityMonitoring) OnWrite(instanceId int, resourceId int) (bool) {
    return true
}


/*
Network Bearer              0           0
Available Network Bearer    1           0
Radio signal strength       2           92
Link Quality                3           2
IP Addresses                4       0   192.168.0.100
Parent IP Addresses         5       0   192.168.1.1
Link Utilization            6           5
APN                         7       0   internet
*/