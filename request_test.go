package betwixt

import (
	"github.com/stretchr/testify/assert"
	"github.com/zubairhamed/canopus"
	"testing"
)

func TestDefaultRequestObject(t *testing.T) {
	coapReq := canopus.NewRequest(canopus.MessageConfirmable, canopus.CoapCodeChanged, 12345)
	coapReq.SetRequestURI("/rd")

	def := Default(coapReq, OPERATIONTYPE_REGISTER)

	assert.NotNil(t, def.GetMessage())
	assert.Equal(t, uint16(12345), def.GetMessage().MessageID)
	assert.Equal(t, "/rd", def.GetMessage().GetURIPath())
	assert.Equal(t, OPERATIONTYPE_REGISTER, def.GetOperationType())

	assert.NotNil(t, def.GetCoapRequest())
}

func TestNilRequestObject(t *testing.T) {
	nil := Nil(OPERATIONTYPE_REGISTER)

	assert.Nil(t, nil.GetMessage())
	assert.Nil(t, nil.GetCoapRequest())
	assert.Equal(t, nil.GetOperationType(), OPERATIONTYPE_REGISTER)
	assert.Equal(t, "", nil.GetPath())
}
