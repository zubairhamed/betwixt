package lwm2m

import (
    . "github.com/zubairhamed/goap"
    "net"
    "log"
    "errors"
    "github.com/zubairhamed/lwm2m/core"
    . "github.com/zubairhamed/lwm2m/api"
)

func NewLWM2MClient(local string, remote string) (*LWM2MClient) {
    localAddr, err := net.ResolveUDPAddr("udp", local)
    IfErrFatal(err)

    remoteAddr, err := net.ResolveUDPAddr("udp", remote)
    IfErrFatal(err)

    coapServer := NewCoapServer(localAddr, remoteAddr)

    return &LWM2MClient{
        coapServer: coapServer,
        enabledObjects: make(map[LWM2MObjectType]ObjectEnabler),
    }
}

type FnOnStartup func()
type FnOnRead func()
type FnOnWrite func()
type FnOnExecute func()
type FnOnRegistered func(string)
type FnOnUnregistered func()
type FnOnError func()

type LWM2MClient struct {
    coapServer          *CoapServer
    registry            Registry
    enabledObjects      map[LWM2MObjectType] ObjectEnabler

    // Events
    evtOnStartup        FnOnStartup
    evtOnRead           FnOnRead
    evtOnWrite          FnOnWrite
    evtOnExecute        FnOnExecute
    evtOnRegistered     FnOnRegistered
    evtOnUnregistered   FnOnUnregistered
    evtOnError          FnOnError
}

// Operations
func (c *LWM2MClient) Register(name string) (string) {
    req := NewRequest(TYPE_CONFIRMABLE, POST, GenerateMessageId())

    req.SetStringPayload(core.BuildModelResourceStringPayload(c.enabledObjects))
    req.SetRequestURI("rd")
    req.SetUriQuery("ep", name)
    resp, err := c.coapServer.Send(req)

    path := ""
    if err != nil {
        log.Println(err)
    } else {
        PrintMessage(resp.GetMessage())

        path = resp.GetMessage().GetLocationPath()
    }

//    CallEvent(c.evtOnRegistered, EmptyEventPayload())

    return path
}

func (c *LWM2MClient) GetEnabledObjects() (map[LWM2MObjectType] ObjectEnabler) {
    return c.enabledObjects
}

func (c *LWM2MClient) GetRegistry() Registry {
    return c.registry
}

func (c *LWM2MClient) Unregister() {

}

func (c *LWM2MClient) Update() {

}

func (c *LWM2MClient) AddResource() {

}

func (c *LWM2MClient) AddObject() {

}

func (c *LWM2MClient) UseRegistry(reg Registry) {
    c.registry = reg
}

func (c *LWM2MClient) EnableObject(t LWM2MObjectType, e RequestHandler) (error) {
    if c.enabledObjects[t] == nil {

        en := &core.DefaultObjectEnabler{
            Handler: e,
            Instances: []ObjectInstance{},
        }
        c.enabledObjects[t] = en

        return nil
    } else {
        return errors.New("Object already enabled")
    }
}

func (c *LWM2MClient) AddObjectInstance(instance ObjectInstance) (error) {
    if instance != nil {
        o := c.GetObjectInstance(instance.GetTypeId(), instance.GetId())
        if o == nil {
            c.enabledObjects[instance.GetTypeId()].SetObjectInstances(append(c.enabledObjects[instance.GetTypeId()].GetObjectInstances(), instance))

            return nil
        } else {
            return errors.New("Instance already exists. Use UpdateObjectInstance instead")
        }
    } else {
        return errors.New("Attempting to add a nil instance")
    }

}

func (c *LWM2MClient) AddObjectInstances (instances ... ObjectInstance) {
    for _, o := range instances {
        c.AddObjectInstance(o)
    }
}

func (c *LWM2MClient) GetObjectEnabler(n LWM2MObjectType) (ObjectEnabler) {
    return c.enabledObjects[n]
}

