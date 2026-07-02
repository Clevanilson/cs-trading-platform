package pkgqueue

type MockQueue struct {
	PublishCalls              int
	PublishCalledWithExchange string
	PublishCalledWithPayload  []byte
}

func NewMockQueue() *MockQueue {
	return &MockQueue{}
}

func (q *MockQueue) Publish(exchange string, payload []byte) error {
	q.PublishCalls++
	q.PublishCalledWithExchange = exchange
	q.PublishCalledWithPayload = payload
	return nil
}

func (q *MockQueue) Consume(queue string, callback ConsumeCallback) error {
	return nil
}

func (q *MockQueue) Close() error {
	return nil
}

func (q *MockQueue) SetupQueue(name string) error {
	return nil
}
