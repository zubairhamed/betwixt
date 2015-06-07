package tests

import (
    . "github.com/zubairhamed/betwixt"
    "time"
    "github.com/zubairhamed/betwixt/response"
    "github.com/zubairhamed/betwixt/values"
    "github.com/zubairhamed/betwixt/objdefs/oma"
)

type TestDeviceObject struct {
    Model       ObjectDefinition
    currentTime time.Time
    utcOffset   string
    timeZone    string
}

func (o *TestDeviceObject) OnExecute(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
    return response.Changed()
}

func (o *TestDeviceObject) OnCreate(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
    return response.Created()
}

func (o *TestDeviceObject) OnDelete(instanceId int, req Lwm2mRequest) Lwm2mResponse {
    return response.Deleted()
}

func (o *TestDeviceObject) OnRead(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
    if resourceId == -1 {
        // Read Object Instance
    } else {
        // Read Resource Instance
        var val ResponseValue

        resource := o.Model.GetResource(resourceId)
        switch resourceId {
            case 0:
            val = values.String("Open Mobile Alliance")
            break

            case 1:
            val = values.String("Lightweight M2M Client")
            break

            case 2:
            val = values.String("345000123")
            break

            case 3:
            val = values.String("1.0")
            break

            case 6:
            val, _ = values.TlvPayloadFromIntResource(resource, []int{oma.POWERSOURCE_INTERNAL, oma.POWERSOURCE_USB})
            break

            case 7:
            val, _ = values.TlvPayloadFromIntResource(resource, []int{3800, 5000})
            break

            case 8:
            val, _ = values.TlvPayloadFromIntResource(resource, []int{125, 900})
            break

            case 9:
            val = values.Integer(100)
            break

            case 10:
            val = values.Integer(15)
            break

            case 11:
            val, _ = values.TlvPayloadFromIntResource(resource, []int{0})
            break

            case 13:
            val = values.Time(o.currentTime)
            break

            case 14:
            val = values.String(o.utcOffset)
            break

            case 15:
            val = values.String(o.timeZone)
            break

            case 16:
            val = values.String(string(BINDINGMODE_UDP))
            break

            default:
            break
        }
        return response.Content(val)
    }
    return response.NotFound()
}

func (o *TestDeviceObject) OnWrite(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
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
        return response.NotFound()
    }
    return response.Changed()
}

func (o *TestDeviceObject) Reboot() ResponseValue {
    return values.Empty()
}

func (o *TestDeviceObject) FactoryReset() ResponseValue {
    return values.Empty()
}

func (o *TestDeviceObject) ResetErrorCode() string {
    return ""
}

func NewTestDeviceObject(reg Registry) *TestDeviceObject {
    return &TestDeviceObject{
        Model:       reg.GetModel(oma.OBJECT_LWM2M_DEVICE),
        currentTime: time.Unix(1367491215, 0),
        utcOffset:   "+02:00",
        timeZone:    "+02:00",
    }
}
