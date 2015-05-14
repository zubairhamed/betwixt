package core
import "log"

func TlvPayloadFromObjects(enabled *ObjectEnabler) (*TlvPayload, error) {
    log.Println("TLV < Payload from Objects")
    return NewTlvPayload(), nil
}

func TlvPayloadFromObjectInstance(o *ObjectInstance) (*TlvPayload, error) {
    log.Println("TLV < Payload from Object Instance")
    return NewTlvPayload(), nil
}

func TlvPayloadFromResourceInstances(o *ResourceInstance) (*TlvPayload, error) {
    log.Println("TLV < Payload from Resource Instance")

    return NewTlvPayload(), nil
}