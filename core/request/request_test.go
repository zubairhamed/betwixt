package request
import (
	"testing"
	"github.com/zubairhamed/canopus"
	"github.com/zubairhamed/betwixt"
	"github.com/stretchr/testify/assert"
)

func TestDefaultRequestObject(t *testing.T) {
	coapReq := canopus.NewRequest(canopus.TYPE_CONFIRMABLE, canopus.COAPCODE_204_CHANGED, 12345)
	coapReq.SetRequestURI("/rd")

	def := Default(coapReq, betwixt.OPERATIONTYPE_REGISTER)

	assert.NotNil(t, def.GetMessage())
	assert.Equal(t, uint16(12345), def.GetMessage().MessageId)
	assert.Equal(t, "/rd", def.GetPath())
	assert.Equal(t, betwixt.OPERATIONTYPE_REGISTER, def.GetOperationType())

	assert.NotNil(t, def.GetCoapRequest())
}

func TestNilRequestObject(t *testing.T) {
	nil := Nil(betwixt.OPERATIONTYPE_REGISTER)

	assert.Nil(t, nil.GetMessage())
	assert.Nil(t, nil.GetCoapRequest())
	assert.Equal(t, nil.GetOperationType(), betwixt.OPERATIONTYPE_REGISTER)
	assert.Equal(t, "", nil.GetPath())
}
