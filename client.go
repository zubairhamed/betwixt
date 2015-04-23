package lwm2m

import (
    "net"
    "log"
    . "github.com/zubairhamed/goap"
    "bytes"
    "fmt"
    "strconv"
)

func NewLwm2mClient(local string, remote string) (*LWM2MClient) {
    localAddr, err := net.ResolveUDPAddr("udp", local)
    IfErrFatal(err)

    remoteAddr, err := net.ResolveUDPAddr("udp", remote)
    IfErrFatal(err)

    coapServer := NewCoapServer(localAddr, remoteAddr)

    repo := NewModelRepository()
    repo.Register(&LWM2MCoreObjects{})
    repo.Register(&IPSOSmartObjects{})

    return &LWM2MClient{
        coapServer: coapServer,
        resources:  []*LWM2MResource{},
        repository: repo,
    }
}

type LWM2MClient struct {
    coapServer          *CoapServer
    resources           []*LWM2MResource
    repository          *ModelRepository

    evtOnStartup        EventHandler
    evtOnRead           EventHandler
    evtOnWrite          EventHandler
    evtOnExecute        EventHandler
    evtOnRegistered     EventHandler
    evtOnUnregistered   EventHandler
}

func (c *LWM2MClient) Start() {
    svr := c.coapServer
    svr.OnStartup(func(evt *Event) {
        CallEvent(c.evtOnStartup, EmptyEventPayload())
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
    svr.NewRoute("{obj}/{inst}/{rsrc}", GET, func(req *CoapRequest) *CoapResponse {
        log.Println("Got READ Request")
        log.Println(req.GetAttribute("obj"), req.GetAttribute("inst"), req.GetAttribute("rsrc"))

        msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
        msg.SetStringPayload("")
        msg.Code = COAPCODE_205_CONTENT
        msg.Token = req.GetMessage().Token

        resp := NewResponseWithMessage(msg)

        objV, _ := strconv.Atoi(req.GetAttribute("obj"))

        CallEvent(c.evtOnRead, map[string] interface{}{
            "objectModel": c.repository.GetModel(objV),
        })

        return resp
    })

    svr.NewRoute("{obj}/{inst}", GET, func(req *CoapRequest) *CoapResponse {
        log.Println("Got READ Request")
        log.Println(req.GetAttribute("obj"), req.GetAttribute("inst"), req.GetAttribute("rsrc"))

        msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
        msg.SetStringPayload("")
        msg.Code = COAPCODE_205_CONTENT
        msg.Token = req.GetMessage().Token

        resp := NewResponseWithMessage(msg)

        return resp
    })


    /*
        ## Write
        PUT / POST
        /{Object ID}/{Object Instance ID}/{Resource ID}
        > 2.04 Changed
        < 4.00 Bad Request, 4.04 Not Found, 4.01 Unauthorized, 4.05 Method Not Allowed

        ## Write Attributes
        PUT
        /{Object ID}/{Object Instance ID}/{Resource ID}?pmin={minimum period}&pmax={maximum period}&gt={greater than}&lt={less than}&st={step}&cancel
        > 2.04 Changed
        < 4.00 Bad Request, 4.04 Not Found, 4.01 Unauthorized, 4.05 Method Not Allowed

    */
    svr.NewRoute("{obj}/{inst}/{rsrc}", PUT, func(req *CoapRequest) *CoapResponse {
        msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
        msg.Token = req.GetMessage().Token

        resp := NewResponseWithMessage(msg)
        return resp
    })

    svr.NewRoute("{obj}/{inst}/{rsrc}", POST, func(req *CoapRequest) *CoapResponse {
        msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
        msg.Token = req.GetMessage().Token

        resp := NewResponseWithMessage(msg)
        return resp
    })


    /*
        ## Execute
        POST
        /{Object ID}/{Object Instance ID}/{Resource ID}
        > 2.04 Changed
        < 4.00 Bad Request, 4.01 Unauthorized, 4.04 Not Found, 4.05 Method Not Allowed
    */


    /*
        ## Create
        POST
        /{Object ID}/{Object Instance ID}
        > 2.01 Created
        < 4.00 Bad Request, 4.01 Unauthorized, 4.04 Not Found, 4.05 Method Not Allowed
    */
    svr.NewRoute("{obj}/{inst}/", POST, func(req *CoapRequest) *CoapResponse {
        msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
        msg.Token = req.GetMessage().Token

        resp := NewResponseWithMessage(msg)
        return resp
    })

    /*
        ## Delete
        DELETE
        /{Object ID}/{Object Instance ID}
        > 2.02 Deleted
        < 4.01 Unauthorized, 4.04 Not Found, 4.05 Method Not Allowed
    */
    svr.NewRoute("{obj}/{inst}/{DELETE}", POST, func(req *CoapRequest) *CoapResponse {
        msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
        msg.Token = req.GetMessage().Token

        resp := NewResponseWithMessage(msg)
        return resp
    })

    /*
        ## Cancel Observe
        Reset message
    */

    c.coapServer.Start()
}

func (c *LWM2MClient) OnStartup (eh EventHandler) {
    c.evtOnStartup = eh
}

func (c *LWM2MClient) OnRegistered (eh EventHandler) {
    c.evtOnRegistered = eh
}

func (c *LWM2MClient) OnUnregistered (eh EventHandler) {
    c.evtOnUnregistered = eh
}

func (c *LWM2MClient) OnRead (eh EventHandler) {
    c.evtOnRead = eh
}

func (c *LWM2MClient) OnWrite (eh EventHandler) {
    c.evtOnWrite = eh
}

func (c *LWM2MClient) OnExecute (eh EventHandler) {
    c.evtOnExecute = eh
}

func BuildModelResourceStringPayload (resources []*LWM2MResource) string {
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
}

func (c *LWM2MClient) Register(name string) (string) {
    req := NewRequest(TYPE_CONFIRMABLE, POST, GenerateMessageId())

    req.SetStringPayload(BuildModelResourceStringPayload(c.resources))
    log.Println(BuildModelResourceStringPayload(c.resources))
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

    CallEvent(c.evtOnRegistered, EmptyEventPayload())

    return path
}

func (c *LWM2MClient) Unregister() {
    CallEvent(c.evtOnUnregistered, EmptyEventPayload())
}

func (c *LWM2MClient) Update() {

}

func (c *LWM2MClient) AddResource(resources ... int) {
    var lwr *LWM2MResource
    if len(resources) > 1 {
        lwr = NewLWM2MResource(c.repository.GetModel(resources[0]))
    } else {
        lwr = NewLWM2MResource(c.repository.GetModel(resources[0]), resources[1:]...)
    }
    c.resources = append(c.resources, lwr)
}
