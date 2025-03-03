package server

import (
	"fmt"
	"log"
	"net"

	"github.com/p4elkab35t/salyte_backend/services/social/pkg/logic"
	grpc_handler "github.com/p4elkab35t/salyte_backend/services/social/pkg/server/grpc_handlers"
	proto "github.com/p4elkab35t/salyte_backend/services/social/pkg/server/proto"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	profileService *logic.ProfileService
}

func NewGRPCServer(profileService *logic.ProfileService) *GRPCServer {
	return &GRPCServer{profileService: profileService}
}

func (s *GRPCServer) Start() {
	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen on port 50053: %v", err)
	}

	grpcServer := grpc.NewServer()
	profileHandler := grpc_handler.NewProfileHandler(s.profileService)
	proto.RegisterSocialServiceServer(grpcServer, profileHandler)

	fmt.Println("gRPC Server is running on port 50053...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
