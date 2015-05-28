package lwm2m

import (
	"errors"
	. "github.com/zubairhamed/go-lwm2m/api"
	"github.com/zubairhamed/go-lwm2m/core"
	. "github.com/zubairhamed/goap"
	"log"
	"net"
)

func NewLWM2MClient(local string, remote string) *DefaultClient {
	localAddr, err := net.ResolveUDPAddr("udp", local)
	IfErrFatal(err)

	remoteAddr, err := net.ResolveUDPAddr("udp", remote)
	IfErrFatal(err)

	coapServer := NewCoapServer(localAddr, remoteAddr)

	return &DefaultClient{
		coapServer:     coapServer,
		enabledObjects: make(map[LWM2MObjectType]ObjectEnabler),
	}
}

type DefaultClient struct {
	coapServer     *CoapServer
	registry       Registry
	enabledObjects map[LWM2MObjectType]ObjectEnabler
	path           string

	// Events
	evtOnStartup      FnOnStartup
	evtOnRead         FnOnRead
	evtOnWrite        FnOnWrite
	evtOnExecute      FnOnExecute
	evtOnRegistered   FnOnRegistered
	evtOnDeregistered FnOnDeregistered
	evtOnError        FnOnError
}

// Operations
func (c *DefaultClient) Register(name string) string {
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

	c.path = path

	return path
}

func (c *DefaultClient) GetEnabledObjects() map[LWM2MObjectType]ObjectEnabler {
	return c.enabledObjects
}

func (c *DefaultClient) GetRegistry() Registry {
	return c.registry
}

func (c *DefaultClient) Deregister() {
	req := NewRequest(TYPE_CONFIRMABLE, DELETE, GenerateMessageId())

	req.SetRequestURI(c.path)
	resp, err := c.coapServer.Send(req)

	if err != nil {
		log.Println(err)
	} else {
		PrintMessage(resp.GetMessage())
	}
}

func (c *DefaultClient) Update() {

}

func (c *DefaultClient) AddResource() {

}

func (c *DefaultClient) AddObject() {

}

func (c *DefaultClient) UseRegistry(reg Registry) {
	c.registry = reg
}

func (c *DefaultClient) EnableObject(t LWM2MObjectType, e RequestHandler) error {
	if c.enabledObjects[t] == nil {

		if c.registry == nil {
			return errors.New("No registry found/set")
		}

		model := c.registry.GetModel(t)
		en := &core.DefaultObjectEnabler{
			Handler:   e,
			Instances: []ObjectInstance{},
			Model: model,
		}
		c.enabledObjects[t] = en

		return nil
	} else {
		return errors.New("Object already enabled")
	}
}

