package request

import (
	. "github.com/zubairhamed/go-lwm2m/api"
	"github.com/zubairhamed/goap"
)

func Default(coap *goap.CoapRequest, op OperationType) Request {
	return &DefaultRequest{
		coap: coap,
		op:   op,
	}
}

type DefaultRequest struct {
	coap *goap.CoapRequest
	op   OperationType
}

func (r *DefaultRequest) GetPath() string {
	return r.coap.GetMessage().GetUriPath()
}

func (r *DefaultRequest) GetMessage() *goap.Message {
	return r.coap.GetMessage()
}

func (r *DefaultRequest) GetOperationType() OperationType {
	return r.op
}

func (r *DefaultRequest) GetCoapRequest() *goap.CoapRequest {
	return r.coap
}

func Nil(op OperationType) Request {
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

func (r *NilRequest) GetMessage() *goap.Message {
	return nil
}

func (r *NilRequest) GetOperationType() OperationType {
	return r.op
}

func (r *NilRequest) GetCoapRequest() *goap.CoapRequest {
	return nil
}
