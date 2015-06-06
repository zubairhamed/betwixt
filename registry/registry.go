package registry

import (
	. "github.com/zubairhamed/betwixt/api"
	"github.com/zubairhamed/betwixt/objects/ipso"
	"github.com/zubairhamed/betwixt/objects/oma"
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

func (m *ObjectRegistry) GetModel(n LWM2MObjectType) ObjectModel {
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

func (m *ObjectRegistry) GetMandatory() []ObjectModel {
	mandatory := []ObjectModel{}

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