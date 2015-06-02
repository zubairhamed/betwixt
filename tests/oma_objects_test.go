package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/zubairhamed/betwixt/api"
	"github.com/zubairhamed/betwixt/client"
	"github.com/zubairhamed/betwixt/core/request"
	"github.com/zubairhamed/betwixt/examples/obj/basic"
	"github.com/zubairhamed/betwixt/objects/oma"
	"github.com/zubairhamed/betwixt/registry"
	"testing"
	"time"
)

func TestExampleObjects(t *testing.T) {
	cli := client.NewDefaultClient(":0", "localhost:5683")

	reg := registry.NewDefaultObjectRegistry()
	cli.UseRegistry(reg)

	cli.EnableObject(oma.OBJECT_LWM2M_DEVICE, basic.NewExampleDeviceObject(reg))
	cli.EnableObject(oma.OBJECT_LWM2M_SECURITY, basic.NewExampleSecurityObject(reg))
	cli.EnableObject(oma.OBJECT_LWM2M_ACCESS_CONTROL, basic.NewExampleAccessControlObject(reg))
	cli.EnableObject(oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING, basic.NewExampleConnectivityMonitoringObject(reg))
	cli.EnableObject(oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS, basic.NewExampleConnectivityStatisticsObject(reg))
	cli.EnableObject(oma.OBJECT_LWM2M_FIRMWARE_UPDATE, basic.NewExampleFirmwareUpdateObject(reg))
	cli.EnableObject(oma.OBJECT_LWM2M_LOCATION, basic.NewExampleLocationObject(reg))
	cli.EnableObject(oma.OBJECT_LWM2M_SERVER, basic.NewExampleConnectivityMonitoringObject(reg))

	instDevice := reg.CreateObjectInstance(oma.OBJECT_LWM2M_DEVICE, 0)
	instSec := reg.CreateObjectInstance(oma.OBJECT_LWM2M_SECURITY, 0)
	instAccCtrl := reg.CreateObjectInstance(oma.OBJECT_LWM2M_ACCESS_CONTROL, 0)
	instConnMon := reg.CreateObjectInstance(oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING, 0)
	instConnStats := reg.CreateObjectInstance(oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS, 0)
	instFwUpdate := reg.CreateObjectInstance(oma.OBJECT_LWM2M_FIRMWARE_UPDATE, 0)
	instLocation := reg.CreateObjectInstance(oma.OBJECT_LWM2M_LOCATION, 0)
	instServer := reg.CreateObjectInstance(oma.OBJECT_LWM2M_SERVER, 0)

	cli.AddObjectInstances(
		instDevice,
		instSec,
		instAccCtrl,
		instConnMon,
		instConnStats,
		instFwUpdate,
		instLocation,
		instServer,
	)

	// Check added enablers
	test_enablers := []struct {
		input api.LWM2MObjectType
	}{
		{oma.OBJECT_LWM2M_DEVICE},
		{oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS},
		{oma.OBJECT_LWM2M_ACCESS_CONTROL},
		{oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING},
		{oma.OBJECT_LWM2M_FIRMWARE_UPDATE},
		{oma.OBJECT_LWM2M_LOCATION},
		{oma.OBJECT_LWM2M_SECURITY},
		{oma.OBJECT_LWM2M_SERVER},
	}

	for _, c := range test_enablers {
		assert.NotNil(t, cli.GetObjectEnabler(c.input), "Enabler returned nil", c.input)
	}

	// Device Object
	test_obj_1 := []struct {
		instanceId int
		resourceId int
		expected   interface{}
		typeId     api.LWM2MObjectType
	}{
		{0, 0, "Open Mobile Alliance", oma.OBJECT_LWM2M_DEVICE},
		{0, 1, "Lightweight M2M Client", oma.OBJECT_LWM2M_DEVICE},
		{0, 2, "345000123", oma.OBJECT_LWM2M_DEVICE},
		{0, 3, "1.0", oma.OBJECT_LWM2M_DEVICE},
		// {0, 6, []int{1, 5}, oma.OBJECT_LWM2M_DEVICE},
		// {0, 7, []int{3800, 5000}, oma.OBJECT_LWM2M_DEVICE},
		// {0, 8, []int{125, 900}, oma.OBJECT_LWM2M_DEVICE},
		{0, 10, 15, oma.OBJECT_LWM2M_DEVICE},
		// {0, 11, []int{0}, oma.OBJECT_LWM2M_DEVICE},
		{0, 13, time.Unix(1367491215, 0), oma.OBJECT_LWM2M_DEVICE},
		{0, 14, "+02:00", oma.OBJECT_LWM2M_DEVICE},
		{0, 15, "+02:00", oma.OBJECT_LWM2M_DEVICE},
		{0, 16, "U", oma.OBJECT_LWM2M_DEVICE},
	}

	for _, c := range test_obj_1 {
		en := cli.GetObjectEnabler(c.typeId)
		lwReq := request.Nil(api.OPERATIONTYPE_READ)
		response := en.OnRead(c.instanceId, c.resourceId, lwReq)
		val := response.GetResponseValue().GetValue()

		assert.Equal(t, val, c.expected, "Unexpected value returned for enabler OnRead: ", val, "vs", c.expected)
	}
}
