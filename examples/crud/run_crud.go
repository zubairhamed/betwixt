package main

import (
	"github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/betwixt/examples/objects"
	"log"
)

func main() {

	// Create Server
	server := CreateServer()
	server.OnRegistered(func(c betwixt.RegisteredClient) {
		log.Println("Registered a client with name", c.GetName())

		// Get all resources
		// v, err := c.ReadResource(3, 0, 0)
		// log.Println("Read Resource", v, err)

		// Create and validate resource

		// Delete and validate resource
	})

	go server.Serve()

	// Create Client
	client := CreateClient()

	client.OnStartup(func() {
		client.Register("betwixt")
	})
	go client.Start()

	<-make(chan struct{})
}

func CreateServer() *betwixt.LWM2MServer {
	store := betwixt.NewInMemoryStore()

	cfg := map[string]string{}
	server := betwixt.NewLwm2mServer("Betwixt LWM2M Server", store, cfg)
	registry := betwixt.NewDefaultObjectRegistry()
	server.UseRegistry(registry)

	return server
}

func CreateClient() betwixt.LWM2MClient {
	reg := betwixt.NewDefaultObjectRegistry()

	client := betwixt.NewLwm2mClient("Default Client", ":0", "localhost:5683", reg)

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

	return client
}
