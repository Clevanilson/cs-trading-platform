package errorc

import "fmt"

type domain struct {
	resource string
}

func NewDomain(resource string) *domain {
	return &domain{resource}
}

func (e *domain) Error() string {
	return fmt.Sprintf("Invalid %v", e.resource)
}

func (e *domain) Code() ErrorCode {
	return DomainErrorCode
}
