package registry

import (
	. "github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/betwixt/objdefs/ipso"
	"github.com/zubairhamed/betwixt/objdefs/oma"
)

func NewDefaultObjectRegistry() Registry {
	reg := NewObjectRegistry(&oma.LWM2MCoreObjects{}, &ipso.IPSOSmartObjects{})

	return reg
}

func NewObjectRegistry(s ...ObjectSource) Registry {
	reg := &ObjectRegistry{}
	reg.sources = []ObjectSource{}

	for _, o := range s {
		reg.Register(o)
	}
	return reg
}

type ObjectRegistry struct {
	sources []ObjectSource
}

func (m *ObjectRegistry) GetModel(n LWM2MObjectType) ObjectDefinition {
	for _, s := range m.sources {
		if s != nil {
			o := s.GetObject(n)
			if o != nil {
				return o
			}
		}
	}
	return nil
}

func (m *ObjectRegistry) Register(s ObjectSource) {
	s.Initialize()
	m.sources = append(m.sources, s)
}

func (m *ObjectRegistry) GetMandatory() []ObjectDefinition {
	mandatory := []ObjectDefinition{}

	for _, s := range m.sources {
		objs := s.GetObjects()
		for _, o := range objs {
			if o.IsMandatory() {
				mandatory = append(mandatory, o)
			}
		}
	}
	return mandatory
}