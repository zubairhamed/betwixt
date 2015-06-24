package main

import (
	. "github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/betwixt/client"
	"github.com/zubairhamed/betwixt/core/registry"
	. "github.com/zubairhamed/betwixt/examples/client/basic"
	"github.com/zubairhamed/betwixt/objectdefs/oma"
)

func main() {
	registry := registry.NewDefaultObjectRegistry()
	c := client.NewDefaultClient(":0", "localhost:5683", registry)

	setupResources(c, registry)

	c.OnStartup(func() {
		c.Register("betwixt")
	})

	c.Start()
}

func setupResources(client LWM2MClient, reg Registry) {
	client.SetEnabler(oma.OBJECT_LWM2M_SECURITY, NewExampleSecurityObject(reg))
	client.AddObjectInstances(oma.OBJECT_LWM2M_SECURITY, 0, 1, 2)

	client.SetEnabler(oma.OBJECT_LWM2M_SERVER, NewExampleServerObject(reg))
	client.AddObjectInstance(oma.OBJECT_LWM2M_SERVER, 1)

	client.SetEnabler(oma.OBJECT_LWM2M_DEVICE, NewExampleDeviceObject(reg))
	client.AddObjectInstance(oma.OBJECT_LWM2M_DEVICE, 0)

	client.EnableObject(oma.OBJECT_LWM2M_ACCESS_CONTROL, NewExampleAccessControlObject(reg))
	client.AddObjectInstances(oma.OBJECT_LWM2M_ACCESS_CONTROL, 0, 1, 2)

	client.EnableObject(oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING, NewExampleConnectivityMonitoringObject(reg))
	client.AddObjectInstance(oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING, 0)

	client.EnableObject(oma.OBJECT_LWM2M_FIRMWARE_UPDATE, NewExampleFirmwareUpdateObject(reg))
	client.AddObjectInstance(oma.OBJECT_LWM2M_FIRMWARE_UPDATE, 0)

	client.EnableObject(oma.OBJECT_LWM2M_LOCATION, NewExampleLocationObject(reg))
	client.EnableObject(oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS, NewExampleConnectivityStatisticsObject(reg))
}
