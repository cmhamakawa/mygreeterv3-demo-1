package demoserver

import (
	"context"
	"fmt"
	"net"

	pb "dev.azure.com/service-hub-flg/service_hub_validation/_git/service_hub_validation_service.git/mygreeterv3/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedMyGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	println("Running SayHello")

	return &pb.HelloReply{Message: "Hello from demoserver |" + in.GetName()}, nil
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Serve(options Options) error { // this has en error
	port := options.Port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	println("Listening on port: ", port)

	grpcServer := grpc.NewServer()
	pb.RegisterMyGreeterServer(grpcServer, s)
	reflection.Register(grpcServer)

	fmt.Printf("server listening at %v\n", lis.Addr())
	return grpcServer.Serve(lis)
}

func (s *Server) Cleanup() {
	// Implement any cleanup logic if needed
}


func (s *Server) Init(options Options) {
}
