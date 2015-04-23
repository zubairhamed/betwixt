package lwm2m

import . "github.com/zubairhamed/goap"

func NewLWM2MServer() (*LWM2MServer) {
    repo := NewModelRepository()

    repo.Register(&LWM2MCoreObjects{})
    repo.Register(&IPSOSmartObjects{})

    s := &LWM2MServer{
        modelRepository: repo,
    }
    return s
}

type LWM2MServer struct {
    coapServer       *CoapServer
    modelRepository  *ModelRepository
}

func (s *LWM2MServer) GetModel(n int) *ObjectModel {
    return s.modelRepository.GetModel( n)
}