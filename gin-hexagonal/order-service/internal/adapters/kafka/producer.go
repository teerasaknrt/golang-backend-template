// adapters/kafka/producer.go
package kafka

import (
	"context"
	"encoding/json"
	"order-service/internal/domain"

	"github.com/segmentio/kafka-go"
)

type OrderProducer struct {
	Writer *kafka.Writer
}

func NewOrderProducer(brokers []string, topic string) *OrderProducer {
	w := &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	return &OrderProducer{Writer: w}
}

func (p *OrderProducer) PublishOrderCreated(order domain.Order) error {
	data, _ := json.Marshal(order)
	msg := kafka.Message{
		Key:   []byte(order.ID),
		Value: data,
	}
	return p.Writer.WriteMessages(context.Background(), msg)
}
