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
	deletePostCommands "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/deletePost/v1/commands"
	deletePostDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/deletePost/v1/dto"

	getPostDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPost/v1/dto"
	getPostQuery "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPost/v1/queries"

	testhelper "github.com/phanhotboy/nien-su-viet/apps/post/test/shared/helper"
)

func TestDeletePost(t *testing.T) {
	var (
		createPostHandler commands.CreatePostHandler
		deletePostHandler deletePostCommands.DeletePostHandler
		getPostHandler    getPostQuery.GetPostHandler
		log               testlogger.TestLogger
	)

	testhelper.GetDIServices(t, &createPostHandler, &deletePostHandler, &getPostHandler, &log)

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

	t.Run("Delete a post successfully", func(t *testing.T) {
		postID := createTestPost(t, "Post to delete", "Content that will be deleted", "post-to-delete", "author-001")

		ctx := context.Background()

		cmd, err := deletePostCommands.NewDeletePostCommand(
			&deletePostDto.DeletePostRequest{
				ID: postID,
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create delete post command: %v", err)
		}

		res, err := deletePostHandler.Handle(ctx, cmd)
		if err != nil {
			log.TestFatalf("Expected post to be deleted, but got error: %v", err)
		}

		if !res.Data.Success {
			log.TestErrorf("Expected success to be true, but got false")
		}

		// Verify the post was deleted
		getCmd, err := getPostQuery.NewGetPostQuery(&getPostDto.GetPostQueryReq{
			IDOrSlug: postID,
		})
		if err != nil {
			log.TestFatalf("Failed to create get post query: %v", err)
		}

		_, err = getPostHandler.Handle(ctx, getCmd)
		if err != nil {
			if grpcStatus := grpcerrors.ParseError(err).GetStatus(); grpcStatus != codes.NotFound {
				log.TestErrorf("Expected error code to be NotFound after deletion, but got %s: %v", grpcStatus, err)
			}
			return
		}

		log.TestErrorf("Expected NotFound error after deleting post, but post was still found")
	})

	t.Run("Delete post with empty ID should fail", func(t *testing.T) {
		_, err := deletePostCommands.NewDeletePostCommand(
			&deletePostDto.DeletePostRequest{
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

	t.Run("Delete post with invalid ID format should fail", func(t *testing.T) {
		_, err := deletePostCommands.NewDeletePostCommand(
			&deletePostDto.DeletePostRequest{
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

	t.Run("Delete non-existent post should fail", func(t *testing.T) {
		ctx := context.Background()

		cmd, err := deletePostCommands.NewDeletePostCommand(
			&deletePostDto.DeletePostRequest{
				ID: "550e8400-e29b-41d4-a716-446655440000",
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create delete post command: %v", err)
		}

		_, err = deletePostHandler.Handle(ctx, cmd)
		if err != nil {
			if grpcStatus := grpcerrors.ParseError(err).GetStatus(); grpcStatus != codes.NotFound {
				log.TestErrorf("Expected error code to be NotFound, but got %s: %v", grpcStatus, err)
			}
			return
		}

		log.TestErrorf("Expected error when deleting non-existent post, but deletion succeeded")
	})

	t.Run("Delete post response contains correct ID", func(t *testing.T) {
		postID := createTestPost(t, "Post for ID verification", "Content for ID verification", "post-id-delete-verification", "author-002")

		ctx := context.Background()

		cmd, err := deletePostCommands.NewDeletePostCommand(
			&deletePostDto.DeletePostRequest{
				ID: postID,
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create delete post command: %v", err)
		}

		res, err := deletePostHandler.Handle(ctx, cmd)
		if err != nil {
			log.TestFatalf("Expected post to be deleted, but got error: %v", err)
		}

		if res.Data.ID != postID {
			log.TestErrorf("Expected response ID to be %s, but got %s", postID, res.Data.ID)
		}
	})
}
