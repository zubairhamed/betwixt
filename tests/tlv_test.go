package tests

import (
    "testing"
    "github.com/zubairhamed/lwm2m/core"
    "github.com/zubairhamed/lwm2m"
    "github.com/zubairhamed/lwm2m/objects"
    "github.com/zubairhamed/goap"
    "github.com/zubairhamed/lwm2m/examples/obj/basic"
    "github.com/zubairhamed/lwm2m/objects/oma"
)

func TestResourceInstanceToTlv(t *testing.T) {
    client := createTestingClient()

    object := client.GetEnabledObjects()[oma.OBJECT_LWM2M_DEVICE]
    instance := object.GetObjectInstance(0)
    resource := instance.GetResource(1)

    _, err := core.TlvPayloadFromResourceInstances(resource)
    if err != nil {
        t.Error("Error thrown attempting to convert Resource Instance to TLV")
    }
}

func TestTlvToResourceInstance(t *testing.T) {

}

func TestObjectInstanceToTlv(t *testing.T) {
    client := createTestingClient()

    object := client.GetEnabledObjects()[oma.OBJECT_LWM2M_DEVICE]
    instance := object.GetObjectInstance(0)

    _, err := core.TlvPayloadFromObjectInstance(instance)
    if err != nil {
        t.Error("Error thrown attempting to convert Object Instance to TLV")
    }
}

func TestTlvToObjectInstance(t *testing.T) {

}

func TestObjectToTlv(t *testing.T) {
    client := createTestingClient()

    object := client.GetEnabledObjects()[oma.OBJECT_LWM2M_DEVICE]

    _, err := core.TlvPayloadFromObjects(object)

    if err != nil {
        t.Error("Error thrown attempting to convert Object to TLV")
    }


    // core.TlvPayloadFromObjects(c.enabledObjects[t])
}

func TestTlvToObject(t *testing.T) {

}

func createTestingClient() (*lwm2m.LWM2MClient) {
    client := lwm2m.NewLWM2MClient(":0", "localhost:5683")

    reg := objects.NewDefaultObjectRegistry()
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