package core
import "time"

type ResponseValue interface {
    GetPayloadValue() ([]byte)
}


type StringResponseValue struct {
    value   string
}

func (rv *StringResponseValue) GetPayloadValue() ([]byte) {
    return []byte(rv.value)
}

type IntResponseValue struct {
    value   int
}

type TimeResponseValue struct {
    value   time.Time
}

type NoResponseValue struct {

}

func (rv *NoResponseValue) GetPayloadValue() ([]byte) {
    return nil
}

func NoResponse() *NoResponseValue {
    return &NoResponseValue{}
}

func NewStringResponseValue (v string) (ResponseValue) {
    return &StringResponseValue{
        value: v,
    }
}

func NewIntResponseValue (v int) (ResponseValue) {
    return &IntResponseValue{
        value: v,
    }
}

func NewTimeResponseValue (v time.Time) (ResponseValue) {
    return &TimeResponseValue{
        value: v,
    }
}




