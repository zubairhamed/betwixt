package objects

import (
    "github.com/zubairhamed/lwm2m/core"
    "time"
)

type TestDevice struct {
    Model       *core.ObjectModel
}

func (o *TestDevice) OnRead(instanceId int, resourceId int) (core.ResourceValue) {
    if resourceId == -1 {
        // Read Object Instance
    } else {
        // Read Resource Instance
        var val core.ResourceValue

        resource := o.Model.GetResource(resourceId)
        switch resourceId {
            case 0:
            val = core.NewStringValue(o.GetManufacturer())
            break

            case 1:
            val = core.NewStringValue(o.GetModelNumber())
            break

            case 2:
            val = core.NewStringValue(o.GetSerialNumber())
            break

            case 3:
            val = core.NewStringValue(o.GetFirmwareVersion())
            break

            case 6:
            val, _ = core.TlvPayloadFromIntResource(resource, o.GetAvailablePowerSources())
            break

            case 7:
            val, _ = core.TlvPayloadFromIntResource(resource, o.GetPowerSourceVoltage())
            break

            case 8:
            val, _ = core.TlvPayloadFromIntResource(resource, o.GetPowerSourceCurrent())
            break

            case 9:
            val = core.NewIntegerValue(o.GetBatteryLevel())
            break

            case 10:
            val = core.NewIntegerValue(o.GetMemoryFree())
            break

            case 11:
            val, _ = core.TlvPayloadFromIntResource(resource, o.GetErrorCode())
            break

            case 13:
            val = core.NewTimeValue(o.GetCurrentTime())
            break

            case 15:
            val = core.NewStringValue(o.GetTimezone())
            break

            case 16:
            val = core.NewStringValue(o.GetSupportedBindingMode())
            break

            default:
            break
        }
        return val
    }
    return core.NewEmptyValue()
}

func (o TestDevice) GetManufacturer() string {
    return "Open Mobile Alliance"
}

func (o TestDevice) GetModelNumber() string {
    return "Lightweight M2M Client"
}

func (o TestDevice) GetSerialNumber() string {
    return "345000123"
}

func (o TestDevice) GetFirmwareVersion() string {
    return "1.0"
}

func (o TestDevice) GetAvailablePowerSources() []int {
    return []int{ 1, 5 }
}

func (o TestDevice) GetPowerSourceVoltage() []int {
    return  []int{ 3800, 5000 }
}

func (o TestDevice) GetPowerSourceCurrent() []int {
    return []int{ 125, 900 }
}

func (o TestDevice) GetBatteryLevel() int {
    return 100
}

func (o TestDevice) GetMemoryFree() int {
    return 15
}

func (o TestDevice) GetErrorCode() []int {
    return []int{ 0 }
}

func (o TestDevice) GetCurrentTime() time.Time {
    return time.Now()
}

func (o TestDevice) GetTimezone() string {
    return "+2:00"
}

func (o TestDevice) GetSupportedBindingMode() string {
    return "U"
}
