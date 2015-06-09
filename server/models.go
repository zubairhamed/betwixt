package server

import (
	"github.com/zubairhamed/betwixt"
	"time"
)

type ServerStatistics struct {
	requestsCount int
}

func (s *ServerStatistics) IncrementCoapRequestsCount() {
	s.requestsCount++
}

func (s *ServerStatistics) GetRequestsCount() int {
	return s.requestsCount
}

// Returns a new instance of DefaultRegisteredClient implementing RegisteredClient
func NewRegisteredClient(ep string, id string, addr string) betwixt.RegisteredClient {
	return &DefaultRegisteredClient{
		name:       ep,
		id:         id,
		addr:       addr,
		regDate:    time.Now(),
		updateDate: time.Now(),
	}
}

type DefaultRegisteredClient struct {
	id             string
	name           string
	lifetime       int
	version        string
	bindingMode    betwixt.BindingMode
	smsNumber      string
	addr           string
	regDate        time.Time
	updateDate     time.Time
	enabledObjects map[betwixt.LWM2MObjectType]betwixt.Object
}

func (c *DefaultRegisteredClient) GetId() string {
	return c.id
}

func (c *DefaultRegisteredClient) GetName() string {
	return c.name
}

func (c *DefaultRegisteredClient) GetLifetime() int {
	return c.lifetime
}

func (c *DefaultRegisteredClient) GetVersion() string {
	return c.version
}

func (c *DefaultRegisteredClient) GetBindingMode() betwixt.BindingMode {
	return c.bindingMode
}

func (c *DefaultRegisteredClient) GetSmsNumber() string {
	return c.smsNumber
}

func (c *DefaultRegisteredClient) GetRegistrationDate() time.Time {
	return c.regDate
}

func (c *DefaultRegisteredClient) Update() {
	c.updateDate = time.Now()
}

func (c *DefaultRegisteredClient) LastUpdate() time.Time {
	return c.updateDate
}

func (c *DefaultRegisteredClient) SetObjects(objects map[betwixt.LWM2MObjectType]betwixt.Object) {
	c.enabledObjects = objects
}

func (c *DefaultRegisteredClient) GetObjects() map[betwixt.LWM2MObjectType]betwixt.Object {
	return c.enabledObjects
}

func (c *DefaultRegisteredClient) GetObject(t betwixt.LWM2MObjectType)(betwixt.Object) {
	return c.enabledObjects[t]
}
