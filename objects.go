package betwixt

import (
	"bytes"
)

// DefaultObjectDefinition
type DefaultObjectDefinition struct {
	Id          LWM2MObjectType
	Name        string
	Description string
	Multiple    bool
	Mandatory   bool
	Resources   []ResourceDefinition
}

func (o *DefaultObjectDefinition) GetName() string {
	return o.Name
}

func (o *DefaultObjectDefinition) GetType() LWM2MObjectType {
	return o.Id
}

func (o *DefaultObjectDefinition) AllowMultiple() bool {
	return o.Multiple
}

func (o *DefaultObjectDefinition) IsMandatory() bool {
	return o.Mandatory
}

func (o *DefaultObjectDefinition) GetDescription() string {
	return o.Description
}

func (o *DefaultObjectDefinition) SetResources(r []ResourceDefinition) {
	o.Resources = r
}

func (o *DefaultObjectDefinition) GetResource(n uint16) ResourceDefinition {
	for _, rsrc := range o.Resources {
		if rsrc.GetId() == n {
			return rsrc
		}
	}
	return nil
}

func (o *DefaultObjectDefinition) GetResources() []ResourceDefinition {
	return o.Resources
}

// DefaultObjectInstance
type DefaultObjectInstance struct {
	Id     int
	TypeId LWM2MObjectType
}

func (o *DefaultObjectInstance) GetId() int {
	return o.Id
}

func (o *DefaultObjectInstance) GetTypeId() LWM2MObjectType {
	return o.TypeId
}

// DefaultObject
func NewObject(t LWM2MObjectType, enabler ObjectEnabler, reg Registry) Object {
	def := reg.GetDefinition(t)

	if enabler == nil {
		enabler = NewNullEnabler()
	}

	return &DefaultObject{
		definition: def,
		typeId:     t,
		enabler:    enabler,
		instances:  make(map[int]bool),
	}
}

type DefaultObject struct {
	typeId     LWM2MObjectType
	definition ObjectDefinition
	enabler    ObjectEnabler
	instances  map[int]bool
}

func (o *DefaultObject) GetDefinition() ObjectDefinition {
	return o.definition
}

func (o *DefaultObject) GetType() LWM2MObjectType {
	return o.typeId
}

func (o *DefaultObject) GetEnabler() ObjectEnabler {
	return o.enabler
}

func (o *DefaultObject) AddInstance(n int) {
	o.instances[n] = true
}

func (o *DefaultObject) RemoveInstance(n int) {
	o.instances[n] = false
}

func (o *DefaultObject) GetInstances() []int {
	instances := []int{}

	for k, v := range o.instances {
		if v {
			instances = append(instances, k)
		}
	}
	return instances
}

func (o *DefaultObject) SetEnabler(e ObjectEnabler) {
	o.enabler = e
}

type DefaultResourceDefinition struct {
	Id             uint16
	Name           string
	Operations     OperationCode
	Multiple       bool
	Mandatory      bool
	ResourceType   ValueTypeCode
	Units          string
	RangeOrEnums   string
	Description    string
	ValueValidator Validator
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

type ObjectValue struct {
	instanceId uint16
	typeId     LWM2MObjectType
	resources  []Value
}

func NewResourceValue(id uint16, value Value) Value {
	return &ResourceValue{
		id:    id,
		value: value,
	}
}

type ResourceValue struct {
	id    uint16
	value Value
}

func (v ResourceValue) GetId() uint16 {
	return v.id
}

func (v ResourceValue) GetBytes() []byte {
	return v.value.GetBytes()
}

func (v ResourceValue) GetContainedType() ValueTypeCode {
	return VALUETYPE_RESOURCE
}

func (v ResourceValue) GetType() ValueTypeCode {
	return VALUETYPE_RESOURCE
}

func (v ResourceValue) GetStringValue() string {
	return v.value.GetStringValue()
}

func (v ResourceValue) GetValue() interface{} {
	return v.value.GetValue()
}

func NewMultipleResourceValue(id uint16, value []*ResourceValue) Value {
	return &MultipleResourceValue{
		id:        id,
		instances: value,
	}
}

type MultipleResourceValue struct {
	id        uint16
	instances []*ResourceValue
}

func (v MultipleResourceValue) GetBytes() []byte {
	return []byte{}
}

func (v MultipleResourceValue) GetContainedType() ValueTypeCode {
	return VALUETYPE_RESOURCE
}

func (v MultipleResourceValue) GetType() ValueTypeCode {
	return VALUETYPE_MULTIRESOURCE
}

func (v MultipleResourceValue) GetStringValue() string {
	var buf bytes.Buffer

	for _, res := range v.instances {
		buf.WriteString(res.GetStringValue())
		buf.WriteString(",")
	}
	return buf.String()
}

func (v MultipleResourceValue) GetValue() interface{} {
	return v.instances
}
