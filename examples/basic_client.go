package main

import (
	. "github.com/zubairhamed/betwixt/api"
	"github.com/zubairhamed/betwixt/client"
	. "github.com/zubairhamed/betwixt/examples/obj/basic"
	"github.com/zubairhamed/betwixt/objects/oma"
	"github.com/zubairhamed/betwixt/registry"
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
	client.AddObjectInstances(oma.OBJECT_LWM2M_SECURITY, 0, 1, 2)

	client.EnableObject(oma.OBJECT_LWM2M_SERVER, server)
	client.AddObjectInstance(oma.OBJECT_LWM2M_SERVER, 1)

	client.EnableObject(oma.OBJECT_LWM2M_ACCESS_CONTROL, accessControl)
	client.AddObjectInstances(oma.OBJECT_LWM2M_ACCESS_CONTROL, 0, 1, 2)

	client.EnableObject(oma.OBJECT_LWM2M_DEVICE, device)
	client.AddObjectInstance(oma.OBJECT_LWM2M_DEVICE, 0)

	client.EnableObject(oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING, connMon)
	client.AddObjectInstance(oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING, 0)

	client.EnableObject(oma.OBJECT_LWM2M_FIRMWARE_UPDATE, fwUpdate)
	client.AddObjectInstance(oma.OBJECT_LWM2M_FIRMWARE_UPDATE, 0)

	client.EnableObject(oma.OBJECT_LWM2M_LOCATION, location)
	client.EnableObject(oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS, connStats)
}
