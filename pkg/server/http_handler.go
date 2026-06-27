package server

type Handler func(input *HandlerInput) (*Response, error)

type HandlerInput struct {
	Params map[string]string
	Query  map[string]string
	Body   []byte
}
