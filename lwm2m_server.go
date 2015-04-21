package lwm2m

import . "github.com/zubairhamed/goap"

func NewLWM2MServer() (*LWM2MServer) {
    s := &LWM2MServer{
        modelRepository: NewModelsRepository(),
    }

    return s
}

type LWM2MServer struct {
    coapServer       *CoapServer
    modelRepository *ModelsRepository
}