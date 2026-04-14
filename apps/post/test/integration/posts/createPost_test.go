//go:build integration

package v1

import (
	"context"
	"testing"

	"google.golang.org/grpc/codes"

	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	testlogger "github.com/phanhotboy/nien-su-viet/libs/pkg/logger/test"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/createPost/v1/commands"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/createPost/v1/dto"

	getPostDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPost/v1/dto"
	getPostQuery "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPost/v1/queries"

	testhelper "github.com/phanhotboy/nien-su-viet/apps/post/test/integration/shared/helper"
)

func TestCreatePost(t *testing.T) {
	var (
		createPostHandler commands.CreatePostHandler
		getPostHandler    getPostQuery.GetPostHandler
		log               testlogger.TestLogger
	)

	testhelper.GetDIServices(t, &createPostHandler, &getPostHandler, &log)

	t.Run("Create post with all required fields", func(t *testing.T) {
		cmd, err := commands.NewCreatePostCommand(
			&dto.CreatePostRequest{
				Title:    "Test Post Title",
				Content:  "This is test content for the post.",
				Slug:     "test-post-title",
				AuthorId: "author-123",
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create post command: %v", err)
		}

		ctx := context.Background()
		res, err := createPostHandler.Handle(ctx, cmd)
		if err != nil {
			log.TestFatalf("Expected post to be created, but got error: %v", err)
		}

		// Verify response contains ID
		if res.Data.ID == "" {
			log.TestErrorf("Expected post ID to be generated, but got empty ID")
		}

		// Verify success flag
		if !res.Data.Success {
			log.TestErrorf("Expected success to be true, but got false")
		}

		// Get the post to verify it was created correctly
		getCmd, err := getPostQuery.NewGetPostQuery(&getPostDto.GetPostQueryReq{
			IDOrSlug: res.Data.ID,
		})
		if err != nil {
			log.TestFatalf("Failed to create get post query: %v", err)
		}

		getRes, err := getPostHandler.Handle(ctx, getCmd)
		if err != nil {
			log.TestFatalf("Expected to get created post, but got error: %v", err)
		}

		if getRes.Data.Title != cmd.Title {
			log.TestErrorf("Expected title to be %s, but got %s", cmd.Title, getRes.Data.Title)
		}

		if getRes.Data.Content != cmd.Content {
			log.TestErrorf("Expected content to be %s, but got %s", cmd.Content, getRes.Data.Content)
		}

		if getRes.Data.Slug != cmd.Slug {
			log.TestErrorf("Expected slug to be %s, but got %s", cmd.Slug, getRes.Data.Slug)
		}

		if getRes.Data.AuthorID != cmd.AuthorId {
			log.TestErrorf("Expected author ID to be %s, but got %s", cmd.AuthorId, getRes.Data.AuthorID)
		}

		// Verify post is not published by default
		if getRes.Data.Published {
			log.TestErrorf("Expected post to be unpublished by default, but got published=true")
		}
	})

	t.Run("Create post with optional fields", func(t *testing.T) {
		description := "This is a test description"
		cmd, err := commands.NewCreatePostCommand(
			&dto.CreatePostRequest{
				Title:    "Test Post With Description",
				Content:  "This is test content with description.",
				Slug:     "test-post-with-description",
				AuthorId: "author-456",
				Summary:  &description,
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create post command: %v", err)
		}

		ctx := context.Background()
		res, err := createPostHandler.Handle(ctx, cmd)
		if err != nil {
			log.TestFatalf("Expected post to be created, but got error: %v", err)
		}

		if res.Data.ID == "" {
			log.TestErrorf("Expected post ID to be generated, but got empty ID")
		}

		// Get the post to verify the description was saved
		getCmd, err := getPostQuery.NewGetPostQuery(&getPostDto.GetPostQueryReq{
			IDOrSlug: res.Data.ID,
		})
		if err != nil {
			log.TestFatalf("Failed to create get post query: %v", err)
		}

		getRes, err := getPostHandler.Handle(ctx, getCmd)
		if err != nil {
			log.TestFatalf("Expected to get created post, but got error: %v", err)
		}

		if *getRes.Data.Summary != description {
			log.TestErrorf("Expected description to be %s, but got %s", description, *getRes.Data.Summary)
		}
	})

	t.Run("Create post with empty title should fail", func(t *testing.T) {
		_, err := commands.NewCreatePostCommand(
			&dto.CreatePostRequest{
				Title:    "",
				Content:  "This is test content.",
				Slug:     "empty-title-post",
				AuthorId: "author-789",
			},
		)

		if err != nil {
			if grpcStatus := grpcerrors.ParseError(err).GetStatus(); grpcStatus != codes.InvalidArgument {
				log.TestErrorf("Expected error code to be InvalidArgument, but got %s: %v", grpcStatus, err)
			}
			return
		}

		log.TestErrorf("Expected validation error for empty title, but post was created")
	})

	t.Run("Create post with empty content should fail", func(t *testing.T) {
		_, err := commands.NewCreatePostCommand(
			&dto.CreatePostRequest{
				Title:    "Empty Content Post",
				Content:  "",
				Slug:     "empty-content-post",
				AuthorId: "author-012",
			},
		)
		if err != nil {
			if grpcStatus := grpcerrors.ParseError(err).GetStatus(); grpcStatus != codes.InvalidArgument {
				log.TestErrorf("Expected error code to be InvalidArgument from command creation, but got %s: %v", grpcStatus, err)
			}
			return
		}

		log.TestErrorf("Expected validation error for empty content, but command validation passed")
	})

	t.Run("Create post with empty slug should fail", func(t *testing.T) {
		_, err := commands.NewCreatePostCommand(
			&dto.CreatePostRequest{
				Title:    "Empty Slug Post",
				Content:  "This is test content.",
				Slug:     "",
				AuthorId: "author-345",
			},
		)
		if err != nil {
			if grpcStatus := grpcerrors.ParseError(err).GetStatus(); grpcStatus != codes.InvalidArgument {
				log.TestErrorf("Expected error code to be InvalidArgument from command creation, but got %s: %v", grpcStatus, err)
			}
			return
		}

		log.TestErrorf("Expected validation error for empty slug, but command validation passed")
	})

	t.Run("Create post with empty author ID should fail", func(t *testing.T) {
		_, err := commands.NewCreatePostCommand(
			&dto.CreatePostRequest{
				Title:    "Empty Author Post",
				Content:  "This is test content.",
				Slug:     "empty-author-post",
				AuthorId: "",
			},
		)
		if err != nil {
			if grpcStatus := grpcerrors.ParseError(err).GetStatus(); grpcStatus != codes.InvalidArgument {
				log.TestErrorf("Expected error code to be InvalidArgument from command creation, but got %s: %v", grpcStatus, err)
			}
			return
		}

		log.TestErrorf("Expected validation error for empty author ID, but command validation passed")
	})

	t.Run("Create multiple posts with different slugs", func(t *testing.T) {
		ctx := context.Background()

		// Create first post
		cmd1, err := commands.NewCreatePostCommand(
			&dto.CreatePostRequest{
				Title:    "First Post",
				Content:  "Content of first post.",
				Slug:     "first-post-unique",
				AuthorId: "author-001",
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create first post command: %v", err)
		}

		res1, err := createPostHandler.Handle(ctx, cmd1)
		if err != nil {
			log.TestFatalf("Expected first post to be created, but got error: %v", err)
		}

		// Create second post
		cmd2, err := commands.NewCreatePostCommand(
			&dto.CreatePostRequest{
				Title:    "Second Post",
				Content:  "Content of second post.",
				Slug:     "second-post-unique",
				AuthorId: "author-002",
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create second post command: %v", err)
		}

		res2, err := createPostHandler.Handle(ctx, cmd2)
		if err != nil {
			log.TestFatalf("Expected second post to be created, but got error: %v", err)
		}

		// Verify both posts have different IDs
		if res1.Data.ID == res2.Data.ID {
			log.TestErrorf("Expected posts to have different IDs, but both have ID: %s", res1.Data.ID)
		}

		if res1.Data.ID == "" || res2.Data.ID == "" {
			log.TestErrorf("Expected both posts to have valid IDs")
		}
	})

	t.Run("Create post with duplicate slug should fail", func(t *testing.T) {
		ctx := context.Background()
		slug := "duplicate-slug-test"

		// Create first post with specific slug
		cmd1, err := commands.NewCreatePostCommand(
			&dto.CreatePostRequest{
				Title:    "First Post",
				Content:  "Content of first post.",
				Slug:     slug,
				AuthorId: "author-001",
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create first post command: %v", err)
		}

		_, err = createPostHandler.Handle(ctx, cmd1)
		if err != nil {
			log.TestFatalf("Expected first post to be created, but got error: %v", err)
		}

		// Try to create another post with the same slug
		cmd2, err := commands.NewCreatePostCommand(
			&dto.CreatePostRequest{
				Title:    "Duplicate Slug Post",
				Content:  "Content of duplicate post.",
				Slug:     slug,
				AuthorId: "author-002",
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create second post command: %v", err)
		}

		_, err = createPostHandler.Handle(ctx, cmd2)
		if err != nil {
			if grpcStatus := grpcerrors.ParseError(err).GetStatus(); grpcStatus != codes.AlreadyExists {
				log.TestErrorf("Expected error code to be AlreadyExists, but got %s: %v", grpcStatus, err)
			}
			return
		}

		log.TestErrorf("Expected error for duplicate slug, but post was created")
	})
}
