package core

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