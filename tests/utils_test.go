package tests

import (
    "testing"
    "github.com/zubairhamed/lwm2m/core"
    "time"
    "github.com/stretchr/testify/assert"
)

func TestGetValueByteLength(t *testing.T) {
    var v uint32
    var err error

    v, _ = core.GetValueByteLength(-128)
    assert.Equal(t, v, uint32(1), "Wrong type size returned")

    v, _ = core.GetValueByteLength(127)
    assert.Equal(t, v, uint32(1), "Wrong type size returned")

    v, _ = core.GetValueByteLength(-32768)
    assert.Equal(t, v, uint32(2), "Wrong type size returned")

    v, _ = core.GetValueByteLength(32767)
    assert.Equal(t, v, uint32(2), "Wrong type size returned")

    v, _ = core.GetValueByteLength(-2147483648)
    assert.Equal(t, v, uint32(4), "Wrong type size returned")

    v, _ = core.GetValueByteLength(2147483647)
    assert.Equal(t, v, uint32(4), "Wrong type size returned")

    v, _ = core.GetValueByteLength(-9223372036854775808)
    assert.Equal(t, v, uint32(8), "Wrong type size returned")

    v, _ = core.GetValueByteLength(9223372036854775807)
    assert.Equal(t, v, uint32(8), "Wrong type size returned")

    v, _ = core.GetValueByteLength(-3.4E+38)
    assert.Equal(t, v, uint32(4), "Wrong type size returned")

    v, _ = core.GetValueByteLength(+3.4E+38)
    assert.Equal(t, v, uint32(4), "Wrong type size returned")

    v, _ = core.GetValueByteLength(-1.7E+308)
    assert.Equal(t, v, uint32(8), "Wrong type size returned")

    v, _ = core.GetValueByteLength(+1.7E+308)
    assert.Equal(t, v, uint32(8), "Wrong type size returned")

    v, _ = core.GetValueByteLength("this is a string")
    assert.Equal(t, v, uint32(16), "Wrong type size returned")

    v, _ = core.GetValueByteLength(true)
    assert.Equal(t, v, uint32(1), "Wrong type size returned")

    v, _ = core.GetValueByteLength(false)
    assert.Equal(t, v, uint32(1), "Wrong type size returned")

    v, _ = core.GetValueByteLength(time.Now())
    assert.Equal(t, v, uint32(8), "Wrong type size returned")

    v, _ = core.GetValueByteLength([]byte{})
    assert.Equal(t, v, uint32(0), "Wrong type size returned")

    _, err = core.GetValueByteLength(uint(1))
    assert.NotNil(t, err, "An error should be returned")

    _, err = core.GetValueByteLength(uint16(1))
    assert.NotNil(t, err, "An error should be returned")

    _, err = core.GetValueByteLength(uint32(1))
    assert.NotNil(t, err, "An error should be returned")

    _, err = core.GetValueByteLength(uint64(1))
    assert.NotNil(t, err, "An error should be returned")
}

func TestObjectData(t *testing.T) {
    data := &core.ObjectsData{
        Data: make(map[string] interface{}),
    }

    data.Put("/0/0", 1)
    assert.Equal(t, data.Get("/0/0"), 1, "Value get not equal to put")

    data.Put("/0/1", 0)
    assert.Equal(t, data.Get("/0/1"), 0, "Value get not equal to put")

    data.Put("/0/2/101", []byte{0, 15})
    assert.Equal(t, data.Get("/0/2/101"), []byte{0, 15}, "Value get not equal to put")

    data.Put("/0/3", 101)
    assert.Equal(t, data.Get("/0/3"), 101, "Value get not equal to put")

    data.Put("/1/0", 1)
    assert.Equal(t, data.Get("/1/0"), 1, "Value get not equal to put")

    data.Put("/1/1", 1)
    assert.Equal(t, data.Get("/1/1"), 1, "Value get not equal to put")

    data.Put("/1/2/102", []byte{0, 15})
    assert.Equal(t, data.Get("/1/2/102"), []byte{0, 15}, "Value get not equal to put")

    data.Put("/1/3", 102)
    assert.Equal(t, data.Get("/1/3"), 102, "Value get not equal to put")

    data.Put("/2/0", 3)
    assert.Equal(t, data.Get("/2/0"), 3, "Value get not equal to put")

    data.Put("/2/1", 0)
    assert.Equal(t, data.Get("/2/1"), 0, "Value get not equal to put")

    data.Put("/2/2/101", []byte{0, 15})
    assert.Equal(t, data.Get("/2/2/101"), []byte{0, 15}, "Value get not equal to put")

    data.Put("/2/2/102", []byte{0, 1})
    assert.Equal(t, data.Get("/2/2/102"), []byte{0, 1}, "Value get not equal to put")

    data.Put("/2/3",  101)
    assert.Equal(t, data.Get("/2/3"), 101, "Value get not equal to put")

    data.Put("/3/0", 4)
    assert.Equal(t, data.Get("/3/0"), 4, "Value get not equal to put")

    data.Put("/3/1", 0)
    assert.Equal(t, data.Get("/3/1"), 0, "Value get not equal to put")

    data.Put("/3/2/101", []byte{0, 1})
    assert.Equal(t, data.Get("/3/2/101"), []byte{0, 1}, "Value get not equal to put")

    data.Put("/3/2/0", []byte{0, 1})
    assert.Equal(t, data.Get("/3/2/0"), []byte{0, 1}, "Value get not equal to put")

    data.Put("/3/3", 101)
    assert.Equal(t, data.Get("/3/3"), 101, "Value get not equal to put")

    data.Put("/4/0", 5)
    assert.Equal(t, data.Get("/4/0"), 5, "Value get not equal to put")

    data.Put("/4/1", 65535)
    assert.Equal(t, data.Get("/4/1"), 65535, "Value get not equal to put")

    data.Put("/4/2/101", []byte{0, 16})
    assert.Equal(t, data.Get("/4/2/101"), []byte{0, 16}, "Value get not equal to put")

    data.Put("/4/3", 65535)
    assert.Equal(t, data.Get("/4/3"), 65535, "Value get not equal to put")
}
