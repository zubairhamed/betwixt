package basic

import (
	. "github.com/zubairhamed/go-lwm2m/api"
	"github.com/zubairhamed/go-lwm2m/core"
	"github.com/zubairhamed/go-lwm2m/objects/oma"
	"github.com/zubairhamed/goap"
	"time"
)

type Device struct {
	Model ObjectModel
	Data  *core.ObjectsData
}

func (o *Device) OnExecute(instanceId int, resourceId int) goap.CoapCode {
	return goap.COAPCODE_204_CHANGED
}

func (o *Device) OnCreate(instanceId int, resourceId int) goap.CoapCode {
	return goap.COAPCODE_201_CREATED
}

func (o *Device) OnDelete(instanceId int) goap.CoapCode {
	return goap.COAPCODE_202_DELETED
}

func (o *Device) OnRead(instanceId int, resourceId int) (ResponseValue, goap.CoapCode) {
	if resourceId == -1 {
		// Read Object Instance
	} else {
		// Read Resource Instance
		var val ResponseValue

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
		return val, goap.COAPCODE_205_CONTENT
	}
	return core.NewEmptyValue(), goap.COAPCODE_404_NOT_FOUND
}

func (o *Device) OnWrite(instanceId int, resourceId int) goap.CoapCode {
	return goap.COAPCODE_404_NOT_FOUND
}

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

func (o *Device) Reboot() ResponseValue {
	return core.NewEmptyValue()
}

func (o *Device) FactoryReset() ResponseValue {
	return core.NewEmptyValue()
}

func (o *Device) GetAvailablePowerSources() []int {
	return []int{1, 5}
}

func (o *Device) GetPowerSourceVoltage() []int {
	return []int{3800, 5000}
}

func (o *Device) GetPowerSourceCurrent() []int {
	return []int{125, 900}
}

func (o *Device) GetBatteryLevel() int {
	return 100
}

func (o *Device) GetMemoryFree() int {
	return 15
}

func (o *Device) GetErrorCode() []int {
	return []int{0}
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

func NewExampleDeviceObject(reg Registry) *Device {
	data := &core.ObjectsData{
		Data: make(map[string]interface{}),
	}

	data.Put("/0/0", "Open Mobile Alliance")
	data.Put("/0/1", "Lightweight M2M Client")
	data.Put("/0/2", 345000123)
	data.Put("/0/3", "1.0")
	data.Put("/0/6/0", 1)
	data.Put("/0/6/1", 5)
	data.Put("/0/7/0", 3800)
	data.Put("/0/7/1", 5000)
	data.Put("/0/8/0", 125)
	data.Put("/0/8/1", 900)
	data.Put("/0/9", 100)
	data.Put("/0/10", 15)
	data.Put("/0/11/0", 0)
	data.Put("/0/13", 1367491215)
	data.Put("/0/14", "+02:00")
	data.Put("/0/15", "U")

	return &Device{
		Model: reg.GetModel(oma.OBJECT_LWM2M_DEVICE),
		Data:  data,
	}
}

/*

*/
