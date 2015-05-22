package basic

import (
    "github.com/zubairhamed/lwm2m/core"
    . "github.com/zubairhamed/lwm2m/api"
    "github.com/zubairhamed/goap"
)

type ConnectivityMonitoring struct {

}

func (o *ConnectivityMonitoring) OnExecute(instanceId int, resourceId int) (goap.CoapCode) {
    return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *ConnectivityMonitoring) OnCreate(instanceId int, resourceId int) (goap.CoapCode) {
    return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *ConnectivityMonitoring) OnDelete(instanceId int) (goap.CoapCode) {
    return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *ConnectivityMonitoring) OnRead(instanceId int, resourceId int) (ResponseValue, goap.CoapCode) {
    return core.NewEmptyValue(), goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *ConnectivityMonitoring) OnWrite(instanceId int, resourceId int) (goap.CoapCode) {
    return goap.COAPCODE_405_METHOD_NOT_ALLOWED
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