package betwixt

import (
	"bytes"
	"encoding/binary"
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
func CreateTlvTypeField(identType byte, value interface{}, ident uint16) byte {
	var typeField byte

	valueTypeLength, _ := GetValueByteLength(value)

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

// CreateTlvIdentifierField Creates a TLV Identifier field
func CreateTlvIdentifierField(ident uint16) []byte {
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

// CreateTlvLengthField creates a TLV Length Field
func CreateTlvLengthField(value interface{}) []byte {
	valueTypeLength, _ := GetValueByteLength(value)

	if valueTypeLength > 7 {
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.BigEndian, valueTypeLength)

		return bytes.Trim(buf.Bytes(), "\x00")
	}
	return []byte{}
}

// CreateTlvValueField creates a TLV Value Field
func CreateTlvValueField(value int) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, uint64(value))
	if value == 0 {
		return []byte{0}
	} else {
		return bytes.Trim(buf.Bytes(), "\x00")
	}
}
