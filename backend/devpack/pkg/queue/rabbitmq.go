package pkgqueue

import (
	"strings"

	"github.com/streadway/amqp"
)

type rabbitAdapter struct {
	conneection *amqp.Connection
	channel     *amqp.Channel
	queues      map[string]*amqp.Queue
}

func NewRabbitAdapter() (*rabbitAdapter, error) {
	connection, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672")
	if err != nil {
		return nil, err
	}
	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}
	return &rabbitAdapter{
		conneection: connection,
		channel:     channel,
		queues:      make(map[string]*amqp.Queue),
	}, nil
}

func (q *rabbitAdapter) SetupExchange(name string) error {
	return q.channel.ExchangeDeclare(name, "direct", true, false, false, false, nil)
}

func (q *rabbitAdapter) SetupQueue(name string) error {
	queue, err := q.channel.QueueDeclare(name, true, false, false, false, nil)
	if err != nil {
		return err
	}
	q.queues[name] = &queue
	return nil
}

func (q *rabbitAdapter) Close() error {
	err := q.conneection.Close()
	if err != nil {
		return err
	}
	return q.channel.Close()
}

func (q *rabbitAdapter) Publish(name string, payload []byte) error {
	return q.channel.Publish(name, "", true, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        payload,
	})
}

func (q *rabbitAdapter) Consume(name string, callback ConsumeCallback) {
	callback([]byte(`{}`))

}

func (q *rabbitAdapter) BindQueue(queue string) error {
	segments := strings.Split(queue, ".")
	return q.channel.QueueBind(queue, "", segments[0], false, nil)
}
