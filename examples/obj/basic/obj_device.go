package basic

import (
	. "github.com/zubairhamed/go-lwm2m/api"
	"github.com/zubairhamed/go-lwm2m/core"
	"github.com/zubairhamed/go-lwm2m/objects/oma"
	"time"
	"github.com/zubairhamed/go-lwm2m/core/response"
	"github.com/zubairhamed/go-lwm2m/core/values"
)

type Device struct {
	Model ObjectModel
	Data  *core.ObjectsData
}

func (o *Device) OnExecute(instanceId int, resourceId int, req Request) Response {
	return response.Changed()
}

func (o *Device) OnCreate(instanceId int, resourceId int, req Request) Response {
	return response.Created()
}

func (o *Device) OnDelete(instanceId int, req Request) Response {
	return response.Deleted()
}

func (o *Device) OnRead(instanceId int, resourceId int, req Request) Response {
	if resourceId == -1 {
		// Read Object Instance
	} else {
		// Read Resource Instance
		var val ResponseValue

		resource := o.Model.GetResource(resourceId)
		switch resourceId {
		case 0:
			val = values.String(o.GetManufacturer())
			break

		case 1:
			val = values.String(o.GetModelNumber())
			break

		case 2:
			val = values.String(o.GetSerialNumber())
			break

		case 3:
			val = values.String(o.GetFirmwareVersion())
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
			val = values.Integer(o.GetBatteryLevel())
			break

		case 10:
			val = values.Integer(o.GetMemoryFree())
			break

		case 11:
			val, _ = core.TlvPayloadFromIntResource(resource, o.GetErrorCode())
			break

		case 13:
			val = values.Time(o.GetCurrentTime())
			break

		case 14:
			val = values.String(o.GetUtcOffset())
			break

		case 15:
			val = values.String(o.GetTimezone())
			break

		case 16:
			val = values.String(o.GetSupportedBindingMode())
			break

		default:
			break
		}
		return response.Content(val)
	}
	return response.NotFound()
}

func (o *Device) OnWrite(instanceId int, resourceId int, req Request) Response {
	val := req.GetMessage().Payload

	switch resourceId {
	case 13:
		break

	case 14:
		o.Data.Put("/0/14", val.String())
		break

	case 15:
		o.Data.Put("/0/15", val.String())
		break

	default:
		return response.NotFound()
	}
	return response.Changed()
}

func (o *Device) GetManufacturer() string {
	return o.Data.Get("/0/0").(string)
}

func (o *Device) GetModelNumber() string {
	return o.Data.Get("/0/1").(string)
}

func (o *Device) GetSerialNumber() string {
	return o.Data.Get("/0/2").(string)
}

func (o *Device) GetFirmwareVersion() string {
	return o.Data.Get("/0/3").(string)
}

func (o *Device) Reboot() ResponseValue {
	return values.Empty()
}

func (o *Device) FactoryReset() ResponseValue {
	return values.Empty()
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
	return o.Data.Get("/0/9").(int)
}

func (o *Device) GetMemoryFree() int {
	return o.Data.Get("/0/10").(int)
}

func (o *Device) GetErrorCode() []int {
	return []int{0}
}

func (o *Device) ResetErrorCode() string {
	return ""
}

func (o *Device) GetCurrentTime() time.Time {
	return o.Data.Get("/0/13").(time.Time)
}

func (o *Device) GetUtcOffset() string {
	return o.Data.Get("/0/14").(string)
}

func (o *Device) GetTimezone() string {
	return o.Data.Get("/0/15").(string)
}

func (o *Device) GetSupportedBindingMode() string {
	return o.Data.Get("/0/16").(string)
}

func NewExampleDeviceObject(reg Registry) *Device {
	data := &core.ObjectsData{
		Data: make(map[string]interface{}),
	}

	data.Put("/0/0", "Open Mobile Alliance")
	data.Put("/0/1", "Lightweight M2M Client")
	data.Put("/0/2", "345000123")
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
	data.Put("/0/13", time.Unix(1367491215, 0))
	data.Put("/0/14", "+02:00")
	data.Put("/0/15", "")
	data.Put("/0/16", "U")

	return &Device{
		Model: reg.GetModel(oma.OBJECT_LWM2M_DEVICE),
		Data:  data,
	}
}

/*

*/
