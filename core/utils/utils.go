package utils

import (
	"bytes"
	"fmt"
	. "github.com/zubairhamed/betwixt"
	"sort"
	"github.com/zubairhamed/go-commons/typeval"
	. "github.com/zubairhamed/betwixt/core/values/tlv"
	"log"
	"encoding/binary"
	"errors"
	"github.com/zubairhamed/go-commons/logging"
)



const(
	TLV_FIELD_IDENTIFIER_TYPE 	= 192
	TLV_FIELD_IDENTIFIER_LENGTH = 32
	TLV_FIELD_TYPE_OF_LENGTH 	= 24
	TLV_FIELD_LENGTH_OF_VALUE 	= 7
)

const (
	TYPEFIELD_TYPE_OBJECTINSTANCE 	= 0		// 00
	TYPEFIELD_TYPE_RESOURCEINSTANCE = 64	// 01
	TYPEFIELD_TYPE_MULTIPLERESOURCE = 128	// 10
	TYPEFIELD_TYPE_RESOURCE 		= 192	// 11
)

/*
	MultipleValue
		[]ObjectValue
 */

type ObjectValue struct {
	instanceId 		uint16
	typeId 			LWM2MObjectType
	resources 		[]typeval.Value
}

func NewResourceValue(id uint16, value typeval.Value) typeval.Value {
	return &ResourceValue{
		id: id,
		value: value,
	}
}

type ResourceValue struct {
	id 		uint16
	value 	typeval.Value
}

func (v ResourceValue) GetId() (uint16) {
	return v.id
}

func (v ResourceValue) GetBytes() ([]byte) {
	return v.value.GetBytes()
}

func (v ResourceValue) GetContainedType() (typeval.ValueTypeCode) {
	return typeval.VALUETYPE_RESOURCE
}

func (v ResourceValue) GetType() (typeval.ValueTypeCode) {
	return typeval.VALUETYPE_RESOURCE
}

func (v ResourceValue) GetStringValue() (string) {
	return ""
}

func (v ResourceValue) GetValue() interface{} {
	return v.value.GetValue()
}

func NewMultipleResourceValue(id uint16, value []*ResourceValue) typeval.Value {
	return &MultipleResourceValue{
		id: id,
		instances: value,
	}
}

type MultipleResourceValue struct {
	id 			uint16
	instances 	[]*ResourceValue
}

func (v MultipleResourceValue) GetBytes() ([]byte) {
	return []byte{}
}

func (v MultipleResourceValue) GetContainedType() (typeval.ValueTypeCode) {
	return typeval.VALUETYPE_RESOURCE
}

func (v MultipleResourceValue) GetType() (typeval.ValueTypeCode) {
	return typeval.VALUETYPE_MULTIRESOURCE
}

func (v MultipleResourceValue) GetStringValue() (string) {
	return ""
}

func (v MultipleResourceValue) GetValue() interface{} {
	return v.instances
}


func DecodeTypeField(typeField byte)(byte, byte, byte, byte) {
	typeOfIdentifier := typeField & TLV_FIELD_IDENTIFIER_TYPE
	lengthOfIdentifier := typeField & TLV_FIELD_IDENTIFIER_LENGTH
	typeOfLength := typeField & TLV_FIELD_TYPE_OF_LENGTH
	lengthOfValue := typeField & TLV_FIELD_LENGTH_OF_VALUE

	return typeOfIdentifier, lengthOfIdentifier, typeOfLength, lengthOfValue
}

func ValueFromBytes(b []byte, v typeval.ValueTypeCode) typeval.Value {
	log.Println("ValueFromBytes", b)
	if len(b) == 0 {
		return typeval.Empty()
	}

	switch v {
	case typeval.VALUETYPE_STRING:
		return typeval.String(string(b))

	case typeval.VALUETYPE_INTEGER:
		intLen := len(b)
		if intLen == 1 {
			return typeval.Integer(int(b[0]))
		} else
		if intLen == 2 {
			return typeval.Integer(int(b[1]) | (int(b[0]) << 8))
		} else
		if intLen == 4 {
			return typeval.Integer(int(b[3]) | (int(b[2]) << 8) | (int(b[1]) << 16) | (int(b[0]) << 24))
		} else
		if intLen == 8 {
			return typeval.Integer(int(b[7]) | (int(b[6]) << 8) | (int(b[5]) << 16) | (int(b[4]) << 24) | (int(b[3]) << 32) | (int(b[2]) << 40) | (int(b[1]) << 48) | (int(b[0]) << 56))
		} else {
			// Error
		}
	}
	return typeval.String("")
}

