package tests

import (
	. "github.com/zubairhamed/betwixt"
	"errors"
)

func NewMockClient() LWM2MClient {
	return &MockClient{
		enabledObjects: make(map[LWM2MObjectType]Object),
	}
}

type MockClient struct {
	enabledObjects map[LWM2MObjectType]Object
}


func NewMockObject(t LWM2MObjectType, enabler ObjectEnabler, reg Registry) Object {
	return &MockObject {

	}
}

func (c *MockClient) AddObjectInstance(LWM2MObjectType, int) error {
	return nil
}
func (c *MockClient) AddObjectInstances(LWM2MObjectType, ...int) {}
func (c *MockClient) AddResource() {}
func (c *MockClient) AddObject() {}
func (c *MockClient) Register(string) string {
	return ""
}
func (c *MockClient) Deregister() {}
func (c *MockClient) Update() {}
func (c *MockClient) UseRegistry(Registry) {}

func (c *MockClient) EnableObject(t LWM2MObjectType, e ObjectEnabler) error {
	_, ok := c.enabledObjects[t]
	if !ok {
		c.enabledObjects[t] = NewMockObject(t, e, nil)

		return nil
	} else {
		return errors.New("Object already enabled")
	}
}

func (c *MockClient) SetEnabler(LWM2MObjectType, ObjectEnabler) {}
func (c *MockClient) GetRegistry() Registry {
	return nil
}
func (c *MockClient) GetEnabledObjects() map[LWM2MObjectType]Object {
	return c.enabledObjects
}

func (c *MockClient) Start() {}
func (c *MockClient) OnStartup(FnOnStartup) {}
func (c *MockClient) OnRead(FnOnRead) {}
func (c *MockClient) OnWrite(FnOnWrite) {}
func (c *MockClient) OnExecute(FnOnExecute) {}
func (c *MockClient) OnRegistered(FnOnRegistered) {}
func (c *MockClient) OnDeregistered(FnOnDeregistered) {}
func (c *MockClient) OnError(FnOnError) {}