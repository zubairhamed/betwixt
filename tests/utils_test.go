package tests

import (
    "testing"
    "github.com/zubairhamed/lwm2m/core"
    "time"
)

func TestGetValueLength(t *testing.T) {
    core.GetValueLength(-128)
    core.GetValueLength(127)
    core.GetValueLength(-32768)
    core.GetValueLength(32767)
    core.GetValueLength(-2147483648)
    core.GetValueLength(2147483647)
    core.GetValueLength(-9223372036854775808)
    core.GetValueLength(9223372036854775807)

    core.GetValueLength(-3.4E+38)
    core.GetValueLength(+3.4E+38)

    core.GetValueLength(-1.7E+308)
    core.GetValueLength(+1.7E+308)

    core.GetValueLength("this is a string")

    core.GetValueLength(true)
    core.GetValueLength(false)

    core.GetValueLength(time.Now())

    core.GetValueLength([]byte{})
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