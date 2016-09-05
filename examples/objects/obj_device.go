package objects

import (
	. "github.com/zubairhamed/betwixt"
	"time"
)

type DeviceObject struct {
	Model       ObjectDefinition
	currentTime time.Time
	utcOffset   string
	timeZone    string
}

func (o *DeviceObject) OnExecute(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Changed()
}

func (o *DeviceObject) OnCreate(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Created()
}

func (o *DeviceObject) OnDelete(instanceId int, req Lwm2mRequest) Lwm2mResponse {
	return Deleted()
}

func (o *DeviceObject) OnRead(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	if resourceId == -1 {
		// Read Object Instance
	} else {
		// Read Resource Instance
		var val Value

		switch resourceId {
		case 0:
			val = String("Open Mobile Alliance")
			break

		case 1:
			val = String("Lightweight M2M Client")
			break

		case 2:
			val = String("345000123")
			break

		case 3:
			val = String("1.0")
			break

		case 6:
			val = Integer(POWERSOURCE_INTERNAL, POWERSOURCE_USB)
			break

		case 7:
			val = Integer(3800, 5000)
			break

		case 8:
			val = Integer(125, 900)
			break

		case 9:
			val = Integer(100)
			// val = Integer(1000)
			// val = Integer(10000)
			// val = Integer(100000)
			// val = Integer(1000000)
			break

		case 10:
			val = Integer(15)
			break

		case 11:
			val = MultipleIntegers(Integer(0))
			// val, _ = tlv.TlvPayloadFromIntResource(resource, []int{0})
			break

		case 13:
			val = Time(o.currentTime)
			break

		case 14:
			val = String(o.utcOffset)
			break

		case 15:
			val = String(o.timeZone)
			break

		case 16:
			val = String(string(BINDINGMODE_UDP))
			break

		default:
			break
		}
		return Content(val)
	}
	return NotFound()
}

func (o *DeviceObject) OnWrite(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	val := req.GetMessage().Payload

	switch resourceId {
	case 13:
		break

	case 14:
		o.utcOffset = val.String()
		break

	case 15:
		o.timeZone = val.String()
		break

	default:
		return NotFound()
	}
	return Changed()
}

func (o *DeviceObject) Reboot() Value {
	return Empty()
}

func (o *DeviceObject) FactoryReset() Value {
	return Empty()
}

func (o *DeviceObject) ResetErrorCode() string {
	return ""
}

func NewExampleDeviceObject(reg Registry) *DeviceObject {
	return &DeviceObject{
		Model:       reg.GetDefinition(OMA_OBJECT_LWM2M_DEVICE),
		currentTime: time.Unix(1367491215, 0),
		utcOffset:   "+02:00",
		timeZone:    "+02:00",
	}
}
