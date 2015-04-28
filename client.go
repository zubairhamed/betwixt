package lwm2m

import (
    . "github.com/zubairhamed/goap"
    "net"
    "log"
    "errors"
)

func NewLWM2MClient(local string, remote string) (*LWM2MClient) {
    localAddr, err := net.ResolveUDPAddr("udp", local)
    IfErrFatal(err)

    remoteAddr, err := net.ResolveUDPAddr("udp", remote)
    IfErrFatal(err)

    coapServer := NewCoapServer(localAddr, remoteAddr)

    return &LWM2MClient{
        coapServer: coapServer,
        enabledObjects: make(LWM2MObjectInstances),
    }
}

type FnOnStartup func()
type FnOnRead func()
type FnOnWrite func()
type FnOnExecute func()
type FnOnRegistered func(string)
type FnOnUnregistered func()
type FnOnError func()

type LWM2MObjectInstances map[LWM2MObjectType][]*ObjectInstance

type LWM2MClient struct {
    coapServer          *CoapServer
    registry            *ObjectRegistry
    // instances           []*ObjectInstance
    enabledObjects      LWM2MObjectInstances

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
        if c.evtOnStartup != nil {
            c.evtOnStartup()
        }
    })
    c.coapServer.Start()
}

func (c *LWM2MClient) Register(name string) (string) {
    req := NewRequest(TYPE_CONFIRMABLE, POST, GenerateMessageId())

    req.SetStringPayload(BuildModelResourceStringPayload(c.enabledObjects))
    log.Println(BuildModelResourceStringPayload(c.enabledObjects))
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
    if c.enabledObjects[t] != nil {
        c.enabledObjects[t] = append(c.enabledObjects[t], NewObjectInstance(t))
    }
}

func (c *LWM2MClient) AddObjectInstance(instance *ObjectInstance) (error) {
    if instance != nil {
        o := c.GetObjectInstance(instance.TypeId, instance.Id)
        if o == nil {
            c.enabledObjects[instance.TypeId] = append(c.enabledObjects[instance.TypeId], o)

            return nil
        } else {
            return errors.New("Instance already exists. Use UpdateObjectInstance instead")
        }
    } else {
        return errors.New("Attempting to add a nil instance")
    }

}

func (c *LWM2MClient) AddObjectInstances (instances ... *ObjectInstance) {
    for _, o := range instances {
        log.Println("Add Object Instance")
        log.Println(o)
        c.AddObjectInstance(o)
    }
}

func (c *LWM2MClient) GetObjectInstance(n LWM2MObjectType, instance int) (*ObjectInstance) {
    obj := c.enabledObjects[n]
    log.Println("!!!!")
    log.Println(c.enabledObjects)

    if obj != nil {
        log.Println(len(obj))
        log.Println(obj)
        if len(obj) > 0 {
            for _, o := range obj {
                if o.Id == instance && o.TypeId == n {
                    return o
                }
            }
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
func BuildModelResourceStringPayload(instances LWM2MObjectInstances) (string) {
    // var buf bytes.Buffer

    for k := range instances {
        log.Println(instances[k])
    }
    /*
    for _, instance := range instances {
        typeId := instance.TypeId
        if len(instance.Resources) > 0 {
            for _, res := range instance.Resources {
                buf.WriteString(fmt.Sprintf("</%d/%d>,", typeId, i))
            }
        }
    }
    */

    return ""
}


/*
    var buf bytes.Buffer

    for _, r := range resources {
        log.Println (r.model)
        resourceId := r.model.Id
        if len(r.instances) > 0 {
            for _, i := range r.instances {
                buf.WriteString(fmt.Sprintf("</%d/%d>,", resourceId, i))
            }
        } else {
            buf.WriteString(fmt.Sprintf("</%d>,", resourceId))
        }
    }

    return buf.String()
*/