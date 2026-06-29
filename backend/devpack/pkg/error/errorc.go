package pkgerror

type ErrorC interface {
	Code() ErrorCode
	Error() string
}

type ErrorCode string

const (
	DomainErrorCode     ErrorCode = "DomainError"
	NotFoundErrorCode   ErrorCode = "NotFoundError"
	UnexpectedErrorCode ErrorCode = "UnexpectedError"
)
