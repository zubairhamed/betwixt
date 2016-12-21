package main

import (
	"github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/betwixt/examples/shared"
	"github.com/zubairhamed/betwixt/examples/shared/obj"
	"github.com/zubairhamed/canopus"
)

func main() {
	registry := betwixt.NewDefaultObjectRegistry()
	client, err := betwixt.NewClient(shared.LWM2M_CLIENT_ID, registry, ":0")
	if err != nil {
		panic(err.Error())
	}
	setupResources(client, registry)

	event := make(chan canopus.ObserveMessage)
	conn, err := betwixt.Dial(shared.LWM2M_SERVER, eventChannel)
	if err != nil {
		panic(err.Error())
	}

	for {
		select {
		case msg.Type() == betwixt.EVENT_CONNECTED, open == true := <-event:
			conn.Register()
			break

		case msg.Type() == betwixt.EVENT_REGISTERED, open == true := <-event:
			break

		case msg.Type() == betwixt.EVENT_OBJECTQUERY, open == true := <-event:
			break

		case _, open == false := event:
			if err := conn.Shutdown(); err != nil {
				panic(err.Error())
			}
			// Graceful Exit
			break
		}
	}
	c.Start()
}

func setupResources(client betwixt.LWM2MClient, reg betwixt.Registry) {
	client.SetEnabler(betwixt.OMA_OBJECT_LWM2M_SECURITY, obj.NewExampleSecurityObject(reg))
	client.AddObjectInstances(betwixt.OMA_OBJECT_LWM2M_SECURITY, 0, 1, 2)

	client.SetEnabler(betwixt.OMA_OBJECT_LWM2M_SERVER, obj.NewExampleServerObject(reg))
	client.AddObjectInstance(betwixt.OMA_OBJECT_LWM2M_SERVER, 1)

	client.SetEnabler(betwixt.OMA_OBJECT_LWM2M_DEVICE, obj.NewExampleDeviceObject(reg))
	client.AddObjectInstance(betwixt.OMA_OBJECT_LWM2M_DEVICE, 0)

	client.EnableObject(betwixt.OMA_OBJECT_LWM2M_ACCESS_CONTROL, obj.NewExampleAccessControlObject(reg))
	client.AddObjectInstances(betwixt.OMA_OBJECT_LWM2M_ACCESS_CONTROL, 0, 1, 2)

	client.EnableObject(betwixt.OMA_OBJECT_LWM2M_CONNECTIVITY_MONITORING, obj.NewExampleConnectivityMonitoringObject(reg))
	client.AddObjectInstance(betwixt.OMA_OBJECT_LWM2M_CONNECTIVITY_MONITORING, 0)

	client.EnableObject(betwixt.OMA_OBJECT_LWM2M_FIRMWARE_UPDATE, obj.NewExampleFirmwareUpdateObject(reg))
	client.AddObjectInstance(betwixt.OMA_OBJECT_LWM2M_FIRMWARE_UPDATE, 0)

	client.EnableObject(betwixt.OMA_OBJECT_LWM2M_LOCATION, obj.NewExampleLocationObject(reg))
	client.EnableObject(betwixt.OMA_OBJECT_LWM2M_CONNECTIVITY_STATISTICS, obj.NewExampleConnectivityStatisticsObject(reg))
}
