package basic

import (
    . "github.com/zubairhamed/lwm2m/api"
    "github.com/zubairhamed/lwm2m/core"
)

type Server struct {

}

func (o *Server) OnRead(r *ResourceModel, resourceId int) ResourceValue {
    return core.NewEmptyValue()
}
