package betwixt

import (
	"bytes"
	"encoding/binary"
	"errors"
	"strconv"
	"time"
)

type ValueTypeCode byte

const (
	VALUETYPE_EMPTY         ValueTypeCode = 0
	VALUETYPE_MULTIPLE      ValueTypeCode = 1
	VALUETYPE_STRING        ValueTypeCode = 2
	VALUETYPE_BYTE          ValueTypeCode = 3
	VALUETYPE_INTEGER       ValueTypeCode = 4
	VALUETYPE_INTEGER32     ValueTypeCode = 5
	VALUETYPE_INTEGER64     ValueTypeCode = 6
	VALUETYPE_FLOAT         ValueTypeCode = 7
	VALUETYPE_FLOAT64       ValueTypeCode = 8
	VALUETYPE_BOOLEAN       ValueTypeCode = 9
	VALUETYPE_OPAQUE        ValueTypeCode = 10
	VALUETYPE_TIME          ValueTypeCode = 11
	VALUETYPE_OBJECTLINK    ValueTypeCode = 12
	VALUETYPE_OBJECT        ValueTypeCode = 13
	VALUETYPE_RESOURCE      ValueTypeCode = 14
	VALUETYPE_MULTIRESOURCE ValueTypeCode = 15
)

// ResponseValue interface represents response to a server request
// Typical response could be plain text, TLV Binary, TLV JSON
type Value interface {
	GetBytes() []byte
	GetType() ValueTypeCode
	GetContainedType() ValueTypeCode
	GetValue() interface{}
	GetStringValue() string
}

type MultipleValue struct {
	values        []Value
	containedType ValueTypeCode
}

func (v *MultipleValue) GetContainedType() ValueTypeCode {
	return v.containedType
}

func (v *MultipleValue) GetBytes() []byte {
	return []byte("")
}

func (v *MultipleValue) GetType() ValueTypeCode {
	return VALUETYPE_MULTIPLE
}

func (v *MultipleValue) GetValue() interface{} {
	return v.values
}

func (v *MultipleValue) GetStringValue() string {
	return ""
}

type StringValue struct {
	value string
}

func (v *StringValue) GetBytes() []byte {
	return []byte(v.value)
}

func (v *StringValue) GetType() ValueTypeCode {
	return VALUETYPE_STRING
}

func (v *StringValue) GetContainedType() ValueTypeCode {
	return VALUETYPE_STRING
}

func (v *StringValue) GetValue() interface{} {
	return v.value
}

func (v *StringValue) GetStringValue() string {
	return v.value
}

type IntegerValue struct {
	value int
}

func (v *IntegerValue) GetBytes() []byte {
	sz, _ := GetValueByteLength(v.value)
	buf := new(bytes.Buffer)
	if sz == 1 {
		binary.Write(buf, binary.LittleEndian, int8(v.value))
	} else if sz == 2 {
		binary.Write(buf, binary.LittleEndian, int16(v.value))
	} else if sz == 4 {
		binary.Write(buf, binary.LittleEndian, int32(v.value))
	} else if sz == 8 {
		binary.Write(buf, binary.LittleEndian, int64(v.value))
	}
	return buf.Bytes()
}

func (v *IntegerValue) GetType() ValueTypeCode {
	return VALUETYPE_INTEGER
}

func (v *IntegerValue) GetContainedType() ValueTypeCode {
	return VALUETYPE_INTEGER
}

func (v *IntegerValue) GetValue() interface{} {
	return v.value
}

func (v *IntegerValue) GetStringValue() string {
	return strconv.Itoa(v.value)
}

type TimeValue struct {
	value time.Time
}

func (v *TimeValue) GetBytes() []byte {
	return []byte(strconv.FormatInt(v.value.Unix(), 10))
}

func (v *TimeValue) GetType() ValueTypeCode {
	return VALUETYPE_TIME
}

func (v *TimeValue) GetContainedType() ValueTypeCode {
	return VALUETYPE_TIME
}

func (v *TimeValue) GetValue() interface{} {
	return v.value
}

func (v *TimeValue) GetStringValue() string {
	return strconv.FormatInt(v.value.Unix(), 10)
}

type FloatValue struct {
	value float32
}

func (v *FloatValue) GetBytes() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, v.value)

	return buf.Bytes()
}

func (v *FloatValue) GetType() ValueTypeCode {
	return VALUETYPE_FLOAT
}

func (v *FloatValue) GetContainedType() ValueTypeCode {
	return VALUETYPE_FLOAT
}

func (v *FloatValue) GetValue() interface{} {
	return v.value
}

func (v *FloatValue) GetStringValue() string {
	return strconv.FormatFloat(float64(v.value), 'g', 1, 32)
}

// Float64
type Float64Value struct {
	value float64
}

func (v *Float64Value) GetBytes() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, v.value)

	return buf.Bytes()
}

func (v *Float64Value) GetType() ValueTypeCode {
	return VALUETYPE_FLOAT64
}

func (v *Float64Value) GetContainedType() ValueTypeCode {
	return VALUETYPE_FLOAT64
}

func (v *Float64Value) GetValue() interface{} {
	return v.value
}

func (v *Float64Value) GetStringValue() string {
	return strconv.FormatFloat(v.value, 'g', 1, 32)
}

