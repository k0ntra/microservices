package ports

import (
	"github.com/k0ntra/microservices/order/internal/application/core/domain"
)

type PaymentPort interface {
	Charge(*domain.Order) error
}
