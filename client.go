package lwm2m

import (
    . "github.com/zubairhamed/goap"
    "net"
    "log"
    "errors"
    "bytes"
    "fmt"
    "github.com/zubairhamed/lwm2m/core"
    "github.com/zubairhamed/lwm2m/objects"
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
        enabledObjects: make(map[core.LWM2MObjectType]*core.ObjectEnabler),
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
    registry            *objects.ObjectRegistry
    enabledObjects      map[core.LWM2MObjectType] *core.ObjectEnabler

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

func (c *LWM2MClient) GetEnabledObjects() (map[core.LWM2MObjectType] *core.ObjectEnabler) {
    return c.enabledObjects
}

func (c *LWM2MClient) GetRegistry() *objects.ObjectRegistry {
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

func (c *LWM2MClient) UseRegistry(reg *objects.ObjectRegistry) {
    c.registry = reg
}

func (c *LWM2MClient) EnableObject(t core.LWM2MObjectType, e core.ObjectHandler) (error) {
    if c.enabledObjects[t] == nil {

        en := &core.ObjectEnabler{
            Handler: e,
            Instances: []*core.ObjectInstance{},
        }
        c.enabledObjects[t] = en

        return nil
    } else {
        return errors.New("Object already enabled")
    }
}

func (c *LWM2MClient) AddObjectInstance(instance *core.ObjectInstance) (error) {
    if instance != nil {
        o := c.GetObjectInstance(instance.TypeId, instance.Id)
        if o == nil {
            c.enabledObjects[instance.TypeId].Instances = append(c.enabledObjects[instance.TypeId].Instances, instance)
            // c.enabledObjects[instance.TypeId] = append(c.enabledObjects[instance.TypeId], instance)

            return nil
        } else {
            return errors.New("Instance already exists. Use UpdateObjectInstance instead")
        }
    } else {
        return errors.New("Attempting to add a nil instance")
    }

}

func (c *LWM2MClient) AddObjectInstances (instances ... *core.ObjectInstance) {
    for _, o := range instances {
        c.AddObjectInstance(o)
    }
}

func (c *LWM2MClient) GetObjectEnabler(n core.LWM2MObjectType) (*core.ObjectEnabler) {
    return c.enabledObjects[n]
}

func (c *LWM2MClient) GetObjectInstance(n core.LWM2MObjectType, instance int) (*core.ObjectInstance) {
    enabler := c.enabledObjects[n]

    if enabler != nil {
        instances := enabler.Instances
        if len(instances) > 0 {
            for _, o := range instances {
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
    // Object ID
    // Object Instance ID
    // Resource ID

    attrObject := req.GetAttribute("obj")
    attrInstance := req.GetAttribute("inst")
    attrResource := req.GetAttribute("rsrc")

    objectId := req.GetAttributeAsInteger("obj")
    instanceId := req.GetAttributeAsInteger("inst")

    var resourceId = -1

    if attrResource != "" {
        resourceId, _ := strconv.Atoi(attrInstance)
    }

    t := core.LWM2MObjectType(objectId)
    enabler := c.GetObjectEnabler(t)

    if enabler != nil {
        if enabler.Handler != nil {
            msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
            msg.SetStringPayload("")
            msg.Code = COAPCODE_205_CONTENT
            msg.Token = req.GetMessage().Token

            val := enabler.Handler.OnRead(objectId, instanceId)
            msg.Payload = val.GetBytes()
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
    // if url has parameters
    // else

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

func (c *LWM2MClient)  handleDeleteRequest(req *CoapRequest) *CoapResponse {
    // DELETE  /0/0

    msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
    msg.SetStringPayload("")
    msg.Code = COAPCODE_205_CONTENT
    msg.Token = req.GetMessage().Token

    resp := NewResponseWithMessage(msg)

    return resp
}

func (c *LWM2MClient)  handlePostRequest(req *CoapRequest) *CoapResponse {
    // if has resource, execute
    // else create
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
func BuildModelResourceStringPayload(instances core.LWM2MObjectInstances) (string) {
    var buf bytes.Buffer

    for k, v := range instances {
        inst := v.Instances
        if len(inst) > 0 {
            for _, j := range inst {
                buf.WriteString(fmt.Sprintf("</%d/%d>,", k, j.Id))
            }
        } else {
            buf.WriteString(fmt.Sprintf("</%d>,", k))
        }
    }
    return buf.String()
}

/*
    /*
    // Returned payload



    enabler := c.GetObjectEnabler(t)
    if enabler != nil {
        if enabler.Handler != nil {
            val := enabler.Handler.OnRead(objectId, )

            // val := enabler.Handler.OnRead(objectId, resourceId)
            // msg.Payload = val.GetBytes()

            if obj != "" {
                model := c.registry.GetModel(t)

                if inst != "" {
                    instInt, _ = strconv.Atoi(inst)

                    if rsrc != "" {
                        // Querying resource, so call Handlers
                        rsrcInt, _ = strconv.Atoi(rsrc)
                        rsrcObj := model.GetResource(rsrcInt)

                        // Multiple Resources
                        ret := enabler.Handler.OnRead(rsrcObj, instInt)

                        if rsrcObj.Multiple {
                            core.TlvPayloadFromResource(ret.GetValue().(*core.MultipleResourceInstanceValue), rsrcObj, c.enabledObjects[t].GetObjectInstance(instInt).GetResource(rsrcInt))
                        } else {
                            // Single value resource
                            msg.Payload = NewPlainTextPayload(ret.GetStringValue())
                        }
                    } else {
                        // Instance of object
                        core.TlvPayloadFromObjectInstance((c.enabledObjects[t].GetObjectInstance(objInt)))
                    }
                } else {
                    // Object
                    core.TlvPayloadFromObjects(c.enabledObjects[t])
                }
            }
            resp := NewResponseWithMessage(msg)

            return resp
        }

    } else {
        log.Println("Enabler not found.")
    }
    */
*/