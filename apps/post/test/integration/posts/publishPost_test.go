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
	publishPostCommands "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/publishPost/v1/commands"
	publishPostDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/publishPost/v1/dto"

	getPostDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPost/v1/dto"
	getPostQuery "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPost/v1/queries"

	testhelper "github.com/phanhotboy/nien-su-viet/apps/post/test/shared/helper"
)

func TestPublishPost(t *testing.T) {
	var (
		createPostHandler  commands.CreatePostHandler
		publishPostHandler publishPostCommands.PublishPostHandler
		getPostHandler     getPostQuery.GetPostHandler
		log                testlogger.TestLogger
	)

	testhelper.GetDIServices(t, &createPostHandler, &publishPostHandler, &getPostHandler, &log)

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

	t.Run("Publish an unpublished post successfully", func(t *testing.T) {
		postID := createTestPost(t, "Post to publish", "Content that will be published", "post-to-publish", "author-001")

		ctx := context.Background()

		cmd, err := publishPostCommands.NewPublishPostCommand(
			&publishPostDto.PublishPostRequest{
				ID: postID,
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create publish post command: %v", err)
		}

		res, err := publishPostHandler.Handle(ctx, cmd)
		if err != nil {
			log.TestFatalf("Expected post to be published, but got error: %v", err)
		}

		if !res.Data.Success {
			log.TestErrorf("Expected success to be true, but got false")
		}

		// Verify the post is published
		getCmd, err := getPostQuery.NewGetPostQuery(&getPostDto.GetPostQueryReq{
			IDOrSlug: postID,
		})
		if err != nil {
			log.TestFatalf("Failed to create get post query: %v", err)
		}

		getRes, err := getPostHandler.Handle(ctx, getCmd)
		if err != nil {
			log.TestFatalf("Expected to get published post, but got error: %v", err)
		}

		if !getRes.Data.Published {
			log.TestErrorf("Expected published to be true, but got false")
		}

		// Verify published_at timestamp is set
		if getRes.Data.PublishedAt == nil {
			log.TestErrorf("Expected published_at to be set, but got nil")
		}
	})

	t.Run("Publish post with empty ID should fail", func(t *testing.T) {
		_, err := publishPostCommands.NewPublishPostCommand(
			&publishPostDto.PublishPostRequest{
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

	t.Run("Publish post with invalid ID format should fail", func(t *testing.T) {
		_, err := publishPostCommands.NewPublishPostCommand(
			&publishPostDto.PublishPostRequest{
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

	t.Run("Publish non-existent post should fail", func(t *testing.T) {
		ctx := context.Background()

		cmd, err := publishPostCommands.NewPublishPostCommand(
			&publishPostDto.PublishPostRequest{
				ID: "550e8400-e29b-41d4-a716-446655440000",
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create publish post command: %v", err)
		}

		_, err = publishPostHandler.Handle(ctx, cmd)
		if err != nil {
			if grpcStatus := grpcerrors.ParseError(err).GetStatus(); grpcStatus != codes.NotFound {
				log.TestErrorf("Expected error code to be NotFound, but got %s: %v", grpcStatus, err)
			}
			return
		}

		log.TestErrorf("Expected error when publishing non-existent post, but publish succeeded")
	})

	t.Run("Publish already published post should succeed", func(t *testing.T) {
		postID := createTestPost(t, "Post already published", "Content for already published post", "post-already-published", "author-002")
		ctx := context.Background()

		// Publish the post first
		cmd, err := publishPostCommands.NewPublishPostCommand(
			&publishPostDto.PublishPostRequest{
				ID: postID,
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create publish post command: %v", err)
		}

		res, err := publishPostHandler.Handle(ctx, cmd)
		if err != nil {
			log.TestFatalf("Failed to publish test post: %v", err)
		}

		if !res.Data.Success {
			log.TestErrorf("Expected success to be true, but got false")
		}

		// Try to publish again (should succeed)
		cmd2, err := publishPostCommands.NewPublishPostCommand(
			&publishPostDto.PublishPostRequest{
				ID: postID,
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create second publish post command: %v", err)
		}

		res2, err := publishPostHandler.Handle(ctx, cmd2)
		if err != nil {
			log.TestFatalf("Expected to publish already published post, but got error: %v", err)
		}

		if !res2.Data.Success {
			log.TestErrorf("Expected success to be true, but got false")
		}

		// Verify the post is still published
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

		if !getRes.Data.Published {
			log.TestErrorf("Expected published to be true, but got false")
		}
	})
}
