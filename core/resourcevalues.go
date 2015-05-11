package core

import (
    "time"
    "encoding/binary"
    "bytes"
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

func (v *StringValue) GetLength() (int) {
    return len(v.value)
}

type IntegerValue struct {
    value       int
}

func (v *IntegerValue) GetBytes() ([]byte) {
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, v.value)

    return buf.Bytes()
}

func (v *IntegerValue) GetType() (ValueTypeCode) {
    return TYPE_INTEGER
}

func (v *IntegerValue) GetValue() (interface{}) {
    return v.value
}

func (v *IntegerValue) GetLength() (int) {
    return 0
}

type TimeValue struct {
    value       time.Time
}

func (v *TimeValue) GetBytes() ([]byte) {
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, v.value.Unix())

    return buf.Bytes()
}

func (v *TimeValue) GetType() (ValueTypeCode) {
    return TYPE_TIME
}

func (v *TimeValue) GetValue() (interface{}) {
    return v.value
}

func (v *TimeValue) GetLength() (int) {
    return 0
}

type FloatValue struct {
    value       float32
}

func (v *FloatValue) GetBytes() ([]byte) {
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, v.value)

    return buf.Bytes()
}

func (v *FloatValue) GetType() (ValueTypeCode) {
    return TYPE_FLOAT
}

func (v *FloatValue) GetValue() (interface{}) {
    return v.value
}

func (v *FloatValue) GetLength() (int) {
    return 0
}

type BooleanValue struct {
    value       bool
}

func (v *BooleanValue) GetBytes() ([]byte) {
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.LittleEndian, v.value)

    return buf.Bytes()
}

func (v *BooleanValue) GetType() (ValueTypeCode) {
    return TYPE_BOOLEAN
}

func (v *BooleanValue) GetValue() (interface{}) {
    return v.value
}

func (v *BooleanValue) GetLength() (int) {
    return 0
}


type EmptyValue struct {

}

func (v *EmptyValue) GetBytes() ([]byte) {
    return []byte("")
}

func (v *EmptyValue) GetType() (ValueTypeCode) {
    return TYPE_STRING
}

func (v *EmptyValue) GetValue() (interface{}) {
    return ""
}

func (v *EmptyValue) GetLength() (int) {
    return 0
}


//
func NewStringValue(v string) ResourceValue {
    return &StringValue{
        value: v,
    }
}

func NewIntegerValue(v int) ResourceValue {
    return &IntegerValue{
        value: v,
    }
}

func NewTimeValue(v time.Time) ResourceValue {
    return &TimeValue{
        value: v,
    }
}

func NewFloatValue(v float32) ResourceValue {
    return &FloatValue{
        value: v,
    }
}

func NewBooleanValue(v bool) ResourceValue {
    return &BooleanValue{
        value: v,
    }
}

func NewEmptyValue() ResourceValue {
    return &EmptyValue{}
}
