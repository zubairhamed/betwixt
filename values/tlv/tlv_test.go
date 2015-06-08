package tlv

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/zubairhamed/betwixt/tests"
	"github.com/zubairhamed/betwixt/objdefs/oma"
)

func TestObjectInstancesToTlv(t *testing.T) {
	omaObjects := &oma.LWM2MCoreObjects{}
	reg := tests.NewMockRegistry(omaObjects)
	deviceModel := omaObjects.GetObject(oma.OBJECT_LWM2M_DEVICE)

	device := tests.NewTestDeviceObject(deviceModel)
	obj := tests.NewMockObject(3, device, reg)
	obj.AddInstance(0)

	_, err := TlvPayloadFromObjects(obj, reg)

	assert.Nil(t, err, "Error thrown attempting to convert Object instance to TLV")
}
