package errorc

import "fmt"

type notFound struct {
	resource string
}

func NewNotFound(resource string) *notFound {
	return &notFound{resource}
}

func (e *notFound) Error() string {
	return fmt.Sprintf("%v not found", e.resource)
}

func (e *notFound) Code() ErrorCode {
	return NotFoundErrorCode
}
