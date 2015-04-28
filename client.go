package lwm2m

import (
    . "github.com/zubairhamed/goap"
    "net"
    "log"
)

func NewLWM2MClient(local string, remote string) (*LWM2MClient) {
    localAddr, err := net.ResolveUDPAddr("udp", local)
    IfErrFatal(err)

    remoteAddr, err := net.ResolveUDPAddr("udp", remote)
    IfErrFatal(err)

    coapServer := NewCoapServer(localAddr, remoteAddr)

    return &LWM2MClient{
        coapServer: coapServer,
        instances: []*ObjectInstance{},
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
    registry            *ObjectRegistry
    instances           []*ObjectInstance

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
func (c *LWM2MClient) Start() {
    s := c.coapServer
    s.OnStartup(func(evt *Event) {
        // CallEvent(c.evtOnStartup, EmptyEventPayload())
    })

    c.coapServer.Start()
}

func (c *LWM2MClient) Register(name string) (string) {
    req := NewRequest(TYPE_CONFIRMABLE, POST, GenerateMessageId())

    req.SetStringPayload(BuildModelResourceStringPayload(c.instances))
    log.Println(BuildModelResourceStringPayload(c.instances))
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

func (c *LWM2MClient) Unregister() {

}

func (c *LWM2MClient) Update() {

}

func (c *LWM2MClient) AddResource() {

}

func (c *LWM2MClient) AddObject() {

}

func (c *LWM2MClient) UseRegistry(reg *ObjectRegistry) {
    c.registry = reg
}

func (c *LWM2MClient) EnableObject(t LWM2MObjectType) {
    c.instances = append(c.instances, NewObjectInstance(t))
}

func (c *LWM2MClient) AddObjectInstances (instances ... *ObjectInstance) {
    for _, o := range instances {
        inst := c.GetObjectInstance(o.TypeId, o.Id)

        if inst == nil {
            c.instances = append(c.instances, o)
        } else {
            // TODO: Throw Error
        }

    }
}

func (c *LWM2MClient) GetObjectInstance(n LWM2MObjectType, inst int) (*ObjectInstance) {
    for _, o := range c.instances {
        if o.Id == inst && o.TypeId == n {
            return o
        }
    }
    return nil
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

// Functions
func BuildModelResourceStringPayload(inst []*ObjectInstance) string {
    return ""
}