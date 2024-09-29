package ports

import (
	"github.com/k0ntra/microservices/order/internal/application/core/domain"
)

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
	GetOrder(int64) (domain.Order, error)
}
