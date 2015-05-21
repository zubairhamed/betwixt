package tests

import (
    "testing"
    "github.com/zubairhamed/lwm2m/core"
    "github.com/zubairhamed/lwm2m"
    "github.com/zubairhamed/lwm2m/examples/obj/basic"
    "github.com/zubairhamed/lwm2m/objects/oma"
    "github.com/zubairhamed/lwm2m/registry"
    "github.com/stretchr/testify/assert"
)


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

    assert.Nil(t, err, "Error thrown attempting to convert Object instance to TLV")
}

