package lwm2m

import (
    . "github.com/zubairhamed/goap"
    "net"
    "log"
    "errors"
    "bytes"
    "fmt"
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

    setupRoutes(s)

    c.coapServer.Start()
}

func setupRoutes(s *CoapServer) {
    s.NewRoute("{obj}/{inst}/{rsrc}", GET, handleGetRequest)
    s.NewRoute("{obj}/{inst}", GET, handleGetRequest)
    s.NewRoute("{obj}", GET, handleGetRequest)

    s.NewRoute("{obj}/{inst}/{rsrc}", PUT, handlePutRequest)
    s.NewRoute("{obj}/{inst}", PUT, handlePutRequest)

    s.NewRoute("{obj}/{inst}", DELETE, handleDeleteRequest)

    s.NewRoute("{obj}/{inst}/{rsrc}", POST, handlePostRequest)
    s.NewRoute("{obj}/{inst}", POST, handlePostRequest)
}

func handleGetRequest(req *CoapRequest) *CoapResponse {
    log.Println(req)

    /*
READ        GET     /0/0
READ        GET     /0/0/0
DISCOVER    GET     /0      +Accept: application/link format
DISCOVER    GET     /0/0    +Accept: application/link format
DISCOVER    GET     /0/0/0  +Accept: application/link format
OBSERVE     GET     /0      +Observe
OBSERVE     GET     /0/0    +Observe
OBSERVE     GET     /0/0/0  +Observe
    */

    msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
    msg.SetStringPayload("")
    msg.Code = COAPCODE_205_CONTENT
    msg.Token = req.GetMessage().Token

    resp := NewResponseWithMessage(msg)

    return resp
}

func handlePutRequest(req *CoapRequest) *CoapResponse {
    log.Println(req)

    /*
WRITE       PUT     /0/0
WRITE       PUT     /0/0/0
WRITE ATTR  PUT     /0/0/0  +?pmin={minimum period}&pmax={maximum period}&gt={greater than}&lt={less than}&st={step}&cancel
    */

    msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
    msg.SetStringPayload("")
    msg.Code = COAPCODE_205_CONTENT
    msg.Token = req.GetMessage().Token

    resp := NewResponseWithMessage(msg)

    return resp
}

func handleDeleteRequest(req *CoapRequest) *CoapResponse {
    log.Println(req)

    // DELETE  /0/0

    msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
    msg.SetStringPayload("")
    msg.Code = COAPCODE_205_CONTENT
    msg.Token = req.GetMessage().Token

    resp := NewResponseWithMessage(msg)

    return resp
}

func handlePostRequest(req *CoapRequest) *CoapResponse {
    log.Println(req)

    /*
EXECUTE     POST    /0/0/0
CREATE      POST    /0/<id>
    */

    msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
    msg.SetStringPayload("")
    msg.Code = COAPCODE_205_CONTENT
    msg.Token = req.GetMessage().Token

    resp := NewResponseWithMessage(msg)

    return resp
}

/*
GET     /0/0/0
GET     /0/0
GET     /0

PUT     /0/0/0
PUT     /0/0

DELETE  /0/0

POST    /0/0/0
POST    /0/0

-----
READ        GET     /0/0
READ        GET     /0/0/0
DISCOVER    GET     /0      +Accept: application/link format
DISCOVER    GET     /0/0    +Accept: application/link format
DISCOVER    GET     /0/0/0  +Accept: application/link format
OBSERVE     GET     /0      +Observe
OBSERVE     GET     /0/0    +Observe
OBSERVE     GET     /0/0/0  +Observe

WRITE       PUT     /0/0
WRITE       PUT     /0/0/0
WRITE ATTR  PUT     /0/0/0  +?pmin={minimum period}&pmax={maximum period}&gt={greater than}&lt={less than}&st={step}&cancel

DELETE      DELETE  /0/0

EXECUTE     POST    /0/0/0
CREATE      POST    /0/<id>


@@@@@@@@@@ INCOMING @@@@@@@@@@
## READ
GET, CON
/0/0
/0/0/0
- handleReadResource
- handleReadInstance

## DISCOVER
GET, CON
/0
/0/0
/0/0/0
Accept: application/link format
- handleDiscoverResources

## WRITE
PUT, CON
/0/0
/0/0/0
Content-Format: 1542
- handleWriteResource
- handleWriteInstance

## WRITE ATTRIBUTES
PUT, CON
/0/0/0
?pmin={minimum period}&pmax={maximum period}&gt={greater than}&lt={less than}&st={step}&cancel
handleWriteResourceAttributes

## DELETE
DELETE, CON
/0/0
- handleDeleteInstance

## EXECUTE
POST, CON
/0/0/0
- handleExecuteResource

## OBSERVE
GET, CON
/0
/0/0
/0/0/0
OPTION OBSERVE = 0
- handleObserveResources
// Related parameters for “Observe” operation are described in 5.3.4

## CANCEL OBSERVATION
(via Write Attribute with Cancel Param)
(Respond to a Notify with a Cancel Observation)

## CREATE
POST, CON
/0/<NEWID>
Content-Format: 1542
- handleCreateInstance

@@@@@@@@@@ OUTGOING @@@@@@@@@@
Register
CON, POST, /rd?ep=DEVKIT&lt=60&b=U

Update
CON, PUT, /rd/<returnedId>&lt=xxx&b=U

Notify


=======

Values:
Plain Text, Opaque, JSON, TLV
*/

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
