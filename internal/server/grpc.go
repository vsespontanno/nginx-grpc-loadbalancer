package server

import (
	"context"
	"log"

	"github.com/vsespontanno/nginx-grpc-loadbalancer/internal/proto"
)

type GRPCServer struct {
	proto.UnimplementedPingPongerServer
}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
}

func (s *GRPCServer) PingPong(ctx context.Context, in *proto.PingRequest) (*proto.PongResponse, error) {
	log.Printf("Ping Received, msg: %s", in.Msg)
	return &proto.PongResponse{Msg: "Pong"}, nil
}
