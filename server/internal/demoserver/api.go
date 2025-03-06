package demoserver

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "dev.azure.com/service-hub-flg/service_hub_validation/_git/service_hub_validation_service.git/mygreeterv3/api/v1"
	"google.golang.org/grpc"
	log "log/slog"
)

type Demoserver struct {
	pb.UnimplementedMyGreeterServer
}

func NewDemoserver() *Demoserver {
	return &Demoserver{}
}

func (s *Demoserver) Init(options Options) {
	logger := log.New(log.NewTextHandler(os.Stdout, nil))
	if options.JsonLog {
		logger = log.New(log.NewJSONHandler(os.Stdout, nil))
	}
	log.SetDefault(logger)
}

func (s *Demoserver) Serve(options Options) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", options.Port))
	if err != nil {
		log.Error("failed to listen: " + err.Error())
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMyGreeterServer(grpcServer, s)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Error("failed to serve: " + err.Error())
			os.Exit(1)
		}
	}()

	log.Info(fmt.Sprintf("demoserver listening at %v", lis.Addr()))

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	grpcServer.GracefulStop()
}

func (s *Demoserver) Cleanup() {
	// Add any necessary cleanup logic here
}

func (s *Demoserver) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Info("Received SayHello request: " + in.String())
	return &pb.HelloReply{Message: "Hello " + in.GetName() + " from demoserver"}, nil
}
