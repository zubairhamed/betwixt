package tests

import . "github.com/zubairhamed/betwixt"

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
