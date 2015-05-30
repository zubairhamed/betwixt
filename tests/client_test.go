package tests

import (
	"github.com/stretchr/testify/assert"
	. "github.com/zubairhamed/go-lwm2m/api"
	"github.com/zubairhamed/go-lwm2m/objects/oma"
	"github.com/zubairhamed/go-lwm2m/registry"
	"testing"
	"github.com/zubairhamed/go-lwm2m/client"
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
	}{
		{oma.OBJECT_LWM2M_SERVER},
		{oma.OBJECT_LWM2M_ACCESS_CONTROL},
		{oma.OBJECT_LWM2M_DEVICE},
		{oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING},
		{oma.OBJECT_LWM2M_FIRMWARE_UPDATE},
		{oma.OBJECT_LWM2M_LOCATION},
		{oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS},
	}

	for _, c := range cases1 {
		err := cli.EnableObject(c.in, nil)

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
		assert.NotNil(t, cli.GetObjectEnabler(c.in), "Error getting object enabler: ", c)
	}

	inst1 := registry.CreateObjectInstance(oma.OBJECT_LWM2M_SECURITY, 0)
	inst2 := registry.CreateObjectInstance(oma.OBJECT_LWM2M_SECURITY, 1)
	inst3 := registry.CreateObjectInstance(oma.OBJECT_LWM2M_SECURITY, 2)

	assert.NotNil(t, inst1, "Error instantiating go-lwm2m object")
	assert.NotNil(t, inst2, "Error instantiating go-lwm2m object")
	assert.NotNil(t, inst3, "Error instantiating go-lwm2m object")

	cli.AddObjectInstances(inst1, inst2, inst3)

	cases3 := []struct {
		ot LWM2MObjectType
		oi int
	}{
		{oma.OBJECT_LWM2M_SECURITY, 0},
		{oma.OBJECT_LWM2M_SECURITY, 1},
		{oma.OBJECT_LWM2M_SECURITY, 2},
	}

	for _, c := range cases3 {
		assert.NotNil(t, cli.GetObjectInstance(c.ot, c.oi), "Object instance", c.oi, "not found")
	}
}

func TestRegistry(t *testing.T) {
	reg := registry.NewDefaultObjectRegistry()

	cases := []struct {
		o LWM2MObjectType
	}{
		{oma.OBJECT_LWM2M_SECURITY},
		{oma.OBJECT_LWM2M_SERVER},
		{oma.OBJECT_LWM2M_ACCESS_CONTROL},
		{oma.OBJECT_LWM2M_DEVICE},
		{oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING},
		{oma.OBJECT_LWM2M_FIRMWARE_UPDATE},
		{oma.OBJECT_LWM2M_LOCATION},
		{oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS},
	}

	for _, c := range cases {
		assert.NotNil(t, reg.CreateObjectInstance(c.o, 0), "Created an LWM2M Object: ", c.o)
	}
	assert.Nil(t, reg.CreateObjectInstance(LWM2MObjectType(-1), 0), "Created an unknown LWM2M Object")
}
