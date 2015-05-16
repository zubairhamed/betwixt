package core
import (
    "log"
    "encoding/binary"
    "unsafe"
)

func TlvPayloadFromObjects(enabled *ObjectEnabler) (ResourceValue, error) {
    return NewTlvValue(make([]byte, 0)), nil
}

func TlvPayloadFromObjectInstance(o *ObjectEnabler) (ResourceValue, error) {

    return NewTlvValue(make([]byte, 0)), nil
}

func TlvPayloadFromIntResource(model *ResourceModel, values []int) (ResourceValue, error) {
    log.Println("TLV < Payload from Resource", model, values)

    // Resource Instances TLV
    var resourceInstanceBytes []byte
    if len(values) > 0 {
        // Type Byte
        var typeVal byte
        for i, o := range values {
            // Bit 7-6: identifier
            typeVal |= 64

            // Bit 5
            if i > 255 {
                typeVal |= 32
            }

            // Bit 4-3
            valSize := o
            log.Println("valSize", o)
            if valSize > 7 {
                if valSize < 256 {
                    typeVal |= 8
                } else  {
                    if valSize < 65535 {
                        typeVal |= 16
                    } else {
                        if valSize > 16777215 {
                            // Error, size exceeds allowed (> 16.7MB)
                        } else {
                            // Size is 16777215 or less
                            typeVal |= 24
                        }
                    }
                }
            } else {
                // Set bit 2-0 instead
                // typeVal |= byte(valSize)
                typeVal |= byte(unsafe.Sizeof(valSize)/8)
            }
            resourceInstanceBytes = append(resourceInstanceBytes, typeVal)

            // Identifier Byte
            if i > 255 {
                // 16-Bit
                bs := make([]byte, 2)
                binary.LittleEndian.PutUint16(bs, uint16(i))

                resourceInstanceBytes = append(resourceInstanceBytes, bs[0])
                resourceInstanceBytes = append(resourceInstanceBytes, bs[1])
            } else {
                // 8-Bit
                resourceInstanceBytes = append(resourceInstanceBytes, byte(i))
            }

            // Value Length
            if valSize > 7 {

            }

            // Value Byte
            if o > 255 {
                bs := make([]byte, 2)
                binary.LittleEndian.PutUint16(bs, uint16(o))

                resourceInstanceBytes = append(resourceInstanceBytes, bs[0])
                resourceInstanceBytes = append(resourceInstanceBytes, bs[1])
            } else {
                resourceInstanceBytes = append(resourceInstanceBytes, byte(o))
            }
        }
    }

    // Resource Root TLV
    var resourceTlv []byte
    var typeVal byte
    // Byte 7-6: identifier
    typeVal |= 128

    // Bit 5
    if model.Id > 255 {
        typeVal |= 32
    }

    // Bit 4-3
    resourceInstanceTlvSize := len(resourceInstanceBytes)
    log.Println("Resource Instance TLV Size/Length", resourceInstanceTlvSize)
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
    resourceTlv = append(resourceTlv, typeVal)

    // Identifier
    if model.Id > 255 {
        // 16-Bit
        bs := make([]byte, 2)
        binary.LittleEndian.PutUint16(bs, uint16(model.Id))

        resourceTlv = append(resourceTlv, bs[0])
        resourceTlv = append(resourceTlv, bs[1])
    } else {
        // 8-Bit
        log.Println("Model ID", model.Id)
        resourceTlv = append(resourceTlv, byte(model.Id))
    }

    // Value Length
    if resourceInstanceTlvSize > 7 {

    }

    // Append Resource Instances TLV to Resource TLV
    for _, b := range resourceInstanceBytes {
        resourceTlv = append(resourceTlv, b)
    }

    return NewTlvValue(resourceTlv), nil

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