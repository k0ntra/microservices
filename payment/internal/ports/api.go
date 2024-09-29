package ports

import (
	"github.com/k0ntra/microservices/payment/internal/application/core/domain"
)

type APIPort interface {
	Charge(domain.Payment) (domain.Payment, error)
}
