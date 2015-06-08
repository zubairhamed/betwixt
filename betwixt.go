package betwixt

import (
	. "github.com/zubairhamed/canopus"
	"time"
)

type LWM2MObjectType int
type LWM2MObjectInstances map[LWM2MObjectType]Object

type FnOnStartup func()
type FnOnRead func()
type FnOnWrite func()
type FnOnExecute func()
type FnOnRegistered func(string)
type FnOnDeregistered func()
type FnOnError func()

type ValueTypeCode byte
type OperationCode int
type IdentifierType byte
type BindingMode string
type OperationType byte

const (
	VALUETYPE_STRING     ValueTypeCode = 0
	VALUETYPE_INTEGER    ValueTypeCode = 1
	VALUETYPE_FLOAT      ValueTypeCode = 2
	VALUETYPE_BOOLEAN    ValueTypeCode = 3
	VALUETYPE_OPAQUE     ValueTypeCode = 4
	VALUETYPE_TIME       ValueTypeCode = 5
	VALUETYPE_OBJECTLINK ValueTypeCode = 6
	VALUETYPE_MULTIPLE   ValueTypeCode = 6
	VALUETYPE_TLV        ValueTypeCode = 7
)

const (
	OPERATION_NONE OperationCode = 0
	OPERATION_R    OperationCode = 1
	OPERATION_W    OperationCode = 2
	OPERATION_RW   OperationCode = 3
	OPERATION_E    OperationCode = 4
	OPERATION_RE   OperationCode = 5
	OPERATION_WE   OperationCode = 6
	OPERATION_RWE  OperationCode = 7
)

const (
	IDENTIFIER_OBJECT_INSTANCE     IdentifierType = 0
	IDENTIFIER_RESOURCE_INSTANCE   IdentifierType = 1
	IDENTIFIER_RESOURCES           IdentifierType = 2
	IDENTIFIER_RESOURCE_WITH_VALUE IdentifierType = 3
)

const (
	BINDINGMODE_UDP                         BindingMode = "U"
	BINDINGMODE_UDP_WITH_QUEUE_MODE         BindingMode = "UQ"
	BINDINGMODE_SMS                         BindingMode = "S"
	BINDINGMODE_SMS_WITH_QUEUE_MODE         BindingMode = "SQ"
	BINDINGMODE_UDP_AND_SMS                 BindingMode = "US"
	BINDINGMODE_UDP_WITH_QUEUE_MODE_AND_SMS BindingMode = "UQS"
)

const (
	OPERATIONTYPE_REGISTER         OperationType = 0
	OPERATIONTYPE_UPDATE           OperationType = 1
	OPERATIONTYPE_DEREGISTER       OperationType = 2
	OPERATIONTYPE_READ             OperationType = 3
	OPERATIONTYPE_DISCOVER         OperationType = 4
	OPERATIONTYPE_WRITE            OperationType = 5
	OPERATIONTYPE_WRITE_ATTRIBUTES OperationType = 6
	OPERATIONTYPE_EXECUTE          OperationType = 7
	OPERATIONTYPE_CREATE           OperationType = 8
	OPERATIONTYPE_DELETE           OperationType = 9
	OPERATIONTYPE_OBSERVE          OperationType = 10
	OPERATIONTYPE_NOTIFY           OperationType = 11
	OPERATIONTYPE_CANCEL_OBSERVE   OperationType = 12
)

// ResponseValue interface represents response to a server request
// Typical response could be plain text, TLV Binary, TLV JSON
type ResponseValue interface {
	GetBytes() []byte
	GetType() ValueTypeCode
	GetValue() interface{}
	GetStringValue() string
}

// ObjectEnabler interface to handler any incoming requests from a server for a given object
type ObjectEnabler interface {
	OnRead(int, int, Lwm2mRequest) Lwm2mResponse
	OnDelete(int, Lwm2mRequest) Lwm2mResponse
	OnWrite(int, int, Lwm2mRequest) Lwm2mResponse
	OnCreate(int, int, Lwm2mRequest) Lwm2mResponse
	OnExecute(int, int, Lwm2mRequest) Lwm2mResponse
}

// ObjectSource interface representing a source consumed by a Registry to resolve and retrieve
// LWM2M object definitions
type ObjectSource interface {
	Initialize()
	GetObject(n LWM2MObjectType) ObjectDefinition
	GetObjects() map[LWM2MObjectType]ObjectDefinition
	AddObject(m ObjectDefinition, res ...ResourceDefinition)
}

// Registry interface represents a source from which LWM2M object definitions can be looked up/resolved or
// stored
type Registry interface {
	GetDefinition(LWM2MObjectType) ObjectDefinition
	Register(ObjectSource)
	GetMandatory() []ObjectDefinition
}

// ObjectDefinition interface defines a LWM2M Object
type ObjectDefinition interface {
	GetType() LWM2MObjectType
	GetDescription() string
	SetResources([]ResourceDefinition)
	GetResources() []ResourceDefinition
	GetResource(n int) ResourceDefinition
	AllowMultiple() bool
	IsMandatory() bool
}

// ResourceDefinition interface defines a LWM2M Resource
type ResourceDefinition interface {
	GetId() int
	MultipleValuesAllowed() bool
	GetResourceType() ValueTypeCode
	GetOperations() OperationCode
}

// LWM2MClient interface defining a LWM2M Client
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

// Lwm2mRequest interface represents an incoming request from a server
type Lwm2mRequest interface {
	GetPath() string
	GetMessage() *Message
	GetOperationType() OperationType
	GetCoapRequest() *CoapRequest
}

// Lwm2mResponse interface represents an outgoing response to a server
type Lwm2mResponse interface {
	GetResponseCode() CoapCode
	GetResponseValue() ResponseValue
}

// Server interface defines a LWM2M Server
type Server interface {
	UseRegistry(Registry)
	Start()
}

// RegisteredClient interface is an instance of a client registered on a server
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
	SetObjects(map[LWM2MObjectType]Object)
}

// An Object interface represents an Object used on a client or Objects supported by a Registered Client on a server
// Not to be confused with ObjectDefinition, which represents the definition of an Object
type Object interface {
	AddInstance(int)
	RemoveInstance(int)
	GetInstances() []int
	GetEnabler() ObjectEnabler
	GetType() LWM2MObjectType
	GetDefinition() ObjectDefinition
	SetEnabler(ObjectEnabler)
}
