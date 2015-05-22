package lwm2m

import (
    . "github.com/zubairhamed/goap"
    "github.com/zubairhamed/go-lwm2m/objects/oma"
    "github.com/zubairhamed/go-lwm2m/objects/ipso"
    "github.com/zubairhamed/go-lwm2m/registry"
    . "github.com/zubairhamed/go-lwm2m/api"
)

func NewLWM2MServer() (*LWM2MServer) {
    reg := registry.NewObjectRegistry()
    reg.Register(&oma.LWM2MCoreObjects{})
    reg.Register(&ipso.IPSOSmartObjects{})

    s := &LWM2MServer{
        registry :reg,
    }
    return s
}

type LWM2MServer struct {
    coapServer      *CoapServer
    registry        Registry
}

func (s *LWM2MServer) GetModel(n LWM2MObjectType) ObjectModel {
    return s.registry.GetModel(n)
}