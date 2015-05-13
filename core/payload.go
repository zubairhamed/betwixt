package core

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

func NewTlvPayload() *TlvPayload {
    return nil
}

type TlvPayload struct {

}

func (p *TlvPayload) GetBytes() ([]byte) {
    return make([]byte, 0)
}

func (p *TlvPayload) Length() (int) {
    return 0
}

func (p *TlvPayload) String() (string) {
    return ""
}

type JsonSenmlPayload struct {

}

func (p *JsonSenmlPayload) GetBytes() ([]byte) {
    return make([]byte, 0)
}

func (p *JsonSenmlPayload) Length() (int) {
    return 0
}

func (p *JsonSenmlPayload) String() (string) {
    return ""
}