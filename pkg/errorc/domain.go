package errorc

import "fmt"

type domain struct {
	resource string
}

func Newdomain(resource string) *domain {
	return &domain{resource}
}

func (e *domain) Error() string {
	return fmt.Sprintf("Inválid %v", e.resource)
}
