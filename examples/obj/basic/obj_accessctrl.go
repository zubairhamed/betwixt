package basic

import (
    "github.com/zubairhamed/lwm2m/core"
    . "github.com/zubairhamed/lwm2m/api"
)

type AccessControl struct {
    Model       ObjectModel
}

func (o *AccessControl) OnDelete(instanceId int) (bool) {
    return true
}

func (o *AccessControl) OnRead(instanceId int, resourceId int) (ResponseValue) {
    return core.NewEmptyValue()
}

func (o *AccessControl) OnWrite(instanceId int, resourceId int) (bool) {
    return true
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