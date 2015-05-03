package lwm2m

import (
    . "github.com/zubairhamed/goap"
    . "github.com/zubairhamed/lwm2m/objects"
    . "github.com/zubairhamed/lwm2m/core"
    "github.com/zubairhamed/lwm2m/objects/oma"
    "github.com/zubairhamed/lwm2m/objects/ipso"
)

func NewLWM2MServer() (*LWM2MServer) {
    reg := NewObjectRegistry()
    reg.Register(&oma.LWM2MCoreObjects{})
    reg.Register(&ipso.IPSOSmartObjects{})

    s := &LWM2MServer{
        registry :reg,
    }
    return s
}

type LWM2MServer struct {
    coapServer      *CoapServer
    registry        *ObjectRegistry
}

func (s *LWM2MServer) GetModel(n LWM2MObjectType) *ObjectModel {
    return s.registry.GetModel(n)
}