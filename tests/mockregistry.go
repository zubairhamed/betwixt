package tests

import (
	. "github.com/zubairhamed/betwixt"
)

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
