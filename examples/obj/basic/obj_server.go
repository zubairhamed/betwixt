package basic

import (
    . "github.com/zubairhamed/lwm2m/api"
    "github.com/zubairhamed/lwm2m/core"
)

type Server struct {

}

func (o *Server) OnExecute(instanceId int, resourceId int) (bool, int) {
    return true, 0
}

func (o *Server) OnCreate(instanceId int, resourceId int) (bool, int) {
    return true, 0
}

func (o *Server) OnDelete(instanceId int) (bool) {
    return true
}

func (o *Server) OnRead(instanceId int, resourceId int) (ResponseValue) {
    return core.NewEmptyValue()
}

func (o *Server) OnWrite(instanceId int, resourceId int) (bool) {
    return true
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