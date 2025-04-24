// ports/order_event_port.go
package ports

import "order-service/internal/domain"

type OrderEventPort interface {
	PublishOrderCreated(order domain.Order) error
}
