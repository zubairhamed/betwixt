package tests

import (
    "testing"
    "github.com/zubairhamed/lwm2m/core"
    "time"
)

func TestGetValueByteLength(t *testing.T) {
    core.GetValueByteLength(-128)
    core.GetValueByteLength(127)
    core.GetValueByteLength(-32768)
    core.GetValueByteLength(32767)
    core.GetValueByteLength(-2147483648)
    core.GetValueByteLength(2147483647)
    core.GetValueByteLength(-9223372036854775808)
    core.GetValueByteLength(9223372036854775807)

    core.GetValueByteLength(-3.4E+38)
    core.GetValueByteLength(+3.4E+38)

    core.GetValueByteLength(-1.7E+308)
    core.GetValueByteLength(+1.7E+308)

    core.GetValueByteLength("this is a string")

    core.GetValueByteLength(true)
    core.GetValueByteLength(false)

    core.GetValueByteLength(time.Now())

    core.GetValueByteLength([]byte{})

    core.GetValueByteLength(uint(1))
    core.GetValueByteLength(uint16(1))
    core.GetValueByteLength(uint32(1))
    core.GetValueByteLength(uint64(1))

}


/*
if _, ok := val.(float32); ok {
log.Println("float32")
}

if _, ok := val.(float64); ok {
log.Println("float64")
}

if _, ok := val.(time.Time); ok {
log.Println("time")
}
*/