//func BytesToInt(b []byte) int {
//	return int(b[0]) | int(b[1] <<8 )
//}

func ValidResourceTypeField(b []byte) (error){
	typeField := b[0]

	typeFieldTypeOfIdentifier, _, _, _ := DecodeTypeField(typeField)

	if typeFieldTypeOfIdentifier != TYPEFIELD_TYPE_RESOURCEINSTANCE && typeFieldTypeOfIdentifier != TYPEFIELD_TYPE_MULTIPLERESOURCE && typeFieldTypeOfIdentifier != TYPEFIELD_TYPE_RESOURCE {
		return errors.New("Invalid identifier. Expecting a resource identifier")
	}
	return nil
}

func DecodeIdentifierField(b []byte, pos int)(identifier uint16, typeLength int) {
	_, typeFieldLengthOfIdentifier, _, _ := DecodeTypeField(b[0])

	if typeFieldLengthOfIdentifier == 0 {
		_identifier, _ := binary.Uvarint(b[pos:pos+1])
		identifier = uint16(_identifier)
		typeLength = 1
	} else {
		_identifier, _ := binary.Uvarint(b[pos:pos+2])
		identifier = uint16(_identifier)
		typeLength = 2
	}
	return
}

func DecodeLengthField(b []byte, pos int)(valueLength uint64, typeLength int) {
	_, _, typeFieldTypeOfLength, typeFieldLengthOfValue := DecodeTypeField(b[0])

	typeLength = 0
	if typeFieldTypeOfLength == 0 {
		valueLength = uint64(typeFieldLengthOfValue)
	} else
	if typeFieldTypeOfLength == 8 {
		valueLength, _ = binary.Uvarint(b[pos:pos+1])
		typeLength = 1
	} else
	if typeFieldTypeOfLength == 16 {
		valueLength, _ = binary.Uvarint(b[pos:pos+2])
		typeLength = 2
	} else
	if typeFieldTypeOfLength == 24 {
		valueLength, _ = binary.Uvarint(b[pos:pos+3])
		typeLength = 3
	} else {
		// Invalid type of Length}
	}
	return
}

func DecodeResourceValue(resourceId uint16, b []byte, resourceDef ResourceDefinition) (typeval.Value, error) {
	log.Println("Parsing", b)
	if resourceDef.MultipleValuesAllowed() {

		err := ValidResourceTypeField(b)
		if err != nil {
			logging.LogError(err)
		}

		typeFieldTypeOfIdentifier, _, _, _ := DecodeTypeField(b[0])

		if typeFieldTypeOfIdentifier == TYPEFIELD_TYPE_MULTIPLERESOURCE {
			valueOffset := 1
			identifier, identifierTypeLength := DecodeIdentifierField(b, valueOffset)
			valueOffset += identifierTypeLength

			_, valueTypeLength := DecodeLengthField(b, valueOffset)
			valueOffset += valueTypeLength
			bytesValue := b[valueOffset:]

			bytesLeft := bytesValue
			resourceBytes := [][]byte{}
			for len(bytesLeft) > 0 {
				err := ValidResourceTypeField(bytesLeft)
				if err != nil {
					logging.LogError(err)
				}

				valueOffset := 1
				_, identifierTypeLength := DecodeIdentifierField(bytesLeft, valueOffset)
				valueOffset += identifierTypeLength

				valueFieldLength, valueTypeLength := DecodeLengthField(bytesLeft, valueOffset)
				valueOffset += valueTypeLength

				actualValueLength := (uint64(valueOffset) + valueFieldLength)

				bytesValue := bytesLeft[:actualValueLength]
				resourceBytes = append(resourceBytes, bytesValue)
				bytesLeft = bytesLeft[actualValueLength:]
			}

			decodedValues := []*ResourceValue{}
			for _, r := range resourceBytes {
				v, _ := DecodeResourceValue(identifier, r, resourceDef)
				decodedValues = append(decodedValues, v.(*ResourceValue))
			}

			return NewMultipleResourceValue(identifier, decodedValues), nil
		} else {
			valueOffset := 1
			_, identifierTypeLength := DecodeIdentifierField(b, valueOffset)
			valueOffset += identifierTypeLength

			_, valueTypeLength := DecodeLengthField(b, valueOffset)
			valueOffset += valueTypeLength

			bytesValue := b[valueOffset:]
			return NewResourceValue(resourceId, ValueFromBytes(bytesValue, resourceDef.GetResourceType())), nil
		}
	} else {
		return NewResourceValue(resourceId, ValueFromBytes(b, resourceDef.GetResourceType())), nil
	}
}

