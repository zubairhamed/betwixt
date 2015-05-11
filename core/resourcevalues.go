package core

import (
    "time"
    "encoding/binary"
)



type ResourceValue interface {
    GetBytes() []byte
    GetType() ValueTypeCode
    GetValue()  interface{}
    GetLength() int
}

type StringValue struct {
    value       string
}

func (v *StringValue) GetBytes() ([]byte) {
    return []byte(v.value)
}

func (v *StringValue) GetType() (ValueTypeCode) {
    return TYPE_STRING
}

func (v *StringValue) GetValue() (interface{}) {
    return v.value
}

type IntegerValue struct {
    value       int
}

func (v *IntegerValue) GetBytes() ([]byte) {
    bs := make([]byte, 4)
    binary.LittleEndian.PutUint16(bs, v.value)

    return bs
}

func (v *IntegerValue) GetType() (ValueTypeCode) {
    return TYPE_INTEGER
}

func (v *IntegerValue) GetValue() (interface{}) {
    return v.value
}

type TimeValue struct {
    value       time.Time
}

func (v *TimeValue) GetBytes() ([]byte) {
    return []byte(v.value)
}

func (v *TimeValue) GetType() (ValueTypeCode) {
    return TYPE_TIME
}

func (v *TimeValue) GetValue() (interface{}) {
    return v.value
}

type FloatValue struct {
    value       float32
}

func (v *FloatValue) GetBytes() ([]byte) {
    return []byte(v.value)
}

func (v *FloatValue) GetType() (ValueTypeCode) {
    return TYPE_FLOAT
}

func (v *FloatValue) GetValue() (interface{}) {
    return v.value
}

type BooleanValue struct {
    value       bool
}

func (v *BooleanValue) GetBytes() ([]byte) {
    return []byte(v.value)
}

func (v *BooleanValue) GetType() (ValueTypeCode) {
    return TYPE_BOOLEAN
}

func (v *BooleanValue) GetValue() (interface{}) {
    return v.value
}

type EmptyResponseValue struct {

}

func (v *EmptyResponseValue) GetBytes() ([]byte) {
    return []byte("")
}

func (v *EmptyResponseValue) GetType() (ValueTypeCode) {
    return TYPE_STRING
}

func (v *EmptyResponseValue) GetValue() (interface{}) {
    return ""
}

//
func NewStringValue(v string) ResourceValue {
    return &StringValue{
        value: v,
    }
}

func NewIntegerResponseValue(v int) ResourceValue {
    return &IntegerValue{
        value: v,
    }
}

func NewTimeResponseValue(v time.Time) ResourceValue {
    return &TimeValue{
        value: v,
    }
}

func NewFloatResponseValue(v float32) ResourceValue {
    return &FloatValue{
        value: v,
    }
}

func NewBooleanResponseValue(v bool) ResourceValue {
    return &BooleanValue{
        value: v,
    }
}

func NoResponse() ResourceValue {
    return &NoResponseValue{}
}
