package api

import (
	. "github.com/zubairhamed/canopus"
	"time"
)

// ResponseValue interface represents response to a server request
// Typical response could be plain text, TLV Binary, TLV JSON
type ResponseValue interface {
	GetBytes() []byte
	GetType() ValueTypeCode
	GetValue() interface{}
	GetStringValue() string
}

//
type ObjectEnabler interface {
	OnRead(int, int, Lwm2mRequest) Lwm2mResponse
	OnDelete(int, Lwm2mRequest) Lwm2mResponse
	OnWrite(int, int, Lwm2mRequest) Lwm2mResponse
	OnCreate(int, int, Lwm2mRequest) Lwm2mResponse
	OnExecute(int, int, Lwm2mRequest) Lwm2mResponse
}

type ObjectSource interface {
	Initialize()
	GetObject(n LWM2MObjectType) ObjectModel
	GetObjects() map[LWM2MObjectType]ObjectModel
	AddObject(m ObjectModel, res ...ResourceModel)
}

type Registry interface {
	GetModel(LWM2MObjectType) ObjectModel
	Register(ObjectSource)
	GetMandatory() []ObjectModel
}

type ObjectModel interface {
	GetType() LWM2MObjectType
	GetDescription() string
	SetResources([]ResourceModel)
	GetResources() []ResourceModel
	GetResource(n int) ResourceModel
	AllowMultiple() bool
	IsMandatory() bool
}

type ResourceModel interface {
	GetId() int
	MultipleValuesAllowed() bool
	GetResourceType() ValueTypeCode
	GetOperations() OperationCode
}

type LWM2MClient interface {
	AddObjectInstance(LWM2MObjectType, int) error
	AddObjectInstances(LWM2MObjectType, ...int)
	AddResource()
	AddObject()
	Register(string) string
	Deregister()
	Update()
	UseRegistry(Registry)
	EnableObject(LWM2MObjectType, ObjectEnabler) error
	SetEnabler(LWM2MObjectType, ObjectEnabler)
	GetRegistry() Registry
	GetEnabledObjects() map[LWM2MObjectType]Object
	Start()

	// Events
	OnStartup(FnOnStartup)
	OnRead(FnOnRead)
	OnWrite(FnOnWrite)
	OnExecute(FnOnExecute)
	OnRegistered(FnOnRegistered)
	OnDeregistered(FnOnDeregistered)
	OnError(FnOnError)
}

type Lwm2mRequest interface {
	GetPath() string
	GetMessage() *Message
	GetOperationType() OperationType
	GetCoapRequest() *CoapRequest
}

type Lwm2mResponse interface {
	GetResponseCode() CoapCode
	GetResponseValue() ResponseValue
}

type Server interface {
	UseRegistry(Registry)
	Start()
}

type RegisteredClient interface {
	GetId() string
	GetName() string
	GetLifetime() int
	GetVersion() string
	GetBindingMode() BindingMode
	GetSmsNumber() string
	GetRegistrationDate() time.Time
	Update()
	LastUpdate() time.Time
}

type Object interface {
	AddInstance(int)
	RemoveInstance(int)
	GetInstances()[]int
	GetEnabler() ObjectEnabler
	GetType() LWM2MObjectType
	GetModel() ObjectModel
	SetEnabler(ObjectEnabler)
}
