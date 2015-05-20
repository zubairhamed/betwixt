package api

type RequestHandler interface {
    OnRead(instanceId int, resourceId int)(ResourceValue)
}

type ResourceValue interface {
    GetBytes() []byte
    GetType() ValueTypeCode
    GetValue()  interface{}
    GetStringValue() string
}

type ObjectEnabler interface {
    GetObjectInstance(idx int) (ObjectInstance)
    GetObjectInstances() ([]ObjectInstance)
    SetObjectInstances([]ObjectInstance)
    GetHandler() RequestHandler

    OnRead(instanceId int, resourceId int)(ResourceValue)
}

type ObjectInstance interface {
    GetResource(id int) (Resource)
    GetId() (int)
    GetTypeId() (LWM2MObjectType)
}

type Resource interface {

}

type ModelSource interface {
    Initialize()
    Get(LWM2MObjectType) ObjectModel
    Add(ObjectModel, ...ResourceModel)
}

type Registry interface {
    CreateObjectInstance(t LWM2MObjectType, n int) (ObjectInstance)
    GetModel(n LWM2MObjectType) ObjectModel
    Register(ModelSource)
    CreateHandler(t LWM2MObjectType)
}

type ObjectModel interface {
    GetId() (LWM2MObjectType)
    SetResources([]ResourceModel)
    GetResources()([]ResourceModel)
    GetResource(n int) (ResourceModel)
}

type ResourceModel interface {
    GetId() (int)
    MultipleValuesAllowed() (bool)
    GetResourceType()(ValueTypeCode)

    IsExecutable() (bool)
    IsReadable() (bool)
    IsWritable() (bool)

}