package registry

import (
	"github.com/stretchr/testify/assert"
	. "github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/betwixt/objectdefs/oma"
	"testing"
)

func TestRegistry(t *testing.T) {
	reg := NewDefaultObjectRegistry()
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

	assert.Equal(t, 3, len(reg.GetMandatory()))
	assert.Equal(t, 26, len(reg.GetDefinitions()))

	for _, c := range cases {
		assert.NotNil(t, reg.GetDefinition(c.o), "Created an LWM2M Object: ", c.o)
	}

	assert.Nil(t, reg.GetDefinition(LWM2MObjectType(9999)))
}
