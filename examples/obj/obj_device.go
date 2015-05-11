package obj

import (
    "github.com/zubairhamed/lwm2m/core"
    "time"
)

type Device struct {
    Serial      string
}

/*
exec
case 4:
case 5:
case 12:

*/

func (o *Device) OnRead(t core.LWM2MObjectType, m *core.ObjectModel, i *core.ObjectInstance, r *core.ResourceModel) core.ResourceValue {

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
        val = core. NewIntegerValue(o.GetAvailablePowerSources())
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

func (o *Device) GetManufacturer() string {
    return "GOLWM2M"
}

func (o *Device) GetModelNumber() string {
    return "0.1"
}

func (o *Device) GetSerialNumber() string {
    return o.Serial
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
