package utils

import (
	"bytes"
	"fmt"
	. "github.com/zubairhamed/betwixt"
	"sort"
	"github.com/zubairhamed/go-commons/typeval"
	. "github.com/zubairhamed/betwixt/core/values/tlv"
)

func BytesFromValue(resourceDef ResourceDefinition, v typeval.Value) []byte {
	if v.GetType() == typeval.VALUETYPE_MULTIPLE {
		typeOfMultipleValue := v.GetContainedType()
		if typeOfMultipleValue == typeval.VALUETYPE_INTEGER {

			// Resource Instances TLV
			resourceInstanceBytes := bytes.NewBuffer([]byte{})
			if resourceDef.MultipleValuesAllowed() {
				intValues := v.GetValue().([]typeval.Value)
				for i, intValue := range intValues {
					value := intValue.GetValue().(int)

					// Type Field Byte
					typeField := CreateTlvTypeField(64, value, i)
					resourceInstanceBytes.Write([]byte{typeField})

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
			}

			// Resource Root TLV
			resourceTlv := bytes.NewBuffer([]byte{})

			// Byte 7-6: identifier
			typeField := CreateTlvTypeField(128, resourceInstanceBytes.Bytes(), resourceDef.GetId())
			resourceTlv.Write([]byte{typeField})

			// Identifier Field
			identifierField := CreateTlvIdentifierField(resourceDef.GetId())
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
