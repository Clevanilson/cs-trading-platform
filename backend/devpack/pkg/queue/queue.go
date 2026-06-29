package pkgqueue

type Queue interface {
	Publish(exchange string, payload []byte)
	Consume(queue string, callback ConsumeCallback)
	Close() error
	SetupExchange(name string) error
	SetupQueue(name string) error
}

type ConsumeCallback func(data []byte) error
