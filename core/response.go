package core

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

}

type CoreResponseValue struct {

}

func NewStringResponseValue (s string) (ResponseValue) {
    return &StringResponseValue{
        value: s,
    }
}