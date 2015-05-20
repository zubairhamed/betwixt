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

        log.Println("Instance ID:", oi.Id)

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
    log.Println("TlvPayloadFromObjectInstance", o.Resources)

    for i, r := range o.Resources {
        log.Println(i, r)
    }

    return NewTlvValue(make([]byte, 0)), nil
}

// TODO: Heavy refactoring needed
func TlvPayloadFromIntResource(model *ResourceModel, values []int) (ResourceValue, error) {
    // log.Println("TLV < Payload from Resource", model, values)

    // Resource Instances TLV
    resourceInstanceBytes := bytes.NewBuffer([]byte{})
    if len(values) > 0 {

        // Type Byte
        var typeVal byte
        for i, value := range values {

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
                } else  {
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
                log.Println("2-0", b)

                typeVal |= b
                log.Println("typeVal", typeVal)
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
                binary.Write(buf, binary.LittleEndian, valueTypeLength)
                resourceInstanceBytes.Write(bytes.Trim(buf.Bytes(), "\x00"))

                // Value
                resourceInstanceBytes.Write([]byte{byte(value)})
            } else {
                resourceInstanceBytes.Write([]byte{byte(value )})
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
    log.Println("resourceInstanceTlvSize", resourceInstanceTlvSize)
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
        log.Println("Set bit 2-0 instead", resourceInstanceTlvSize)
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

    /*                      Type                ID Byte(s)      Length Byte(s)      Value
         Available Power     0b10 0 00 110       0x06            (6 byte)            The next two rows
             Multiple Resources
             8 Bits Identifier Field
             No Length Field, the value immediately follows the Identifier field in is of the length indicated by Bits 2-0 of this field
             Length of Value = 6 Bytes

         Available Power[0]  0b01 0 00 001       0x00            (1 byte)            0X01 [8-bit Integer]
             Resource Instance with Value for use within a multiple Resource TLV
             8 Bits Identifier Field
             No Length Field, the value immediately follows the Identifier field in is of the length indicated by Bits 2-0 of this field
             Length of Value = 1 byte
             Value 1

         Available Power[1]  0b01 0 00 001       0x01            (1 byte)            0X05 [8-bit Integer]
             Resource Instance with Value for use within a multiple Resource TLV
             8 Bits Identifier Field
             No Length Field, the value immediately follows the Identifier field in is of the length indicated by Bits 2-0 of this field
             Length of Value = 1 byte
             Value 5

         Type Byte
             7-6         identifier
             5           0 = 8 bits, 1 = 16 bits
             4-3         00 = bit 2-0 is the length field
                         01 = length field is 8 bits, ignore 2-0
                         10 = length field is 16 bits, ignore 2-0
                         11 = length field is 24 bits, ignore 2-0
             2-0         length of value

         identifier
             8-bit or 16-bit

         length
             0 - 24 bit
         value
             sequence of bytes in Length

     */

    // ((binval & 0xC0) >> 6)





    /*
    if len(o.Instances) > 0 {
        // Root Resource with Instances for Values
    } else {
        // Root Resource with Multiple Values
    }
    */

}

/*
    | type | identifier | lenght |

    type:
        00  object instance
        01  resource instance with value
        10  multiple resource
        11  resource with value

*/


/*
func IntResourceToTlv(model *ResourceModel, values []int) (ResourceValue, error) {

}

func Int16ResourceToTlv(model *ResourceModel, values []int16) (ResourceValue, error) {

}

func Int32ResourceToTlv(model *ResourceModel, values []int32) (ResourceValue, error) {

}

func Int64ResourceToTlv(model *ResourceModel, values []int64) (ResourceValue, error) {

}
*/

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
