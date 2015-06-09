package resources

import (
	. "github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/betwixt/utils"
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

func (o *DefaultResourceDefinition) GetName() string {
	return o.Name
}

func (o *DefaultResourceDefinition) GetDescription() string {
	return o.Description
}

func (o *DefaultResourceDefinition) GetUnits() string {
	return o.Units
}

func (o *DefaultResourceDefinition) GetRangeOrEnums() string {
	return o.RangeOrEnums
}

func (o *DefaultResourceDefinition) IsMandatory() bool {
	return o.Mandatory
}

func (o *DefaultResourceDefinition) IsExecutable() bool {
	return utils.IsExecutableResource(o)
}

func (o *DefaultResourceDefinition) IsWritable() bool {
	return utils.IsWritableResource(o)
}

func (o *DefaultResourceDefinition) IsReadable() bool {
	return utils.IsReadableResource(o)
}
