package main

import (
	"github.com/k0ntra/microservicess/payment/config"
	"github.com/k0ntra/microservicess/payment/internal/adapters/db"
	"github.com/k0ntra/microservicess/payment/internal/adapters/grpc"
	"github.com/k0ntra/microservicess/payment/internal/application/core/api"

	log "github.com/sirupsen/logrus"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
