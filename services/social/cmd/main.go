package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/p4elkab35t/salyte_backend/services/social/config"
	"github.com/p4elkab35t/salyte_backend/services/social/pkg/logic"
	"github.com/p4elkab35t/salyte_backend/services/social/pkg/server/handlers"
	"github.com/p4elkab35t/salyte_backend/services/social/pkg/storage/database"
	"github.com/p4elkab35t/salyte_backend/services/social/pkg/storage/repository"
)

func main() {
	var cfg config.Config
	config.LoadConfig(&cfg)
	ctx := context.Background()

	r := mux.NewRouter()

	pgConfig := cfg.Database.Postgres
	dbConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		pgConfig.Host,
		pgConfig.Port,
		pgConfig.User,
		pgConfig.Password,
		pgConfig.DBName,
		pgConfig.SSLMode,
	)

	// Connect to the PostgreSQL database
	db, err := database.NewPostgreSQL(ctx, dbConnStr)
	if err != nil {
		log.Fatalf("unable to connect to the database: %v", err)
	}

	// Load repositories
	profileRepo := repository.NewPostgresProfileRepositorySQL(db.GetDB())
	communityRepo := repository.NewPostgresCommunityRepositorySQL(db.GetDB())
	postRepo := repository.NewPostgresPostRepositorySQL(db.GetDB())
	commentRepo := repository.NewPostgresCommentRepositorySQL(db.GetDB())
	interactionRepo := repository.NewPostgresInteractionRepositorySQL(db.GetDB())

	followRepo := repository.NewPostgresFollowRepositorySQL(db.GetDB())

	// Load services
	profileService := logic.NewProfileService(profileRepo)
	communityService := logic.NewCommunityService(communityRepo)
	postService := logic.NewPostService(postRepo)
	interactionService := logic.NewInteractionService(interactionRepo)
	followService := logic.NewFollowService(followRepo)
	commentService := logic.NewCommentService(commentRepo)

	// Load handlers
	profileHandler := handlers.NewProfileHandler(profileService)
	communityHandler := handlers.NewCommunityHandler(communityService)
	postHandler := handlers.NewPostHandler(postService)
	interactionHandler := handlers.NewInteractionHandler(interactionService)
	followHandler := handlers.NewFollowHandler(followService)
	commentHandler := handlers.NewCommentHandler(commentService)

	// Middleware to inject user profile into request context

	profileMiddleware := InjectProfileMiddleware(profileService)

	// Group protected routes under middleware
	protectedRoutes := r.PathPrefix("/social").Subrouter()
	protectedRoutes.Use(profileMiddleware)

	publicRoutes := r.PathPrefix("/social").Subrouter()

	// Enable routes

	r.HandleFunc("/social", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the social service"))
	})

	// Profile routes
	publicRoutes.HandleFunc("/profile", profileHandler.CreateProfile).Methods("POST")
	publicRoutes.HandleFunc("/profile", profileHandler.GetProfile).Methods("GET")
	protectedRoutes.HandleFunc("/profile", profileHandler.UpdateProfile).Methods("PUT")
	protectedRoutes.HandleFunc("/profile/settings", profileHandler.GetProfileSettings).Methods("GET")
	protectedRoutes.HandleFunc("/profile/settings", profileHandler.UpdateProfileSettings).Methods("PUT")

	// Follow routes

	protectedRoutes.HandleFunc("/follow", followHandler.FollowProfile).Methods("POST")
	protectedRoutes.HandleFunc("/follow", followHandler.UnfollowProfile).Methods("DELETE")
	protectedRoutes.HandleFunc("/following", followHandler.GetFollowing).Methods("GET")
	protectedRoutes.HandleFunc("/followers", followHandler.GetFollowers).Methods("GET")
	protectedRoutes.HandleFunc("/friends", followHandler.GetFriends).Methods("GET")
	protectedRoutes.HandleFunc("/friends", followHandler.MakeFriend).Methods("POST")
	protectedRoutes.HandleFunc("/friends", followHandler.Unfriend).Methods("DELETE")

	// Community routes
	protectedRoutes.HandleFunc("/community", communityHandler.CreateCommunity).Methods("POST")
	protectedRoutes.HandleFunc("/community", communityHandler.GetCommunity).Methods("GET")
	protectedRoutes.HandleFunc("/community", communityHandler.UpdateCommunity).Methods("PUT")
	protectedRoutes.HandleFunc("/community/members", communityHandler.GetCommunityMembers).Methods("GET")

	// Post routes
	protectedRoutes.HandleFunc("/post", postHandler.CreatePost).Methods("POST")
	publicRoutes.HandleFunc("/post", postHandler.GetPost).Methods("GET")
	protectedRoutes.HandleFunc("/post", postHandler.UpdatePost).Methods("PUT")
	protectedRoutes.HandleFunc("/post", postHandler.DeletePost).Methods("DELETE")
	protectedRoutes.HandleFunc("/post/community", postHandler.GetPostsByCommunity).Methods("GET")
	publicRoutes.HandleFunc("/post/user", postHandler.GetPostsByUser).Methods("GET")

	// Comment routes
	protectedRoutes.HandleFunc("/post/comment", commentHandler.CreateComment).Methods("POST")
	protectedRoutes.HandleFunc("/post/comment", commentHandler.UpdateComment).Methods("PUT")
	protectedRoutes.HandleFunc("/post/comment", commentHandler.DeleteComment).Methods("DELETE")
	protectedRoutes.HandleFunc("/post/comments", commentHandler.GetCommentsByPostID).Methods("GET")

	// Interaction routes
	// r.HandleFunc("/post/comments", interactionHandler.GetPostComments).Methods("GET")
	protectedRoutes.HandleFunc("/post/likes", interactionHandler.GetPostLikes).Methods("GET")
	protectedRoutes.HandleFunc("/post/like", interactionHandler.LikePost).Methods("POST")
	protectedRoutes.HandleFunc("/post/like", interactionHandler.UnlikePost).Methods("DELETE")

	// Start the HTTP server
	log.Println("Server is running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))

}

// Middleware to inject user profile into request context
func InjectProfileMiddleware(profileService *logic.ProfileService) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID := r.URL.Query().Get("userID") // Example: Get from header (modify as needed)

			if userID == "" {
				http.Error(w, `{"error": "missing user ID"}`, http.StatusUnauthorized)
				return
			}

			// Fetch profile from database
			profile, err := profileService.GetProfileByUserID(r.Context(), userID)
			if err != nil {
				http.Error(w, `{"error": "profile not found"}`, http.StatusNotFound)
				return
			}

			// Store profile in request context
			ctx := context.WithValue(r.Context(), "profileID", profile.ProfileID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
