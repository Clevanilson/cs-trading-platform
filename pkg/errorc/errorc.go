package errorc

type ErrorC interface {
	Code() string
	Error() string
}
