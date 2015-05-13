package core
import "log"

func TlvPayloadFromObjects(enabled *ObjectEnabler) *TlvPayload {
    log.Println("TLV < Payload from Objects")
    return NewTlvPayload()
}

func TlvPayloadFromObjectInstance(o *ObjectInstance) *TlvPayload {
    log.Println("TLV < Payload from Object Instance")
    return NewTlvPayload()
}

func TlvPayloadFromResourceInstances(o *ResourceInstance) *TlvPayload {
    log.Println("TLV < Payload from Resource Instance")
    return NewTlvPayload()
}