package server

import (
	"fmt"
	"log"
	"net"

	"github.com/p4elkab35t/salyte_backend/services/message/pkg/logic"
	grpc_handler "github.com/p4elkab35t/salyte_backend/services/message/pkg/server/grpc_handlers"
	proto "github.com/p4elkab35t/salyte_backend/services/message/pkg/server/proto"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	messageService  *logic.MessageService
	chatService     *logic.ChatService
	reactionService *logic.ReactionService
}

func NewGRPCServer(messageService *logic.MessageService, reactionService *logic.ReactionService) *GRPCServer {
	return &GRPCServer{messageService: messageService, reactionService: reactionService}
}

func (s *GRPCServer) Start() {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen on port 50052: %v", err)
	}

	grpcServer := grpc.NewServer()
	messageHandler := grpc_handler.NewMessageHandler(s.messageService)
	reactionHandler := grpc_handler.NewReactionHandler(s.reactionService)
	proto.RegisterMessagingServiceServer(grpcServer, messageHandler)
	proto.RegisterMessagingServiceServer(grpcServer, reactionHandler)

	// authHandler := grpc_handler.NewAuthHandler(s.authLogic)
	// proto.RegisterAuthServiceServer(grpcServer, authHandler)
	// securityLogHandler := grpc_handler.NewSecurityLogHandler(s.securityLogLogic)
	// proto.RegisterSecurityLogsServiceServer(grpcServer, securityLogHandler)

	fmt.Println("gRPC Server is running on port 50052...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
