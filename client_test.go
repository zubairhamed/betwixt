package betwixt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient(t *testing.T) {

	registry := NewDefaultObjectRegistry()
	cli := NewDefaultClient(":0", "localhost:5683", registry)
	assert.NotNil(t, cli, "Error instantiating client")
	assert.NotNil(t, registry, "Error instantiating registry")

	cases1 := []struct {
		in LWM2MObjectType
	}{
		{OMA_OBJECT_LWM2M_SERVER},
		{OMA_OBJECT_LWM2M_DEVICE},
		{OMA_OBJECT_LWM2M_SECURITY},
	}

	for _, c := range cases1 {
		err := cli.EnableObject(c.in, nil)

		assert.NotNil(t, err, "Object should already be enabled: ", c.in)
	}

	cases2 := []struct {
		in LWM2MObjectType
		en ObjectEnabler
	}{
		{OMA_OBJECT_LWM2M_ACCESS_CONTROL, NewNullEnabler()},
		{OMA_OBJECT_LWM2M_CONNECTIVITY_MONITORING, NewNullEnabler()},
		{OMA_OBJECT_LWM2M_FIRMWARE_UPDATE, NewNullEnabler()},
		{OMA_OBJECT_LWM2M_LOCATION, NewNullEnabler()},
		{OMA_OBJECT_LWM2M_CONNECTIVITY_STATISTICS, NewNullEnabler()},
	}

	for _, c := range cases2 {
		err := cli.EnableObject(c.in, c.en)

		assert.Nil(t, err, "Error enabling object: ", c.in)
	}

	cases3 := []struct {
		in LWM2MObjectType
	}{
		{OMA_OBJECT_LWM2M_SERVER},
		{OMA_OBJECT_LWM2M_ACCESS_CONTROL},
		{OMA_OBJECT_LWM2M_DEVICE},
		{OMA_OBJECT_LWM2M_CONNECTIVITY_MONITORING},
		{OMA_OBJECT_LWM2M_FIRMWARE_UPDATE},
		{OMA_OBJECT_LWM2M_LOCATION},
		{OMA_OBJECT_LWM2M_CONNECTIVITY_STATISTICS},
	}

	for _, c := range cases3 {
		o := cli.GetObject(c.in)
		assert.NotNil(t, o, "Error getting object: ", c)
		assert.NotNil(t, o.GetEnabler(), "Error getting object enabler: ", c)
	}

	cli.AddObjectInstances(OMA_OBJECT_LWM2M_SECURITY, 0, 1, 2)

	assert.Equal(t, len(cli.GetObject(OMA_OBJECT_LWM2M_SECURITY).GetInstances()), 3)
}
