// app/order_service.go
package app

import (
	"order-service/internal/domain"
	"order-service/internal/ports"
)

type OrderService struct {
	EventPublisher ports.OrderEventPort
	// OrderRepo      ports.OrderRepositoryPort
}

func (s *OrderService) CreateOrder(order domain.Order) error {
	// 1. Save order
	// err := s.OrderRepo.Save(order)
	// if err != nil {
	// 	return err
	// }

	// 2. Publish event
	return s.EventPublisher.PublishOrderCreated(order)
}
