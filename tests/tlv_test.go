package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/zubairhamed/betwixt/client"
	"github.com/zubairhamed/betwixt/core"
	"github.com/zubairhamed/betwixt/examples/obj/basic"
	"github.com/zubairhamed/betwixt/objects/oma"
	"github.com/zubairhamed/betwixt/registry"
	"testing"
)

func TestObjectInstancesToTlv(t *testing.T) {
	cli := client.NewDefaultClient(":0", "localhost:5683")

	reg := registry.NewDefaultObjectRegistry()
	cli.UseRegistry(reg)

	device := basic.NewExampleDeviceObject(reg)

	cli.EnableObject(oma.OBJECT_LWM2M_DEVICE, device)
	cli.AddObjectInstance(oma.OBJECT_LWM2M_DEVICE, 0)

	obj := cli.GetObject(oma.OBJECT_LWM2M_DEVICE)
	_, err := core.TlvPayloadFromObjects(obj, cli.GetRegistry())

	assert.Nil(t, err, "Error thrown attempting to convert Object instance to TLV")
}
