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
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Serve(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMyGreeterServer(grpcServer, s)
	reflection.Register(grpcServer)

	fmt.Printf("server listening at %v\n", lis.Addr())
	return grpcServer.Serve(lis)
}

func (s *Server) Cleanup() {
	// Implement any cleanup logic if needed
}
