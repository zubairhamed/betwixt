## Betwixt - A LWM2M Client and Server in Go
[![GoDoc](https://godoc.org/github.com/zubairhamed/betwixt?status.svg)](https://godoc.org/github.com/zubairhamed/betwixt)
[![Build Status](https://drone.io/github.com/zubairhamed/betwixt/status.png)](https://drone.io/github.com/zubairhamed/betwixt/latest)
[![Coverage Status](https://coveralls.io/repos/zubairhamed/betwixt/badge.svg?branch=master)](https://coveralls.io/r/zubairhamed/betwixt?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/zubairhamed/betwixt)](https://goreportcard.com/report/github.com/zubairhamed/betwixt)

#### Betwixt is a Lightweight M2M implementation written in Go
OMA Lightweight M2M is a protocol from the Open Mobile Alliance for M2M or IoT device management. Lightweight M2M enabler defines the application layer communication protocol between a LWM2M Server and a LWM2M Client, which is located in a LWM2M Device. 

The OMA Lightweight M2M enabler includes device management and service enablement for LWM2M Devices. The target LWM2M Devices for this enabler are mainly resource constrained devices. Therefore, this enabler makes use of a light and compact protocol as well as an efficient resource data model. It provides a choice for the M2M Service Provider to deploy a M2M system to provide service to the M2M User. 

### Basic Client Example
```go
package main

import (
	"github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/betwixt/examples"
	"github.com/zubairhamed/betwixt/examples/objects"
)

func main() {
	cli := examples.StandardCommandLineFlags()

	registry := betwixt.NewDefaultObjectRegistry()
	c := betwixt.NewLwm2mClient("TestClient", ":0", cli.Server, registry)

	setupResources(c, registry)

	c.OnStartup(func() {
		c.Register(cli.Name)
	})

	c.Start()
}

func setupResources(client betwixt.LWM2MClient, reg betwixt.Registry) {
	client.SetEnabler(betwixt.OMA_OBJECT_LWM2M_SECURITY, objects.NewExampleSecurityObject(reg))
	client.AddObjectInstances(betwixt.OMA_OBJECT_LWM2M_SECURITY, 0, 1, 2)

	client.SetEnabler(betwixt.OMA_OBJECT_LWM2M_SERVER, objects.NewExampleServerObject(reg))
	client.AddObjectInstance(betwixt.OMA_OBJECT_LWM2M_SERVER, 1)

	client.SetEnabler(betwixt.OMA_OBJECT_LWM2M_DEVICE, objects.NewExampleDeviceObject(reg))
	client.AddObjectInstance(betwixt.OMA_OBJECT_LWM2M_DEVICE, 0)

	client.EnableObject(betwixt.OMA_OBJECT_LWM2M_ACCESS_CONTROL, objects.NewExampleAccessControlObject(reg))
	client.AddObjectInstances(betwixt.OMA_OBJECT_LWM2M_ACCESS_CONTROL, 0, 1, 2)

	client.EnableObject(betwixt.OMA_OBJECT_LWM2M_CONNECTIVITY_MONITORING, objects.NewExampleConnectivityMonitoringObject(reg))
	client.AddObjectInstance(betwixt.OMA_OBJECT_LWM2M_CONNECTIVITY_MONITORING, 0)

	client.EnableObject(betwixt.OMA_OBJECT_LWM2M_FIRMWARE_UPDATE, objects.NewExampleFirmwareUpdateObject(reg))
	client.AddObjectInstance(betwixt.OMA_OBJECT_LWM2M_FIRMWARE_UPDATE, 0)

	client.EnableObject(betwixt.OMA_OBJECT_LWM2M_LOCATION, objects.NewExampleLocationObject(reg))
	client.EnableObject(betwixt.OMA_OBJECT_LWM2M_CONNECTIVITY_STATISTICS, objects.NewExampleConnectivityStatisticsObject(reg))
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