// Boolean
type BooleanValue struct {
	value bool
}

func (v *BooleanValue) GetBytes() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, v.value)

	return buf.Bytes()
}

func (v *BooleanValue) GetType() ValueTypeCode {
	return VALUETYPE_BOOLEAN
}

func (v *BooleanValue) GetContainedType() ValueTypeCode {
	return VALUETYPE_BOOLEAN
}

func (v *BooleanValue) GetValue() interface{} {
	return v.value
}

func (v *BooleanValue) GetStringValue() string {
	if v.value {
		return "1"
	} else {
		return "0"
	}
}

func Empty() Value {
	return &EmptyValue{}
}

type EmptyValue struct {
}

func (v *EmptyValue) GetBytes() []byte {
	return []byte("")
}

func (v *EmptyValue) GetType() ValueTypeCode {
	return VALUETYPE_EMPTY
}

func (v *EmptyValue) GetContainedType() ValueTypeCode {
	return VALUETYPE_EMPTY
}

func (v *EmptyValue) GetValue() interface{} {
	return ""
}

func (v *EmptyValue) GetStringValue() string {
	return ""
}

func String(v ...string) Value {
	if len(v) > 1 {
		vs := []Value{}

		for _, o := range v {
			vs = append(vs, String(o))
		}
		return Multiple(VALUETYPE_STRING, vs...)
	} else {
		return &StringValue{
			value: v[0],
		}
	}
}

func Integer(v ...int) Value {
	if len(v) > 1 {
		vs := []Value{}

		for _, o := range v {
			vs = append(vs, Integer(o))
		}
		return Multiple(VALUETYPE_INTEGER, vs...)
	} else {
		return &IntegerValue{
			value: v[0],
		}
	}
}

func Time(v ...time.Time) Value {
	if len(v) > 1 {
		vs := []Value{}

		for _, o := range v {
			vs = append(vs, Time(o))
		}
		return Multiple(VALUETYPE_TIME, vs...)
	} else {
		return &TimeValue{
			value: v[0],
		}
	}
}

func Float(v ...float32) Value {
	if len(v) > 1 {
		vs := []Value{}

		for _, o := range v {
			vs = append(vs, Float(o))
		}
		return Multiple(VALUETYPE_FLOAT, vs...)
	} else {
		return &FloatValue{
			value: v[0],
		}
	}
}

func Float64(v ...float64) Value {
	if len(v) > 1 {
		vs := []Value{}

		for _, o := range v {
			vs = append(vs, Float64(o))
		}
		return Multiple(VALUETYPE_FLOAT64, vs...)
	} else {
		return &Float64Value{
			value: v[0],
		}
	}
}

func Boolean(v ...bool) Value {
	if len(v) > 1 {
		vs := []Value{}

		for _, o := range v {
			vs = append(vs, Boolean(o))
		}
		return Multiple(VALUETYPE_BOOLEAN, vs...)
	} else {
		return &BooleanValue{
			value: v[0],
		}
	}
}

func Multiple(ct ValueTypeCode, v ...Value) Value {
	return &MultipleValue{
		values:        v,
		containedType: ct,
	}
}

func MultipleIntegers(v ...Value) Value {
	return &MultipleValue{
		values:        v,
		containedType: VALUETYPE_INTEGER,
	}
}

func ValueByType(t ValueTypeCode, val []byte) Value {
	var value Value

	switch t {
	case VALUETYPE_STRING:
		value = String(string(val))
		break
	}

	return value
}

func GetValueByteLength(val interface{}) (uint32, error) {
	if _, ok := val.(int); ok {
		v := val.(int)
		if v > 127 || v < -128 {
			if v > 32767 || v < -32768 {
				if v > 2147483647 || v < -2147483648 {
					return 8, nil
				} else {
					return 4, nil
				}
			} else {
				return 2, nil
			}
		} else {
			return 1, nil
		}
	} else if _, ok := val.(bool); ok {
		return 1, nil
	} else if _, ok := val.(string); ok {
		v := val.(string)

		return uint32(len(v)), nil
	} else if _, ok := val.(float64); ok {
		v := val.(float64)

		if v > +3.4E+38 || v < -3.4E+38 {
			return 8, nil
		} else {
			return 4, nil
		}
	} else if _, ok := val.(time.Time); ok {
		return 8, nil
	} else if _, ok := val.([]byte); ok {
		v := val.([]byte)
		return uint32(len(v)), nil
	} else {
		return 0, errors.New("Unknown type")
	}
}

func BytesToIntegerValue(b []byte) (conv Value) {
	intLen := len(b)

	if intLen == 1 {
		conv = Integer(int(b[0]))
	} else if intLen == 2 {
		conv = Integer(int(b[1]) | (int(b[0]) << 8))
	} else if intLen == 4 {
		conv = Integer(int(b[3]) | (int(b[2]) << 8) | (int(b[1]) << 16) | (int(b[0]) << 24))
	} else if intLen == 8 {
		conv = Integer(int(b[7]) | (int(b[6]) << 8) | (int(b[5]) << 16) | (int(b[4]) << 24) | (int(b[3]) << 32) | (int(b[2]) << 40) | (int(b[1]) << 48) | (int(b[0]) << 56))
	} else {
		// Error
	}
	return
}