func (c *LWM2MClient) GetObjectInstance(n LWM2MObjectType, instance int) (ObjectInstance) {
    enabler := c.enabledObjects[n]

    if enabler != nil {
        instances := enabler.GetObjectInstances()
        if len(instances) > 0 {
            for _, o := range instances {
                if o.GetId() == instance && o.GetTypeId() == n {
                    return o
                }
            }
        }
    }
    return nil
}

func (c *LWM2MClient) Start() {
    s := c.coapServer
    s.OnStartup(func(evt *Event) {
        if c.evtOnStartup != nil {
            c.evtOnStartup()
        }
    })

    s.NewRoute("{obj}/{inst}/{rsrc}", GET, c.handleGetRequest)
    s.NewRoute("{obj}/{inst}", GET, c.handleGetRequest)
    s.NewRoute("{obj}", GET, c.handleGetRequest)

    s.NewRoute("{obj}/{inst}/{rsrc}", PUT, c.handlePutRequest)
    s.NewRoute("{obj}/{inst}", PUT, c.handlePutRequest)

    s.NewRoute("{obj}/{inst}", DELETE, c.handleDeleteRequest)

    s.NewRoute("{obj}/{inst}/{rsrc}", POST, c.handlePostRequest)
    s.NewRoute("{obj}/{inst}", POST, c.handlePostRequest)

    c.coapServer.Start()
}


func (c *LWM2MClient) handleGetRequest(req *CoapRequest) *CoapResponse {
    attrResource := req.GetAttribute("rsrc")
    objectId := req.GetAttributeAsInt("obj")
    instanceId := req.GetAttributeAsInt("inst")

    var resourceId = -1

    if attrResource != "" {
        resourceId = req.GetAttributeAsInt("rsrc")
    }

    t := LWM2MObjectType(objectId)
    enabler := c.GetObjectEnabler(t)


    if enabler != nil {
        if enabler.GetHandler() != nil {
            msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
            msg.SetStringPayload("")
            msg.Code = COAPCODE_205_CONTENT
            msg.Token = req.GetMessage().Token

            val := enabler.OnRead(instanceId, resourceId)
            msg.Payload = NewBytesPayload(val.GetBytes())

            return NewResponseWithMessage(msg)
        }
    }
    return nil
}

func (c *LWM2MClient)  handleDiscoverRequest() {

}

func (c *LWM2MClient)  handleObserveRequest() {

}

func (c *LWM2MClient)  handleReadRequest() {

}

func (c *LWM2MClient)  handlePutRequest(req *CoapRequest) *CoapResponse {
    msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
    msg.SetStringPayload("")
    msg.Code = COAPCODE_205_CONTENT
    msg.Token = req.GetMessage().Token

    resp := NewResponseWithMessage(msg)

    return resp
}

func (c *LWM2MClient)  handleDeleteRequest(req *CoapRequest) *CoapResponse {
    msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
    msg.SetStringPayload("")
    msg.Code = COAPCODE_205_CONTENT
    msg.Token = req.GetMessage().Token

    resp := NewResponseWithMessage(msg)

    return resp
}

func (c *LWM2MClient)  handlePostRequest(req *CoapRequest) *CoapResponse {
    msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
    msg.SetStringPayload("")
    msg.Code = COAPCODE_205_CONTENT
    msg.Token = req.GetMessage().Token

    resp := NewResponseWithMessage(msg)

    return resp
}

// Events
func (c *LWM2MClient) OnStartup(fn FnOnStartup) {
    c.evtOnStartup = fn
}

func (c *LWM2MClient) OnRead(fn FnOnRead) {
    c.evtOnRead = fn
}

func (c *LWM2MClient) OnWrite(fn FnOnWrite) {
    c.evtOnWrite = fn
}

func (c *LWM2MClient) OnExecute(fn FnOnExecute) {
    c.evtOnExecute = fn
}

func (c *LWM2MClient) OnRegistered(fn FnOnRegistered) {
    c.evtOnRegistered = fn
}

func (c *LWM2MClient) OnUnregistered(fn FnOnUnregistered) {
    c.evtOnUnregistered = fn
}

func (c *LWM2MClient) OnError (fn FnOnError) {
    c.evtOnError = fn
}

