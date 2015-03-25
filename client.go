package golwm2m

import "github.com/zubairhamed/goap"
import "log"

func NewClient() (*Client) {
	return &Client{
        coapClient: goap.NewClient(),
    }
}

type Client struct {
	coapClient	*goap.Client
}

func (c *Client) Dial(nwNet string, host string, port int) {
    c.coapClient.Dial(nwNet, host, port)
}

func (c *Client) Bootstrap(ep string, fn goap.MessageHandler) {
    msg := goap.NewMessageOfType(goap.TYPE_CONFIRMABLE, 12345)
    msg.Code = goap.POST
    msg.AddOptions(goap.NewPathOptions("bs"))
    msg.AddOption(goap.OPTION_URI_QUERY, "ep=" + ep)

    err := c.coapClient.SendAsync(msg, fn)
    if err != nil {
        log.Println (err)
    }
}

// TODO: Return EndPoint URL
func (c *Client) Register(ep string, fn goap.MessageHandler) {
    msg := goap.NewMessageOfType(goap.TYPE_CONFIRMABLE, goap.GenerateMessageId())
    msg.Code = goap.POST
    msg.AddOptions(goap.NewPathOptions("rd"))
    msg.AddOption(goap.OPTION_URI_QUERY, "ep=" + ep)
    msg.Payload = []byte("</3/0>")

    err := c.coapClient.SendAsync(msg, fn)
    if err != nil {
        log.Println (err)
    }

    // /rd?ep={Endpoint Client Name}&lt={Lifetime}&sms={MSISDN} &lwm2m={version}&b={binding}
}

func (c *Client) Deregister(ins string, fn goap.MessageHandler) {
    msg := goap.NewMessageOfType(goap.TYPE_CONFIRMABLE, goap.GenerateMessageId())
    msg.Code = goap.DELETE
    msg.AddOptions(goap.NewPathOptions(ins))
    err := c.coapClient.SendAsync(msg, fn)
    if err != nil {
        log.Println (err)
    }
}

/*
    Update
    PUT
    /{location}?lt={Lifetime}&sms={MSISDN} &b={binding}

    ---
    Deregister
    DELETE
    /{location}

    ---
    Read
    GET
    /{Object ID}/{Object Instance ID}/{Resource ID}

    ---
    Discover
    GET Accept: application/link- format
    /{Object ID}/{Object Instance ID}/{Resource ID}

    ---
    Write
    PUT/POST
    /{Object ID}/{Object Instance ID}/{Resource ID}

    ---
    Write Attributes
    PUT
    /{Object ID}/{Object Instance ID}/{Resource ID}?pmin={minimum period}&pmax={maximum period}&gt={greater than}&lt={less than}&st={step}&cancel

    ---
    Execute
    POST
    /{Object ID}/{Object Instance ID}/{Resource ID}

    ---
    Create
    POST
    /{Object ID}/{Object Instance ID}

    ---
    Delete
    DELETE
    /{Object ID}/{Object Instance ID}

    ---
    Observe
    GET with Observe Option
    /{Object ID}/{Object Instance ID}/{Resource ID}

    ---
    Cancel Observation
    Reset message

    ---
    Notify
    Asynchronous Response

    ---




*/

