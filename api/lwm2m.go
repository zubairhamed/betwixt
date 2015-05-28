package api

type LWM2MObjectType int
type LWM2MObjectInstances map[LWM2MObjectType]ObjectEnabler

type FnOnStartup func()
type FnOnRead func()
type FnOnWrite func()
type FnOnExecute func()
type FnOnRegistered func(string)
type FnOnDeregistered func()
type FnOnError func()

type ValueTypeCode byte
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

type OperationCode int
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

type IdentifierType byte
const (
	IDENTIFIER_OBJECT_INSTANCE     IdentifierType = 0
	IDENTIFIER_RESOURCE_INSTANCE   IdentifierType = 1
	IDENTIFIER_RESOURCES           IdentifierType = 2
	IDENTIFIER_RESOURCE_WITH_VALUE IdentifierType = 3
)

type BindingMode string
const (
	BINDINGMODE_UDP 						BindingMode = "U"
	BINDINGMODE_UDP_WITH_QUEUE_MODE 		BindingMode = "UQ"
	BINDINGMODE_SMS 						BindingMode = "S"
	BINDINGMODE_SMS_WITH_QUEUE_MODE 		BindingMode = "SQ"
	BINDINGMODE_UDP_AND_SMS 				BindingMode = "US"
	BINDINGMODE_UDP_WITH_QUEUE_MODE_AND_SMS BindingMode = "UQS"
)

type OperationType byte
const (
	OPERATIONTYPE_REGISTER			OperationType = 0
	OPERATIONTYPE_UPDATE			OperationType = 1
	OPERATIONTYPE_DEREGISTER		OperationType = 2
	OPERATIONTYPE_READ				OperationType = 3
	OPERATIONTYPE_DISCOVER			OperationType = 4
	OPERATIONTYPE_WRITE				OperationType = 5
	OPERATIONTYPE_WRITE_ATTRIBUTES	OperationType = 6
	OPERATIONTYPE_EXECUTE			OperationType = 7
	OPERATIONTYPE_CREATE			OperationType = 8
	OPERATIONTYPE_DELETE			OperationType = 9
	OPERATIONTYPE_OBSERVE 			OperationType = 10
	OPERATIONTYPE_NOTIFY 			OperationType = 11
	OPERATIONTYPE_CANCEL_OBSERVE 	OperationType = 12
)