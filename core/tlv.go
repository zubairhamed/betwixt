package core

import (
    "log"
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

func TlvPayloadFromObjectInstance(o *ObjectInstance) (ResourceValue, error) {
    for i, r := range o.Resources {
        log.Println(i, r)
    }

    return NewTlvValue(make([]byte, 0)), nil
}

// TODO: Heavy refactoring needed
func TlvPayloadFromIntResource(model *ResourceModel, values []int) (ResourceValue, error) {

    // Resource Instances TLV
    resourceInstanceBytes := bytes.NewBuffer([]byte{})
    if model.Multiple {
        // Type Byte
        for i, value := range values {
            var typeVal byte

            valueTypeLength, _ := GetValueByteLength(value)
            // Bit 7-6: identifier
            typeVal |= 64

            // Bit 5
            if i > 255 {
                typeVal |= 32
            }

            // Bit 4-3
            if valueTypeLength > 7 {
                if valueTypeLength < 256 {
                    typeVal |= 8
                } else {
                    if valueTypeLength < 65535 {
                        typeVal |= 16
                    } else {
                        if valueTypeLength > 16777215 {
                            // Error, size exceeds allowed (> 16.7MB)
                        } else {
                            // Size is 16777215 or less
                            typeVal |= 24
                        }
                    }
                }
            } else {
                // Set bit 2-0 instead
                b := byte(valueTypeLength)
                typeVal |= b
            }
            resourceInstanceBytes.Write([]byte{typeVal})

            // Identifier Byte
            if i > 255 {
                // 16-Bit
                bs := make([]byte, 2)
                binary.LittleEndian.PutUint16(bs, uint16(i))

                resourceInstanceBytes.Write(bs)
            } else {
                // 8-Bit
                resourceInstanceBytes.Write([]byte{byte(i)})
            }

            // Value Length & Value Byte
            if valueTypeLength > 7 {
                buf := new(bytes.Buffer)
                binary.Write(buf, binary.BigEndian, valueTypeLength)
                resourceInstanceBytes.Write(bytes.Trim(buf.Bytes(), "\x00"))
            }

            // Value
            buf := new(bytes.Buffer)
            binary.Write(buf, binary.BigEndian, uint64(value))
            if value == 0 {
                resourceInstanceBytes.Write([]byte{ 0 })
            } else {
                resourceInstanceBytes.Write(bytes.Trim(buf.Bytes(), "\x00"))
            }

        }
    }

    // Resource Root TLV
    resourceTlv := bytes.NewBuffer([]byte{})
    var typeVal byte

    // Byte 7-6: identifier
    typeVal |= 128

    // Bit 5
    if model.Id > 255 {
        typeVal |= 32
    }

    // Bit 4-3
    resourceInstanceTlvSize := len(resourceInstanceBytes.Bytes())
    if resourceInstanceTlvSize > 7 {
        if resourceInstanceTlvSize < 256 {
            typeVal |= 8
        } else {
            if resourceInstanceTlvSize < 65535 {
                typeVal |= 16
            } else {
                if resourceInstanceTlvSize > 16777215 {
                    // Error, size exceeds allowed (> 16.7MB)
                } else {
                    typeVal |= 24
                }
            }
        }
    } else {
        // Set bit 2-0 instead
        typeVal |= byte(resourceInstanceTlvSize)
    }
    resourceTlv.Write([]byte{typeVal})

    // Identifier
    if model.Id > 255 {
        // 16-Bit
        bs := make([]byte, 2)
        binary.LittleEndian.PutUint16(bs, uint16(model.Id))

        resourceTlv.Write(bs)
    } else {
        // 8-Bit
        resourceTlv.Write([]byte{byte(model.Id)})
    }

    // Value Length
    if resourceInstanceTlvSize > 7 {
        resourceTlv.Write([]byte{byte(resourceInstanceTlvSize)})
    }

    // Append Resource Instances TLV to Resource TLV
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
