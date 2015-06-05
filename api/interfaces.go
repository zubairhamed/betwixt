package api

import (
	. "github.com/zubairhamed/canopus"
	"time"
)

type RequestHandler interface {
	OnRead(int, int, Lwm2mRequest) Lwm2mResponse
	OnDelete(int, Lwm2mRequest) Lwm2mResponse
	OnWrite(int, int, Lwm2mRequest) Lwm2mResponse
	OnCreate(int, int, Lwm2mRequest) Lwm2mResponse
	OnExecute(int, int, Lwm2mRequest) Lwm2mResponse
}

type ResponseValue interface {
	GetBytes() []byte
	GetType() ValueTypeCode
	GetValue() interface{}
	GetStringValue() string
}

type ObjectEnabler interface {
	OnRead(int, int, Lwm2mRequest) Lwm2mResponse
	OnDelete(int, Lwm2mRequest) Lwm2mResponse
	OnWrite(int, int, Lwm2mRequest) Lwm2mResponse
	OnCreate(int, int, Lwm2mRequest) Lwm2mResponse
	OnExecute(int, int, Lwm2mRequest) Lwm2mResponse
}

//type ObjectInstance interface {
//	GetId() int
//	GetTypeId() LWM2MObjectType
//}

type ModelSource interface {
	Initialize()
	Get(LWM2MObjectType) ObjectModel
	Add(ObjectModel, ...ResourceModel)
}

type Registry interface {
	// CreateObjectInstance(LWM2MObjectType, int) ObjectInstance
	GetModel(LWM2MObjectType) ObjectModel
	Register(ModelSource)
	CreateHandler(LWM2MObjectType)
}

type ObjectModel interface {
	GetId() LWM2MObjectType
	SetResources([]ResourceModel)
	GetResources() []ResourceModel
	GetResource(n int) ResourceModel
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
	SetObjects(map[string][]string)
}

type Object interface {
	AddInstance(int)
	RemoveInstance(int)
	GetInstances()[]int
	GetEnabler() ObjectEnabler
	GetType() LWM2MObjectType
	GetModel() ObjectModel
}

