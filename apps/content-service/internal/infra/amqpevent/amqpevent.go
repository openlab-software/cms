package amqpevent

import (
	"github.com/patrickdevbr-portfolio/cms/libs/go-common/event"
	"github.com/patrickdevbr-portfolio/cms/libs/go-common/rabbitmq"
)

type RabbitMQEventPublisher struct {
	event.Publisher
	rabbitmq *rabbitmq.RabbitMQPublisher
}

func NewRabbitMQEventPublisher(rabbitmq *rabbitmq.RabbitMQPublisher) event.Publisher {
	return &RabbitMQEventPublisher{rabbitmq: rabbitmq}
}

func (p *RabbitMQEventPublisher) Publish(e event.Event) error {
	return p.rabbitmq.Publish(e.Event, e)
}
