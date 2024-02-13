package rpc

import (
	"log"
	"net"

	"hex/config"
	"hex/internal/adapters/inbound/grpc/pb"
	"hex/internal/ports/inbound"
	"hex/utils/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// Adapter implements the GRPCPort interface
type Adapter struct {
	configs *config.Config
	logger  logger.Logger
	api     inbound.APIPort
	pb.UnimplementedUrlShortnerServiceServer
}

// NewAdapter creates a new Adapter
func NewAdapter(configs *config.Config, api inbound.APIPort, logger logger.Logger) *Adapter {
	return &Adapter{
		configs: configs,
		logger:  logger,
		api:     api,
	}
}

// Run registers the ArithmeticServiceServer to a grpcServer and serves on
// the specified port
func (grpca *Adapter) Run() {
	var err error
	address := ":" + grpca.configs.AppPort
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen on port %v: %v", address, err)
	}

	// arithmeticServiceServer := grpca
	grpcServer := grpc.NewServer()

	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(grpcServer, healthServer)
	pb.RegisterUrlShortnerServiceServer(grpcServer, grpca)
	reflection.Register(grpcServer)
	// pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)
	log.Println("starting server ", address)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server over port %v: %v", address, err)
	}
	log.Println("stoping server")
}
