package core

import "time"

// Generic PayloadValue to return for a request
type PayloadValue interface {
    GetValue() ([]byte)
    //GetValueAsTLV()
//    GetValueAsPlainText()
//    GetValueAsOpaque()
    //GetValueAsJSON()
}

// A UTF-8 string, the minimum and/or maximum length of the String MAY be defined.
// Text: Represented as a UTF-8 string
// TLV: Represented as a UTF-8 string of Length bytes.
type StringValue struct {
    value   string
}

func (rv *StringValue) GetValue() ([]byte) {
    return []byte(rv.value)
}

func NewStringValue (v string) (PayloadValue) {
    return &StringValue{
        value: v,
    }
}

// An 8, 16, 32 or 64-bit signed integer. The valid range of the value for a Resource
// SHOULD be defined. This data type is also used for the purpose of enumeration.
// Text: Represented as an ASCII signed integer.
// TLV: Represented as a binary signed integer in network byte order, where the first (most significant) bit is 0 for a
// positive integer and 1 for a negative integer. The value may be 1 (8-bit), 2 (16-bit), 4 (32-bit) or 8 (64-bit) bytes
// long as indicated by the Length field.
type IntegerValue struct {
    value   int
}

func (rv *IntegerValue) GetPayloadValue() ([]byte) {
    return []byte(string(rv.value))
}

type Integer8Value struct {
    value   int8
}

func (rv *Integer8Value) GetValue() ([]byte) {
    return []byte(string(rv.value))
}

func NewIntegerValue (v int) (PayloadValue) {
    return &IntegerValue{
        value: v,
    }
}

type Integer16Value struct {
    value   int16
}

func (rv *Integer16Value) GetValue() ([]byte) {
    return []byte(string(rv.value))
}


type Integer32Value struct {
    value   int32
}

func (rv *Integer32Value) GetValue() ([]byte) {
    return []byte(string(rv.value))
}


type Integer64Value struct {
    value   int64
}

func (rv *Integer64Value) GetValue() ([]byte) {
    return []byte(string(rv.value))
}


// Unix Time. A signed integer representing the number of seconds since Jan 1st, 1970 in the UTC time zone.
// Text: Represented as an ASCII integer.
// TLV: Same representation as Integer.
type TimeValue struct {
    value   time.Time
}

func (rv *TimeValue) GetValue() ([]byte) {
    return []byte(string(rv.value.Unix()))
}

func NewTimePayloadValue (v time.Time) (PayloadValue) {
    return &TimeValue{
        value: v,
    }
}

// A 32 or 64-bit floating point value. The valid range of the value for a Resource SHOULD be defined.
// Text: Represented as an ASCII signed decimal.
// TLV: Represented as an [IEEE 754-2008] [FLOAT] binary floating point value. The value may use the binary32 (4 byte
// Length) or binary64 (8 byte Length) format as indicated by the Length field.
type Float32Value struct {
    value   float32
}

type Float64Value struct {
    value   float64
}

// An integer with the value 0 for False and the value 1 for True.
// Text: Represented as the ASCII value 0 or 1.
// TLV: Represented as an Integer with value 0, or 1. The Length of a Boolean value MUST always be 1.
type BooleanValue struct {

}

// A sequence of binary octets, the minimum and/or maximum length of the String MAY be defined.
// TLV: Represented as a sequence of binary data of Length bytes.
type OpaqueValue struct {

}

// Object Link. The object link is used to refer an Instance of a given Object. An Object link value is composed of
// two concatenated 16-bits unsigned integers following the Network Byte Order convention. The Most Significant
// Halfword is an ObjectID, the Least Significant Hafword is an ObjectInstance ID.
// An Object Link referencing no Object Instance will contain the concatenation of 2 MAX-ID values (null link)
// Text: Represented as a UTF-8 string containing 2 16-bits ASCII integers separated by a ‘:’ ASCII character.
// TLV: Represented as a UTF-8 string containing 2 16-bits ASCII integers separated by a ‘:’ ASCII character.
type ObjectLinkPayloadValue struct {

}

// In cases where a request doesn't warrant a response
type NoPayloadValue struct {

}

func (rv *NoPayloadValue) GetPayloadValue() ([]byte) {
    return nil
}

func NoResponse() *NoPayloadValue {
    return &NoPayloadValue{}
}

