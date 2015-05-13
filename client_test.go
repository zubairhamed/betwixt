package lwm2m
import (
    "testing"
    "github.com/zubairhamed/lwm2m/objects/oma"
    "github.com/zubairhamed/lwm2m/objects"
    "github.com/zubairhamed/lwm2m/core"
)

func TestClient(t *testing.T) {
    client := NewLWM2MClient(":0", "localhost:5683")
    if client == nil {
        t.Error("Error instantiating client")
    }

    if client.EnableObject(oma.OBJECT_LWM2M_SECURITY, nil) != nil {
        t.Error("Error enabling object")
    }

    if client.EnableObject(oma.OBJECT_LWM2M_SECURITY, nil) == nil {
        t.Error("Object should already be enabled")
    }

    if client.EnableObject(oma.OBJECT_LWM2M_SERVER, nil) != nil {
        t.Error("Error enabling object")
    }

    if client.EnableObject(oma.OBJECT_LWM2M_ACCESS_CONTROL, nil) != nil {
        t.Error("Error enabling object")
    }

    if client.EnableObject(oma.OBJECT_LWM2M_DEVICE, nil) != nil {
        t.Error("Error enabling object")
    }

    if client.EnableObject(oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING, nil) != nil {
        t.Error("Error enabling object")
    }

    if client.EnableObject(oma.OBJECT_LWM2M_FIRMWARE_UPDATE, nil) != nil {
        t.Error("Error enabling object")
    }

    if client.EnableObject(oma.OBJECT_LWM2M_LOCATION, nil) != nil {
        t.Error("Error enabling object")
    }

    if client.EnableObject(oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS, nil) != nil {
        t.Error("Error enabling object")
    }

    registry := objects.NewDefaultObjectRegistry()
    if registry == nil {
        t.Error("Error instantiating registry")
    }

    client.UseRegistry(registry)

    inst1 := registry.CreateObjectInstance(oma.OBJECT_LWM2M_SECURITY, 0)
    inst2 := registry.CreateObjectInstance(oma.OBJECT_LWM2M_SECURITY, 1)
    inst3 := registry.CreateObjectInstance(oma.OBJECT_LWM2M_SECURITY, 2)

    if inst1 == nil || inst2 == nil || inst3 == nil {
        t.Error("Error instantiating lwm2m object")
    }

    client.AddObjectInstances(inst1, inst2, inst3)

    if client.GetObjectInstance(oma.OBJECT_LWM2M_SECURITY, 0) == nil {
        t.Error("Object instance 1 not found")
    }

    if client.GetObjectInstance(oma.OBJECT_LWM2M_SECURITY, 1) == nil {
        t.Error("Object instance 2 not found")
    }

    if client.GetObjectInstance(oma.OBJECT_LWM2M_SECURITY, 2) == nil {
        t.Error("Object instance 3 not found")
    }

    if client.AddObjectInstance(inst1) == nil {
        t.Error("Error should be thrown for adding duplicate instance")
    }
}

func TestRegistry(t *testing.T) {
    reg := objects.NewDefaultObjectRegistry()

    if reg.CreateObjectInstance(core.LWM2MObjectType(-1), 0) != nil {
        t.Error("Created an unknown LWM2M Object")
    }

    if reg.CreateObjectInstance(oma.OBJECT_LWM2M_SECURITY, 0) == nil {
        t.Error("Error creating lwm2m object")
    }

    if reg.CreateObjectInstance(oma.OBJECT_LWM2M_SERVER, 0) == nil {
        t.Error("Error creating lwm2m object")
    }

    if reg.CreateObjectInstance(oma.OBJECT_LWM2M_ACCESS_CONTROL, 0) == nil {
        t.Error("Error creating lwm2m object")
    }

    if reg.CreateObjectInstance(oma.OBJECT_LWM2M_DEVICE, 0) == nil {
        t.Error("Error creating lwm2m object")
    }

    if reg.CreateObjectInstance(oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING, 0) == nil {
        t.Error("Error creating lwm2m object")
    }

    if reg.CreateObjectInstance(oma.OBJECT_LWM2M_FIRMWARE_UPDATE, 0) == nil {
        t.Error("Error creating lwm2m object")
    }

    if reg.CreateObjectInstance(oma.OBJECT_LWM2M_LOCATION, 0) == nil {
        t.Error("Error creating lwm2m object")
    }

    if reg.CreateObjectInstance(oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS, 0) == nil {
        t.Error("Error creating lwm2m object")
    }
}