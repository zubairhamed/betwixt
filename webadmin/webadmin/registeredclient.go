package webadmin

import (
	"fmt"
	"github.com/zubairhamed/betwixt"
	. "github.com/zubairhamed/canopus"
	"net"
	"time"
)

// Returns a new instance of DefaultRegisteredClient implementing RegisteredClient
func NewRegisteredClient(ep string, id string, addr string) betwixt.RegisteredClient {
	return &DefaultRegisteredClient{
		name:       ep,
		id:         id,
		addr:       addr,
		regDate:    time.Now(),
		updateDate: time.Now(),
	}
}

type DefaultRegisteredClient struct {
	id             string
	name           string
	lifetime       int
	version        string
	bindingMode    betwixt.BindingMode
	smsNumber      string
	addr           string
	regDate        time.Time
	updateDate     time.Time
	enabledObjects map[betwixt.LWM2MObjectType]betwixt.Object
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

func (c *DefaultRegisteredClient) GetBindingMode() betwixt.BindingMode {
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

func (c *DefaultRegisteredClient) SetObjects(objects map[betwixt.LWM2MObjectType]betwixt.Object) {
	c.enabledObjects = objects
}

func (c *DefaultRegisteredClient) GetObjects() map[betwixt.LWM2MObjectType]betwixt.Object {
	return c.enabledObjects
}

func (c *DefaultRegisteredClient) GetObject(t betwixt.LWM2MObjectType) betwixt.Object {
	return c.enabledObjects[t]
}

func (c *DefaultRegisteredClient) ReadObject(obj uint16, inst uint16) (betwixt.Value, error) {
	return nil, nil
}

func (c *DefaultRegisteredClient) ReadResource(obj uint16, inst uint16, rsrc uint16) (betwixt.Value, error) {
	rAddr, _ := net.ResolveUDPAddr("udp", c.addr)
	lAddr, _ := net.ResolveUDPAddr("udp", ":0")

	conn, _ := net.DialUDP("udp", lAddr, rAddr)

	uri := fmt.Sprintf("/%d/%d/%d", obj, inst, rsrc)
	req := NewRequest(MessageConfirmable, Get, GenerateMessageID())
	req.SetRequestURI(uri)

	resourceDefinition := c.GetObject(betwixt.LWM2MObjectType(obj)).GetDefinition().GetResource(rsrc)
	if resourceDefinition.MultipleValuesAllowed() {
		req.SetMediaType(MediaTypeTlvVndOmaLwm2m)
	} else {
		req.SetMediaType(MediaTypeTextPlainVndOmaLwm2m)
	}

	response, _ := SendMessage(req.GetMessage(), NewUDPConnection(conn))
	PrintMessage(response.GetMessage())
	responseValue, _ := betwixt.DecodeResourceValue(rsrc, response.GetMessage().Payload.GetBytes(), resourceDefinition)

	return responseValue, nil
}

func (c *DefaultRegisteredClient) Delete(int, int) {

}

func (c *DefaultRegisteredClient) Execute(int, int, int) {

}