func EncodeValue(resourceId uint16, allowMultipleValues bool, v typeval.Value) []byte {
	if v.GetType() == typeval.VALUETYPE_MULTIPLE {
		typeOfMultipleValue := v.GetContainedType()
		if typeOfMultipleValue == typeval.VALUETYPE_INTEGER {

			// Resource Instances TLV
			resourceInstanceBytes := bytes.NewBuffer([]byte{})
			intValues := v.GetValue().([]typeval.Value)
			for i, intValue := range intValues {
				value := intValue.GetValue().(int)

				// Type Field Byte
				if allowMultipleValues {
					typeField := CreateTlvTypeField(TYPEFIELD_TYPE_RESOURCEINSTANCE, value, uint16(i))
					resourceInstanceBytes.Write([]byte{typeField})
				} else {
					typeField := CreateTlvTypeField(TYPEFIELD_TYPE_RESOURCE, value, uint16(i))
					resourceInstanceBytes.Write([]byte{typeField})
				}

				// Identifier Field
				identifierField := CreateTlvIdentifierField(uint16(i))
				resourceInstanceBytes.Write(identifierField)

				// Length Field
				lengthField := CreateTlvLengthField(value)
				resourceInstanceBytes.Write(lengthField)

				// Value Field
				valueField := CreateTlvValueField(value)
				resourceInstanceBytes.Write(valueField)
			}

			// Resource Root TLV
			resourceTlv := bytes.NewBuffer([]byte{})

			// Byte 7-6: identifier
			typeField := CreateTlvTypeField(128, resourceInstanceBytes.Bytes(), resourceId)
			resourceTlv.Write([]byte{typeField})

			// Identifier Field
			identifierField := CreateTlvIdentifierField(resourceId)
			resourceTlv.Write(identifierField)

			// Length Field
			lengthField := CreateTlvLengthField(resourceInstanceBytes.Bytes())
			resourceTlv.Write(lengthField)

			// Value Field, Append Resource Instances TLV to Resource TLV
			resourceTlv.Write(resourceInstanceBytes.Bytes())

			log.Println("NOT SINGLE!!");
			return resourceTlv.Bytes()
		}
	} else {
		return v.GetBytes()
	}
	return nil
}

func BuildModelResourceStringPayload(instances LWM2MObjectInstances) string {
	var buf bytes.Buffer

	var keys []int
	for k := range instances {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	for _, k := range keys {
		v := instances[LWM2MObjectType(k)]
		inst := v.GetInstances()
		if len(inst) > 0 {
			for _, j := range inst {
				buf.WriteString(fmt.Sprintf("</%d/%d>,", k, j))
			}
		} else {
			buf.WriteString(fmt.Sprintf("</%d>,", k))
		}
	}
	return buf.String()
}

func IsExecutableResource(m ResourceDefinition) bool {
	op := m.GetOperations()
	return (op == OPERATION_E || op == OPERATION_RE || op == OPERATION_RWE || op == OPERATION_WE)
}

func IsReadableResource(m ResourceDefinition) bool {
	op := m.GetOperations()
	return (op == OPERATION_RE || op == OPERATION_R || op == OPERATION_RWE || op == OPERATION_RW)
}

func IsWritableResource(m ResourceDefinition) bool {
	op := m.GetOperations()
	return (op == OPERATION_RW || op == OPERATION_RWE || op == OPERATION_WE || op == OPERATION_W)
}

func CallEvent(e EventType, fn FnEvent) {
	if fn != nil {
		go fn()
	}
}
