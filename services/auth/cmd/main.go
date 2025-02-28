package main

import (
	"context"
	"fmt"
	"log"

	"github.com/p4elkab35t/salyte_backend/services/auth/config"
	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/logic"

	// "github.com/p4elkab35t/salyte_backend/services/auth/pkg/server/handlers"
	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/server"
	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/storage/database"
	"github.com/p4elkab35t/salyte_backend/services/auth/pkg/storage/repository"
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

	//redis, err := database.NewRedis(ctx, redisConnString, redisConfig.Password, redisConfig.DB)
	if err != nil {
		log.Fatalf("unable to connect to the redis: %v", err)
	}

	// load repositories

	userRepository := repository.NewPostgresUserRepositorySQL(db.GetDB())
	postgresSessionRepository := repository.NewPostgresSessionRepositorySQL(db.GetDB())

	// TODO: Implement Redis session repository
	//redisSessionRepository := repository.NewRedisSessionRepository(redis.GetClient())

	securityLogRepository := repository.NewPostgresSecurityLogRepositorySQL(db.GetDB())

	// load services

	// Logger service first (it will be injected into other services to log security importnant actions)

	securityLoggerService := logic.NewSecurityLogLogicService(securityLogRepository)

	authService := logic.NewAuthLogic(userRepository, postgresSessionRepository, securityLoggerService) //, redisSessionRepository)

	// load grpc server

	grpcServer := server.NewGRPCServer(authService, securityLoggerService)
	grpcServer.Start()

	// load handlers

	// signupHandler := handlers.NewSignUpHandler(authService)
	// signinHandler := handlers.NewSignInHandler(authService)
	// tokenHandler := handlers.NewTokenHandler(authService)

	// // enable routes

	// http.HandleFunc("/auth/signin", signinHandler.SignIn)
	// http.HandleFunc("/auth/signup", signupHandler.SignUp)
	// http.HandleFunc("/auth/verify", tokenHandler.VerifyToken)
	// http.HandleFunc("/auth/signout", tokenHandler.SignOut)
	// log.Fatal(http.ListenAndServe(":8081", nil))
	// log.Println("Server is running on port 8080")
}
