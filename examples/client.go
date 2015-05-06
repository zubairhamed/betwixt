package main

import (
    . "github.com/zubairhamed/lwm2m"
    . "github.com/zubairhamed/lwm2m/objects"
    "github.com/zubairhamed/lwm2m/objects/oma"
    "github.com/zubairhamed/lwm2m/core"
    "log"
    "github.com/zubairhamed/goap"
    "time"
)

func main() {
    client := NewLWM2MClient(":0", "localhost:5683")

    registry := NewDefaultObjectRegistry()
    client.UseRegistry(registry)

    serial := setupResources(client, registry)

    client.OnStartup(func(){
        client.Register("GO-" + serial)
    })

    client.Start()
}

func setupResources (client *LWM2MClient, reg *ObjectRegistry) (string) {
    accessControl := &AccessControl{}
    device := &Device{
        serial: goap.GenerateToken(5),
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

    return device.GetSerialNumber()
}

type AccessControl struct {

}

func (o *AccessControl) OnRead(t core.LWM2MObjectType, m *core.ObjectModel, i *core.ObjectInstance, r *core.ResourceModel) core.ResponseValue {
    return nil
}

// -----

type Device struct {
    serial  string
}

/*
exec
case 4:
case 5:
case 12:

*/
func (o *Device) OnRead(t core.LWM2MObjectType, m *core.ObjectModel, i *core.ObjectInstance, r *core.ResourceModel) core.ResponseValue{

    var val core.ResponseValue
    switch r.Id {
        case 0:
        val = core.NewStringResponseValue(o.GetManufacturer())
        break

        case 1:
        val = core.NewStringResponseValue(o.GetModelNumber())
        break

        case 2:
        val = core.NewStringResponseValue(o.GetSerialNumber())
        break

        case 3:
        val = core.NewStringResponseValue(o.GetFirmwareVersion())
        break

        case 6:
        val = core.NewIntResponseValue(o.GetAvailablePowerSources())
        break

        case 7:
        val = core.NewIntResponseValue(o.GetPowerSourceVoltage())
        break

        case 8:
        val = core.NewIntResponseValue(o.GetPowerSourceCurrent())
        break

        case 9:
        val = core.NewIntResponseValue(o.GetBatteryLevel())
        break

        case 10:
        val = core.NewIntResponseValue(o.GetMemoryFree())
        break

        case 11:
        val = core.NewIntResponseValue(o.GetErrorCode())
        break

        case 13:
        val = core.NewTimeResponseValue(o.GetCurrentTime())
        break

        case 14:
        val = core.NewStringResponseValue(o.GetUtcOffset())
        break

        case 15:
        val = core.NewStringResponseValue(o.GetTimezone())
        break

        case 16:
        val = core.NewStringResponseValue(o.GetSupportedBindingMode())
        break

        default:
        break
    }
    return val
}

func (o *Device) GetManufacturer() string {
    return "GOLWM2M"
}

func (o *Device) GetModelNumber() string {
    return "0.1"
}

func (o *Device) GetSerialNumber() string {
    return o.serial
}

func (o *Device) GetFirmwareVersion() string {
    return "1.0"
}

func (o *Device) Reboot() *core.NoResponseValue {
    return core.NoResponse()
}

func (o *Device) FactoryReset() *core.NoResponseValue {
    return core.NoResponse()
}

func (o *Device) GetAvailablePowerSources() int {
    return 0
}

func (o *Device) GetPowerSourceVoltage() int {
    return 0
}

func (o *Device) GetPowerSourceCurrent() int {
    return 0
}

func (o *Device) GetBatteryLevel() int {
    return 0
}

func (o *Device) GetMemoryFree() int {
    return 0
}

func (o *Device) GetErrorCode() int {
    return 0
}

func (o *Device) ResetErrorCode() string {
    return ""
}

func (o *Device) GetCurrentTime() time.Time {
    return time.Now()
}

func (o *Device) GetUtcOffset() string {
    return "+8:00"
}

func (o *Device) GetTimezone() string {
    return ""
}

func (o *Device) GetSupportedBindingMode() string {
    return ""
}
