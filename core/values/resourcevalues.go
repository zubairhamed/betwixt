package values

import (
	"bytes"
	"encoding/binary"
	. "github.com/zubairhamed/go-lwm2m/api"
	"strconv"
	"time"
)

type MultipleResourceInstanceValue struct {
	values []ResponseValue
}

func (v *MultipleResourceInstanceValue) GetBytes() []byte {
	return []byte("")
}

func (v *MultipleResourceInstanceValue) GetType() ValueTypeCode {
	return VALUETYPE_MULTIPLE
}

func (v *MultipleResourceInstanceValue) GetValue() interface{} {
	return v.values
}

func (v *MultipleResourceInstanceValue) GetStringValue() string {
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
	return []byte(strconv.Itoa(v.value))
}

func (v *IntegerValue) GetType() ValueTypeCode {
	return VALUETYPE_INTEGER
}

func (v *IntegerValue) GetValue() interface{} {
	return v.value
}

func (v *IntegerValue) GetStringValue() string {
	return string(v.value)
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

func (v *TimeValue) GetValue() interface{} {
	return v.value
}

func (v *TimeValue) GetStringValue() string {
	return strconv.FormatInt(v.value.Unix(), 10)
}

type FloatValue struct {
	value float64
}

func (v *FloatValue) GetBytes() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, v.value)

	return buf.Bytes()
}

func (v *FloatValue) GetType() ValueTypeCode {
	return VALUETYPE_FLOAT
}

func (v *FloatValue) GetValue() interface{} {
	return v.value
}

func (v *FloatValue) GetStringValue() string {
	return strconv.FormatFloat(v.value, 'g', 1, 32)
}

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

type EmptyValue struct {
}

func (v *EmptyValue) GetBytes() []byte {
	return []byte("")
}

func (v *EmptyValue) GetType() ValueTypeCode {
	return VALUETYPE_STRING
}

func (v *EmptyValue) GetValue() interface{} {
	return ""
}

func (v *EmptyValue) GetStringValue() string {
	return ""
}

func String(v ...string) ResponseValue {
	if len(v) > 1 {
		vs := []ResponseValue{}

		for _, o := range v {
			vs = append(vs, String(o))
		}
		return MultipleResourceInstance(vs)
	} else {
		return &StringValue{
			value: v[0],
		}
	}
}

func Integer(v ...int) ResponseValue {
	if len(v) > 1 {
		vs := []ResponseValue{}

		for _, o := range v {
			vs = append(vs, Integer(o))
		}
		return MultipleResourceInstance(vs)
	} else {
		return &IntegerValue{
			value: v[0],
		}
	}
}

func Time(v ...time.Time) ResponseValue {
	if len(v) > 1 {
		vs := []ResponseValue{}

		for _, o := range v {
			vs = append(vs, Time(o))
		}
		return MultipleResourceInstance(vs)
	} else {
		return &TimeValue{
			value: v[0],
		}
	}
}

func Float(v ...float64) ResponseValue {
	if len(v) > 1 {
		vs := []ResponseValue{}

		for _, o := range v {
			vs = append(vs, Float(o))
		}
		return MultipleResourceInstance(vs)
	} else {
		return &FloatValue{
			value: v[0],
		}
	}
}

func Boolean(v ...bool) ResponseValue {
	if len(v) > 1 {
		vs := []ResponseValue{}

		for _, o := range v {
			vs = append(vs, Boolean(o))
		}
		return MultipleResourceInstance(vs)
	} else {
		return &BooleanValue{
			value: v[0],
		}
	}
}

func Empty() ResponseValue {
	return &EmptyValue{}
}

func MultipleResourceInstance(v []ResponseValue) ResponseValue {
	return &MultipleResourceInstanceValue{
		values: v,
	}
}

////////////////////////////////////////////////////////////////////
func Tlv(b []byte) ResponseValue {
	return &TlvValue{
		content: b,
	}
}

type TlvValue struct {
	content []byte
}

func (p *TlvValue) GetBytes() []byte {
	return p.content
}

func (p *TlvValue) Length() int {
	return len(p.content)
}

func (p *TlvValue) String() string {
	return ""
}

func (p *TlvValue) GetStringValue() string {
	return ""
}

func (v *TlvValue) GetType() ValueTypeCode {
	return VALUETYPE_TLV
}

func (v *TlvValue) GetValue() interface{} {
	return v.content
}
