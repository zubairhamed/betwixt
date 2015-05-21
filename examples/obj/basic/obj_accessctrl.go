package basic

import (
    "github.com/zubairhamed/lwm2m/core"
    . "github.com/zubairhamed/lwm2m/api"
    "github.com/zubairhamed/lwm2m/objects/oma"
)

type AccessControl struct {
    Model       ObjectModel
    Data        *ObjectsData
}

func (o *AccessControl) OnDelete(instanceId int) (bool, int) {
    return true, 0
}

func (o *AccessControl) OnRead(instanceId int, resourceId int) (ResponseValue, int) {
    /*

    resource := o.Model.GetResource(resourceId)

    val := o.Data.Get(instanceId, resourceId)


    switch resourceId {
        case 0:
        val = core.NewIntegerValue(instanceId)
        break

        case 1:
        val = core.NewIntegerValue(instanceId)
        break

        case 2:
        break

        case 3:
        break

    }
    */
    return core.NewEmptyValue(), 0
}

func (o *AccessControl) OnWrite(instanceId int, resourceId int) (bool, int) {
    return true, 0
}

//////////////////////////////////////////////////////
type ObjectsData struct {
    data    map[string] interface{}
}

func (o *ObjectsData) Put(path string, value interface{}) {
    o.data[path] = value
}

func (o *ObjectsData) Get(path ...int) (interface{}) {
    return o.data[path]
}

//////////////////////////////////////////////////////


func NewExampleAccessControlObject(reg Registry) (*AccessControl) {
    data := &ObjectsData{}

    // Object 0
    data.Put("/0/0", 1)
    data.Put("/0/1", 0)
    data.Put("/0/2/101", []byte{0, 15})
    data.Put("/0/3", 101)

    data.Put("1/0", 1)
    data.Put("1/1", 1)
    data.Put("1/2/102", []byte{0, 15})
    data.Put("1/3", 102)

    data.Put("2/0", 3)
    data.Put("2/1", 0)
    data.Put("2/2/101",  []byte{0, 15})
    data.Put("2/2/102",  []byte{0, 1})
    data.Put("2/3",  101)

    data.Put("3/0", 4)
    data.Put("3/1", 0)
    data.Put("3/2/101", []byte{0, 1})
    data.Put("3/2/0", []byte{0, 1})
    data.Put("3/3", 101)

    data.Put("4/0", 5)
    data.Put("4/1", 65535)
    data.Put("4/2/101", []byte{0, 16})
    data.Put("4/3", 65535)

    return &AccessControl{
        Model: reg.GetModel(oma.OBJECT_LWM2M_SECURITY),
        Data: data,
    }
}



/*
[0] - LWM2M Server Object
Object ID               0           1
Object Instance ID      1           0
ACL                     2   101     0b0000000000001111
Access Control Owner    3           101

[1] -LWM2M Server Object
Object ID               0           1
Object Instance ID      1           1
ACL                     2   102     0b0000000000001111
Access Control Owner    3           102

[2] - Device Object
Object ID               0           3
Object Instance ID      1           0
ACL                     2   101     0b0000000000001111
ACL                     2   102     0b0000000000000001
Access Control Owner    3           101

[3] - Connectivity Monitoring Object
Object ID               0           4
Object Instance ID      1           0
ACL                     2   101     0b0000000000000001
ACL                     2   0       0b0000000000000001
Access Control Owner    3           101

[4] - Firmware Update Object
Object ID               0           5
Object Instance ID      1           65535
ACL                     2   101     0b0000000000010000
Access Control Owner    3           65535

*/