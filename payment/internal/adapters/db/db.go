package db

import (
	"fmt"

	"github.com/k0ntra/microservices/payment/internal/application/core/domain"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderID    int64
	TotalPrice float32
}

type Adapter struct {
	db *gorm.DB
}

func (a Adapter) Get(id string) (domain.Payment, error) {
	var paymentEntity Payment
	res := a.db.First(&paymentEntity, id)
	payment := domain.Payment{
		ID:         int64(paymentEntity.ID),
		CustomerID: paymentEntity.CustomerID,
		Status:     paymentEntity.Status,
		OrderId:    paymentEntity.OrderID,
		TotalPrice: paymentEntity.TotalPrice,
		CreatedAt:  paymentEntity.CreatedAt.UnixNano(),
	}
	return payment, res.Error
}

func (a Adapter) Save(payment *domain.Payment) error {
	orderModel := Payment{
		CustomerID: payment.CustomerID,
		Status:     payment.Status,
		OrderID:    payment.OrderId,
		TotalPrice: payment.TotlaPrice,
	}
	res := a.db.Create(&orderModel)
	if res.Error == nil {
		payment.ID = int64(orderMode.ID)
	}
	return res.Error
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}

	if err := db.Use(otelgorm.NewPlugin(otelgorm.WithDBName("payment"))); err != nil {
		return nil, fmt.Errorf("db otel plugin error: %v", err)
	}

	err := db.AutoMigrate(&Payment{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil
}