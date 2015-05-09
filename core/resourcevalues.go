package core

import "time"

type ResourceValueType interface {
    GetValue()  interface{}
}

type StringValue struct {
    value       string
}

func (v *StringValue) GetValue() (string) {
    return v.value
}

type IntegerValue struct {
    value       int
}

type TimeValue struct {
    value       time.Time
}

type FloatValue struct {
    value       float32
}

type BooleanValue struct {
    value       bool
}

type NoResponseValue struct {

}

//
func NewStringValue(v string) ResourceValueType {
    return &StringValue{
        value: v,
    }
}

func NewIntegerResponseValue(v int) ResourceValueType {
    return &IntegerValue{
        value: v,
    }
}

func NewTimeResponseValue(v time.Time) ResponseValue {
    return &TimeValue{
        value: v,
    }
}

func NewFloatResponseValue(v float32) ResponseValue {
    return &FloatValue{
        value: v,
    }
}

func NewBooleanResponseValue(v bool) ResponseValue {
    return &BooleanValue{
        value: v,
    }
}

func NoResponse() ResponseValue {
    return &NoResponseValue{}
}
