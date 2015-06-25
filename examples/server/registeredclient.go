package server

import (
	"fmt"
	"github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/betwixt/core/utils"
	. "github.com/zubairhamed/canopus"
	"github.com/zubairhamed/go-commons/network"
	"github.com/zubairhamed/go-commons/typeval"
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

func (c *DefaultRegisteredClient) ReadObject(obj uint16, inst uint16) (typeval.Value, error) {
	return nil, nil
}

func (c *DefaultRegisteredClient) ReadResource(obj uint16, inst uint16, rsrc uint16) (typeval.Value, error) {
	rAddr, _ := net.ResolveUDPAddr("udp", c.addr)
	lAddr, _ := net.ResolveUDPAddr("udp", ":0")

	conn, _ := net.DialUDP("udp", lAddr, rAddr)

	uri := fmt.Sprintf("%d/%d/%d", obj, inst, rsrc)
	req := NewRequest(TYPE_CONFIRMABLE, GET, GenerateMessageId())
	req.SetRequestURI(uri)

	resourceDefinition := c.GetObject(betwixt.LWM2MObjectType(obj)).GetDefinition().GetResource(rsrc)
	if resourceDefinition.MultipleValuesAllowed() {
		req.SetMediaType(network.MEDIATYPE_TLV_VND_OMA_LWM2M)
	} else {
		req.SetMediaType(network.MEDIATYPE_TEXT_PLAIN_VND_OMA_LWM2M)
	}

	response, _ := SendMessage(req.GetMessage(), conn)
	responseValue, _ := utils.DecodeResourceValue(rsrc, response.GetMessage().Payload.GetBytes(), resourceDefinition)

	return responseValue, nil
}

func (c *DefaultRegisteredClient) Delete(int, int) {

}

func (c *DefaultRegisteredClient) Execute(int, int, int) {

}
