package betwixt

import (
	"fmt"
	"log"
	"net"
	"time"

	. "github.com/zubairhamed/canopus"
)

// Returns a new instance of DefaultRegisteredClient implementing RegisteredClient
func NewRegisteredClient(ep string, id string, addr string, coapServer CoapServer) RegisteredClient {
	return &DefaultRegisteredClient{
		name:       ep,
		id:         id,
		addr:       addr,
		regDate:    time.Now(),
		updateDate: time.Now(),
		coapServer: coapServer,
	}
}

type DefaultRegisteredClient struct {
	id             string
	name           string
	lifetime       int
	version        string
	bindingMode    BindingMode
	smsNumber      string
	addr           string
	regDate        time.Time
	updateDate     time.Time
	coapServer     CoapServer
	enabledObjects map[LWM2MObjectType]Object
}

func (c *DefaultRegisteredClient) GetAddress() string {
	return c.addr
}

func (c *DefaultRegisteredClient) GetId() string {
	return c.id
}

func (c *DefaultRegisteredClient) GetName() string {
	return c.name
}

func (c *DefaultRegisteredClient) GetLifetime() int {
	return c.lifetime
}

func (c *DefaultRegisteredClient) GetVersion() string {
	return c.version
}

func (c *DefaultRegisteredClient) GetBindingMode() BindingMode {
	return c.bindingMode
}

func (c *DefaultRegisteredClient) GetSmsNumber() string {
	return c.smsNumber
}

func (c *DefaultRegisteredClient) GetRegistrationDate() time.Time {
	return c.regDate
}

func (c *DefaultRegisteredClient) Update() {
	c.updateDate = time.Now()
}

func (c *DefaultRegisteredClient) LastUpdate() time.Time {
	return c.updateDate
}

func (c *DefaultRegisteredClient) SetObjects(objects map[LWM2MObjectType]Object) {
	c.enabledObjects = objects
}

func (c *DefaultRegisteredClient) GetObjects() map[LWM2MObjectType]Object {
	return c.enabledObjects
}

func (c *DefaultRegisteredClient) GetObject(t LWM2MObjectType) Object {
	return c.enabledObjects[t]
}

func (c *DefaultRegisteredClient) ReadObject(obj uint16, inst uint16) (Value, error) {
	return nil, nil
}

func (c *DefaultRegisteredClient) ReadResource(obj uint16, inst uint16, rsrc uint16) (Value, error) {
	rAddr, _ := net.ResolveUDPAddr("udp", c.addr)

	uri := fmt.Sprintf("/%d/%d/%d", obj, inst, rsrc)
	req := NewRequest(MessageConfirmable, Get, GenerateMessageID())
	req.SetRequestURI(uri)

	resourceDefinition := c.GetObject(LWM2MObjectType(obj)).GetDefinition().GetResource(LWM2MResourceType(rsrc))
	if resourceDefinition.MultipleValuesAllowed() {
		req.SetMediaType(MediaTypeTlvVndOmaLwm2m)
	} else {
		req.SetMediaType(MediaTypeTextPlainVndOmaLwm2m)
	}

	response, err := c.coapServer.SendTo(req, rAddr)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	responseValue, _ := DecodeResourceValue(LWM2MResourceType(rsrc), response.GetMessage().GetPayload().GetBytes(), resourceDefinition)

	return responseValue, nil
}

func (c *DefaultRegisteredClient) Delete(int, int) {

}

func (c *DefaultRegisteredClient) Execute(int, int, int) {

}
