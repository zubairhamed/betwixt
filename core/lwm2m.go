package core


type OperationCode   int
type LWM2MObjectInstances map[LWM2MObjectType] *ObjectEnabler

type ValueTypeCode byte
const (
    TYPE_STRING         ValueTypeCode = 0
    TYPE_INTEGER        ValueTypeCode = 1
    TYPE_FLOAT          ValueTypeCode = 2
    TYPE_BOOLEAN        ValueTypeCode = 3
    TYPE_OPAQUE         ValueTypeCode = 4
    TYPE_TIME           ValueTypeCode = 5
    TYPE_OBJECTLINK     ValueTypeCode = 6
)

const (
    OPERATION_NONE  OperationCode = 0
    OPERATION_R     OperationCode = 1
    OPERATION_W     OperationCode = 2
    OPERATION_RW    OperationCode = 3
    OPERATION_E     OperationCode = 4
    OPERATION_RE    OperationCode = 5
    OPERATION_WE    OperationCode = 6
    OPERATION_RWE   OperationCode = 7
)

type IdentifierType byte
const (
    OBJECT_INSTANCE     IdentifierType = 0
    RESOURCE_INSTANCE   IdentifierType = 1
    RESOURCES           IdentifierType = 2
    RESOURCE            IdentifierType = 3
)

// Enablers
type ObjectHandler interface {
    OnRead(LWM2MObjectType, *ObjectModel, *ObjectInstance, *ResourceModel) (ResourceValue)
}

type ObjectEnabler struct {
    Handler     ObjectHandler
    Instances   []*ObjectInstance
}



