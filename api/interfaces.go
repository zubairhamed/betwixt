package api

import . "github.com/zubairhamed/goap"

type RequestHandler interface {
	OnRead(int, int, Request) (ResponseValue, CoapCode)
	OnDelete(int, Request) CoapCode
	OnWrite(int, int, Request) CoapCode
	OnCreate(int, int, Request) CoapCode
	OnExecute(int, int, Request) CoapCode
}

/*
type RequestValue interface {
	GetBytes() []byte
	GetType() ValueTypeCode
	GetValue() interface{}
	GetStringValue() string
}
*/

type ResponseValue interface {
	GetBytes() []byte
	GetType() ValueTypeCode
	GetValue() interface{}
	GetStringValue() string
}

type ObjectEnabler interface {
	GetObjectInstance(int) ObjectInstance
	GetObjectInstances() []ObjectInstance
	SetObjectInstances([]ObjectInstance)
	GetHandler() RequestHandler
	GetModel() ObjectModel

	OnRead(int, int, Request) (ResponseValue, CoapCode)
	OnDelete(int, Request) CoapCode
	OnWrite(int, int, Request) CoapCode
	OnCreate(int, int, Request) CoapCode
	OnExecute(int, int, Request) CoapCode
}

type ObjectInstance interface {
	GetId() int
	GetTypeId() LWM2MObjectType
}

type ModelSource interface {
	Initialize()
	Get(LWM2MObjectType) ObjectModel
	Add(ObjectModel, ...ResourceModel)
}

type Registry interface {
	CreateObjectInstance(LWM2MObjectType, int) ObjectInstance
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
	AddObjectInstance(ObjectInstance) error
	AddObjectInstances(...ObjectInstance)
	AddResource()
	AddObject()
	Register(string) string
	Deregister()
	Update()
	UseRegistry(Registry)
	EnableObject(LWM2MObjectType, RequestHandler) error
	GetRegistry() Registry
	GetEnabledObjects() map[LWM2MObjectType]ObjectEnabler
	GetObjectEnabler(LWM2MObjectType) ObjectEnabler
	GetObjectInstance(LWM2MObjectType, int) ObjectInstance
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

type Request interface {
	GetPath() string
	GetMessage() *Message
	GetOperationType() OperationType
	GetCoapRequest() *CoapRequest
}

type Response interface {
	GetResponseCode() CoapCode
}
