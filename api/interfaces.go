package api
import "github.com/zubairhamed/goap"

type RequestHandler interface {
    OnRead(int, int)(ResponseValue, goap.CoapCode)
    OnDelete(int)(goap.CoapCode)
    OnWrite(int, int)(goap.CoapCode)
    OnCreate(int, int)(goap.CoapCode)
    OnExecute(int, int)(goap.CoapCode)
}

/*
type ResourceValue interface {
    GetBytes() ([]byte)
    GetType() (ValueTypeCode)
    GetValue() (interface{})
    GetStringValue() (string)
}
*/

type RequestValue interface {
    GetBytes() ([]byte)
    GetType() (ValueTypeCode)
    GetValue() (interface{})
    GetStringValue() (string)
}

type ResponseValue interface {
    GetBytes() ([]byte)
    GetType() (ValueTypeCode)
    GetValue() (interface{})
    GetStringValue() (string)
}


type ObjectEnabler interface {
    GetObjectInstance(int) (ObjectInstance)
    GetObjectInstances() ([]ObjectInstance)
    SetObjectInstances([]ObjectInstance)
    GetHandler() RequestHandler

    OnRead(int, int)(RequestValue, goap.CoapCode)
    OnDelete(int)(goap.CoapCode)
    OnWrite(int, int)(goap.CoapCode)
    OnCreate(int, int)(goap.CoapCode)
    OnExecute(int, int)(goap.CoapCode)
}

type ObjectInstance interface {
    GetResource(int) (Resource)
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
    CreateObjectInstance(LWM2MObjectType, int) (ObjectInstance)
    GetModel(LWM2MObjectType) ObjectModel
    Register(ModelSource)
    CreateHandler(LWM2MObjectType)
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
    GetOperations() (OperationCode)
}

type LWM2MClient interface {
    AddObjectInstance(ObjectInstance) (error)
    AddObjectInstances (... ObjectInstance)
    AddResource()
    AddObject()
    Register(string) (string)
    Deregister()
    Update()
    UseRegistry(Registry)
    EnableObject(LWM2MObjectType, RequestHandler) (error)
    GetRegistry() Registry
    GetEnabledObjects() (map[LWM2MObjectType] ObjectEnabler)
    GetObjectEnabler(LWM2MObjectType) (ObjectEnabler)
    GetObjectInstance(LWM2MObjectType, int) (ObjectInstance)
    Start()

    // Events
    OnStartup(FnOnStartup)
    OnRead(FnOnRead)
    OnWrite(FnOnWrite)
    OnExecute(FnOnExecute)
    OnRegistered(FnOnRegistered)
    OnDeregistered(FnOnDeregistered)
    OnError (FnOnError)
}