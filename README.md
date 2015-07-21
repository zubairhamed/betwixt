## Betwixt - A LWM2M Client and Server in Go
[![GoDoc](https://godoc.org/github.com/zubairhamed/betwixt?status.svg)](https://godoc.org/github.com/zubairhamed/betwixt)
[![Build Status](https://drone.io/github.com/zubairhamed/betwixt/status.png)](https://drone.io/github.com/zubairhamed/betwixt/latest)
[![Coverage Status](https://coveralls.io/repos/zubairhamed/betwixt/badge.svg?branch=master)](https://coveralls.io/r/zubairhamed/betwixt?branch=master)

#### Betwixt is a Lightweight M2M implementation written in Go 

### Basic Client Example
```go
package main

import (
	. "github.com/zubairhamed/betwixt"
	. "github.com/zubairhamed/betwixt/examples/client/basic/objects"
)

func main() {
	registry := NewDefaultObjectRegistry()
	c, _ := NewDefaultClient(":0", "localhost:5683", registry)

	setupResources(c, registry)

	c.OnStartup(func() {
		// When client has started up, register itself with a LWM2M server
		c.Register("betwixt")

		// TODO: Randomly fire change events for values changed
	})

	// Start client's CoAP listen
	c.Start()
}

func setupResources(client LWM2MClient, reg Registry) {
	client.SetEnabler(OMA_OBJECT_LWM2M_SECURITY, NewExampleSecurityObject(reg))
	client.AddObjectInstances(OMA_OBJECT_LWM2M_SECURITY, 0, 1, 2)

	client.SetEnabler(OMA_OBJECT_LWM2M_SERVER, NewExampleServerObject(reg))
	client.AddObjectInstance(OMA_OBJECT_LWM2M_SERVER, 1)

	client.SetEnabler(OMA_OBJECT_LWM2M_DEVICE, NewExampleDeviceObject(reg))
	client.AddObjectInstance(OMA_OBJECT_LWM2M_DEVICE, 0)

	client.EnableObject(OMA_OBJECT_LWM2M_ACCESS_CONTROL, NewExampleAccessControlObject(reg))
	client.AddObjectInstances(OMA_OBJECT_LWM2M_ACCESS_CONTROL, 0, 1, 2)

	client.EnableObject(OMA_OBJECT_LWM2M_CONNECTIVITY_MONITORING, NewExampleConnectivityMonitoringObject(reg))
	client.AddObjectInstance(OMA_OBJECT_LWM2M_CONNECTIVITY_MONITORING, 0)

	client.EnableObject(OMA_OBJECT_LWM2M_FIRMWARE_UPDATE, NewExampleFirmwareUpdateObject(reg))
	client.AddObjectInstance(OMA_OBJECT_LWM2M_FIRMWARE_UPDATE, 0)

	client.EnableObject(OMA_OBJECT_LWM2M_LOCATION, NewExampleLocationObject(reg))
	client.EnableObject(OMA_OBJECT_LWM2M_CONNECTIVITY_STATISTICS, NewExampleConnectivityStatisticsObject(reg))
}
```

### Implementing a LWM2M Object (LWM2M 'Device' Object)
```go
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
			break

		case 10:
			val = Integer(15)
			break

		case 11:
			val = MultipleIntegers(Integer(0))
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
```

### Minimal LWM2M Server (See /examples/server)
```go
package main

import (
	"github.com/zubairhamed/betwixt"
	. "github.com/zubairhamed/betwixt/examples/server"
)

func main() {
	s := NewDefaultServer("8081")

	registry := betwixt.NewDefaultObjectRegistry()

	s.UseRegistry(registry)

	s.Start()
}

```

## Limitations
- No dTLS support.

## LWM2M - Short of it
- Device Management Standard out of OMA

- Lightweight and compact binary protocol based on CoAP

- Targets as light as 8-bit MCUs

## Links
[A primer on LWM2M](http://www.slideshare.net/zdshelby/oma-lightweightm2-mtutorial)

[Specifications and Technical Information](http://technical.openmobilealliance.org/Technical/technical-information/release-program/current-releases/oma-lightweightm2m-v1-0)

[Leshan - A fairly complete Java-based LWM2M implementation](https://github.com/eclipse/leshan)


