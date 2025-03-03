package main

import (
	"context"
	"fmt"
	"log"

	"github.com/p4elkab35t/salyte_backend/services/message/config"
	"github.com/p4elkab35t/salyte_backend/services/message/pkg/logic"
	"github.com/p4elkab35t/salyte_backend/services/message/pkg/server"
	"github.com/p4elkab35t/salyte_backend/services/message/pkg/storage/database"
	"github.com/p4elkab35t/salyte_backend/services/message/pkg/storage/repository"
)

func main() {
	var cfg config.Config
	config.LoadConfig(&cfg)
	ctx := context.Background()

	pgConfig := cfg.Database.Postgres
	dbConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		pgConfig.Host,
		pgConfig.Port,
		pgConfig.User,
		pgConfig.Password,
		pgConfig.DBName,
		pgConfig.SSLMode,
	)

	//redisConfig := cfg.Redis

	// redisConnString := fmt.Sprintf("%s:%d",
	// 	redisConfig.Host,
	// 	redisConfig.Port,
	// )

	// Connect to the databases
	db, err := database.NewPostgreSQL(ctx, dbConnStr)
	if err != nil {
		log.Fatalf("unable to connect to the database: %v", err)
	}

	// Load repositories
	messageRepo := repository.NewPostgresMessageRepositorySQL(db.GetDB())

	// Load services
	messageService := logic.NewMessageService(messageRepo)
	// chatService := logic.NewChatService(messageRepo)
	reactionService := logic.NewReactionService(messageRepo)

	// Start gRPC server

	grpcServer := server.NewGRPCServer(messageService, reactionService)
	grpcServer.Start()

}
