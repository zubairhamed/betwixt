package main

import (
    . "github.com/zubairhamed/lwm2m"
    . "github.com/zubairhamed/lwm2m/objects"
    "github.com/zubairhamed/lwm2m/objects/oma"
    "github.com/zubairhamed/lwm2m/core"
    "log"
)

func main() {
    client := NewLWM2MClient(":0", "localhost:5683")

    registry := NewDefaultObjectRegistry()
    client.UseRegistry(registry)

    setupResources(client, registry)

    client.OnStartup(func(){
        client.Register("GOCLIENT")
    })

    // client.OnRead(func(evt *Event, m *ObjectInstance, i *ResourceInstance) (*LWM2MResponse) {
    client.OnRead(func() {
        // log.Println(evt.Data["objectModel"].(*ObjectModel))
    })

    /*
    client.OnRegistered(func(evt *Event, path string){
        log.Println("Client is registered")
    })

    client.OnUnregistered(func(evt *Event){
        log.Println("Client is Unregistered")
    })

    client.OnExecute(func(evt *Event, m *ObjectInstance, i *ResourceInstance) (*LWM2MResponse) {

    })

    client.OnWrite(func(evt *Event, m *ObjectInstance, i *ResourceInstance, value interface{}) (*LWM2MResponse) {

    })

    client.OnCreate(func (evt *Event, m *ObjectInstance, i *ResourceInstance) (*LWM2MResponse) {

    })
    */

    client.Start()
}

func setupResources (client *LWM2MClient, reg *ObjectRegistry) {


    accessControl := &AccessControl{}
    device := &Device{}

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

type AccessControl struct {

}

func (o *AccessControl) OnRead(t core.LWM2MObjectType, m *core.ObjectModel, i *core.ObjectInstance, r *core.ResourceModel) {

}

// -----

type Device struct {

}

func (o *Device) OnRead(t core.LWM2MObjectType, m *core.ObjectModel, i *core.ObjectInstance, r *core.ResourceModel) {
    log.Println("OnRead Invoked", t, m, i, r)
}