package tests

import . "github.com/zubairhamed/betwixt"

type MockObject struct {
}

func (o *MockObject) AddInstance(int)    {}
func (o *MockObject) RemoveInstance(int) {}
func (o *MockObject) GetInstances() []int {
	return nil
}
func (o *MockObject) GetEnabler() ObjectEnabler {
	return nil
}
func (o *MockObject) GetType() LWM2MObjectType {
	return LWM2MObjectType(0)
}
func (o *MockObject) GetModel() ObjectDefinition {
	return nil
}
func (o *MockObject) SetEnabler(ObjectEnabler) {}
