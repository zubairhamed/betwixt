package core

type OperationCode   int
type LWM2MObjectInstances map[LWM2MObjectType] *ObjectEnabler

type ValueTypeCode byte
const (
    VALUETYPE_STRING         ValueTypeCode = 0
    VALUETYPE_INTEGER        ValueTypeCode = 1
    VALUETYPE_FLOAT          ValueTypeCode = 2
    VALUETYPE_BOOLEAN        ValueTypeCode = 3
    VALUETYPE_OPAQUE         ValueTypeCode = 4
    VALUETYPE_TIME           ValueTypeCode = 5
    VALUETYPE_OBJECTLINK     ValueTypeCode = 6
    VALUETYPE_MULTIPLE       ValueTypeCode = 6
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
    // OnRead(*ResourceModel, int) (ResourceValue)
    OnRead(objectId int, instanceId int)(ResourceValue)
}

type ObjectEnabler struct {
    Handler     ObjectHandler
    Instances   []*ObjectInstance
}

func (en *ObjectEnabler) GetObjectInstance(idx int) (*ObjectInstance) {
    for _, o := range en.Instances {
        if o.Id == idx {
            return o
        }
    }
    return nil
}


