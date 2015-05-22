package basic


import (
    . "github.com/zubairhamed/go-lwm2m/api"
    "github.com/zubairhamed/go-lwm2m/core"
    "github.com/zubairhamed/goap"
    "github.com/zubairhamed/go-lwm2m/objects/oma"
)


type Security struct {
    Model       ObjectModel
    Data        *core.ObjectsData
}

func (o *Security) OnExecute(instanceId int, resourceId int) (goap.CoapCode) {
    return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *Security) OnCreate(instanceId int, resourceId int) (goap.CoapCode) {
    return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *Security) OnDelete(instanceId int) (goap.CoapCode) {
    return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *Security) OnRead(instanceId int, resourceId int) (ResponseValue, goap.CoapCode) {
    return core.NewEmptyValue(),  goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *Security) OnWrite(instanceId int, resourceId int) (goap.CoapCode) {
    return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func NewExampleSecurityObject(reg Registry) (*Security) {
    return &Security{
        Model: reg.GetModel(oma.OBJECT_LWM2M_SECURITY),
    }
}


/*
[0]
LWM2M Server URI        0   coap://bootstrap.example.com
Bootstrap Server        1   true
Security Mode           2   0
Public Key or Identity  3   [identity string]
Secret Key              4   [secret key data]
Short Server ID         10  0
Client Hold Off Time    11  3600

[1]
LWM2M Server URI        0   coap://server1.example.com
Bootstrap Server        1   false
Security Mode           2   0
Public Key or Identity  3   [identity string]
Secret Key              4   [secret key data]
Short Server ID         10  101
Client Hold Off Time    11  0

[2]
LWM2M Server URI        0   coap://server2.example.com
Bootstrap Server        1   false
Security Mode           2   0
Public Key or Identity  3   [identity string]
Secret Key              4   [secret key data]
Short Server ID         10  102
Client Hold Off Time    11  0
*/