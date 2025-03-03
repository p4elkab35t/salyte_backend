package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"net/http"

	"github.com/gorilla/mux"

	"github.com/p4elkab35t/salyte_backend/services/message/config"
	"github.com/p4elkab35t/salyte_backend/services/message/pkg/logic"
	"github.com/p4elkab35t/salyte_backend/services/message/pkg/server"
	"github.com/p4elkab35t/salyte_backend/services/message/pkg/server/handlers"
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
	chatService := logic.NewChatService(messageRepo)
	reactionService := logic.NewReactionService(messageRepo)

	// Start gRPC server

	var wg sync.WaitGroup
	wg.Add(2)

	go startGRPCServer(&wg, messageService, reactionService)
	go startHTTPServer(&wg, messageService, chatService, reactionService)

	wg.Wait()
}

func startGRPCServer(wg *sync.WaitGroup, messageService *logic.MessageService, reactionService *logic.ReactionService) {
	defer wg.Done()

	grpcServer := server.NewGRPCServer(messageService, reactionService)
	grpcServer.Start()
}

func startHTTPServer(wg *sync.WaitGroup, messageService *logic.MessageService, chatService *logic.ChatService, reactionService *logic.ReactionService) {
	defer wg.Done()

	r := mux.NewRouter()

	// Load handlers

	messageHandler := handlers.NewMessageHandler(messageService)
	chatHandler := handlers.NewChatHandler(chatService)
	reactionHandler := handlers.NewreactionLogicHandler(reactionService)

	// Register handlers

	// message routes
	r.HandleFunc("/message/getallbychat", messageHandler.GetMessagesByChatID).Methods("GET")
	r.HandleFunc("/message/unread", messageHandler.GetUnreadMessages).Methods("GET")
	r.HandleFunc("/message/deleteall", messageHandler.DeleteAllMessagesByChatID).Methods("POST")

	// chat routes
	r.HandleFunc("/chat/get", chatHandler.GetChat).Methods("GET")
	r.HandleFunc("/chat/create", chatHandler.CreateChat).Methods("POST")
	r.HandleFunc("/chat/getall", chatHandler.GetAllChats).Methods("GET")
	r.HandleFunc("/chat/adduser", chatHandler.AddUserToChat).Methods("POST")
	r.HandleFunc("/chat/removeuser", chatHandler.RemoveUserFromChat).Methods("POST")
	r.HandleFunc("/chat/members", chatHandler.GetChatMembers).Methods("GET")
	r.HandleFunc("/chat/messages", chatHandler.GetChatByID).Methods("GET")

	// reaction routes
	r.HandleFunc("/reaction/get", reactionHandler.GetReactions).Methods("GET")

	log.Println("Server is running on port 8083")
	log.Fatal(http.ListenAndServe(":8083", r))
}
