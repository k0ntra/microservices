package ports

import (
	"github.com/k0ntra/microservices/payment/internal/application/core/domain"
)

type DBPort interface {
	Get(string) (domain.Payment, error)
	Save(*domain.Payment) error
}
