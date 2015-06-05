package tests

import (
	"github.com/stretchr/testify/assert"
	. "github.com/zubairhamed/betwixt/api"
	"github.com/zubairhamed/betwixt/client"
	"github.com/zubairhamed/betwixt/objects/oma"
	"github.com/zubairhamed/betwixt/registry"
	"testing"
	"github.com/zubairhamed/betwixt/examples/obj/basic"
)

func TestClient(t *testing.T) {

	cli := client.NewDefaultClient(":0", "localhost:5683")
	assert.NotNil(t, cli, "Error instantiating client")

	assert.NotNil(t, cli.EnableObject(oma.OBJECT_LWM2M_SERVER, nil), "Error should be thrown - registry not set")

	registry := registry.NewDefaultObjectRegistry()
	assert.NotNil(t, registry, "Error instantiating registry")

	cli.UseRegistry(registry)

	cases1 := []struct {
		in LWM2MObjectType
		en ObjectEnabler
	}{
		{oma.OBJECT_LWM2M_SERVER, basic.NewExampleServerObject(registry)},
		{oma.OBJECT_LWM2M_ACCESS_CONTROL, basic.NewExampleAccessControlObject(registry)},
		{oma.OBJECT_LWM2M_DEVICE, basic.NewExampleDeviceObject(registry)},
		{oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING, basic.NewExampleConnectivityMonitoringObject(registry)},
		{oma.OBJECT_LWM2M_FIRMWARE_UPDATE, basic.NewExampleFirmwareUpdateObject(registry)},
		{oma.OBJECT_LWM2M_LOCATION, basic.NewExampleLocationObject(registry)},
		{oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS, basic.NewExampleConnectivityStatisticsObject(registry)},
	}

	for _, c := range cases1 {
		err := cli.EnableObject(c.in, c.en)

		assert.Nil(t, err, "Error enabling object: ", c.in)
	}

	assert.Nil(t, cli.EnableObject(oma.OBJECT_LWM2M_SECURITY, nil), "Error enabling object")
	assert.NotNil(t, cli.EnableObject(oma.OBJECT_LWM2M_SECURITY, nil), "Object should already be enabled")

	cases2 := []struct {
		in LWM2MObjectType
	}{
		{oma.OBJECT_LWM2M_SERVER},
		{oma.OBJECT_LWM2M_ACCESS_CONTROL},
		{oma.OBJECT_LWM2M_DEVICE},
		{oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING},
		{oma.OBJECT_LWM2M_FIRMWARE_UPDATE},
		{oma.OBJECT_LWM2M_LOCATION},
		{oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS},
	}

	for _, c := range cases2 {
		o := cli.GetObject(c.in)
		assert.NotNil(t, o, "Error getting object: ", c)
		assert.NotNil(t, o.GetEnabler(), "Error getting object enabler: ", c)
	}

	cli.AddObjectInstances(oma.OBJECT_LWM2M_SECURITY, 0, 1, 2)

	assert.Equal(t, len(cli.GetObject(oma.OBJECT_LWM2M_SECURITY).GetInstances()), 3)
}

//func TestRegistry(t *testing.T) {
//	reg := registry.NewDefaultObjectRegistry()
//
//	cases := []struct {
//		o LWM2MObjectType
//	}{
//		{oma.OBJECT_LWM2M_SECURITY},
//		{oma.OBJECT_LWM2M_SERVER},
//		{oma.OBJECT_LWM2M_ACCESS_CONTROL},
//		{oma.OBJECT_LWM2M_DEVICE},
//		{oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING},
//		{oma.OBJECT_LWM2M_FIRMWARE_UPDATE},
//		{oma.OBJECT_LWM2M_LOCATION},
//		{oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS},
//	}
//
//	for _, c := range cases {
//		assert.NotNil(t, reg.CreateObjectInstance(c.o, 0), "Created an LWM2M Object: ", c.o)
//	}
//	assert.Nil(t, reg.CreateObjectInstance(LWM2MObjectType(-1), 0), "Created an unknown LWM2M Object")
//}
