package core

import (
	. "github.com/zubairhamed/betwixt/api"
	"github.com/zubairhamed/betwixt/core/response"
)

type DefaultObjectModel struct {
	Id          LWM2MObjectType
	Name        string
	Description string
	Multiple    bool
	Mandatory   bool
	Resources   []ResourceModel
}

func (o *DefaultObjectModel) GetType() LWM2MObjectType {
	return o.Id
}

func (o *DefaultObjectModel) AllowMultiple() bool {
	return o.Multiple
}

func (o *DefaultObjectModel) IsMandatory() bool {
	return o.Mandatory
}

func (o *DefaultObjectModel) GetDescription() string {
	return o.Description
}

func (o *DefaultObjectModel) SetResources(r []ResourceModel) {
	o.Resources = r
}

func (o *DefaultObjectModel) GetResource(n int) ResourceModel {
	for _, rsrc := range o.Resources {
		if rsrc.GetId() == n {
			return rsrc
		}
	}
	return nil
}

func (o *DefaultObjectModel) GetResources() []ResourceModel {
	return o.Resources
}

type DefaultResourceModel struct {
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

func (o *DefaultResourceModel) GetId() int {
	return o.Id
}

func (o *DefaultResourceModel) GetOperations() OperationCode {
	return o.Operations
}

func (o *DefaultResourceModel) MultipleValuesAllowed() bool {
	return o.Multiple
}

func (o *DefaultResourceModel) GetResourceType() ValueTypeCode {
	return o.ResourceType
}

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

type DefaultResource struct {
	Id int
}

type DefaultObject struct {
	typeId 		LWM2MObjectType
	model 		ObjectModel
	enabler 	ObjectEnabler
	instances 	map[int]bool
}

func (o *DefaultObject) GetModel() ObjectModel {
	return o.model
}

func (o *DefaultObject) GetType() LWM2MObjectType {
	return o.typeId
}

func (o *DefaultObject) GetEnabler()ObjectEnabler {
	return o.enabler
}

func (o *DefaultObject) AddInstance(n int) {
	o.instances[n] = true
}

func (o *DefaultObject) RemoveInstance(n int) {
	o.instances[n] = false
}

func (o *DefaultObject) GetInstances()[]int {
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

func NewObject(t LWM2MObjectType, enabler ObjectEnabler, reg Registry) Object {
	model := reg.GetModel(t)

	if enabler == nil {
		enabler = &NullEnabler{}
	}

	return &DefaultObject {
		model: model,
		typeId: t,
		enabler: enabler,
		instances: make(map[int]bool),
	}
}

type NullEnabler struct {

}

func (e *NullEnabler) OnRead(int, int, Lwm2mRequest) Lwm2mResponse {
	return response.MethodNotAllowed()
}

func (e *NullEnabler) OnDelete(int, Lwm2mRequest) Lwm2mResponse {
	return response.MethodNotAllowed()
}

func (e *NullEnabler) OnWrite(int, int, Lwm2mRequest) Lwm2mResponse {
	return response.MethodNotAllowed()
}

func (e *NullEnabler) OnCreate(int, int, Lwm2mRequest) Lwm2mResponse {
	return response.MethodNotAllowed()
}

func (e *NullEnabler) OnExecute(int, int, Lwm2mRequest) Lwm2mResponse {
	return response.MethodNotAllowed()
}
