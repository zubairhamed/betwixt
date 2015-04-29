package lwm2m

import (
    . "github.com/zubairhamed/goap"
    "net"
    "log"
    "errors"
    "bytes"
    "fmt"
    "strconv"
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
func (c *LWM2MClient) Register(name string) (string) {
    req := NewRequest(TYPE_CONFIRMABLE, POST, GenerateMessageId())

    req.SetStringPayload(BuildModelResourceStringPayload(c.enabledObjects))
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

func (c *LWM2MClient) EnableObject(t LWM2MObjectType) (error) {
    if c.enabledObjects[t] == nil {
        c.enabledObjects[t] = []*ObjectInstance{}

        return nil
    } else {
        return errors.New("Object already enabled")
    }
}

func (c *LWM2MClient) AddObjectInstance(instance *ObjectInstance) (error) {
    if instance != nil {
        o := c.GetObjectInstance(instance.TypeId, instance.Id)
        if o == nil {
            c.enabledObjects[instance.TypeId] = append(c.enabledObjects[instance.TypeId], instance)

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
        c.AddObjectInstance(o)
    }
}

func (c *LWM2MClient) GetObjectInstance(n LWM2MObjectType, instance int) (*ObjectInstance) {
    obj := c.enabledObjects[n]

    if obj != nil {
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

func (c *LWM2MClient) Start() {
    s := c.coapServer
    s.OnStartup(func(evt *Event) {
        if c.evtOnStartup != nil {
            c.evtOnStartup()
        }
    })


    /*
        ## Observe
        GET + Observe option
        /{Object ID}/{Object Instance ID}/{Resource ID}
        > 2.05 Content with Observe option
        < 4.04 Not Found, 4.05 Method Not Allowed
    */

    /*
        ## Discover
        GET + Accept: application/link- forma
        /{Object ID}/{Object Instance ID}/{Resource ID}
        > 2.05 Content
        < 4.04 Not Found, 4.01 Unauthorized, 4.05 Method Not Allowed
    */
    /*
        svr.NewRoute("{obj}/{inst}/{rsrc}", GET, func(req *CoapRequest) *CoapResponse {
            msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, GenerateMessageId())
            resp := NewResponseWithMessage(msg)

            return resp
        }).BindMediaTypes([]MediaType{ MEDIATYPE_APPLICATION_LINK_FORMAT })
    */

    /*
        ## Read
        GET
        /{Object ID}/{Object Instance ID}/{Resource ID}
        > 2.05 Content
        < 4.01 Unauthorized, 4.04 Not Found, 4.05 Method Not Allowed
    */
    s.NewRoute("{obj}/{inst}/{rsrc}", GET, handleReadResource)
    s.NewRoute("{obj}/{inst}", GET, handleReadInstance)



    c.coapServer.Start()
}

// Handlers

/*
    Read Resource
    GET     /obj/instance/resource
*/
func handleReadResource(req *CoapRequest) *CoapResponse {
    log.Println("Got READ Request")
    log.Println(req.GetAttribute("obj"), req.GetAttribute("inst"), req.GetAttribute("rsrc"))

    msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
    msg.SetStringPayload("")
    msg.Code = COAPCODE_205_CONTENT
    msg.Token = req.GetMessage().Token

    resp := NewResponseWithMessage(msg)

    objV, _ := strconv.Atoi(req.GetAttribute("obj"))

    /*
    CallEvent(c.evtOnRead, map[string] interface{}{
        "objectModel": c.registry.GetModel(objV),
    })
    */
    log.Println(objV)
    log.Println(resp)

    return resp
}

/*
    Read Instance
    GET     /obj/instance
*/
func handleReadInstance(req *CoapRequest) *CoapResponse {
    log.Println("Got READ Request")
    log.Println(req.GetAttribute("obj"), req.GetAttribute("inst"), req.GetAttribute("rsrc"))

    msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
    msg.SetStringPayload("")
    msg.Code = COAPCODE_205_CONTENT
    msg.Token = req.GetMessage().Token

    resp := NewResponseWithMessage(msg)

    objV, _ := strconv.Atoi(req.GetAttribute("obj"))

    /*
    CallEvent(c.evtOnRead, map[string] interface{}{
        "objectModel": c.registry.GetModel(objV),
    })
    */
    log.Println(objV)
    log.Println(resp)

    return resp
}

/*
    Write Replace
    PUT     /obj/instance/resource
*/
func handleWriteReplaceResource() {

}

func handleWriteReplaceInstance() {

}

func handleWriteOverwriteResource() {

}

func handleWriteOverwriteInstance() {

}


func handleExecuteResource() {

}

func handleCreateInstance() {

}

func handleDiscoverResources() {

}

func handleWriteResourceAttributes() {

}

func handleDeleteInstance() {

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
    var buf bytes.Buffer

    for k, v := range instances {
        if len(v) > 0 {
            for _, j := range v {
                buf.WriteString(fmt.Sprintf("</%d/%d>,", k, j.Id))
            }
        } else {
            buf.WriteString(fmt.Sprintf("</%d>,", k))
        }
    }
    return buf.String()
}
