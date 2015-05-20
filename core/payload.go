package core



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