package resources

import (
	. "github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/betwixt/core/values/validators"
	"github.com/zubairhamed/go-commons/typeval"
)

type DefaultResourceDefinition struct {
	Id           	uint16
	Name         	string
	Operations   	OperationCode
	Multiple     	bool
	Mandatory    	bool
	ResourceType 	typeval.ValueTypeCode
	Units        	string
	RangeOrEnums 	string
	Description  	string
	ValueValidator 	validators.Validator
}

func (o *DefaultResourceDefinition) GetId() uint16 {
	return o.Id
}

func (o *DefaultResourceDefinition) GetOperations() OperationCode {
	return o.Operations
}

func (o *DefaultResourceDefinition) MultipleValuesAllowed() bool {
	return o.Multiple
}

func (o *DefaultResourceDefinition) GetResourceType() typeval.ValueTypeCode {
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
