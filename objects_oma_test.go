package betwixt

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestExampleObjects(t *testing.T) {
	omaObjects := &LWM2MCoreObjects{}

	reg := NewMockRegistry(omaObjects)
	cli := NewDefaultClient("0", "localhost:5683", reg)

	deviceModel := omaObjects.GetObject(OMA_OBJECT_LWM2M_DEVICE)

	cli.SetEnabler(OMA_OBJECT_LWM2M_SERVER, NewNullEnabler())
	cli.SetEnabler(OMA_OBJECT_LWM2M_DEVICE, NewTestDeviceObject(deviceModel))
	cli.SetEnabler(OMA_OBJECT_LWM2M_SECURITY, NewNullEnabler())
	cli.EnableObject(OMA_OBJECT_LWM2M_ACCESS_CONTROL, NewNullEnabler())
	cli.EnableObject(OMA_OBJECT_LWM2M_CONNECTIVITY_MONITORING, NewNullEnabler())
	cli.EnableObject(OMA_OBJECT_LWM2M_CONNECTIVITY_STATISTICS, NewNullEnabler())
	cli.EnableObject(OMA_OBJECT_LWM2M_FIRMWARE_UPDATE, NewNullEnabler())
	cli.EnableObject(OMA_OBJECT_LWM2M_LOCATION, NewNullEnabler())

	// Check added enablers
	test_enablers := []struct {
		input LWM2MObjectType
	}{
		{OMA_OBJECT_LWM2M_DEVICE},
		{OMA_OBJECT_LWM2M_CONNECTIVITY_STATISTICS},
		{OMA_OBJECT_LWM2M_ACCESS_CONTROL},
		{OMA_OBJECT_LWM2M_CONNECTIVITY_MONITORING},
		{OMA_OBJECT_LWM2M_FIRMWARE_UPDATE},
		{OMA_OBJECT_LWM2M_LOCATION},
		{OMA_OBJECT_LWM2M_SECURITY},
		{OMA_OBJECT_LWM2M_SERVER},
	}

	for _, c := range test_enablers {
		obj := cli.GetObject(c.input)
		assert.NotNil(t, obj, "Object returned nil", c.input)
		if obj != nil {
			en := obj.GetEnabler()
			assert.NotNil(t, en, "Enabler returned nil", c.input)
		}
	}

	// Device Object
	test_obj_1 := []struct {
		instanceId int
		resourceId int
		expected   interface{}
		typeId     LWM2MObjectType
	}{
		{0, 0, "Open Mobile Alliance", OMA_OBJECT_LWM2M_DEVICE},
		{0, 1, "Lightweight M2M Client", OMA_OBJECT_LWM2M_DEVICE},
		{0, 2, "345000123", OMA_OBJECT_LWM2M_DEVICE},
		{0, 3, "1.0", OMA_OBJECT_LWM2M_DEVICE},
		// {0, 6, []int{1, 5}, OBJECT_LWM2M_DEVICE},
		// {0, 7, []int{3800, 5000}, OBJECT_LWM2M_DEVICE},
		// {0, 8, []int{125, 900}, OBJECT_LWM2M_DEVICE},
		{0, 10, 15, OMA_OBJECT_LWM2M_DEVICE},
		// {0, 11, []int{0}, OBJECT_LWM2M_DEVICE},
		{0, 13, time.Unix(1367491215, 0), OMA_OBJECT_LWM2M_DEVICE},
		{0, 14, "+02:00", OMA_OBJECT_LWM2M_DEVICE},
		{0, 15, "+02:00", OMA_OBJECT_LWM2M_DEVICE},
		{0, 16, "U", OMA_OBJECT_LWM2M_DEVICE},
	}

	for _, c := range test_obj_1 {
		en := cli.GetObject(c.typeId).GetEnabler()
		lwReq := Nil(OPERATIONTYPE_READ)
		response := en.OnRead(c.instanceId, c.resourceId, lwReq)
		val := response.GetResponseValue().GetValue()

		assert.Equal(t, c.expected, val, "Unexpected value returned for enabler OnRead: ", val, "vs", c.expected)
	}
}
