package pkgerror

type domain struct {
	message string
}

func NewDomain(message string) *domain {
	return &domain{message}
}

func (e *domain) Error() string {
	return e.message
}

func (e *domain) Code() ErrorCode {
	return DomainErrorCode
}
