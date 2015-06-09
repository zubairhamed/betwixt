package objects

import (
	. "github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/betwixt/enablers"
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

func (o *DefaultObjectDefinition) GetResource(n int) ResourceDefinition {
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
		enabler = enablers.NewNullEnabler()
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
