package lwm2m

import . "github.com/zubairhamed/goap"

func NewLWM2MServer() (*LWM2MServer) {
    reg := NewObjectRegistry()

    reg.Register(&LWM2MCoreObjects{})
    reg.Register(&IPSOSmartObjects{})

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