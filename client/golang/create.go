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

	req := &order.CreateOrderRequest{
		UserId: 1,
		OrderItems: []*order.OrderItem{
			{
				ProductCode: "sku002",
				UnitPrice:   75,
				Quantity:    3,
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.Create(ctx, req)
	if err != nil {
		log.Fatalf("Could not create order: %v", err)
	}

	log.Printf("Order created: %v", resp)
}
