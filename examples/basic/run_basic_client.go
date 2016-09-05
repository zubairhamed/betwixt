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
