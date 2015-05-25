package tests
import (
    "testing"
    "github.com/zubairhamed/go-lwm2m"
    "github.com/zubairhamed/go-lwm2m/objects/oma"
    "github.com/zubairhamed/go-lwm2m/examples/obj/basic"
    "github.com/zubairhamed/go-lwm2m/registry"
    "github.com/stretchr/testify/assert"
    "time"
)

func TestExampleObjects(t *testing.T) {
    client := lwm2m.NewLWM2MClient(":0", "localhost:5683")

    reg := registry.NewDefaultObjectRegistry()

    client.EnableObject(oma.OBJECT_LWM2M_DEVICE, basic.NewExampleDeviceObject(reg))
    client.EnableObject(oma.OBJECT_LWM2M_SECURITY, basic.NewExampleSecurityObject(reg))
    client.EnableObject(oma.OBJECT_LWM2M_ACCESS_CONTROL, basic.NewExampleAccessControlObject(reg))
    client.EnableObject(oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING, basic.NewExampleConnectivityMonitoringObject(reg))
    client.EnableObject(oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS, basic.NewExampleConnectivityStatisticsObject(reg))
    client.EnableObject(oma.OBJECT_LWM2M_FIRMWARE_UPDATE, basic.NewExampleFirmwareUpdateObject(reg))
    client.EnableObject(oma.OBJECT_LWM2M_LOCATION, basic.NewExampleLocationObject(reg))
    client.EnableObject(oma.OBJECT_LWM2M_SERVER, basic.NewExampleConnectivityMonitoringObject(reg))

    instDevice := reg.CreateObjectInstance(oma.OBJECT_LWM2M_DEVICE, 0)
    instSec := reg.CreateObjectInstance(oma.OBJECT_LWM2M_SECURITY, 0)
    instAccCtrl := reg.CreateObjectInstance(oma.OBJECT_LWM2M_ACCESS_CONTROL, 0)
    instConnMon := reg.CreateObjectInstance(oma.OBJECT_LWM2M_CONNECTIVITY_MONITORING, 0)
    instConnStats := reg.CreateObjectInstance(oma.OBJECT_LWM2M_CONNECTIVITY_STATISTICS, 0)
    instFwUpdate := reg.CreateObjectInstance(oma.OBJECT_LWM2M_FIRMWARE_UPDATE, 0)
    instLocation := reg.CreateObjectInstance(oma.OBJECT_LWM2M_LOCATION, 0)
    instServer := reg.CreateObjectInstance(oma.OBJECT_LWM2M_SERVER, 0)

    client.AddObjectInstances(
        instDevice,
        instSec,
        instAccCtrl,
        instConnMon,
        instConnStats,
        instFwUpdate,
        instLocation,
        instServer,
    )

    en := client.GetObjectEnabler(oma.OBJECT_LWM2M_DEVICE)
    en.OnRead(0, 0)
    // Device Object
    tests := []struct {
        instanceId  int
        resourceId  int
        expected    interface{}
    }{
        {0, 0, "Open Mobile Alliance"},
        {0, 1, "Lightweight M2M Client"},
        {0, 2, "345000123"},
        {0, 3, "1.0"},
        // {0, 6, []int{1, 5}},
        // {0, 7, []int{3800, 5000}},
        // {0, 8, []int{125, 900}},
        {0, 10, 15},
        // {0, 11, []int{0}},
        {0, 13, time.Unix(1367491215, 0)},
        {0, 14, "+02:00"},
        {0, 16, "U"},
    }

    for _, c := range tests {
        ret, _ := en.OnRead(c.instanceId, c.resourceId)
        val := ret.GetValue()

        assert.Equal(t, val, c.expected, "Unexpected value returned for enabler OnRead: ", val, "vs", c.expected)
    }



}