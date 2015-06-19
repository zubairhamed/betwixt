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
)

func DecodeValue(b []byte, resourceDef ResourceDefinition) typeval.Value {
	log.Println("Mutliple Values ALlowed?? ", resourceDef.MultipleValuesAllowed(), resourceDef.GetName())
	var val typeval.Value
	if resourceDef.MultipleValuesAllowed() {
		log.Println("Multiple Values Allowed")
		val = DecodeTlv(b, resourceDef)
	} else {
		switch resourceDef.GetResourceType() {
		case typeval.VALUETYPE_STRING:
			val = typeval.String(string(b[:len(b)]))
			break

		case typeval.VALUETYPE_INTEGER:
			buf := bytes.NewBuffer(b)
			intVal, _ := binary.ReadVarint(buf)
			val = typeval.Integer(int(intVal))
			break
		}
	}
	return val
}

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
	log.Println("ValueTypeCode (ValueFromBytes)", v)

	switch v {
	case typeval.VALUETYPE_STRING:
		return typeval.String(string(b))

	case typeval.VALUETYPE_INTEGER:
		return typeval.Integer(int(b[0]))
	}
	return typeval.String("")
}

func DecodeResourceValue(resourceId uint16, b []byte, resourceDef ResourceDefinition) (typeval.Value, error) {
	if resourceDef.MultipleValuesAllowed() {
		typeField := b[0]
		typeFieldTypeOfIdentifier, typeFieldLengthOfIdentifier, typeFieldTypeOfLength, typeFieldLengthOfValue := DecodeTypeField(typeField)

		if typeFieldTypeOfIdentifier != TYPEFIELD_TYPE_RESOURCEINSTANCE && typeFieldTypeOfIdentifier != TYPEFIELD_TYPE_MULTIPLERESOURCE && typeFieldTypeOfIdentifier != TYPEFIELD_TYPE_RESOURCE {
			return nil, errors.New("Invalid identifier. Expecting a resource identifier")
		}

		if typeFieldTypeOfIdentifier == TYPEFIELD_TYPE_MULTIPLERESOURCE {
			valueOffset := 1
			var identifier uint16
			if typeFieldLengthOfIdentifier == 0 {
				_identifier, _ := binary.Uvarint(b[valueOffset:valueOffset+1])
				identifier = uint16(_identifier)
				valueOffset += 1
			} else {
				_identifier, _ := binary.Uvarint(b[valueOffset:valueOffset+2])
				identifier = uint16(_identifier)
				valueOffset += 2
			}

			totalValueLength := len(b)
			var valueFieldLength uint64
			if typeFieldTypeOfLength == 0 {
				valueFieldLength = uint64(typeFieldLengthOfValue)
			} else {
				valueFieldLength, _ = binary.Uvarint(b[:valueOffset+int(typeFieldTypeOfLength)])
			}

			actualValueLength := (uint64(valueOffset) + valueFieldLength)
			if uint64(totalValueLength) != actualValueLength {
				return nil, errors.New("Resource is of invalid length.")
			}

			bytesValue := b[valueOffset:]

			bytesLeft := bytesValue
			resourceBytes := [][]byte{}
			for len(bytesLeft) > 0 {

				typeField := bytesLeft[0]
				typeFieldTypeOfIdentifier, typeFieldLengthOfIdentifier, typeFieldTypeOfLength, typeFieldLengthOfValue := DecodeTypeField(typeField)

				if typeFieldTypeOfIdentifier != TYPEFIELD_TYPE_RESOURCEINSTANCE && typeFieldTypeOfIdentifier != TYPEFIELD_TYPE_MULTIPLERESOURCE && typeFieldTypeOfIdentifier != TYPEFIELD_TYPE_RESOURCE {
					return nil, errors.New("Invalid identifier. Expecting a resource identifier")
				}

				valueOffset := 1
				if typeFieldLengthOfIdentifier == 0 {
					valueOffset += 1
				} else {
					valueOffset += 2
				}

				var valueFieldLength uint64
				if typeFieldTypeOfLength == 0 {
					valueFieldLength = uint64(typeFieldLengthOfValue)
				} else {
					valueFieldLength, _ = binary.Uvarint(b[:valueOffset+int(typeFieldTypeOfLength)])
				}

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
			if typeFieldLengthOfIdentifier == 0 {
				valueOffset += 1
			} else {
				valueOffset += 2
			}

			totalValueLength := len(b)
			var valueFieldLength uint64
			if typeFieldTypeOfLength == 0 {
				valueFieldLength = uint64(typeFieldLengthOfValue)
			} else {
				valueFieldLength, _ = binary.Uvarint(b[:valueOffset+int(typeFieldTypeOfLength)])
			}

			actualValueLength := (uint64(valueOffset) + valueFieldLength)
			if uint64(totalValueLength) != actualValueLength {
				return nil, errors.New("Resource is of invalid length.")
			}
			bytesValue := b[valueOffset:]
			return NewResourceValue(resourceId, ValueFromBytes(bytesValue, resourceDef.GetResourceType())), nil
		}
	} else {
		return NewResourceValue(resourceId, ValueFromBytes(b, resourceDef.GetResourceType())), nil
	}


	/*
	if resourceDef.MultipleValuesAllowed() {


		if typeOfIdentifier == TYPEFIELD_TYPE_RESOURCEINSTANCE || typeOfIdentifier == TYPEFIELD_TYPE_RESOURCE {
			log.Println("Resource Instance/Value")
			if typeOfLength == 0 {
				if lengthOfIdentifier == 0 {
					if len(b) != 2 {

					} else {
						log.Println("Identifier is ", b[1])
					}
				} else {
					if len(b) != 3 {
						return nil, errors.New("Resource value is of invalid length. Expecting 3 bytes.")
					} else {
						log.Println("Identifier is ", b[1:2])
					}
				}
				log.Println("Value is ", lengthOfValue)
			} else {
				valueOffset := 1
				if lengthOfIdentifier == 0 {
					valueOffset += 1
				} else {
					valueOffset += 2
				}

				var valueLength uint64

				if typeOfLength == 1 {
					// 8-bit length value
					valueLength, _ = binary.Uvarint(b[valueOffset:valueOffset+1])
					valueOffset += 1
				} else
				if typeOfLength == 2 {
					// 16-bit length value
					valueLength, _ = binary.Uvarint(b[valueOffset:valueOffset+2])
					valueOffset += 2
				} else
				if typeOfLength == 3 {
					// 24 bit length value
					valueLength, _ = binary.Uvarint(b[valueOffset:valueOffset+3])
					valueOffset += 3
				}
				log.Println("Value length == ", valueLength)

				// Validate resource value length

				log.Println("Value is ", b[valueOffset:])
			}

			log.Println("Went OK for Single Resource")
			return nil, nil
		} else {
			// parse all byte content for resource
			// create MultiResource
			// call DecodeResourceValue(byte, resourceDef)
			log.Println("Went OK for Multi Resource")
			return nil, nil
		}
	} else {
		return NewResourceValue()
		log.Println("Went OK for Single Resource")
		return nil, nil
	}
	*/
}

func DecodeObjectValue(b []byte, objectDef ObjectDefinition) typeval.Value {
	// if type identifier is not object value, return error

	// parse all byte content
	// call DecodeObjectTlv(bytes)
	return nil
}

func DecodeObjectTlv() {
	// [] Value
	// for each resource
		// parse all byte content for resource
		// call DecodeResourceValue(bytes, resourceDef)

	// create ObjectValue
}

func DecodeResourceTlv() {
	// if
}

func DecodeTlv(b []byte, resourceDef ResourceDefinition) typeval.Value {
	log.Println("Decoding TLV Value")
	typeField := b[0]
	typeOfIdentifier := typeField & TLV_FIELD_IDENTIFIER_TYPE
	lengthOfIdentifier := typeField & TLV_FIELD_IDENTIFIER_LENGTH
	typeOfLength := typeField & TLV_FIELD_TYPE_OF_LENGTH
	lengthOfValue := typeField & TLV_FIELD_LENGTH_OF_VALUE


	if typeOfIdentifier == TYPEFIELD_TYPE_OBJECTINSTANCE {
		// Extract all Object TLV Values

		// Parse ObjectBytes
		// DecodeObjectTlv()
		// return []ObjectValue
	} else
	if typeOfIdentifier == TYPEFIELD_TYPE_RESOURCEINSTANCE {
		// DecodeResourceInstanceTlv([]byte)
		// return ResourceInstanceValue
	} else
	if typeOfIdentifier == TYPEFIELD_TYPE_MULTIPLERESOURCE {

	} else
	if typeOfIdentifier == TYPEFIELD_TYPE_RESOURCE {

	} else {
		// Unknown
	}
	// DescribeTLVField(typeField)
	log.Println(typeOfIdentifier, lengthOfIdentifier, typeOfLength, lengthOfValue)



	// Type Field Byte
		// 7-6 : Type of Identifier
			// 00: Object Instance
			// 01: Resource Instance with Value for use within a multiple Resource TLV
			// 10: Multiple Resource, in which case the Value contains one or more Resource Instance TLVs
			// 11: Resource with Value

		// 5: Length of Identifier
			// 0: Identifier field is 8 bits
			// 1: Identifier field is 16 bits

		// 4-3: Type of Length
			// 00: No length field, the value immediately follows the Identifier field in is of the length indicated by Bits 2-0 of this field
			// 01: The Length field is 8-bits and Bits 2-0 MUST be ignored
			// 10: The Length field is 16-bits and Bits 2-0 MUST be ignored
			// 11: The Length field is 24-bits and Bits 2-0 MUST be ignored

		// 2-0: A 3-bit unsigned integer indicating the Length of the Value.

	// Identifier Byte: 8-bit or 16-bit unsigned integer as indicated by the Type field.
		// The Length of the following field in bytes.

	// Value: Sequence of bytes of Length
		// Value of the tag. The format of the value depends on the Resource’s data type (See Appendix C).
	return nil
}

func EncodeObjectInstanceValue() {

}

func EncodeResourceInstanceValue() {

}

func EncodeMultipleResourceValue() {

}

func EncodeResourceValue(resourceId int, v typeval.Value) {

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
