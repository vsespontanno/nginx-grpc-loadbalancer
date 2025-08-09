package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vsespontanno/nginx-grpc-loadbalancer/internal/client"
	"github.com/vsespontanno/nginx-grpc-loadbalancer/internal/proto"
)

const endpoint = "localhost:2103"

func main() {
	ctx := context.Background()
	grpcClient, err := client.NewGRPCClient(endpoint)
	if err != nil {
		log.Fatal("Error while connecting grpc-client")
	}
	grpcClient.PingPong(ctx, &proto.PingRequest{Msg: "Checking"})
	fmt.Println("Ping is sent!")
}
