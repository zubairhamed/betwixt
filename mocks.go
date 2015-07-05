package betwixt

import (
	"errors"
	"github.com/zubairhamed/canopus"
	"github.com/zubairhamed/go-commons/network"
	"github.com/zubairhamed/go-commons/typeval"
	"time"
	"github.com/zubairhamed/sugoi"
)

type TestDeviceObject struct {
	Model       ObjectDefinition
	currentTime time.Time
	utcOffset   string
	timeZone    string
}

func (o *TestDeviceObject) OnExecute(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Changed()
}

func (o *TestDeviceObject) OnCreate(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	return Created()
}

func (o *TestDeviceObject) OnDelete(instanceId int, req Lwm2mRequest) Lwm2mResponse {
	return Deleted()
}

func (o *TestDeviceObject) OnRead(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	if resourceId == -1 {
		// Read Object Instance
	} else {
		// Read Resource Instance
		var val typeval.Value

		// resource := o.Model.GetResource(resourceId)
		switch resourceId {
		case 0:
			val = typeval.String("Open Mobile Alliance")
			break

		case 1:
			val = typeval.String("Lightweight M2M Client")
			break

		case 2:
			val = typeval.String("345000123")
			break

		case 3:
			val = typeval.String("1.0")
			break

		case 6:
			// val, _ = values.TlvPayloadFromIntResource(resource, []int{1, 5})
			break

		case 7:
			// val, _ = values.TlvPayloadFromIntResource(resource, []int{3800, 5000})
			break

		case 8:
			// val, _ = values.TlvPayloadFromIntResource(resource, []int{125, 900})
			break

		case 9:
			val = typeval.Integer(100)
			break

		case 10:
			val = typeval.Integer(15)
			break

		case 11:
			// val, _ = values.TlvPayloadFromIntResource(resource, []int{0})
			break

		case 13:
			val = typeval.Time(o.currentTime)
			break

		case 14:
			val = typeval.String(o.utcOffset)
			break

		case 15:
			val = typeval.String(o.timeZone)
			break

		case 16:
			val = typeval.String(string(BINDINGMODE_UDP))
			break

		default:
			break
		}
		return Content(val)
	}
	return NotFound()
}

func (o *TestDeviceObject) OnWrite(instanceId int, resourceId int, req Lwm2mRequest) Lwm2mResponse {
	val := req.GetMessage().Payload

	switch resourceId {
	case 13:
		break

	case 14:
		o.utcOffset = val.String()
		break

	case 15:
		o.timeZone = val.String()
		break

	default:
		return NotFound()
	}
	return Changed()
}

func (o *TestDeviceObject) Reboot() typeval.Value {
	return typeval.Empty()
}

func (o *TestDeviceObject) FactoryReset() typeval.Value {
	return typeval.Empty()
}

func (o *TestDeviceObject) ResetErrorCode() string {
	return ""
}

func NewTestDeviceObject(def ObjectDefinition) *TestDeviceObject {
	return &TestDeviceObject{
		Model:       def,
		currentTime: time.Unix(1367491215, 0),
		utcOffset:   "+02:00",
		timeZone:    "+02:00",
	}
}

type MockServerStatistics struct {
}

func (s *MockServerStatistics) IncrementCoapRequestsCount() {

}

func (s *MockServerStatistics) GetRequestsCount() int {
	return 0
}

func NewMockServer() Server {
	return &MockServer{
		stats:      &MockServerStatistics{},
		httpServer: network.NewDefaultHttpServer("8080"),
	}
}

type MockServer struct {
	stats      ServerStatistics
	httpServer *sugoi.SugoiServer
	coapServer *canopus.CoapServer
}

func (server *MockServer) Start() {

}

func (server *MockServer) UseRegistry(reg Registry) {

}

func (server *MockServer) On(e EventType, fn FnEvent) {

}

func (server *MockServer) GetClients() map[string]RegisteredClient {
	return make(map[string]RegisteredClient)
}

func (server *MockServer) GetStats() ServerStatistics {
	return server.stats
}

func (server *MockServer) GetHttpServer() *sugoi.SugoiServer {
	return server.httpServer
}

