package core

type ResponseValue interface {
    GetType() IdentifierType
}

type ObjectResponseValue struct {
    resources   []*ResourceValue
}

type ResourceValue struct {
    resourceInstances   []*ResourceInstanceValue
}

type ResourceInstanceValue struct {
    value       ResourceValueType
}
