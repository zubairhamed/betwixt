package registry

import (
    "github.com/zubairhamed/lwm2m/objects/ipso"
    "github.com/zubairhamed/lwm2m/objects/oma"
    . "github.com/zubairhamed/lwm2m/core"
    . "github.com/zubairhamed/lwm2m/api"
)

func NewDefaultObjectRegistry() (Registry) {
    reg := NewObjectRegistry(&oma.LWM2MCoreObjects{}, &ipso.IPSOSmartObjects{})

    return reg
}

func NewObjectRegistry(s ...ModelSource) (Registry) {
    reg := &ObjectRegistry{}
    reg.sources = []ModelSource{}

    for _, o := range s {
        reg.Register(o)
    }
    return reg
}

type ObjectRegistry struct {
    sources     []ModelSource
}

func (m *ObjectRegistry) CreateObjectInstance(t LWM2MObjectType, n int) (ObjectInstance) {
    o := m.GetModel(t)
    if o != nil {
        obj := NewObjectInstance(n, t)

        return obj
    }
    return nil
}

func (m *ObjectRegistry) GetModel(n LWM2MObjectType) ObjectModel {
    for _, s := range m.sources {
        if s != nil {
            o := s.Get(n)
            if o != nil {
                return o
            }
        }
    }
    return nil
}

func (m *ObjectRegistry) Register(s ModelSource) {
    s.Initialize()
    m.sources = append(m.sources, s)
}

func (m *ObjectRegistry) CreateHandler(t LWM2MObjectType) {

}
