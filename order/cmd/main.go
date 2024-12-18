package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/k0ntra/microservices/order/config"
	"github.com/k0ntra/microservices/order/internal/adapters/db"
	"github.com/k0ntra/microservices/order/internal/adapters/grpc"
	"github.com/k0ntra/microservices/order/internal/adapters/payment"
	"github.com/k0ntra/microservices/order/internal/application/core/api"
)

func main() {
	// Returns internal/adapters/db/db.go:db.Adapter
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}

	// Returns internal/adapters/api/api.go:api.Application
	application := api.NewApplication(dbAdapter, paymentAdapter)

	// Returns internal/adapters/grpc/server.go:grpc.Adapter
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
