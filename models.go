package lwm2m

type LWM2MObjectType int

type ObjectModel struct {
    Id              LWM2MObjectType
    Name            string
    Description     string
    Multiple        bool
    Mandatory       bool
    Resources       []*ResourceModel
}

type ResourceModel struct {
    Id                  int
    Name                string
    Operations          OperationCode
    Multiple            bool
    Mandatory           bool
    ResourceType        TypeCode
    Units               string
    RangeOrEnums        string
    Description         string
    Value               interface{}
}

type TypeCode        int
type OperationCode   int

const (
    TYPE_STRING   TypeCode = 0
    TYPE_INTEGER  TypeCode = 1
    TYPE_FLOAT    TypeCode = 2
    TYPE_BOOLEAN  TypeCode = 3
    TYPE_OPAQUE   TypeCode = 4
    TYPE_TIME     TypeCode = 5
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

type ModelSource interface {
    Initialize()
    Get(LWM2MObjectType) *ObjectModel
    Add(*ObjectModel, ...*ResourceModel)
}

func NewDefaultObjectRegistry() (*ObjectRegistry) {
    reg := NewObjectRegistry()

    reg.Register(&LWM2MCoreObjects{})
    reg.Register(&IPSOSmartObjects{})

    return reg
}

func NewObjectRegistry() (*ObjectRegistry) {
    reg := &ObjectRegistry{}
    reg.sources = []ModelSource{}

    return reg
}

type ObjectRegistry struct {
    sources     []ModelSource
}

func (m *ObjectRegistry) CreateObjectInstance(t LWM2MObjectType, n int) (*ObjectInstance) {
    o := m.GetModel(t)
    if o != nil {
        obj := NewObjectInstance(t)
        obj.Id = n
        obj.TypeId = t

        return obj
    }
    return nil
}

func (m *ObjectRegistry) GetModel(n LWM2MObjectType) *ObjectModel {
    for _, s := range m.sources {
        if s != nil {
            o := s.Get(n)
            if o != nil {
                return o
            }
        }
    }
    return nil
}

func (m *ObjectRegistry) Register(s ModelSource) {
    s.Initialize()
    m.sources = append(m.sources, s)
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

type ResourceInstance struct {
    Id          LWM2MObjectType
    Value       interface{}
}

type LWM2MRequest struct {

}

type LWM2MResponse struct {

}