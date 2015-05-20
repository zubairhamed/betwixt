package basic

import (
    "github.com/zubairhamed/lwm2m/core"
    "time"
    "log"
)

type Device struct {
    Serial      string
    Model       *core.ObjectModel
}

func (o *Device) OnRead(instanceId int, resourceId int) (core.ResourceValue) {
    log.Println("OnRead")
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

            case 14:
            val = core.NewStringValue(o.GetUtcOffset())
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


/*
func (o *Device) OnRead(r *core.ResourceModel, resourceId int) core.ResourceValue {

    var val core.ResourceValue

    switch r.Id {
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
        val = core. NewIntegerValue(o.GetAvailablePowerSources()...)
        break

        case 7:
        val = core.NewIntegerValue(o.GetPowerSourceVoltage())
        break

        case 8:
        val = core.NewIntegerValue(o.GetPowerSourceCurrent())
        break

        case 9:
        val = core.NewIntegerValue(o.GetBatteryLevel())
        break

        case 10:
        val = core.NewIntegerValue(o.GetMemoryFree())
        break

        case 11:
        val = core.NewIntegerValue(o.GetErrorCode())
        break

        case 13:
        val = core.NewTimeValue(o.GetCurrentTime())
        break

        case 14:
        val = core.NewStringValue(o.GetUtcOffset())
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
*/

func (o *Device) GetManufacturer() string {
    return "Open Mobile Alliance"
}

func (o *Device) GetModelNumber() string {
    return "Lightweight M2M Client"
}

func (o *Device) GetSerialNumber() string {
    return "345000123"
}

func (o *Device) GetFirmwareVersion() string {
    return "1.0"
}

func (o *Device) Reboot() core.ResourceValue {
    return core.NewEmptyValue()
}

func (o *Device) FactoryReset() core.ResourceValue {
    return core.NewEmptyValue()
}

func (o *Device) GetAvailablePowerSources() []int {
    return []int{ 1, 5 }
}

func (o *Device) GetPowerSourceVoltage() []int {
    return  []int{ 3800, 5000 }
}

func (o *Device) GetPowerSourceCurrent() []int {
    return []int{ 125, 900 }
}

func (o *Device) GetBatteryLevel() int {
    return 100
}

func (o *Device) GetMemoryFree() int {
    return 15
}

func (o *Device) GetErrorCode() []int {
    return []int{ 0 }
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
    return "+2:00"
}

func (o *Device) GetSupportedBindingMode() string {
    return "U"
}
