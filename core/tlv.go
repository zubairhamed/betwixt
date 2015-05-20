package core

import (
    "encoding/binary"
    "bytes"
    "time"
    "errors"
)

func TlvPayloadFromObjects(en *ObjectEnabler, reg Registry) (ResourceValue, error) {
    buf := bytes.NewBuffer([]byte{})

    for _, oi := range en.Instances {
        m := reg.GetModel(oi.TypeId)

        rsrcBuf := bytes.NewBuffer([]byte{})
        for _, ri := range m.Resources {
            if ri.IsReadable() {
                ret := en.Handler.OnRead(oi.Id, ri.Id)

                if ri.Multiple {
                    rsrcBuf.Write(ret.GetBytes())
                } else {
                    if ri.ResourceType == VALUETYPE_INTEGER {
                        v, _ := TlvPayloadFromIntResource(ri, []int{ret.GetValue().(int)})
                        rsrcBuf.Write(v.GetBytes())
                    }
                }
            }
        }

        if len(en.Instances) > 1 {
            // Create Root TLV Value for Resource

            // Append to Resource Buffer
            // buf.Write(..)
        }
        // Append to Resource TLV to Main Buffer
        buf.Write(rsrcBuf.Bytes())
    }

    return NewTlvValue(buf.Bytes()), nil
}

func CreateTlvTypeField(identType byte, value interface{}, ident int) byte {
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
    valueTypeLength, _ := GetValueByteLength(value)

    if valueTypeLength > 7 {
        buf := new(bytes.Buffer)
        binary.Write(buf, binary.BigEndian, valueTypeLength)

        return bytes.Trim(buf.Bytes(), "\x00")
    }
    return []byte{}
}

func CreateTlvValueField(value int) [] byte {
    buf := new(bytes.Buffer)
    binary.Write(buf, binary.BigEndian, uint64(value))
    if value == 0 {
        return []byte{ 0 }
    } else {
        return bytes.Trim(buf.Bytes(), "\x00")
    }
}

func TlvPayloadFromIntResource(model *ResourceModel, values []int) (ResourceValue, error) {

    // Resource Instances TLV
    resourceInstanceBytes := bytes.NewBuffer([]byte{})

    if model.Multiple {
        for i, value := range values {
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
    typeField := CreateTlvTypeField(128, resourceInstanceBytes.Bytes(), model.Id)
    resourceTlv.Write([]byte{typeField})

    // Identifier Field
    identifierField := CreateTlvIdentifierField(model.Id)
    resourceTlv.Write(identifierField)

    // Length Field
    lengthField := CreateTlvLengthField(resourceInstanceBytes.Bytes())
    resourceTlv.Write(lengthField)

    // Value Field, Append Resource Instances TLV to Resource TLV
    resourceTlv.Write(resourceInstanceBytes.Bytes())

    return NewTlvValue(resourceTlv.Bytes()), nil
}

func GetValueByteLength(val interface{}) (uint32, error) {

    if _, ok := val.(int); ok {
        v := val.(int)
        if v > 127 || v < -128 {
            if v > 32767 || v < -32768 {
                if v > 2147483647 || v < -2147483648 {
                    return 8, nil
                } else {
                    return 4, nil
                }
            } else {
                return 2, nil
            }
        } else {
            return 1, nil
        }
    } else
    if _, ok := val.(bool); ok {
        return 1, nil
    } else
    if _, ok := val.(string); ok {
        v := val.(string)

        return uint32(len(v)), nil
    } else
    if _, ok := val.(float64); ok {
        v := val.(float64)

        if v > +3.4E+38 || v < -3.4E+38 {
            return 8, nil
        } else {
            return 4, nil
        }
    } else
    if _, ok := val.(time.Time); ok {
        return 8, nil
    } else
    if _, ok := val.([]byte); ok {
        v := val.([]byte)
        return uint32(len(v)), nil
    } else {
        return 0, errors.New("Unknown type")
    }
}
