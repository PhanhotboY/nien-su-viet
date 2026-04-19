//go:build integration

package v1

import (
	"context"
	"testing"

	"google.golang.org/grpc/codes"

	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	testlogger "github.com/phanhotboy/nien-su-viet/libs/pkg/logger/test"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/createPost/v1/commands"
	createPostDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/createPost/v1/dto"
	incrementPostLikesCommands "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/incrementPostLikes/v1/commands"
	incrementPostLikesDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/incrementPostLikes/v1/dto"

	getPostDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPost/v1/dto"
	getPostQuery "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPost/v1/queries"

	testhelper "github.com/phanhotboy/nien-su-viet/apps/post/test/integration/shared/helper"
)

func TestIncrementPostLikes(t *testing.T) {
	var (
		createPostHandler         commands.CreatePostHandler
		incrementPostLikesHandler incrementPostLikesCommands.IncrementPostLikesHandler
		getPostHandler            getPostQuery.GetPostHandler
		log                       testlogger.TestLogger
	)

	testhelper.GetDIServices(t, &createPostHandler, &incrementPostLikesHandler, &getPostHandler, &log)

	// Helper function to create a post for testing
	createTestPost := func(t *testing.T, title, content, slug, authorID string) string {
		cmd, err := commands.NewCreatePostCommand(
			&createPostDto.CreatePostRequest{
				Title:    title,
				Content:  content,
				Slug:     slug,
				AuthorId: authorID,
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create test post command: %v", err)
		}

		ctx := context.Background()
		res, err := createPostHandler.Handle(ctx, cmd)
		if err != nil {
			log.TestFatalf("Failed to create test post: %v", err)
		}

		return res.Data.ID
	}

	t.Run("Increment post likes successfully", func(t *testing.T) {
		postID := createTestPost(t, "Post to increment likes", "Content to increment likes", "post-increment-likes", "author-001")

		ctx := context.Background()

		cmd, err := incrementPostLikesCommands.NewIncrementPostLikesCommand(
			&incrementPostLikesDto.IncrementPostLikesRequest{
				ID: postID,
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create increment post likes command: %v", err)
		}

		res, err := incrementPostLikesHandler.Handle(ctx, cmd)
		if err != nil {
			log.TestFatalf("Expected post likes to be incremented, but got error: %v", err)
		}

		if !res.Data.Success {
			log.TestErrorf("Expected success to be true, but got false")
		}

		// Response should show likes as 1 (initial increment)
		if res.Data.Likes != 1 {
			log.TestErrorf("Expected likes in response to be 1, but got %d", res.Data.Likes)
		}

		// Verify the likes were incremented in the database
		getCmd, err := getPostQuery.NewGetPostQuery(&getPostDto.GetPostQueryReq{
			IDOrSlug: postID,
		})
		if err != nil {
			log.TestFatalf("Failed to create get post query: %v", err)
		}

		getRes, err := getPostHandler.Handle(ctx, getCmd)
		if err != nil {
			log.TestFatalf("Expected to get post, but got error: %v", err)
		}

		if getRes.Data.Likes != 1 {
			log.TestErrorf("Expected likes in database to be 1, but got %d", getRes.Data.Likes)
		}
	})

	t.Run("Increment post likes multiple times", func(t *testing.T) {
		postID := createTestPost(t, "Post for multiple increments", "Content for multiple increments", "post-multiple-increments", "author-002")

		ctx := context.Background()

		// Increment likes 5 times
		for i := 1; i <= 5; i++ {
			cmd, err := incrementPostLikesCommands.NewIncrementPostLikesCommand(
				&incrementPostLikesDto.IncrementPostLikesRequest{
					ID: postID,
				},
			)
			if err != nil {
				log.TestFatalf("Failed to create increment post likes command on iteration %d: %v", i, err)
			}

			res, err := incrementPostLikesHandler.Handle(ctx, cmd)
			if err != nil {
				log.TestFatalf("Expected post likes to be incremented on iteration %d, but got error: %v", i, err)
			}

			if !res.Data.Success {
				log.TestErrorf("Expected success to be true on iteration %d, but got false", i)
			}

			if res.Data.Likes != i {
				log.TestErrorf("Expected likes in response on iteration %d to be %d, but got %d", i, i, res.Data.Likes)
			}
		}

		// Verify final like count
		getCmd, err := getPostQuery.NewGetPostQuery(&getPostDto.GetPostQueryReq{
			IDOrSlug: postID,
		})
		if err != nil {
			log.TestFatalf("Failed to create get post query: %v", err)
		}

		getRes, err := getPostHandler.Handle(ctx, getCmd)
		if err != nil {
			log.TestFatalf("Expected to get post, but got error: %v", err)
		}

		if getRes.Data.Likes != 5 {
			log.TestErrorf("Expected final likes to be 5, but got %d", getRes.Data.Likes)
		}
	})

	t.Run("Increment post likes with empty ID should fail", func(t *testing.T) {
		_, err := incrementPostLikesCommands.NewIncrementPostLikesCommand(
			&incrementPostLikesDto.IncrementPostLikesRequest{
				ID: "",
			},
		)

		if err != nil {
			if grpcStatus := grpcerrors.ParseError(err).GetStatus(); grpcStatus != codes.InvalidArgument {
				log.TestErrorf("Expected error code to be InvalidArgument, but got %s: %v", grpcStatus, err)
			}
			return
		}

		log.TestErrorf("Expected validation error for empty ID, but command was created")
	})

	t.Run("Increment post likes with invalid ID format should fail", func(t *testing.T) {
		_, err := incrementPostLikesCommands.NewIncrementPostLikesCommand(
			&incrementPostLikesDto.IncrementPostLikesRequest{
				ID: "not-a-valid-uuid",
			},
		)

		if err != nil {
			if grpcStatus := grpcerrors.ParseError(err).GetStatus(); grpcStatus != codes.InvalidArgument {
				log.TestErrorf("Expected error code to be InvalidArgument, but got %s: %v", grpcStatus, err)
			}
			return
		}

		log.TestErrorf("Expected validation error for invalid UUID format, but command was created")
	})

	t.Run("Increment likes for non-existent post should fail", func(t *testing.T) {
		ctx := context.Background()

		cmd, err := incrementPostLikesCommands.NewIncrementPostLikesCommand(
			&incrementPostLikesDto.IncrementPostLikesRequest{
				ID: "550e8400-e29b-41d4-a716-446655440000",
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create increment post likes command: %v", err)
		}

		_, err = incrementPostLikesHandler.Handle(ctx, cmd)
		if err != nil {
			if grpcStatus := grpcerrors.ParseError(err).GetStatus(); grpcStatus != codes.NotFound {
				log.TestErrorf("Expected error code to be NotFound, but got %s: %v", grpcStatus, err)
			}
			return
		}

		log.TestErrorf("Expected error when incrementing likes for non-existent post, but increment succeeded")
	})

	t.Run("Increment post likes response contains correct ID", func(t *testing.T) {
		postID := createTestPost(t, "Post for ID verification", "Content for ID verification", "post-id-verification", "author-003")

		ctx := context.Background()

		cmd, err := incrementPostLikesCommands.NewIncrementPostLikesCommand(
			&incrementPostLikesDto.IncrementPostLikesRequest{
				ID: postID,
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create increment post likes command: %v", err)
		}

		res, err := incrementPostLikesHandler.Handle(ctx, cmd)
		if err != nil {
			log.TestFatalf("Expected post likes to be incremented, but got error: %v", err)
		}

		if res.Data.ID != postID {
			log.TestErrorf("Expected response ID to be %s, but got %s", postID, res.Data.ID)
		}
	})
}
