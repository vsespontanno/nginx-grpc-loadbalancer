package client

import (
	"context"
	"fmt"

	"github.com/vsespontanno/nginx-grpc-loadbalancer/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	Endpoint       string
	pingPongClient proto.PingPongerClient
}

func NewGRPCClient(endpoint string) (*GRPCClient, error) {
	conn, err := grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error while making new client: ", err)
		return nil, err
	}
	client := proto.NewPingPongerClient(conn)
	return &GRPCClient{Endpoint: endpoint, pingPongClient: client}, nil
}

func (c *GRPCClient) PingPong(ctx context.Context, in *proto.PingRequest) (*proto.PongResponse, error) {
	return c.pingPongClient.PingPong(ctx, in)
}
