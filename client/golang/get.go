package main

import (
	"context"
	"log"
	"time"

	"github.com/k0ntra/microservices-proto/golang/order"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := order.NewOrderClient(conn)

	req := &order.GetOrderRequest{
		OrderId: 1,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.Get(ctx, req)
	if err != nil {
		log.Fatalf("Could not get order: %v", err)
	}

	log.Printf("Order: %v", resp)
}
