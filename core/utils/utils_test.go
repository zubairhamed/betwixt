package utils

import (
	"github.com/stretchr/testify/assert"
	"github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/betwixt/tests"
	"github.com/zubairhamed/go-commons/typeval"
	"testing"
	"time"
)

func TestGetValueByteLength(t *testing.T) {
	test1 := []struct {
		input    interface{}
		expected uint32
	}{
		{-128, 1},
		{127, 1},
		{-32768, 2},
		{-2147483648, 4},
		{2147483647, 4},
		{-9223372036854775808, 8},
		{9223372036854775807, 8},
		{-3.4E+38, 4},
		{3.4E+38, 4},
		{-1.7E+308, 8},
		{1.7E+308, 8},
		{"this is a string", 16},
		{true, 1},
		{false, 1},
		{time.Now(), 8},
		{[]byte{}, 0},
	}

	for _, c := range test1 {
		v, _ := typeval.GetValueByteLength(c.input)
		assert.Equal(t, v, uint32(c.expected), "Wrong expected length returned")
	}

	test2 := []struct {
		input interface{}
	}{
		{uint(1)},
		{uint16(1)},
		{uint32(1)},
		{uint64(1)},
	}

	for _, c := range test2 {
		_, err := typeval.GetValueByteLength(c.input)
		assert.NotNil(t, err, "An error should be returned")
	}
}

func TestDecodeTypeField(t *testing.T) {
	var data []byte

	data = []byte{134, 6, 65, 0, 1, 65, 1, 5}
	typeOfIdentifier, lengthOfIdentifier, typeOfLength, lengthOfValue := DecodeTypeField(data[0])

	assert.Equal(t, byte(TYPEFIELD_TYPE_MULTIPLERESOURCE), typeOfIdentifier)
	assert.Equal(t, byte(0), lengthOfIdentifier)
	assert.Equal(t, byte(0), typeOfLength)
	assert.Equal(t, byte(6), lengthOfValue)

	data = []byte{136, 7, 8, 66, 0, 14, 216, 66, 1, 19, 136}
	typeOfIdentifier, lengthOfIdentifier, typeOfLength, lengthOfValue = DecodeTypeField(data[0])
	assert.Equal(t, byte(TYPEFIELD_TYPE_MULTIPLERESOURCE), typeOfIdentifier)
	assert.Equal(t, byte(0), lengthOfIdentifier)
	assert.Equal(t, byte(8), typeOfLength)
	assert.Equal(t, byte(0), lengthOfValue)

	data = []byte{135, 8, 65, 0, 125, 66, 1, 3, 132}
	typeOfIdentifier, lengthOfIdentifier, typeOfLength, lengthOfValue = DecodeTypeField(data[0])
	assert.Equal(t, byte(TYPEFIELD_TYPE_MULTIPLERESOURCE), typeOfIdentifier)
	assert.Equal(t, byte(0), lengthOfIdentifier)
	assert.Equal(t, byte(0), typeOfLength)
	assert.Equal(t, byte(7), lengthOfValue)
}

func TestValueFromBytes(t *testing.T) {

	var data []byte
	var val typeval.Value

	data = []byte{79, 112, 101, 110, 32, 77, 111, 98, 105, 108, 101, 32, 65, 108, 108, 105, 97, 110, 99, 101}
	val = ValueFromBytes(data, typeval.VALUETYPE_STRING)
	assert.Equal(t, "Open Mobile Alliance", val.GetValue().(string))

	data = []byte{76, 105, 103, 104, 116, 119, 101, 105, 103, 104, 116, 32, 77, 50, 77, 32, 67, 108, 105, 101, 110, 116}
	val = ValueFromBytes(data, typeval.VALUETYPE_STRING)
	assert.Equal(t, "Lightweight M2M Client", val.GetValue().(string))

	data = []byte{49, 46, 48}
	val = ValueFromBytes(data, typeval.VALUETYPE_STRING)
	assert.Equal(t, "1.0", val.GetValue().(string))

	data = []byte{}
	val = ValueFromBytes(data, typeval.VALUETYPE_STRING)
	assert.Equal(t, typeval.VALUETYPE_EMPTY, val.GetType() )

	data =[]byte{ 100 }
	val = ValueFromBytes(data, typeval.VALUETYPE_INTEGER)
	assert.Equal(t, 100, val.GetValue().(int))

	data = []byte{}
	val = ValueFromBytes(data, typeval.VALUETYPE_OBJECTLINK)
	assert.Equal(t, typeval.VALUETYPE_EMPTY, val.GetType() )

}

func TestValidResourceTypeField(t *testing.T) {
	var data []byte
	var err error

	data = []byte{134, 6, 65, 0, 1, 65, 1, 5}
	err = ValidResourceTypeField(data)
	assert.Nil(t, err)

	data = []byte{1, 6, 65, 0, 1, 65, 1, 5}
	err = ValidResourceTypeField(data)
	assert.NotNil(t, err)
}

func TestDecodeIdentifierField(t *testing.T) {

	var data []byte
	data = []byte{136, 7, 8, 66, 0, 14, 216, 66, 1, 19, 136}

	identifier, length := DecodeIdentifierField(data, 1)

	assert.Equal(t, byte(7), byte(identifier))
	assert.Equal(t, 1, length)
}

func TestDecodeLengthField(t *testing.T) {

	var data []byte
	data = []byte{136, 7, 8, 66, 0, 14, 216, 66, 1, 19, 136}
	valueLength, typeLength := DecodeLengthField(data, 2)

	assert.Equal(t, byte(8), byte(valueLength))
	assert.Equal(t, 1, typeLength)
}

func TestDecodeResourceValue(t *testing.T) {
	/*
	var data []byte
	data = []byte{136, 7, 8, 66, 0, 14, 216, 66, 1, 19, 136}

	val, err := DecodeResourceValue(7, data, nil)

	log.Println(val, err)
	*/

	//func DecodeResourceValue(resourceId uint16, b []byte, resourceDef ResourceDefinition) (typeval.Value, error) {
}

func TestEncodeValue(t *testing.T) {
	// func EncodeValue(resourceId uint16, allowMultipleValues bool, v typeval.Value) []byte {
}

func TestResourceOperations(t *testing.T) {
	//	func IsExecutableResource(m ResourceDefinition) bool {
	//	func IsReadableResource(m ResourceDefinition) bool {
	//	func IsWritableResource(m ResourceDefinition) bool {
}

func TestBuildResourceStringPayload(t *testing.T) {
	cli := tests.NewMockClient()

	reg := tests.NewMockRegistry()
	cli.UseRegistry(reg)

	cli.EnableObject(betwixt.LWM2MObjectType(0), nil)
	cli.EnableObject(betwixt.LWM2MObjectType(2), nil)
	cli.EnableObject(betwixt.LWM2MObjectType(4), nil)

	str := BuildModelResourceStringPayload(cli.GetEnabledObjects())
	compare := "</0>,</2>,</4>,"

	assert.Equal(t, str, compare, "Unexpected output building Model Resource String")
}
