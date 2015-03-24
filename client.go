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

func (c *Client) Register(ep string, fn goap.MessageHandler) {
    msg := goap.NewMessageOfType(goap.TYPE_CONFIRMABLE, 12345)
    msg.Code = goap.POST
    msg.AddOptions(goap.NewPathOptions("rd"))
    msg.AddOption(goap.OPTION_URI_QUERY, "ep=" + ep)
    msg.Payload = []byte("</3/0>")

    err := c.coapClient.SendAsync(msg, fn)
    if err != nil {
        log.Println (err)
    }
}
