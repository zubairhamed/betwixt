package core

import "time"

// Generic ResponseValue to return for a request
type ResponseValue interface {
    GetValue() ([]byte)
    GetValueAsTLV()
    GetValueAsPlainText()
    GetValueAsOpaque()
    GetValueAsJSON()

}

// A UTF-8 string, the minimum and/or maximum length of the String MAY be defined.
// Text: Represented as a UTF-8 string
// TLV: Represented as a UTF-8 string of Length bytes.
type StringResponseValue struct {
    value   string
}

func (rv *StringResponseValue) GetPayloadValue() ([]byte) {
    return []byte(rv.value)
}

func NewStringResponseValue (v string) (ResponseValue) {
    return &StringResponseValue{
        value: v,
    }
}

// An 8, 16, 32 or 64-bit signed integer. The valid range of the value for a Resource
// SHOULD be defined. This data type is also used for the purpose of enumeration.
// Text: Represented as an ASCII signed integer.
// TLV: Represented as a binary signed integer in network byte order, where the first (most significant) bit is 0 for a
// positive integer and 1 for a negative integer. The value may be 1 (8-bit), 2 (16-bit), 4 (32-bit) or 8 (64-bit) bytes
// long as indicated by the Length field.
type IntResponseValue struct {
    value   int
}

func (rv *IntResponseValue) GetPayloadValue() ([]byte) {
    return []byte(string(rv.value))
}

type Int8ResponseValue struct {
    value   int8
}

func (rv *Int8ResponseValue) GetPayloadValue() ([]byte) {
    return []byte(string(rv.value))
}

func NewIntResponseValue (v int) (ResponseValue) {
    return &IntResponseValue{
        value: v,
    }
}

type Int16ResponseValue struct {
    value   int16
}

func (rv *Int16ResponseValue) GetPayloadValue() ([]byte) {
    return []byte(string(rv.value))
}


type Int32ResponseValue struct {
    value   int32
}

func (rv *Int32ResponseValue) GetPayloadValue() ([]byte) {
    return []byte(string(rv.value))
}


type Int64ResponseValue struct {
    value   int64
}

func (rv *Int64ResponseValue) GetPayloadValue() ([]byte) {
    return []byte(string(rv.value))
}


// Unix Time. A signed integer representing the number of seconds since Jan 1st, 1970 in the UTC time zone.
// Text: Represented as an ASCII integer.
// TLV: Same representation as Integer.
type TimeResponseValue struct {
    value   time.Time
}

func (rv *TimeResponseValue) GetPayloadValue() ([]byte) {
    return []byte(string(rv.value.Unix()))
}

func NewTimeResponseValue (v time.Time) (ResponseValue) {
    return &TimeResponseValue{
        value: v,
    }
}

// A 32 or 64-bit floating point value. The valid range of the value for a Resource SHOULD be defined.
// Text: Represented as an ASCII signed decimal.
// TLV: Represented as an [IEEE 754-2008] [FLOAT] binary floating point value. The value may use the binary32 (4 byte
// Length) or binary64 (8 byte Length) format as indicated by the Length field.
type Float32ResponseValue struct {
    value   float32
}

type Float64ResponseValue struct {
    value   float64
}

// An integer with the value 0 for False and the value 1 for True.
// Text: Represented as the ASCII value 0 or 1.
// TLV: Represented as an Integer with value 0, or 1. The Length of a Boolean value MUST always be 1.
type BooleanResponseValue struct {

}

// A sequence of binary octets, the minimum and/or maximum length of the String MAY be defined.
// TLV: Represented as a sequence of binary data of Length bytes.
type OpaqueResponseValue struct {

}

// Object Link. The object link is used to refer an Instance of a given Object. An Object link value is composed of
// two concatenated 16-bits unsigned integers following the Network Byte Order convention. The Most Significant
// Halfword is an ObjectID, the Least Significant Hafword is an ObjectInstance ID.
// An Object Link referencing no Object Instance will contain the concatenation of 2 MAX-ID values (null link)
// Text: Represented as a UTF-8 string containing 2 16-bits ASCII integers separated by a ‘:’ ASCII character.
// TLV: Represented as a UTF-8 string containing 2 16-bits ASCII integers separated by a ‘:’ ASCII character.
type ObjectLinkResponseValue struct {

}

// In cases where a request doesn't warrant a response
type NoResponseValue struct {

}

func (rv *NoResponseValue) GetPayloadValue() ([]byte) {
    return nil
}

func NoResponse() *NoResponseValue {
    return &NoResponseValue{}
}

