package core

import . "github.com/zubairhamed/lwm2m/api"

type DefaultObjectModel struct {
    Id              LWM2MObjectType
    Name            string
    Description     string
    Multiple        bool
    Mandatory       bool
    Resources       []ResourceModel
}

func (o *DefaultObjectModel) GetId() (LWM2MObjectType){
    return o.Id
}

func (o *DefaultObjectModel) SetResources(r []ResourceModel) {
    o.Resources = r
}

func (o *DefaultObjectModel) GetResource(n int) (ResourceModel){
    for _,rsrc := range o.Resources {
        if rsrc.GetId() == n {
            return rsrc
        }
    }
    return nil
}

func (o *DefaultObjectModel) GetResources() ([]ResourceModel) {
    return o.Resources
}

type DefaultResourceModel struct {
    Id                  int
    Name                string
    Operations          OperationCode
    Multiple            bool
    Mandatory           bool
    ResourceType        ValueTypeCode
    Units               string
    RangeOrEnums        string
    Description         string
}

func (o *DefaultResourceModel) GetId() (int) {
    return o.Id
}

func (o *DefaultResourceModel) MultipleValuesAllowed() (bool) {
    return o.Multiple
}

func (o *DefaultResourceModel) GetResourceType() (ValueTypeCode) {
    return o.ResourceType
}

func (o *DefaultResourceModel) IsExecutable() (bool) {
    return (o.Operations == OPERATION_E || o.Operations == OPERATION_RE || o.Operations == OPERATION_RWE || o.Operations == OPERATION_WE)
}

func (o *DefaultResourceModel) IsReadable() (bool) {
    return (o.Operations == OPERATION_RE || o.Operations == OPERATION_R || o.Operations == OPERATION_RWE || o.Operations == OPERATION_RW)
}

func (o *DefaultResourceModel) IsWritable() (bool) {
    return (o.Operations ==OPERATION_RW || o.Operations == OPERATION_RWE || o.Operations == OPERATION_WE || o.Operations == OPERATION_W)
}

func NewObjectInstance(id int, t LWM2MObjectType) (ObjectInstance) {
    return &DefaultObjectInstance{
        Id: id,
        TypeId: t,
        Resources: make(map[int]Resource),
    }
}

type DefaultObjectInstance struct {
    Id          int
    TypeId      LWM2MObjectType
    Resources   map[int]Resource
}

func (o *DefaultObjectInstance) GetResource(id int) (Resource) {
    return o.Resources[id]
}

func (o *DefaultObjectInstance) GetId() (int) {
    return o.Id
}

func (o *DefaultObjectInstance) GetTypeId() (LWM2MObjectType) {
    return o.TypeId
}

type DefaultResource struct {
    Id          int
}

type DefaultObjectEnabler struct {
    Handler     RequestHandler
    Instances   []ObjectInstance
}

func (en *DefaultObjectEnabler) GetHandler() RequestHandler {
    return en.Handler
}

func (en *DefaultObjectEnabler) GetObjectInstance(idx int) (ObjectInstance) {
    for _, o := range en.Instances {
        if o.GetId() == idx {
            return o
        }
    }
    return nil
}

func (en *DefaultObjectEnabler) GetObjectInstances() []ObjectInstance {
    return en.Instances
}

func (en *DefaultObjectEnabler) SetObjectInstances(o []ObjectInstance) {
    en.Instances = o
}

func (en *DefaultObjectEnabler) OnRead(instanceId int, resourceId int)(ResourceValue) {
    if en.Handler != nil {
        return en.Handler.OnRead(instanceId, resourceId)
    }
    return nil
}
