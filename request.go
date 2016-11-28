package betwixt

import (
	"github.com/zubairhamed/canopus"
)

// Default is a helper/shortcut method which creates a Default LWM2M Request
func Default(coap canopus.Request, op OperationType) Lwm2mRequest {
	return &DefaultRequest{
		coap: coap,
		op:   op,
	}
}

type DefaultRequest struct {
	coap canopus.Request
	op   OperationType
}

func (r *DefaultRequest) GetPath() string {
	return r.coap.GetMessage().GetURIPath()
}

func (r *DefaultRequest) GetMessage() canopus.Message {
	return r.coap.GetMessage()
}

func (r *DefaultRequest) GetOperationType() OperationType {
	return r.op
}

func (r *DefaultRequest) GetCoapRequest() canopus.Request {
	return r.coap
}

func Nil(op OperationType) Lwm2mRequest {
	return &NilRequest{
		op: op,
	}
}

type NilRequest struct {
	op OperationType
}

func (r *NilRequest) GetPath() string {
	return ""
}

func (r *NilRequest) GetMessage() canopus.Message {
	return nil
}

func (r *NilRequest) GetOperationType() OperationType {
	return r.op
}

func (r *NilRequest) GetCoapRequest() canopus.Request {
	return nil
}
