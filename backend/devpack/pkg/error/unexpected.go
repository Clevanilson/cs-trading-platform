package pkgerror

type unexpected struct {
	cause error
}

func NewUnexpected(cause error) *unexpected {
	return &unexpected{cause}
}

func (e *unexpected) Error() string {
	return "Unexpected error"
}

func (e *unexpected) Code() ErrorCode {
	return UnexpectedErrorCode
}
