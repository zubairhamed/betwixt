package betwixt

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStringValue(t *testing.T) {
	val := String("this is a string")
	assert.Equal(t, VALUETYPE_STRING, val.GetType())
	assert.Equal(t, "this is a string", val.GetValue())
	assert.Equal(t, "this is a string", val.GetStringValue())
	assert.Equal(t, 16, len(val.GetBytes()))
}

func TestIntegerValue(t *testing.T) {
	val := Integer(42)
	assert.Equal(t, VALUETYPE_INTEGER, val.GetType())
	assert.Equal(t, 42, val.GetValue())
	assert.Equal(t, "42", val.GetStringValue())
	assert.Equal(t, 1, len(val.GetBytes()))
}

func TestTimeValue(t *testing.T) {
	tv := time.Unix(1433767779, 0)
	val := Time(tv)
	assert.Equal(t, VALUETYPE_TIME, val.GetType())
	assert.Equal(t, tv, val.GetValue())
	assert.Equal(t, "1433767779", val.GetStringValue())
	assert.Equal(t, 10, len(val.GetBytes()))
}

func TestFloatValue(t *testing.T) {
	val := Float(float32(4.2))
	assert.Equal(t, VALUETYPE_FLOAT, val.GetType())
	assert.Equal(t, float32(4.2), val.GetValue())
	assert.Equal(t, "4", val.GetStringValue())
	assert.Equal(t, 4, len(val.GetBytes()))
}

func TestBooleanValue(t *testing.T) {
	val := Boolean(true)
	assert.Equal(t, VALUETYPE_BOOLEAN, val.GetType())
	assert.Equal(t, true, val.GetValue())
	assert.Equal(t, "1", val.GetStringValue())
	assert.Equal(t, 0, len(val.GetBytes()))
}

func TestEmptyValue(t *testing.T) {
	val := Empty()
	assert.Equal(t, VALUETYPE_EMPTY, val.GetType())
	assert.Equal(t, "", val.GetValue())
	assert.Equal(t, "", val.GetStringValue())
	assert.Equal(t, 0, len(val.GetBytes()))
}

/*
func TestTlvValue(t *testing.T) {
	val := Tlv([]byte{0, 1, 2})
	assert.Equal(t, VALUETYPE_TLV, val.GetType())
	assert.Equal(t, []byte{0, 1, 2}, val.GetValue())
	assert.Equal(t, "", val.GetStringValue())
	assert.Equal(t, 3, len(val.GetBytes()))
}

func TestMultipleResourceInstanceValue(t *testing.T) {

}
*/
