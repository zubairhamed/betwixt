package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/zubairhamed/betwixt/client"
	"github.com/zubairhamed/betwixt/objdefs/oma"
	"github.com/zubairhamed/betwixt/registry"
	"testing"
	"github.com/zubairhamed/betwixt/values"
)

func TestObjectInstancesToTlv(t *testing.T) {
	reg := registry.NewDefaultObjectRegistry()
	cli := client.NewDefaultClient(":0", "localhost:5683", reg)

	device := NewTestDeviceObject(reg)

	cli.SetEnabler(oma.OBJECT_LWM2M_DEVICE, device)
	cli.AddObjectInstance(oma.OBJECT_LWM2M_DEVICE, 0)

	obj := cli.GetObject(oma.OBJECT_LWM2M_DEVICE)
	_, err := values.TlvPayloadFromObjects(obj, cli.GetRegistry())

	assert.Nil(t, err, "Error thrown attempting to convert Object instance to TLV")
}
