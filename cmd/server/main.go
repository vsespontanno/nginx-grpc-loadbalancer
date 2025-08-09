package main

import (
	"fmt"
	"log"
	"net"

	"github.com/vsespontanno/nginx-grpc-loadbalancer/internal/proto"
	serv "github.com/vsespontanno/nginx-grpc-loadbalancer/internal/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	grpcEndpoint := ":2103"
	log.Fatal(makeGRPCTransport(grpcEndpoint))
}

func makeGRPCTransport(endpoint string) error {
	ln, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	server := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	proto.RegisterPingPongerServer(server, serv.NewGRPCServer())
	fmt.Println("GRPC transport running on port", endpoint)
	return server.Serve(ln)
}
