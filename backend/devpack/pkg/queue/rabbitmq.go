package pkgqueue

import (
	"strings"

	pkgerror "github.com/clevanilson/cs-trading-platform/devpack/pkg/error"
	"github.com/streadway/amqp"
)

type rabbitAdapter struct {
	conneection *amqp.Connection
	channel     *amqp.Channel
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
	}, nil
}

func (q *rabbitAdapter) Close() error {
	err := q.conneection.Close()
	if err != nil {
		return err
	}
	return q.channel.Close()
}

func (q *rabbitAdapter) SetupQueue(name string) error {
	segments := strings.Split(name, ".")
	if len(segments) != 2 {
		return pkgerror.NewDomain("Invalid queue name")
	}
	if err := q.setupExchange(segments[0]); err != nil {
		return err
	}
	_, err := q.channel.QueueDeclare(name, true, false, false, false, nil)
	if err != nil {
		return err
	}
	if err := q.channel.QueueBind(name, "", segments[0], false, nil); err != nil {
		return err
	}
	return nil
}

func (q *rabbitAdapter) Publish(name string, payload []byte) error {
	return q.channel.Publish(name, "", true, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        payload,
	})
}

func (q *rabbitAdapter) Consume(name string, callback ConsumeCallback) error {
	messages, err := q.channel.Consume(name, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	for message := range messages {
		go func() {
			if err := callback(message.Body); err != nil {
				err = message.Nack(false, true)
			}
			message.Ack(false)
		}()
	}
	return nil
}

func (q *rabbitAdapter) setupExchange(name string) error {
	return q.channel.ExchangeDeclare(name, "direct", true, false, false, false, nil)
}
