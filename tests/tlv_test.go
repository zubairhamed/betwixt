package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/zubairhamed/go-lwm2m/core"
	"github.com/zubairhamed/go-lwm2m/examples/obj/basic"
	"github.com/zubairhamed/go-lwm2m/objects/oma"
	"github.com/zubairhamed/go-lwm2m/registry"
	"testing"
	"github.com/zubairhamed/go-lwm2m/client"
)

func TestObjectInstancesToTlv(t *testing.T) {
	cli := client.NewDefaultClient(":0", "localhost:5683")

	reg := registry.NewDefaultObjectRegistry()
	cli.UseRegistry(reg)

	device := basic.NewExampleDeviceObject(reg)

	cli.EnableObject(oma.OBJECT_LWM2M_DEVICE, device)
	instanceDevice := reg.CreateObjectInstance(oma.OBJECT_LWM2M_DEVICE, 0)
	cli.AddObjectInstances(instanceDevice)

	en := cli.GetObjectEnabler(oma.OBJECT_LWM2M_DEVICE)
	_, err := core.TlvPayloadFromObjects(en, cli.GetRegistry())

	assert.Nil(t, err, "Error thrown attempting to convert Object instance to TLV")
}
