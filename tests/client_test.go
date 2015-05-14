package tests

import (
    "testing"
    "github.com/zubairhamed/lwm2m/objects/oma"
    "github.com/zubairhamed/lwm2m/objects"
    "github.com/zubairhamed/lwm2m/core"
    "github.com/zubairhamed/lwm2m"
)

func TestClient(t *testing.T) {
    client := lwm2m.NewLWM2MClient(":0", "localhost:5683")
    if client == nil {
        t.Error("Error instantiating client")
    }

    cases1 := []struct {
        in core.LWM2MObjectType
    }{
        {oma.OBJECT_LWM2M_SERVER},
        {oma.OBJECT_LWM2M_ACCESS_CONTROL},
        {oma.OBJECT_LWM2M_DEVICE},
        {oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING},
        {oma.OBJECT_LWM2M_FIRMWARE_UPDATE},
        {oma.OBJECT_LWM2M_LOCATION},
        {oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS},
    }

    for _, c := range cases1 {
        if client.EnableObject(c.in, nil) != nil {
            t.Error("Error enabling object: ", c)
        }
    }

    if client.EnableObject(oma.OBJECT_LWM2M_SECURITY, nil) != nil {
        t.Error("Error enabling object")
    }

    if client.EnableObject(oma.OBJECT_LWM2M_SECURITY, nil) == nil {
        t.Error("Object should already be enabled")
    }

    cases2 := []struct {
        in core.LWM2MObjectType
    }{
        {oma.OBJECT_LWM2M_SERVER},
        {oma.OBJECT_LWM2M_ACCESS_CONTROL},
        {oma.OBJECT_LWM2M_DEVICE},
        {oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING},
        {oma.OBJECT_LWM2M_FIRMWARE_UPDATE},
        {oma.OBJECT_LWM2M_LOCATION},
        {oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS},
    }

    for _, c := range cases2 {
        if client.GetObjectEnabler(c.in) == nil {
            t.Error("Error getting object enabler: ", c)
        }
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

    cases3 := []struct {
        ot   core.LWM2MObjectType
        oi   int
    }{
        {oma.OBJECT_LWM2M_SECURITY, 0},
        {oma.OBJECT_LWM2M_SECURITY, 1},
        {oma.OBJECT_LWM2M_SECURITY, 2},
    }

    for _, c := range cases3 {
        if client.GetObjectInstance(c.ot, c.oi) == nil {
            t.Error("Object instance", c.oi, "not found")
        }
    }
}

func TestRegistry(t *testing.T) {
    reg := objects.NewDefaultObjectRegistry()

    cases := []struct {
        o   core.LWM2MObjectType
    }{
        {oma.OBJECT_LWM2M_SECURITY},
        {oma.OBJECT_LWM2M_SERVER},
        {oma.OBJECT_LWM2M_ACCESS_CONTROL},
        {oma.OBJECT_LWM2M_DEVICE},
        {oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING},
        {oma.OBJECT_LWM2M_FIRMWARE_UPDATE},
        {oma.OBJECT_LWM2M_LOCATION},
        {oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS},
    }

    for _, c := range cases {
        if reg.CreateObjectInstance(c.o, 0) == nil {
            t.Error("Created an LWM2M Object: ", c.o)
        }
    }

    if reg.CreateObjectInstance(core.LWM2MObjectType(-1), 0) != nil {
        t.Error("Created an unknown LWM2M Object")
    }
}

func TestBuildResourceStringPayload(t *testing.T) {
    client := lwm2m.NewLWM2MClient(":0", "localhost:5683")

    client.EnableObject(oma.OBJECT_LWM2M_SECURITY, nil)
    client.EnableObject(oma.OBJECT_LWM2M_ACCESS_CONTROL, nil)
    client.EnableObject(oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING, nil)

    str := lwm2m.BuildModelResourceStringPayload(client.GetEnabledObjects())
    compare := "</0>,</2>,</4>,"
    if str != compare {
        t.Error("Unexpected output building Model Resource String: Expected = ", compare, "Actual = ", str)
    }
}
