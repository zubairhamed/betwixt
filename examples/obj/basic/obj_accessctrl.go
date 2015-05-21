package basic

import (
    "github.com/zubairhamed/lwm2m/core"
    . "github.com/zubairhamed/lwm2m/api"
)

type AccessControl struct {
    Model       ObjectModel
    State       map[int] map[int] interface{}
}

func (o *AccessControl) OnDelete(instanceId int) (bool) {
    return true
}

func (o *AccessControl) OnRead(instanceId int, resourceId int) (ResponseValue) {
    var val ResponseValue

    resource := o.Model.GetResource(resourceId)
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
    return core.NewEmptyValue()
}

func (o *AccessControl) OnWrite(instanceId int, resourceId int) (bool) {
    return true
}

/*
func NewExampleAccessControlObject(reg Registry) (*AccessControl) {
    state := map[int] map[int] interface{} {
        0: {
            0: {
                0: 1,
                1: 0,
                2: "0b0000000000001111",
                3: 101,
            }
        },
        1: {

        },
        2: {

        },
        3: {

        },
        4: {

        },
    }

    return &AccessControl{
        Model: reg.GetModel(oma.OBJECT_LWM2M_SECURITY),
        State: state,
    }
}
*/



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