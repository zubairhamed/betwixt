package tlv

import (
	"github.com/stretchr/testify/assert"
	"github.com/zubairhamed/betwixt/objdefs/oma"
	"github.com/zubairhamed/betwixt/tests"
	"testing"
	"log"
)

func TestObjectInstancesToTlv(t *testing.T) {
	omaObjects := &oma.LWM2MCoreObjects{}
	reg := tests.NewMockRegistry(omaObjects)
	deviceModel := omaObjects.GetObject(oma.OBJECT_LWM2M_DEVICE)

	device := tests.NewTestDeviceObject(deviceModel)
	obj := tests.NewMockObject(3, device, reg)
	obj.AddInstance(0)

	rv, err := TlvPayloadFromObjects(obj, reg)

	log.Println(rv.GetBytes())

	assert.Nil(t, err, "Error thrown attempting to convert Object instance to TLV")
}
