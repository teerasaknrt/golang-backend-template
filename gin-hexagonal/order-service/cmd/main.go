// cmd/main.go
package main

import (
	"order-service/internal/adapters/gin"
	"order-service/internal/adapters/kafka"
	"order-service/internal/app"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	kafkaProducer := kafka.NewOrderProducer([]string{"localhost:9092"}, "order.created")

	svc := &app.OrderService{
		EventPublisher: kafkaProducer,
		// OrderRepo: your DB impl here
	}

	ginadapter.NewOrderHandler(r, svc)

	r.Run(":8080")
}
