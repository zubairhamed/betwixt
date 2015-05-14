package core

func TlvPayloadFromObjects(enabled *ObjectEnabler) (*TlvPayload, error) {
    // log.Println("TLV < Payload from Objects")
    return NewTlvPayload(), nil
}

func TlvPayloadFromObjectInstance(o *ObjectInstance) (*TlvPayload, error) {
    // log.Println("TLV < Payload from Object Instance")
    return NewTlvPayload(), nil
}

func TlvPayloadFromResource(v *MultipleResourceInstanceValue , m *ResourceModel, o *Resource) (*TlvPayload, error) {
    // log.Println("TLV < Payload from Resource", v, m, o)


    /*
    if len(o.Instances) > 0 {
        // Root Resource with Instances for Values
    } else {
        // Root Resource with Multiple Values
    }
    */

    return NewTlvPayload(), nil
}