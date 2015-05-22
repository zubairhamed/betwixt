package main

import (
    . "github.com/zubairhamed/lwm2m"
    "github.com/zubairhamed/lwm2m/objects/oma"
    . "github.com/zubairhamed/lwm2m/examples/obj/basic"
    "github.com/zubairhamed/lwm2m/registry"
    . "github.com/zubairhamed/lwm2m/api"
)

func main() {
    client := NewLWM2MClient(":0", "localhost:5683")

    registry := registry.NewDefaultObjectRegistry()
    client.UseRegistry(registry)

    setupResources(client, registry)

    client.OnStartup(func(){
        client.Register("GOClient")
    })

    client.Start()
}

func setupResources (client LWM2MClient, reg Registry) {
    accessControl := &AccessControl{
        Model: reg.GetModel(oma.OBJECT_LWM2M_ACCESS_CONTROL),
    }

    device := NewExampleDeviceObject(reg)

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
}