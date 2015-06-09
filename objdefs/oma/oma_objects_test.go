package oma

import (
	"github.com/stretchr/testify/assert"
	. "github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/betwixt/client"
	"github.com/zubairhamed/betwixt/enablers"
	"github.com/zubairhamed/betwixt/request"
	"github.com/zubairhamed/betwixt/tests"
	"testing"
	"time"
)

func TestExampleObjects(t *testing.T) {
	objs := tests.NewMockObjectSource()

	reg := tests.NewMockRegistry(objs)
	cli := client.NewDefaultClient(":0", "localhost:5683", reg)

	deviceModel := objs.GetObject(OBJECT_LWM2M_DEVICE)

	cli.SetEnabler(OBJECT_LWM2M_SERVER, enablers.NewNullEnabler())
	cli.SetEnabler(OBJECT_LWM2M_DEVICE, tests.NewTestDeviceObject(deviceModel))
	cli.SetEnabler(OBJECT_LWM2M_SECURITY, enablers.NewNullEnabler())
	cli.EnableObject(OBJECT_LWM2M_ACCESS_CONTROL, enablers.NewNullEnabler())
	cli.EnableObject(OBJECT_LWM2M_CONNECTIVITY_MONITORING, enablers.NewNullEnabler())
	cli.EnableObject(OBJECT_LWM2M_CONNECTIVITY_STATISTICS, enablers.NewNullEnabler())
	cli.EnableObject(OBJECT_LWM2M_FIRMWARE_UPDATE, enablers.NewNullEnabler())
	cli.EnableObject(OBJECT_LWM2M_LOCATION, enablers.NewNullEnabler())

	// Check added enablers
	test_enablers := []struct {
		input LWM2MObjectType
	}{
		{OBJECT_LWM2M_DEVICE},
		{OBJECT_LWM2M_CONNECTIVITY_STATISTICS},
		{OBJECT_LWM2M_ACCESS_CONTROL},
		{OBJECT_LWM2M_CONNECTIVITY_MONITORING},
		{OBJECT_LWM2M_FIRMWARE_UPDATE},
		{OBJECT_LWM2M_LOCATION},
		{OBJECT_LWM2M_SECURITY},
		{OBJECT_LWM2M_SERVER},
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
		{0, 0, "Open Mobile Alliance", OBJECT_LWM2M_DEVICE},
		{0, 1, "Lightweight M2M Client", OBJECT_LWM2M_DEVICE},
		{0, 2, "345000123", OBJECT_LWM2M_DEVICE},
		{0, 3, "1.0", OBJECT_LWM2M_DEVICE},
		// {0, 6, []int{1, 5}, OBJECT_LWM2M_DEVICE},
		// {0, 7, []int{3800, 5000}, OBJECT_LWM2M_DEVICE},
		// {0, 8, []int{125, 900}, OBJECT_LWM2M_DEVICE},
		{0, 10, 15, OBJECT_LWM2M_DEVICE},
		// {0, 11, []int{0}, OBJECT_LWM2M_DEVICE},
		{0, 13, time.Unix(1367491215, 0), OBJECT_LWM2M_DEVICE},
		{0, 14, "+02:00", OBJECT_LWM2M_DEVICE},
		{0, 15, "+02:00", OBJECT_LWM2M_DEVICE},
		{0, 16, "U", OBJECT_LWM2M_DEVICE},
	}

	for _, c := range test_obj_1 {
		en := cli.GetObject(c.typeId).GetEnabler()
		lwReq := request.Nil(OPERATIONTYPE_READ)
		response := en.OnRead(c.instanceId, c.resourceId, lwReq)
		val := response.GetResponseValue().GetValue()

		assert.Equal(t, c.expected, val, "Unexpected value returned for enabler OnRead: ", val, "vs", c.expected)
	}
}
