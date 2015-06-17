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
)

func DecodeValue(b []byte, resourceDef ResourceDefinition) typeval.Value {

	var val typeval.Value
	if resourceDef.MultipleValuesAllowed() {
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


func DecodeTlv(b []byte, resourceDef ResourceDefinition) typeval.Value {
	typeField := b[0]
	typeOfIdentifier := typeField & 192
	log.Println("Type Field", typeField)
	log.Println("Type of Identifier", typeOfIdentifier)

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

	// Identifier Bite: 8-bit or 16-bit unsigned integer as indicated by the Type field.
		// The Length of the following field in bytes.

	// Value: Sequence of bytes of Length
		// Value of the tag. The format of the value depends on the Resource’s data type (See Appendix C).
	return nil
}

const (
	TYPEFIELD_TYPE_OBJECTINSTANCE = 0
	TYPEFIELD_TYPE_RESOURCEINSTANCE = 64
	TYPEFIELD_TYPE_MULTIPLERESOURCE = 128
	TYPEFIELD_TYPE_RESOURCE = 192
)

func EncodeObjectInstanceValue() {

}

func EncodeResourceInstanceValue() {

}

func EncodeMultipleResourceValue() {

}

func EncodeResourceValue(resourceId int, v typeval.Value) {

}

func EncodeValue(resourceId int, allowMultipleValues bool, v typeval.Value) []byte {
	log.Println("typeVal Multiple ? ", v.GetType() == typeval.VALUETYPE_MULTIPLE)
	if v.GetType() == typeval.VALUETYPE_MULTIPLE {
		typeOfMultipleValue := v.GetContainedType()
		if typeOfMultipleValue == typeval.VALUETYPE_INTEGER {

			// Resource Instances TLV
			resourceInstanceBytes := bytes.NewBuffer([]byte{})
			intValues := v.GetValue().([]typeval.Value)
			for i, intValue := range intValues {
				value := intValue.GetValue().(int)

				// Type Field Byte
				log.Println("###### AllowMultipleValues", allowMultipleValues)
				if allowMultipleValues {
					typeField := CreateTlvTypeField(TYPEFIELD_TYPE_RESOURCEINSTANCE, value, i)
					log.Println("Type Field == ", typeField)
					resourceInstanceBytes.Write([]byte{typeField})
				} else {
					typeField := CreateTlvTypeField(TYPEFIELD_TYPE_RESOURCE, value, i)
					log.Println("Type Field == ", typeField)
					resourceInstanceBytes.Write([]byte{typeField})
				}

				// Identifier Field
				identifierField := CreateTlvIdentifierField(i)
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
