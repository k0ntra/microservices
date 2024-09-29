package api

import (
	"github.com/k0ntra/microservices/order/internal/application/core/domain"
	"github.com/k0ntra/microservices/order/internal/ports"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db:      db,
		payment: payment,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
	paymentErr := a.payment.Charge(&order)
	if paymentErr != nil {
		return domain.Order{}, err
	}
	return order, nil
}

func (a Application) GetOrder(id int64) (domain.Order, error) {
	return a.db.Get(id)
}
