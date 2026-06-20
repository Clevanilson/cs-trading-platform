package errorc

type unexpected struct {
	cause error
}

func NewUnexpected(cause error) *unexpected {
	return &unexpected{cause}
}

func (e *unexpected) Error() string {
	return "Unexpected error"
}