func (server *MockServer) GetCoapServer() *canopus.CoapServer {
	return server.coapServer
}

func (server *MockServer) GetClient(id string) RegisteredClient {
	return nil
}

func NewMockRegistry(s ...ObjectSource) Registry {
	reg := &MockRegistry{}

	reg.sources = []ObjectSource{}

	for _, o := range s {
		reg.Register(o)
	}
	return reg
}

type MockRegistry struct {
	sources []ObjectSource
}

func (r *MockRegistry) GetDefinition(t LWM2MObjectType) ObjectDefinition {
	return nil
}

func (m *MockRegistry) Register(s ObjectSource) {
	s.Initialize()
	m.sources = append(m.sources, s)
}

func (r *MockRegistry) GetMandatory() []ObjectDefinition {
	mandatory := []ObjectDefinition{}

	for _, s := range r.sources {
		objs := s.GetObjects()
		for _, o := range objs {
			if o.IsMandatory() {
				mandatory = append(mandatory, o)
			}
		}
	}
	return mandatory
}

func (m *MockRegistry) GetDefinitions() []ObjectDefinition {
	defs := []ObjectDefinition{}

	for _, s := range m.sources {
		if s != nil {
			for _, v := range s.GetObjects() {
				defs = append(defs, v)
			}
		}
	}
	return defs
}

func NewMockObject(t LWM2MObjectType, enabler ObjectEnabler, reg Registry) Object {
	def := reg.GetDefinition(t)

	return &MockObject{
		definition: def,
		typeId:     t,
		enabler:    enabler,
		instances:  make(map[int]bool),
	}
}

type MockObject struct {
	typeId     LWM2MObjectType
	definition ObjectDefinition
	enabler    ObjectEnabler
	instances  map[int]bool
}

func (o *MockObject) AddInstance(int)    {}
func (o *MockObject) RemoveInstance(int) {}
func (o *MockObject) GetInstances() []int {
	return nil
}
func (o *MockObject) GetEnabler() ObjectEnabler {
	return o.enabler
}
func (o *MockObject) GetType() LWM2MObjectType {
	return LWM2MObjectType(0)
}
func (o *MockObject) GetDefinition() ObjectDefinition {
	return nil
}
func (o *MockObject) SetEnabler(ObjectEnabler) {}

func NewMockClient() LWM2MClient {
	return &MockClient{
		enabledObjects: make(map[LWM2MObjectType]Object),
	}
}

type MockClient struct {
	enabledObjects map[LWM2MObjectType]Object
	registry       Registry
}

func (c *MockClient) AddObjectInstance(LWM2MObjectType, int) error {
	return nil
}
func (c *MockClient) Register(string) (string, error) {
	return "", nil
}
func (c *MockClient) UseRegistry(r Registry) {
	c.registry = r
}

func (c *MockClient) EnableObject(t LWM2MObjectType, e ObjectEnabler) error {
	_, ok := c.enabledObjects[t]
	if !ok {
		c.enabledObjects[t] = NewMockObject(t, e, c.GetRegistry())

		return nil
	} else {
		return errors.New("Object already enabled")
	}
}

func (c *MockClient) GetRegistry() Registry {
	return c.registry
}
func (c *MockClient) GetEnabledObjects() map[LWM2MObjectType]Object {
	return c.enabledObjects
}

func (c *MockClient) AddObjectInstances(LWM2MObjectType, ...int) {}
func (c *MockClient) AddResource()                               {}
func (c *MockClient) AddObject()                                 {}
func (c *MockClient) Deregister()                                {}
func (c *MockClient) Update()                                    {}
func (c *MockClient) SetEnabler(LWM2MObjectType, ObjectEnabler)  {}
func (c *MockClient) Start()                                     {}
func (c *MockClient) OnStartup(FnOnStartup)                      {}
func (c *MockClient) OnRead(FnOnRead)                            {}
func (c *MockClient) OnWrite(FnOnWrite)                          {}
func (c *MockClient) OnExecute(FnOnExecute)                      {}
func (c *MockClient) OnRegistered(FnOnRegistered)                {}
func (c *MockClient) OnDeregistered(FnOnDeregistered)            {}
func (c *MockClient) OnError(FnOnError)                          {}
