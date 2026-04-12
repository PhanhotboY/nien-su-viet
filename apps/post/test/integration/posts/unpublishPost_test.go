//go:build integration

package v1

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc/codes"

	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	testlogger "github.com/phanhotboy/nien-su-viet/libs/pkg/logger/test"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/createPost/v1/commands"
	createPostDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/createPost/v1/dto"
	unpublishPostCommands "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/unpublishPost/v1/commands"
	unpublishPostDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/unpublishPost/v1/dto"
	updatePostCommands "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/updatePost/v1/commands"
	updatePostDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/updatePost/v1/dto"

	getPostDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPost/v1/dto"
	getPostQuery "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPost/v1/queries"

	testhelper "github.com/phanhotboy/nien-su-viet/apps/post/test/shared/helper"
)

func TestUnpublishPost(t *testing.T) {
	var (
		createPostHandler    commands.CreatePostHandler
		updatePostHandler    updatePostCommands.UpdatePostHandler
		unpublishPostHandler unpublishPostCommands.UnpublishPostHandler
		getPostHandler       getPostQuery.GetPostHandler
		log                  testlogger.TestLogger
	)

	testhelper.GetDIServices(t, &createPostHandler, &updatePostHandler, &unpublishPostHandler, &getPostHandler, &log)

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

	// Helper function to publish a post
	publishTestPost := func(t *testing.T, postID string) {
		isPublished := true
		publishedAt := time.Now()

		cmd, err := updatePostCommands.NewUpdatePostCommand(
			&updatePostDto.UpdatePostRequest{
				ID:          postID,
				Published:   &isPublished,
				PublishedAt: &publishedAt,
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create update post command: %v", err)
		}

		ctx := context.Background()
		_, err = updatePostHandler.Handle(ctx, cmd)
		if err != nil {
			log.TestFatalf("Failed to publish test post: %v", err)
		}
	}

	t.Run("Unpublish a published post successfully", func(t *testing.T) {
		postID := createTestPost(t, "Post to unpublish", "Content that will be unpublished", "post-to-unpublish", "author-001")

		// First publish the post
		publishTestPost(t, postID)

		ctx := context.Background()

		cmd, err := unpublishPostCommands.NewUnpublishPostCommand(
			&unpublishPostDto.UnpublishPostRequest{
				ID: postID,
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create unpublish post command: %v", err)
		}

		res, err := unpublishPostHandler.Handle(ctx, cmd)
		if err != nil {
			log.TestFatalf("Expected post to be unpublished, but got error: %v", err)
		}

		if !res.Data.Success {
			log.TestErrorf("Expected success to be true, but got false")
		}

		// Verify the post is unpublished
		getCmd, err := getPostQuery.NewGetPostQuery(&getPostDto.GetPostQueryReq{
			IDOrSlug: postID,
		})
		if err != nil {
			log.TestFatalf("Failed to create get post query: %v", err)
		}

		getRes, err := getPostHandler.Handle(ctx, getCmd)
		if err != nil {
			log.TestFatalf("Expected to get unpublished post, but got error: %v", err)
		}

		if getRes.Data.Published {
			log.TestErrorf("Expected published to be false, but got true")
		}

		// Verify published_at timestamp is cleared
		if getRes.Data.PublishedAt != nil {
			log.TestErrorf("Expected published_at to be nil, but got %v", getRes.Data.PublishedAt)
		}
	})

	t.Run("Unpublish post with empty ID should fail", func(t *testing.T) {
		_, err := unpublishPostCommands.NewUnpublishPostCommand(
			&unpublishPostDto.UnpublishPostRequest{
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

	t.Run("Unpublish post with invalid ID format should fail", func(t *testing.T) {
		_, err := unpublishPostCommands.NewUnpublishPostCommand(
			&unpublishPostDto.UnpublishPostRequest{
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

	t.Run("Unpublish non-existent post should fail", func(t *testing.T) {
		ctx := context.Background()

		cmd, err := unpublishPostCommands.NewUnpublishPostCommand(
			&unpublishPostDto.UnpublishPostRequest{
				ID: "550e8400-e29b-41d4-a716-446655440000",
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create unpublish post command: %v", err)
		}

		_, err = unpublishPostHandler.Handle(ctx, cmd)
		if err != nil {
			if grpcStatus := grpcerrors.ParseError(err).GetStatus(); grpcStatus != codes.NotFound {
				log.TestErrorf("Expected error code to be NotFound, but got %s: %v", grpcStatus, err)
			}
			return
		}

		log.TestErrorf("Expected error when unpublishing non-existent post, but unpublish succeeded")
	})

	t.Run("Unpublish already unpublished post should succeed", func(t *testing.T) {
		postID := createTestPost(t, "Post already unpublished", "Content for already unpublished post", "post-already-unpublished", "author-002")
		ctx := context.Background()

		// Post is created but not published, so it's already unpublished
		cmd, err := unpublishPostCommands.NewUnpublishPostCommand(
			&unpublishPostDto.UnpublishPostRequest{
				ID: postID,
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create unpublish post command: %v", err)
		}

		res, err := unpublishPostHandler.Handle(ctx, cmd)
		if err != nil {
			log.TestFatalf("Expected to unpublish already unpublished post, but got error: %v", err)
		}

		if !res.Data.Success {
			log.TestErrorf("Expected success to be true, but got false")
		}

		// Verify the post is still unpublished
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

		if getRes.Data.Published {
			log.TestErrorf("Expected published to be false, but got true")
		}
	})
}
