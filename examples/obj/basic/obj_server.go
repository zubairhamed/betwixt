package basic

import (
    . "github.com/zubairhamed/go-lwm2m/api"
    "github.com/zubairhamed/go-lwm2m/core"
    "github.com/zubairhamed/goap"
    "github.com/zubairhamed/go-lwm2m/objects/oma"
)

type Server struct {
    Model       ObjectModel
    Data        *core.ObjectsData
}

func (o *Server) OnExecute(instanceId int, resourceId int) (goap.CoapCode) {
    return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *Server) OnCreate(instanceId int, resourceId int) (goap.CoapCode) {
    return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *Server) OnDelete(instanceId int) (goap.CoapCode) {
    return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *Server) OnRead(instanceId int, resourceId int) (ResponseValue, goap.CoapCode) {
    return core.NewEmptyValue(), goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (o *Server) OnWrite(instanceId int, resourceId int) (goap.CoapCode) {
    return goap.COAPCODE_405_METHOD_NOT_ALLOWED
}

func NewExampleServerObject(reg Registry) (*Server) {
    return &Server{
        Model: reg.GetModel(oma.OBJECT_LWM2M_SERVER),
    }
}


/*
[1]
Short Server ID                                 0       101
Lifetime                                        1       86400
Default Minimum Period                          2       300
Default Maximum Period                          3       6000
DisableTimeout                                  5       86400
Notification Storing When Disabled or Offline   6       True
Binding Preference                              7       U

[2]
Short Server ID                                 0       102
Lifetime                                        1       86400
Default Minimum Period                          2       60
Default Maximum Period                          3       6000
DisableTimeout                                  5       86400
Notification Storing When Disabled or Offline   6       False
Binding Preference                              7       UQ
*/