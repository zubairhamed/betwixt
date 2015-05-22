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
