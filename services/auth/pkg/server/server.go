package server

import (
	"fmt"
	"log"
	"net"

	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/logic"
	grpc_handler "github.com/p4elkab35t/salyte_backend/services/auth/pkg/server/grpc_handlers"
	proto "github.com/p4elkab35t/salyte_backend/services/auth/pkg/server/proto"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	authLogic        *logic.AuthLogicService
	securityLogLogic *logic.SecurityLogLogicService
}

func NewGRPCServer(authLogic *logic.AuthLogicService, securityLogLogic *logic.SecurityLogLogicService) *GRPCServer {
	return &GRPCServer{authLogic: authLogic, securityLogLogic: securityLogLogic}
}

func (s *GRPCServer) Start() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	authHandler := grpc_handler.NewAuthHandler(s.authLogic)
	proto.RegisterAuthServiceServer(grpcServer, authHandler)
	securityLogHandler := grpc_handler.NewSecurityLogHandler(s.securityLogLogic)
	proto.RegisterSecurityLogsServiceServer(grpcServer, securityLogHandler)

	fmt.Println("gRPC Server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
