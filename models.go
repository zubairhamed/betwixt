package lwm2m

type ObjectModel struct {
    Id              int
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
    RangeOrEnums         string
    Description         string
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
    Get(int) *ObjectModel
    Add(*ObjectModel, ...*ResourceModel)
}

func NewModelRepository() (*ModelRepository) {
    repo := &ModelRepository{}
    repo.sources = make([]ModelSource, 10)

    return repo
}

type ModelRepository struct {
    sources     []ModelSource
}

func (m *ModelRepository) GetModel(n int) *ObjectModel {
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

func (m *ModelRepository) GetModels(n int) []*ObjectModel {
    var models []*ObjectModel

    for _, s := range m.sources {
        if s != nil {
            o := s.Get(n)
            if o != nil {
                models = append(models, o)
            }
        }
    }
    return models
}

func (m *ModelRepository) Register(s ModelSource) {
    s.Initialize()
    m.sources = append(m.sources, s)
}


func NewLWM2MResource (m *ObjectModel, instances ... int) (*LWM2MResource) {
    return &LWM2MResource{
        instances: instances,
        model: m,
    }
}

type LWM2MResource struct {
    instances   []int
    model       *ObjectModel
}

