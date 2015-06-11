package enablers

import (
	. "github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/betwixt/core/response"
)

func NewNullEnabler() ObjectEnabler {
	return &NullEnabler{}
}

type NullEnabler struct {
}

func (e *NullEnabler) OnRead(int, int, Lwm2mRequest) Lwm2mResponse {
	return response.MethodNotAllowed()
}

func (e *NullEnabler) OnDelete(int, Lwm2mRequest) Lwm2mResponse {
	return response.MethodNotAllowed()
}

func (e *NullEnabler) OnWrite(int, int, Lwm2mRequest) Lwm2mResponse {
	return response.MethodNotAllowed()
}

func (e *NullEnabler) OnCreate(int, int, Lwm2mRequest) Lwm2mResponse {
	return response.MethodNotAllowed()
}

func (e *NullEnabler) OnExecute(int, int, Lwm2mRequest) Lwm2mResponse {
	return response.MethodNotAllowed()
}
