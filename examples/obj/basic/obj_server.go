package basic

import "github.com/zubairhamed/lwm2m/core"

type Server struct {

}

func (o *Server) OnRead(r *core.ResourceModel, resourceId int) core.ResourceValue {
    return core.NewEmptyValue()
}
