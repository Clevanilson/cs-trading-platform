package pkgqueue

type Queue interface {
	Publish(exchange string, payload []byte) error
	Consume(queue string, callback ConsumeCallback) error
	Close() error
	SetupQueue(name string) error
}

type ConsumeCallback func(data []byte) error
