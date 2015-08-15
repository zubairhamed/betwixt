package betwixt

// A default enabler which does abosolutely..nothing.
// Its still trying to find its purpose in life. Best of luck to it.
func NewNullEnabler() ObjectEnabler {
	return &NullEnabler{}
}

type NullEnabler struct {
}

func (e *NullEnabler) OnRead(int, int, Lwm2mRequest) Lwm2mResponse {
	return MethodNotAllowed()
}

func (e *NullEnabler) OnDelete(int, Lwm2mRequest) Lwm2mResponse {
	return MethodNotAllowed()
}

func (e *NullEnabler) OnWrite(int, int, Lwm2mRequest) Lwm2mResponse {
	return MethodNotAllowed()
}

func (e *NullEnabler) OnCreate(int, int, Lwm2mRequest) Lwm2mResponse {
	return MethodNotAllowed()
}

func (e *NullEnabler) OnExecute(int, int, Lwm2mRequest) Lwm2mResponse {
	return MethodNotAllowed()
}
