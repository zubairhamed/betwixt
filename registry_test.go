package betwixt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegistry(t *testing.T) {
	reg := NewDefaultObjectRegistry()
	cases := []struct {
		o LWM2MObjectType
	}{
		{OMA_OBJECT_LWM2M_SECURITY},
		{OMA_OBJECT_LWM2M_SERVER},
		{OMA_OBJECT_LWM2M_ACCESS_CONTROL},
		{OMA_OBJECT_LWM2M_DEVICE},
		{OMA_OBJECT_LWM2M_CONNECTIVITY_MONITORING},
		{OMA_OBJECT_LWM2M_FIRMWARE_UPDATE},
		{OMA_OBJECT_LWM2M_LOCATION},
		{OMA_OBJECT_LWM2M_CONNECTIVITY_STATISTICS},
	}

	assert.Equal(t, 3, len(reg.GetMandatory()))
	assert.Equal(t, 26, len(reg.GetDefinitions()))

	for _, c := range cases {
		assert.NotNil(t, reg.GetDefinition(c.o), "Created an LWM2M Object: ", c.o)
	}

	assert.Nil(t, reg.GetDefinition(LWM2MObjectType(9999)))
}
