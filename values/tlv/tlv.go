package tlv

import (
	"bytes"
	"encoding/binary"
	. "github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/betwixt/utils"
	"github.com/zubairhamed/betwixt/values"
	"log"
)

/*
   |-------------||-------------||-------------||-------- ........ -----|
       8-bit        8 or 16 bit     0-24 bit
       Type          Identifier     Length                 Value


   0bxxxxxxxx
   7-6: Type of Identifier
       00: Object Instance
       01: Resource Instance with Value
       10: Multiple Resource
       11: Resource with Value
   5: Length of Identifier
       0: 8 bits
       1: 16 bits
   4-3: Type of Length
       00: NO length field, value immediately follows the identifier field
       01: Length field is 8 bits and Bits 2-0 must be ignored
       10: Length field is 16 bits and Bits 2-0 must be ignored
       11: Length field is 24 bits and Bits 2-0 must be ignored
   2-0: 3 bit unsigned integer indiciating Length of the Value

   ------------------------------------
*/

func TlvPayloadFromObjects(o Object, reg Registry) (ResponseValue, error) {
	buf := bytes.NewBuffer([]byte{})

	m := reg.GetDefinition(o.GetType())
	en := o.GetEnabler()
	instances := o.GetInstances()
	for _, oi := range instances {
		rsrcBuf := bytes.NewBuffer([]byte{})
		for _, ri := range m.GetResources() {
			if utils.IsReadableResource(ri) {
				response := en.OnRead(oi, ri.GetId(), nil)

				val := response.GetResponseValue()
				if ri.MultipleValuesAllowed() {
					rsrcBuf.Write(val.GetBytes())
				} else {
					if ri.GetResourceType() == VALUETYPE_INTEGER {
						v, _ := TlvPayloadFromIntResource(ri, []int{val.GetValue().(int)})
						rsrcBuf.Write(v.GetBytes())
					}
				}
			}
		}

		if len(instances) > 1 {
			// Create Root TLV Value for Resource

			// Append to Resource Buffer
			// buf.Write(..)
		}
		// Append to Resource TLV to Main Buffer
		buf.Write(rsrcBuf.Bytes())
	}

	return values.Tlv(buf.Bytes()), nil
}

func TlvPayloadFromIntResource(model ResourceDefinition, vals []int) (ResponseValue, error) {

	// Resource Instances TLV
	resourceInstanceBytes := bytes.NewBuffer([]byte{})

	if model.MultipleValuesAllowed() {
		for i, value := range vals {
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
	typeField := CreateTlvTypeField(128, resourceInstanceBytes.Bytes(), model.GetId())
	resourceTlv.Write([]byte{typeField})

	// Identifier Field
	identifierField := CreateTlvIdentifierField(model.GetId())
	resourceTlv.Write(identifierField)

	// Length Field
	lengthField := CreateTlvLengthField(resourceInstanceBytes.Bytes())
	resourceTlv.Write(lengthField)

	// Value Field, Append Resource Instances TLV to Resource TLV
	resourceTlv.Write(resourceInstanceBytes.Bytes())

	return values.Tlv(resourceTlv.Bytes()), nil
}

func CreateTlvTypeField(identType byte, value interface{}, ident int) byte {
	var typeField byte

	valueTypeLength, _ := utils.GetValueByteLength(value)

	// Bit 7-6: identifier
	typeField |= identType

	// Bit 5
	if ident > 255 {
		typeField |= 32
	}

	// Bit 4-3
	if valueTypeLength > 7 {
		if valueTypeLength < 256 {
			typeField |= 8
		} else {
			if valueTypeLength < 65535 {
				typeField |= 16
			} else {
				if valueTypeLength > 16777215 {
					// Error, size exceeds allowed (> 16.7MB)
				} else {
					// Size is 16777215 or less
					typeField |= 24
				}
			}
		}
	} else {
		// Set bit 2-0 instead
		b := byte(valueTypeLength)
		typeField |= b
	}

	return typeField
}

func CreateTlvIdentifierField(ident int) []byte {
	// Identifier Byte
	if ident > 255 {
		// 16-Bit
		bs := make([]byte, 2)
		binary.LittleEndian.PutUint16(bs, uint16(ident))

		return bs
	} else {
		// 8-Bit
		return []byte{byte(ident)}
	}
}

func CreateTlvLengthField(value interface{}) []byte {
	valueTypeLength, _ := utils.GetValueByteLength(value)

	if valueTypeLength > 7 {
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.BigEndian, valueTypeLength)

		return bytes.Trim(buf.Bytes(), "\x00")
	}
	return []byte{}
}

func CreateTlvValueField(value int) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, uint64(value))
	if value == 0 {
		return []byte{0}
	} else {
		return bytes.Trim(buf.Bytes(), "\x00")
	}
}
