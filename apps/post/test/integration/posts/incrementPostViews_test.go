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
	incrementPostViewsCommands "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/incrementPostViews/v1/commands"
	incrementPostViewsDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/incrementPostViews/v1/dto"

	getPostDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPost/v1/dto"
	getPostQuery "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPost/v1/queries"

	testhelper "github.com/phanhotboy/nien-su-viet/apps/post/test/shared/helper"
)

func TestIncrementPostViews(t *testing.T) {
	var (
		createPostHandler         commands.CreatePostHandler
		incrementPostViewsHandler incrementPostViewsCommands.IncrementPostViewsHandler
		getPostHandler            getPostQuery.GetPostHandler
		log                       testlogger.TestLogger
	)

	testhelper.GetDIServices(t, &createPostHandler, &incrementPostViewsHandler, &getPostHandler, &log)

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

	t.Run("Increment post views successfully", func(t *testing.T) {
		postID := createTestPost(t, "Post to increment views", "Content to increment views", "post-increment-views", "author-001")

		ctx := context.Background()

		cmd, err := incrementPostViewsCommands.NewIncrementPostViewsCommand(
			&incrementPostViewsDto.IncrementPostViewsRequest{
				ID: postID,
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create increment post views command: %v", err)
		}

		res, err := incrementPostViewsHandler.Handle(ctx, cmd)
		if err != nil {
			log.TestFatalf("Expected post views to be incremented, but got error: %v", err)
		}

		if !res.Data.Success {
			log.TestErrorf("Expected success to be true, but got false")
		}

		// Response should show views as 1 (initial increment)
		if res.Data.Views != 1 {
			log.TestErrorf("Expected views in response to be 1, but got %d", res.Data.Views)
		}

		// Verify the views were incremented in the database
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

		if getRes.Data.Views != 1 {
			log.TestErrorf("Expected views in database to be 1, but got %d", getRes.Data.Views)
		}
	})

	t.Run("Increment post views multiple times", func(t *testing.T) {
		postID := createTestPost(t, "Post for multiple increments", "Content for multiple increments", "post-multiple-increments", "author-002")

		ctx := context.Background()

		// Increment views 5 times
		for i := 1; i <= 5; i++ {
			cmd, err := incrementPostViewsCommands.NewIncrementPostViewsCommand(
				&incrementPostViewsDto.IncrementPostViewsRequest{
					ID: postID,
				},
			)
			if err != nil {
				log.TestFatalf("Failed to create increment post views command on iteration %d: %v", i, err)
			}

			res, err := incrementPostViewsHandler.Handle(ctx, cmd)
			if err != nil {
				log.TestFatalf("Expected post views to be incremented on iteration %d, but got error: %v", i, err)
			}

			if !res.Data.Success {
				log.TestErrorf("Expected success to be true on iteration %d, but got false", i)
			}

			if res.Data.Views != i {
				log.TestErrorf("Expected views in response on iteration %d to be %d, but got %d", i, i, res.Data.Views)
			}
		}

		// Verify final view count
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

		if getRes.Data.Views != 5 {
			log.TestErrorf("Expected final views to be 5, but got %d", getRes.Data.Views)
		}
	})

	t.Run("Increment post views with empty ID should fail", func(t *testing.T) {
		_, err := incrementPostViewsCommands.NewIncrementPostViewsCommand(
			&incrementPostViewsDto.IncrementPostViewsRequest{
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

	t.Run("Increment post views with invalid ID format should fail", func(t *testing.T) {
		_, err := incrementPostViewsCommands.NewIncrementPostViewsCommand(
			&incrementPostViewsDto.IncrementPostViewsRequest{
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

	t.Run("Increment views for non-existent post should fail", func(t *testing.T) {
		ctx := context.Background()

		cmd, err := incrementPostViewsCommands.NewIncrementPostViewsCommand(
			&incrementPostViewsDto.IncrementPostViewsRequest{
				ID: "550e8400-e29b-41d4-a716-446655440000",
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create increment post views command: %v", err)
		}

		_, err = incrementPostViewsHandler.Handle(ctx, cmd)
		if err != nil {
			if grpcStatus := grpcerrors.ParseError(err).GetStatus(); grpcStatus != codes.NotFound {
				log.TestErrorf("Expected error code to be NotFound, but got %s: %v", grpcStatus, err)
			}
			return
		}

		log.TestErrorf("Expected error when incrementing views for non-existent post, but increment succeeded")
	})

	t.Run("Increment post views response contains correct ID", func(t *testing.T) {
		postID := createTestPost(t, "Post for ID verification", "Content for ID verification", "post-id-verification", "author-003")

		ctx := context.Background()

		cmd, err := incrementPostViewsCommands.NewIncrementPostViewsCommand(
			&incrementPostViewsDto.IncrementPostViewsRequest{
				ID: postID,
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create increment post views command: %v", err)
		}

		res, err := incrementPostViewsHandler.Handle(ctx, cmd)
		if err != nil {
			log.TestFatalf("Expected post views to be incremented, but got error: %v", err)
		}

		if res.Data.ID != postID {
			log.TestErrorf("Expected response ID to be %s, but got %s", postID, res.Data.ID)
		}
	})
}
