package server

import (
	"github.com/zubairhamed/betwixt"
	"time"
	"net"
	. "github.com/zubairhamed/canopus"
	"fmt"
	"log"
)

type ServerStatistics struct {
	requestsCount int
}

func (s *ServerStatistics) IncrementCoapRequestsCount() {
	s.requestsCount++
}

func (s *ServerStatistics) GetRequestsCount() int {
	return s.requestsCount
}

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

func (c *DefaultRegisteredClient) Read(obj int, inst int, rsrc int) {
	log.Println("READ!")
	rAddr, err := net.ResolveUDPAddr("udp", c.addr)
	lAddr, err := net.ResolveUDPAddr("udp", ":0")

	conn, _ := net.DialUDP("udp", lAddr, rAddr)

	uri := fmt.Sprintf("/%d/%d/%d", obj, inst, rsrc)
	req := NewRequest(TYPE_CONFIRMABLE, GET, GenerateMessageId())
	req.SetRequestURI(uri)

	log.Println("SendMEssage")
	response, err := SendMessage(req.GetMessage(), conn)

	log.Println(response)
	log.Println(response.GetMessage())
	log.Println(response.GetError())
	log.Println(response.GetPayload())
	log.Println(err)
}

func (c *DefaultRegisteredClient) Delete(int, int) {

}

func (c *DefaultRegisteredClient) Execute(int, int, int) {

}
