package main

import (
	. "github.com/zubairhamed/go-lwm2m/api"
	. "github.com/zubairhamed/go-lwm2m/examples/obj/basic"
	"github.com/zubairhamed/go-lwm2m/objects/oma"
	"github.com/zubairhamed/go-lwm2m/registry"
	"github.com/zubairhamed/go-lwm2m/client"
)

func main() {
	c := client.NewDefaultClient(":0", "localhost:5683")

	registry := registry.NewDefaultObjectRegistry()
	c.UseRegistry(registry)

	setupResources(c, registry)

	c.OnStartup(func() {
		c.Register("GOClient")
	})

	c.Start()
}

func setupResources(client LWM2MClient, reg Registry) {
	accessControl := NewExampleAccessControlObject(reg)
	device := NewExampleDeviceObject(reg)
	sec := NewExampleSecurityObject(reg)
	server := NewExampleServerObject(reg)
	connMon := NewExampleConnectivityMonitoringObject(reg)
	fwUpdate := NewExampleFirmwareUpdateObject(reg)
	location := NewExampleLocationObject(reg)
	connStats := NewExampleConnectivityStatisticsObject(reg)

	client.EnableObject(oma.OBJECT_LWM2M_SECURITY, sec)
	client.EnableObject(oma.OBJECT_LWM2M_SERVER, server)
	client.EnableObject(oma.OBJECT_LWM2M_ACCESS_CONTROL, accessControl)
	client.EnableObject(oma.OBJECT_LWM2M_DEVICE, device)
	client.EnableObject(oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING, connMon)
	client.EnableObject(oma.OBJECT_LWM2M_FIRMWARE_UPDATE, fwUpdate)
	client.EnableObject(oma.OBJECT_LWM2M_LOCATION, location)
	client.EnableObject(oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS, connStats)

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
	instanceFwUpdate := reg.CreateObjectInstance(oma.OBJECT_LWM2M_FIRMWARE_UPDATE, 0)

	client.AddObjectInstances(
		instanceSec1, instanceSec2, instanceSec3,
		instanceServer,
		instanceAccessCtrl1, instanceAccessCtrl2, instanceAccessCtrl3, instanceAccessCtrl4,
		instanceDevice,
		instanceConnMonitoring,
		instanceFwUpdate,
	)
}
