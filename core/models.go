package core

type LWM2MObjectType int

type ObjectModel struct {
    Id              LWM2MObjectType
    Name            string
    Description     string
    Multiple        bool
    Mandatory       bool
    Resources       []*ResourceModel
}

func (o *ObjectModel) GetResource(n int) (*ResourceModel){
    for _,rsrc := range o.Resources {
        if rsrc.Id == n {
            return rsrc
        }
    }
    return nil
}

type ResourceModel struct {
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

type LWM2MResource struct {
    instances   []int
    model       *ObjectModel
}

func NewObjectInstance(t LWM2MObjectType) (*ObjectInstance) {
    return &ObjectInstance{
        TypeId: t,
        Resources: make(map[int]*ResourceInstance),
    }
}

type ObjectInstance struct {
    Id          int
    TypeId      LWM2MObjectType
    Resources   map[int]*ResourceInstance
}

func (o *ObjectInstance) GetResource(id int) (*ResourceInstance) {
    return o.Resources[id]
}

type ResourceInstance struct {
    Id          LWM2MObjectType
    Value       interface{}
}