func (c *DefaultClient) AddObjectInstance(instance ObjectInstance) error {
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

func (c *DefaultClient) AddObjectInstances(instances ...ObjectInstance) {
	for _, o := range instances {
		c.AddObjectInstance(o)
	}
}

func (c *DefaultClient) GetObjectEnabler(n LWM2MObjectType) ObjectEnabler {
	return c.enabledObjects[n]
}

func (c *DefaultClient) GetObjectInstance(n LWM2MObjectType, instance int) ObjectInstance {
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

func (c *DefaultClient) Start() {
	s := c.coapServer
	s.OnStartup(func(evt *Event) {
		if c.evtOnStartup != nil {
			c.evtOnStartup()
		}
	})

	s.NewRoute("{obj}/{inst}/{rsrc}", GET, c.handleReadRequest)
	s.NewRoute("{obj}/{inst}", GET, c.handleReadRequest)
	s.NewRoute("{obj}", GET, c.handleReadRequest)

	s.NewRoute("{obj}/{inst}/{rsrc}", PUT, c.handleWriteRequest)
	s.NewRoute("{obj}/{inst}", PUT, c.handleWriteRequest)

	s.NewRoute("{obj}/{inst}", DELETE, c.handleDeleteRequest)

	s.NewRoute("{obj}/{inst}/{rsrc}", POST, c.handleExecuteRequest)
	s.NewRoute("{obj}/{inst}", POST, c.handleCreateRequest)

	c.coapServer.Start()
}

func (c *DefaultClient) handleCreateRequest(req *CoapRequest) *CoapResponse {
	attrResource := req.GetAttribute("rsrc")
	objectId := req.GetAttributeAsInt("obj")
	instanceId := req.GetAttributeAsInt("inst")

	var resourceId = -1

	if attrResource != "" {
		resourceId = req.GetAttributeAsInt("rsrc")
	}

	t := LWM2MObjectType(objectId)
	enabler := c.GetObjectEnabler(t)

	msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
	msg.Token = req.GetMessage().Token
	msg.Payload = NewEmptyPayload()

	if enabler != nil && enabler.GetHandler() != nil {
        lwReq := core.NewDefaultRequest(req, OPERATIONTYPE_CREATE)
		msg.Code = enabler.OnCreate(instanceId, resourceId, lwReq)
	} else {
		msg.Code = COAPCODE_405_METHOD_NOT_ALLOWED
	}
	return NewResponseWithMessage(msg)
}

func (c *DefaultClient) handleReadRequest(req *CoapRequest) *CoapResponse {
	attrResource := req.GetAttribute("rsrc")
	objectId := req.GetAttributeAsInt("obj")
	instanceId := req.GetAttributeAsInt("inst")

	var resourceId = -1

	if attrResource != "" {
		resourceId = req.GetAttributeAsInt("rsrc")
	}

	t := LWM2MObjectType(objectId)
	enabler := c.GetObjectEnabler(t)

	msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
	msg.Token = req.GetMessage().Token

	if enabler != nil && enabler.GetHandler() != nil {
		model := enabler.GetModel()
		resource := model.GetResource(resourceId)

		if resource == nil {
			msg.Code = COAPCODE_404_NOT_FOUND
		}
		if !core.IsReadableResource(resource) {
			msg.Code = COAPCODE_405_METHOD_NOT_ALLOWED
		} else {
            lwReq := core.NewDefaultRequest(req, OPERATIONTYPE_READ)
			val, _ := enabler.OnRead(instanceId, resourceId, lwReq)
			msg.Code = COAPCODE_205_CONTENT
			msg.Payload = NewBytesPayload(val.GetBytes())
		}
	} else {
		msg.Code = COAPCODE_405_METHOD_NOT_ALLOWED
	}
	return NewResponseWithMessage(msg)
}

func (c *DefaultClient) handleDeleteRequest(req *CoapRequest) *CoapResponse {
	objectId := req.GetAttributeAsInt("obj")
	instanceId := req.GetAttributeAsInt("inst")

	t := LWM2MObjectType(objectId)
	enabler := c.GetObjectEnabler(t)

	msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
	msg.Token = req.GetMessage().Token
	msg.Payload = NewEmptyPayload()

	if enabler != nil && enabler.GetHandler() != nil {
        lwReq := core.NewDefaultRequest(req, OPERATIONTYPE_DELETE)
		msg.Code = enabler.OnDelete(instanceId, lwReq)
	} else {
		msg.Code = COAPCODE_405_METHOD_NOT_ALLOWED
	}
	return NewResponseWithMessage(msg)
}

func (c *DefaultClient) handleDiscoverRequest() {

}

func (c *DefaultClient) handleObserveRequest() {

}

func (c *DefaultClient) handleWriteRequest(req *CoapRequest) *CoapResponse {
	attrResource := req.GetAttribute("rsrc")
	objectId := req.GetAttributeAsInt("obj")
	instanceId := req.GetAttributeAsInt("inst")

	var resourceId = -1

	if attrResource != "" {
		resourceId = req.GetAttributeAsInt("rsrc")
	}

	t := LWM2MObjectType(objectId)
	enabler := c.GetObjectEnabler(t)

	msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
	msg.Token = req.GetMessage().Token
	msg.Payload = NewEmptyPayload()

	if enabler != nil && enabler.GetHandler() != nil {
		model := enabler.GetModel()
		resource := model.GetResource(resourceId)
		if resource == nil {
			msg.Code = COAPCODE_404_NOT_FOUND
		}

		if !core.IsWritableResource(resource) {
			msg.Code = COAPCODE_405_METHOD_NOT_ALLOWED
		} else {
			lwReq := core.NewDefaultRequest(req, OPERATIONTYPE_WRITE)
			msg.Code = enabler.OnWrite(instanceId, resourceId, lwReq)
		}
	} else {
		msg.Code = COAPCODE_404_NOT_FOUND
	}
	return NewResponseWithMessage(msg)
}

func (c *DefaultClient) handleExecuteRequest(req *CoapRequest) *CoapResponse {
	attrResource := req.GetAttribute("rsrc")
	objectId := req.GetAttributeAsInt("obj")
	instanceId := req.GetAttributeAsInt("inst")

	var resourceId = -1

	if attrResource != "" {
		resourceId = req.GetAttributeAsInt("rsrc")
	}

	t := LWM2MObjectType(objectId)
	enabler := c.GetObjectEnabler(t)

	msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
	msg.Token = req.GetMessage().Token
	msg.Payload = NewEmptyPayload()


	if enabler != nil && enabler.GetHandler() != nil {
		model := enabler.GetModel()
		resource := model.GetResource(resourceId)
		if resource == nil {
			msg.Code = COAPCODE_404_NOT_FOUND
		}

		if !core.IsExecutableResource(resource) {
			msg.Code = COAPCODE_405_METHOD_NOT_ALLOWED
		} else {
            lwReq := core.NewDefaultRequest(req, OPERATIONTYPE_EXECUTE)
			msg.Code = enabler.OnExecute(instanceId, resourceId, lwReq)
		}
	} else {
		msg.Code = COAPCODE_404_NOT_FOUND
	}

	return NewResponseWithMessage(msg)
}

// Events
func (c *DefaultClient) OnStartup(fn FnOnStartup) {
	c.evtOnStartup = fn
}

func (c *DefaultClient) OnRead(fn FnOnRead) {
	c.evtOnRead = fn
}

func (c *DefaultClient) OnWrite(fn FnOnWrite) {
	c.evtOnWrite = fn
}

func (c *DefaultClient) OnExecute(fn FnOnExecute) {
	c.evtOnExecute = fn
}

func (c *DefaultClient) OnRegistered(fn FnOnRegistered) {
	c.evtOnRegistered = fn
}

func (c *DefaultClient) OnDeregistered(fn FnOnDeregistered) {
	c.evtOnDeregistered = fn
}

func (c *DefaultClient) OnError(fn FnOnError) {
	c.evtOnError = fn
}
