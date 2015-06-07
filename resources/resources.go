package resources

import (
    . "github.com/zubairhamed/betwixt"
)

type DefaultResourceDefinition struct {
    Id           int
    Name         string
    Operations   OperationCode
    Multiple     bool
    Mandatory    bool
    ResourceType ValueTypeCode
    Units        string
    RangeOrEnums string
    Description  string
}

func (o *DefaultResourceDefinition) GetId() int {
    return o.Id
}

func (o *DefaultResourceDefinition) GetOperations() OperationCode {
    return o.Operations
}

func (o *DefaultResourceDefinition) MultipleValuesAllowed() bool {
    return o.Multiple
}

func (o *DefaultResourceDefinition) GetResourceType() ValueTypeCode {
    return o.ResourceType
}

// DefaultResource
type DefaultResource struct {
    Id int
}