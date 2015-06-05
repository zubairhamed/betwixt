package api

type LWM2MObjectType int
type LWM2MObjectInstances map[LWM2MObjectType]Object

type FnOnStartup func()
type FnOnRead func()
type FnOnWrite func()
type FnOnExecute func()
type FnOnRegistered func(string)
type FnOnDeregistered func()
type FnOnError func()

type ValueTypeCode byte
type OperationCode int
type IdentifierType byte
type BindingMode string
type OperationType byte
