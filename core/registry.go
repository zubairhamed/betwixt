package core

type ModelSource interface {
    Initialize()
    Get(LWM2MObjectType) *ObjectModel
    Add(*ObjectModel, ...*ResourceModel)
}

type Registry interface {
    CreateObjectInstance(t LWM2MObjectType, n int) (*ObjectInstance)
    GetModel(n LWM2MObjectType) *ObjectModel
    Register(ModelSource)
    CreateHandler(t LWM2MObjectType)
}


