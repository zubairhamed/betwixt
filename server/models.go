package server

import "github.com/zubairhamed/go-lwm2m/api"

func NewRegisteredClient(ep string, id string) api.RegisteredClient {
	return &DefaultRegisteredClient{
		name: ep,
		id:   id,
	}
}

type DefaultRegisteredClient struct {
	id          string
	name        string
	lifetime    int
	version     string
	bindingMode api.BindingMode
	smsNumber   string
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

func (c *DefaultRegisteredClient) GetBindingMode() api.BindingMode {
	return c.bindingMode
}

func (c *DefaultRegisteredClient) GetSmsNumber() string {
	return c.smsNumber
}
