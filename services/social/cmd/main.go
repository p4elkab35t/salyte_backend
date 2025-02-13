package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

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
	followRepo := repository.NewPostgresFollowRepositorySQL(db.GetDB())
	communityRepo := repository.NewPostgresCommunityRepositorySQL(db.GetDB())
	postRepo := repository.NewPostgresPostRepositorySQL(db.GetDB())
	commentRepo := repository.NewPostgresCommentRepositorySQL(db.GetDB())
	interactionRepo := repository.NewPostgresInteractionRepositorySQL(db.GetDB())

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

	// Enable routes

	// Profile routes
	http.HandleFunc("/social/profile", profileHandler.GetProfile)
	http.HandleFunc("/social/profile/{userID}", profileHandler.UpdateProfile)
	http.HandleFunc("/social/profile/{userID}/settings", profileHandler.GetProfileSettings)
	http.HandleFunc("/social/profile/{userID}/settings", profileHandler.UpdateProfileSettings)

	// Follow routes

	http.HandleFunc("/social/follow", followHandler.FollowProfile)
	http.HandleFunc("/social/unfollow", followHandler.UnfollowProfile)
	http.HandleFunc("/social/following/{userID}", followHandler.GetFollowing)
	http.HandleFunc("/social/followers/{userID}", followHandler.GetFollowers)
	http.HandleFunc("/social/friends/{userID}", followHandler.GetFriends)
	http.HandleFunc("/social/friends/{userID}/add", followHandler.MakeFriend)
	http.HandleFunc("/social/friends/{userID}/remove", followHandler.Unfriend)

	// Community routes
	http.HandleFunc("/social/community", communityHandler.CreateCommunity)
	http.HandleFunc("/social/community/", communityHandler.GetCommunity)
	http.HandleFunc("/social/community/{communityID}", communityHandler.UpdateCommunity)
	http.HandleFunc("/social/community/{communityID}/members", communityHandler.GetCommunityMembers)

	// Post routes
	http.HandleFunc("/social/post", postHandler.CreatePost)
	http.HandleFunc("/social/post/", postHandler.GetPost)
	http.HandleFunc("/social/post/{postID}", postHandler.UpdatePost)
	http.HandleFunc("/social/post/{postID}", postHandler.DeletePost)
	http.HandleFunc("/social/post/community/{communityID}", postHandler.GetPostsByCommunity)
	http.HandleFunc("/social/post/user/{userID}", postHandler.GetPostsByUser)

	// Comment routes
	http.HandleFunc("/social/post/{postID}/comment", commentHandler.CreateComment)
	http.HandleFunc("/social/post/{postID}/comment/{commentID}", commentHandler.UpdateComment)
	http.HandleFunc("/social/post/{postID}/comment/{commentID}", commentHandler.DeleteComment)
	http.HandleFunc("/social/post/{postID}/comments", commentHandler.GetCommentsByPostID)

	// Interaction routes
	http.HandleFunc("/social/post/{postID}/comments", interactionHandler.GetPostComments)
	http.HandleFunc("/social/post/{postID}/likes", interactionHandler.GetPostLikes)
	http.HandleFunc("/social/post/{postID}/like", interactionHandler.LikePost)
	http.HandleFunc("/social/post/{postID}/unlike", interactionHandler.UnlikePost)

	// Start the HTTP server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
