package core
import "log"

func TlvPayloadFromObjects(enabled *ObjectEnabler) (*TlvPayload, error) {
    // log.Println("TLV < Payload from Objects")
    return NewTlvPayload(), nil
}

func TlvPayloadFromObjectInstance(o *ObjectInstance) (*TlvPayload, error) {
    // log.Println("TLV < Payload from Object Instance")
    return NewTlvPayload(), nil
}

func TlvPayloadFromIntResource(model *ResourceModel, values []int) (*TlvPayload, error) {
    log.Println("TLV < Payload from Resource", model, values)

    var identifier byte
    if len(values) > 0 {
        identifier = 0xb10000110
    } else {
        identifier = 0xb01000001
    }

    /*
    if len(o.Instances) > 0 {
        // Root Resource with Instances for Values
    } else {
        // Root Resource with Multiple Values
    }
    */

    return NewTlvPayload(), nil
}

/*
    | type | identifier | lenght |

    type:
        00  object instance
        01  resource instance with value
        10  multiple resource
        11  resource with value

*/