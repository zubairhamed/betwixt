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

func (o *AccessControl) OnRead(t core.LWM2MObjectType, m *core.ObjectModel, i *core.ObjectInstance, r *core.ResourceModel) core.ResponseValue {
    return nil
}

// -----

type Device struct {

}

func (o *Device) OnRead(t core.LWM2MObjectType, m *core.ObjectModel, i *core.ObjectInstance, r *core.ResourceModel) core.ResponseValue{
    log.Println("OnRead Invoked", t, m, i, r)
    log.Println("Resource Model ID ", r.Id)

    var val core.ResponseValue
    switch r.Id {
        case 0:
            val = core.NewStringResponseValue(o.GetManufacturer())
            break

        case 1:


        default:
            break
    }
    return val
}

func (o *Device) GetManufacturer() string {
    return "GOLWM2M"
}


/*
        GetModelNumber() string
        &ResourceModel{ Id: 1, Name: "Model Number", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: TYPE_STRING, RangeOrEnums: "", Units: "", Description: "" },

        GetSerialNumber() string
        &ResourceModel{ Id: 2, Name: "Serial Number", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: TYPE_STRING, RangeOrEnums: "", Units: "", Description: "" },

        GetFirmwareVersion() string
        &ResourceModel{ Id: 3, Name: "Firmware Version", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: TYPE_STRING, RangeOrEnums: "", Units: "", Description: "" },

        Reboot() string
        &ResourceModel{ Id: 4, Name: "Reboot", Operations: OPERATION_E, Multiple: false, Mandatory: true, ResourceType: TYPE_STRING, RangeOrEnums: "", Units: "", Description: "" },

        FactoryReset() string
        &ResourceModel{ Id: 5, Name: "Factory Reset", Operations: OPERATION_E, Multiple: false, Mandatory: false, ResourceType: TYPE_STRING, RangeOrEnums: "", Units: "", Description: "" },

        GetAvailablePowerSources() int
        &ResourceModel{ Id: 6, Name: "Available Power Sources", Operations: OPERATION_R, Multiple: true, Mandatory: false, ResourceType: TYPE_INTEGER, RangeOrEnums: "0-7", Units: "", Description: "" },

        GetPowerSourceVoltage() int
        &ResourceModel{ Id: 7, Name: "Power Source Voltage", Operations: OPERATION_R, Multiple: true, Mandatory: false, ResourceType: TYPE_INTEGER, RangeOrEnums: "", Units: "mV", Description: "" },

        GetPowerSourceCurrent() int
        &ResourceModel{ Id: 8, Name: "Power Source Current", Operations: OPERATION_R, Multiple: true, Mandatory: false, ResourceType: TYPE_INTEGER, RangeOrEnums: "", Units: "mA", Description: "" },

        GetBatteryLevel() int
        &ResourceModel{ Id: 9, Name: "Battery Level", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: TYPE_INTEGER, RangeOrEnums: "0-100", Units: "%", Description: "" },

        GetMemoryFree() int
        &ResourceModel{ Id: 10, Name: "Memory Free", Operations: OPERATION_R, Multiple: false, Mandatory: false, ResourceType: TYPE_INTEGER, RangeOrEnums: "", Units: "KB", Description: "" },

        GetErrorCode() int
        &ResourceModel{ Id: 11, Name: "Error Code", Operations: OPERATION_R, Multiple: true, Mandatory: true, ResourceType: TYPE_INTEGER, RangeOrEnums: "", Units: "", Description: "" },

        ResetErrorCode() string
        &ResourceModel{ Id: 12, Name: "Reset Error Code", Operations: OPERATION_E, Multiple: false, Mandatory: false, ResourceType: TYPE_STRING, RangeOrEnums: "", Units: "", Description: "" },

        GetCurrentTime() time.Time
        &ResourceModel{ Id: 13, Name: "Current Time", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: TYPE_TIME, RangeOrEnums: "", Units: "", Description: "" },

        GetUtcOffset() string
        &ResourceModel{ Id: 14, Name: "UTC Offset", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: TYPE_STRING, RangeOrEnums: "", Units: "", Description: "" },

        GetTimexone() string
        &ResourceModel{ Id: 15, Name: "Timezone", Operations: OPERATION_RW, Multiple: false, Mandatory: false, ResourceType: TYPE_STRING, RangeOrEnums: "", Units: "", Description: "" },

        GetSupportedBindingMode() string
        &ResourceModel{ Id: 16, Name: "Supported Binding and Modes", Operations: OPERATION_R, Multiple: false, Mandatory: true, ResourceType: TYPE_STRING, RangeOrEnums: "", Units: "", Description: "" },

*.