package betwixt

type Validator interface {
	Valid(val interface{}) bool
}

func NewRangeValidator(from int64, to int64) Validator {
	return &RangeValidator{
		from: from,
		to:   to,
	}
}

type RangeValidator struct {
	from int64
	to   int64
}

func (v *RangeValidator) Valid(val interface{}) bool {
	return true
}

func NewLengthValidator(len uint64) Validator {
	return &LengthValidator{
		len: len,
	}
}

type LengthValidator struct {
	len uint64
}

func (v *LengthValidator) Valid(val interface{}) bool {
	return true
}
