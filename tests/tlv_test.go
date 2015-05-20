package tests

import (
    "testing"
    "github.com/zubairhamed/lwm2m/core"
    "github.com/zubairhamed/lwm2m"
    "github.com/zubairhamed/goap"
    "github.com/zubairhamed/lwm2m/examples/obj/basic"
    "github.com/zubairhamed/lwm2m/objects/oma"
    "github.com/zubairhamed/lwm2m/registry"
)

/*
func TestResourceInstanceToTlv(t *testing.T) {
    client := createTestingClient()
    model := client.GetRegistry().GetModel(oma.OBJECT_LWM2M_DEVICE)
    resourceModel := model.GetResource(6)

    // client.GetObjectEnabler(oma.OBJECT_LWM2M_DEVICE)
    // en.Handler.OnRead(0, 6)

    _, err := core.TlvPayloadFromIntResource(resourceModel, []int{1,5,})
    if err != nil {
        t.Error("Error thrown attempting to convert Resource Instance to TLV")
    }
}
*/

func TestObjectInstancesToTlv(t *testing.T) {
    client := lwm2m.NewLWM2MClient(":0", "localhost:5683")

    reg := registry.NewDefaultObjectRegistry()
    client.UseRegistry(reg)

    device := &basic.Device{
        Model: reg.GetModel(oma.OBJECT_LWM2M_DEVICE),
    }

    client.EnableObject(oma.OBJECT_LWM2M_DEVICE, device)
    instanceDevice := reg.CreateObjectInstance(oma.OBJECT_LWM2M_DEVICE, 0)
    client.AddObjectInstances(instanceDevice)

    en := client.GetObjectEnabler(oma.OBJECT_LWM2M_DEVICE)
    _, err := core.TlvPayloadFromObjects(en, client.GetRegistry())
    if err != nil {
        t.Error("Error thrown attempting to convert Object instance to TLV")
    }
}

func createTestingClient() (*lwm2m.LWM2MClient) {
    client := lwm2m.NewLWM2MClient(":0", "localhost:5683")

    reg := registry.NewDefaultObjectRegistry()
    client.UseRegistry(reg)

    accessControl := &basic.AccessControl{}
    device := &basic.Device{
        Serial: goap.GenerateToken(5),
    }

    client.EnableObject(oma.OBJECT_LWM2M_SECURITY, nil)
    client.EnableObject(oma.OBJECT_LWM2M_SERVER, nil)
    client.EnableObject(oma.OBJECT_LWM2M_ACCESS_CONTROL, accessControl)
    client.EnableObject(oma.OBJECT_LWM2M_DEVICE, device)
    client.EnableObject(oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING, nil)
    client.EnableObject(oma.OBJECT_LWM2M_FIRMWARE_UPDATE, nil)
    client.EnableObject(oma.OBJECT_LWM2M_LOCATION, nil)
    client.EnableObject(oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS, nil)

    instanceSec1 := reg.CreateObjectInstance(oma.OBJECT_LWM2M_SECURITY, 0)
    instanceSec2 := reg.CreateObjectInstance(oma.OBJECT_LWM2M_SECURITY, 1)
    instanceSec3 := reg.CreateObjectInstance(oma.OBJECT_LWM2M_SECURITY, 2)

    instanceServer := reg.CreateObjectInstance(oma.OBJECT_LWM2M_SERVER, 1)

    instanceAccessCtrl1 := reg.CreateObjectInstance(oma.OBJECT_LWM2M_ACCESS_CONTROL, 0)
    instanceAccessCtrl2 := reg.CreateObjectInstance(oma.OBJECT_LWM2M_ACCESS_CONTROL, 1)
    instanceAccessCtrl3 := reg.CreateObjectInstance(oma.OBJECT_LWM2M_ACCESS_CONTROL, 2)
    instanceAccessCtrl4 := reg.CreateObjectInstance(oma.OBJECT_LWM2M_ACCESS_CONTROL, 3)
    instanceDevice := reg.CreateObjectInstance(oma.OBJECT_LWM2M_DEVICE, 0)
    instanceConnMonitoring := reg.CreateObjectInstance(oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING, 0)
    instanceFwUpdate :=  reg.CreateObjectInstance(oma.OBJECT_LWM2M_FIRMWARE_UPDATE, 0)

    client.AddObjectInstances(
        instanceSec1, instanceSec2, instanceSec3,
        instanceServer,
        instanceAccessCtrl1, instanceAccessCtrl2, instanceAccessCtrl3, instanceAccessCtrl4,
        instanceDevice,
        instanceConnMonitoring,
        instanceFwUpdate,
    )

    return client
}

/*
package main
import (
    // "encoding/hex"
    "log"
)

func main() {

    var hexString = "8606410001410105"
    rawMessage, _ := hex.DecodeString(hexString)

    for _, b := range rawMessage {
        log.Println (b)
    }


var binval byte = 85
// 1 2 4 8
// 00000110
// 00000111


//if type == OBJECT_INSTANCE
//if type == RESOURCE_WITH_VALUE
//if type == MULTIPLE_RESOURCE
//if type == RESOURCE_INSTANCE



log.Println ((binval & 0xC0) >> 6)
}


/*
    for bytes {
        read first byte
        if type == object instance

        if type == resource instance
        if type == multiple resources
        if type == resource with value
    }
*/
