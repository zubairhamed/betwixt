package lwm2m
import "log"

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
    RangeOrEnum         string
    Description         string
}

type TypeCode        int
type OperationCode   int

const (
    OBJECTTYPE_STRING   TypeCode = 0
    OBJECTTYPE_INTEGER  TypeCode = 1
    OBJECTTYPE_FLOAT    TypeCode = 2
    OBJECTTYPE_BOOLEAN  TypeCode = 3
    OBJECTTYPE_OPAQUE   TypeCode = 4
    OBJECTTYPE_TIME     TypeCode = 5
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

func NewModelsRepository() (*ModelsRepository) {
    mr := &ModelsRepository{}
    mr.models = make(map[int]*ObjectModel)

    mr.Add(
        &ObjectModel{ Name: "LWM2M Security",  Id: 0, Multiple: true, Mandatory: true, Description: "" },
        &ResourceModel{ Id: 0, Name: "LWM2M  Server URI", Operations: OPERATION_NONE, Multiple: false, Mandatory: true, ResourceType: OBJECTTYPE_STRING, RangeOrEnum: "0-255 bytes", Units: "", Description: "" },
        &ResourceModel{ Id: 1, Name: "Bootstrap Server", Operations: OPERATION_NONE, Multiple: false, Mandatory: true, ResourceType: OBJECTTYPE_BOOLEAN, RangeOrEnum: "", Units: "", Description: "" },
        &ResourceModel{ Id: 2, Name: "Security Mode", Operations: OPERATION_NONE, Multiple: false, Mandatory: true,ResourceType: OBJECTTYPE_INTEGER, RangeOrEnum: "0-3",Units: "", Description: ""},
        &ResourceModel{ Id: 3, Name: "Public Key or Identity",Operations: OPERATION_NONE, Multiple: false, Mandatory: true,ResourceType: OBJECTTYPE_OPAQUE, RangeOrEnum: "", Units: "", Description: "" },
        &ResourceModel{ Id: 4, Name: "Server Public Key or Identity", Operations: OPERATION_NONE, Multiple: false, Mandatory: true, ResourceType: OBJECTTYPE_OPAQUE, RangeOrEnum: "", Units: "", Description: "" },
        &ResourceModel{ Id: 5, Name: "Secret Key", Operations: OPERATION_NONE, Multiple: false, Mandatory: true, ResourceType: OBJECTTYPE_OPAQUE, RangeOrEnum: "", Units: "", Description: "" },
        &ResourceModel{ Id: 6, Name: "SMS Security Mode", Operations: OPERATION_NONE, Multiple: false, Mandatory: true, ResourceType: OBJECTTYPE_INTEGER, RangeOrEnum: "0-255", Units: "", Description: "" },
        &ResourceModel{ Id: 7, Name: "SMS Binding Key Parameters", Operations: OPERATION_NONE, Multiple: false, Mandatory: true, ResourceType: OBJECTTYPE_OPAQUE, RangeOrEnum: "6 bytes", Units: "",  Description: "" },
        &ResourceModel{ Id: 8, Name: "SMS Binding Secret Keys", Operations: OPERATION_NONE, Multiple: false, Mandatory: true, ResourceType: OBJECTTYPE_OPAQUE, RangeOrEnum: "32-48 bytes", Units: "",  Description: "" },
        &ResourceModel{ Id: 9, Name: "LWM2M Server SMS Number", Operations: OPERATION_NONE, Multiple: false, Mandatory: true, ResourceType: OBJECTTYPE_INTEGER, RangeOrEnum: "", Units: "", Description: "" },
        &ResourceModel{ Id: 10, Name: "Short Server ID", Operations: OPERATION_NONE, Multiple: false, Mandatory: false, ResourceType: OBJECTTYPE_INTEGER, RangeOrEnum: "1-65535", Units: "", Description: "" },
        &ResourceModel{ Id: 11, Name: "Client Hold Off Time", Operations: OPERATION_NONE, Multiple: false, Mandatory: true, ResourceType: OBJECTTYPE_INTEGER, RangeOrEnum: "", Units: "s",  Description: "" },
    )

    return mr
}

type ModelsRepository struct {
    models  map[int] *ObjectModel
}

func (mr *ModelsRepository) Register(ms []ObjectModel) {

}

func (mr *ModelsRepository) GetObjectModel(n int) (*ObjectModel) {
    return mr.models[n]
}

func (mr *ModelsRepository) Add(m *ObjectModel, res ...*ResourceModel) {
    log.Println("Model", m)

    for i, o := range res {
        log.Println(i, o)
    }

    m.Resources = res

    log.Println (m)

    mr.models[m.Id] = m

    log.Println (mr.models)
}


