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
	updatePostCommands "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/updatePost/v1/commands"
	updatePostDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/updatePost/v1/dto"

	getPostDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPost/v1/dto"
	getPostQuery "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPost/v1/queries"

	testhelper "github.com/phanhotboy/nien-su-viet/apps/post/test/integration/shared/helper"
)

func TestUpdatePost(t *testing.T) {
	var (
		createPostHandler commands.CreatePostHandler
		updatePostHandler updatePostCommands.UpdatePostHandler
		getPostHandler    getPostQuery.GetPostHandler
		log               testlogger.TestLogger
	)

	testhelper.GetDIServices(t, &createPostHandler, &updatePostHandler, &getPostHandler, &log)

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

	t.Run("Update post with title only", func(t *testing.T) {
		postID := createTestPost(t, "Original Title", "Original content", "original-title", "author-001")
		ctx := context.Background()

		newTitle := "Updated Title"
		cmd, err := updatePostCommands.NewUpdatePostCommand(
			&updatePostDto.UpdatePostRequest{
				ID:    postID,
				Title: &newTitle,
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create update post command: %v", err)
		}

		res, err := updatePostHandler.Handle(ctx, cmd)
		if err != nil {
			log.TestFatalf("Expected post to be updated, but got error: %v", err)
		}

		if !res.Data.Success {
			log.TestErrorf("Expected success to be true, but got false")
		}

		// Verify the update
		getCmd, err := getPostQuery.NewGetPostQuery(&getPostDto.GetPostQueryReq{
			IDOrSlug: postID,
		})
		if err != nil {
			log.TestFatalf("Failed to create get post query: %v", err)
		}

		getRes, err := getPostHandler.Handle(ctx, getCmd)
		if err != nil {
			log.TestFatalf("Expected to get updated post, but got error: %v", err)
		}

		if getRes.Data.Title != newTitle {
			log.TestErrorf("Expected title to be %s, but got %s", newTitle, getRes.Data.Title)
		}

		// Verify other fields remain unchanged
		if getRes.Data.Content != "Original content" {
			log.TestErrorf("Expected content to remain unchanged, but got %s", getRes.Data.Content)
		}
	})

	t.Run("Update post with multiple fields", func(t *testing.T) {
		postID := createTestPost(t, "Title to Update", "Content to update", "content-to-update", "author-002")
		ctx := context.Background()

		newTitle := "New Updated Title"
		newContent := "New updated content"
		newSummary := "This is a new summary for the post"

		cmd, err := updatePostCommands.NewUpdatePostCommand(
			&updatePostDto.UpdatePostRequest{
				ID:      postID,
				Title:   &newTitle,
				Content: &newContent,
				Summary: &newSummary,
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create update post command: %v", err)
		}

		res, err := updatePostHandler.Handle(ctx, cmd)
		if err != nil {
			log.TestFatalf("Expected post to be updated, but got error: %v", err)
		}

		if !res.Data.Success {
			log.TestErrorf("Expected success to be true, but got false")
		}

		// Verify the update
		getCmd, err := getPostQuery.NewGetPostQuery(&getPostDto.GetPostQueryReq{
			IDOrSlug: postID,
		})
		if err != nil {
			log.TestFatalf("Failed to create get post query: %v", err)
		}

		getRes, err := getPostHandler.Handle(ctx, getCmd)
		if err != nil {
			log.TestFatalf("Expected to get updated post, but got error: %v", err)
		}

		if getRes.Data.Title != newTitle {
			log.TestErrorf("Expected title to be %s, but got %s", newTitle, getRes.Data.Title)
		}

		if getRes.Data.Content != newContent {
			log.TestErrorf("Expected content to be %s, but got %s", newContent, getRes.Data.Content)
		}

		if *getRes.Data.Summary != newSummary {
			log.TestErrorf("Expected summary to be %s, but got %s", newSummary, *getRes.Data.Summary)
		}
	})

	t.Run("Update post with empty ID should fail", func(t *testing.T) {
		newTitle := "Some title"
		_, err := updatePostCommands.NewUpdatePostCommand(
			&updatePostDto.UpdatePostRequest{
				ID:    "",
				Title: &newTitle,
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

	t.Run("Update post with invalid ID format should fail", func(t *testing.T) {
		newTitle := "Some title"
		_, err := updatePostCommands.NewUpdatePostCommand(
			&updatePostDto.UpdatePostRequest{
				ID:    "not-a-valid-uuid",
				Title: &newTitle,
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

	t.Run("Update non-existent post should fail", func(t *testing.T) {
		ctx := context.Background()
		newTitle := "Some title"

		// Use a valid UUID but non-existent post ID
		cmd, err := updatePostCommands.NewUpdatePostCommand(
			&updatePostDto.UpdatePostRequest{
				ID:    "550e8400-e29b-41d4-a716-446655440000",
				Title: &newTitle,
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create update post command: %v", err)
		}

		_, err = updatePostHandler.Handle(ctx, cmd)
		if err != nil {
			if grpcStatus := grpcerrors.ParseError(err).GetStatus(); grpcStatus != codes.NotFound {
				log.TestErrorf("Expected error code to be NotFound, but got %s: %v", grpcStatus, err)
			}
			return
		}

		log.TestErrorf("Expected error when updating non-existent post, but update succeeded")
	})

	t.Run("Update post with invalid thumbnail URL should fail", func(t *testing.T) {
		postID := createTestPost(t, "Post for thumbnail test", "Content", "post-thumbnail-test", "author-003")

		invalidURL := "not-a-valid-url"
		_, err := updatePostCommands.NewUpdatePostCommand(
			&updatePostDto.UpdatePostRequest{
				ID:        postID,
				Thumbnail: &invalidURL,
			},
		)

		if err != nil {
			if grpcStatus := grpcerrors.ParseError(err).GetStatus(); grpcStatus != codes.InvalidArgument {
				log.TestErrorf("Expected error code to be InvalidArgument, but got %s: %v", grpcStatus, err)
			}
			return
		}

		log.TestErrorf("Expected validation error for invalid thumbnail URL, but command was created")
	})

	t.Run("Update post with valid thumbnail URL", func(t *testing.T) {
		postID := createTestPost(t, "Post for valid thumbnail", "Content", "post-valid-thumbnail", "author-004")
		ctx := context.Background()

		validURL := "https://example.com/image.jpg"
		cmd, err := updatePostCommands.NewUpdatePostCommand(
			&updatePostDto.UpdatePostRequest{
				ID:        postID,
				Thumbnail: &validURL,
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create update post command: %v", err)
		}

		res, err := updatePostHandler.Handle(ctx, cmd)
		if err != nil {
			log.TestFatalf("Expected post to be updated with thumbnail, but got error: %v", err)
		}

		if !res.Data.Success {
			log.TestErrorf("Expected success to be true, but got false")
		}

		// Verify the thumbnail was set
		getCmd, err := getPostQuery.NewGetPostQuery(&getPostDto.GetPostQueryReq{
			IDOrSlug: postID,
		})
		if err != nil {
			log.TestFatalf("Failed to create get post query: %v", err)
		}

		getRes, err := getPostHandler.Handle(ctx, getCmd)
		if err != nil {
			log.TestFatalf("Expected to get updated post, but got error: %v", err)
		}

		if *getRes.Data.Thumbnail != validURL {
			log.TestErrorf("Expected thumbnail to be %s, but got %v", validURL, *getRes.Data.Thumbnail)
		}
	})

	t.Run("Update post slug to duplicate slug should fail", func(t *testing.T) {
		ctx := context.Background()

		// Create first post
		createTestPost(t, "First Post", "Content 1", "first-post-unique", "author-005")

		// Create second post
		postID2 := createTestPost(t, "Second Post", "Content 2", "second-post-unique", "author-006")

		// Try to update second post's slug to match first post's slug
		firstPostSlug := "first-post-unique"
		cmd, err := updatePostCommands.NewUpdatePostCommand(
			&updatePostDto.UpdatePostRequest{
				ID:   postID2,
				Slug: &firstPostSlug,
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create update post command: %v", err)
		}

		_, err = updatePostHandler.Handle(ctx, cmd)
		if err != nil {
			if grpcStatus := grpcerrors.ParseError(err).GetStatus(); grpcStatus != codes.AlreadyExists {
				log.TestErrorf("Expected error code to be AlreadyExists, but got %s: %v", grpcStatus, err)
			}
			return
		}

		log.TestErrorf("Expected error when updating post slug to duplicate, but update succeeded")
	})

	t.Run("Update post with publish status", func(t *testing.T) {
		postID := createTestPost(t, "Post for publishing", "Content", "post-for-publishing", "author-007")
		ctx := context.Background()

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

		res, err := updatePostHandler.Handle(ctx, cmd)
		if err != nil {
			log.TestFatalf("Expected post to be updated with publish status, but got error: %v", err)
		}

		if !res.Data.Success {
			log.TestErrorf("Expected success to be true, but got false")
		}

		// Verify the published status
		getCmd, err := getPostQuery.NewGetPostQuery(&getPostDto.GetPostQueryReq{
			IDOrSlug: postID,
		})
		if err != nil {
			log.TestFatalf("Failed to create get post query: %v", err)
		}

		getRes, err := getPostHandler.Handle(ctx, getCmd)
		if err != nil {
			log.TestFatalf("Expected to get updated post, but got error: %v", err)
		}

		if !getRes.Data.Published {
			log.TestErrorf("Expected published to be true, but got false")
		}
	})

	t.Run("Update post with invalid category ID should fail", func(t *testing.T) {
		postID := createTestPost(t, "Post for category test", "Content", "post-category-test", "author-008")

		invalidCategoryID := "not-a-valid-uuid"
		_, err := updatePostCommands.NewUpdatePostCommand(
			&updatePostDto.UpdatePostRequest{
				ID:         postID,
				CategoryID: &invalidCategoryID,
			},
		)

		if err != nil {
			if grpcStatus := grpcerrors.ParseError(err).GetStatus(); grpcStatus != codes.InvalidArgument {
				log.TestErrorf("Expected error code to be InvalidArgument, but got %s: %v", grpcStatus, err)
			}
			return
		}

		log.TestErrorf("Expected validation error for invalid category ID, but command was created")
	})

	t.Run("Update post with title having length constraints", func(t *testing.T) {
		postID := createTestPost(t, "Original", "Content", "original-title-constraints", "author-009")

		// Title too long (max 255)
		veryLongTitle := ""
		for i := 0; i < 260; i++ {
			veryLongTitle += "a"
		}

		_, err := updatePostCommands.NewUpdatePostCommand(
			&updatePostDto.UpdatePostRequest{
				ID:    postID,
				Title: &veryLongTitle,
			},
		)

		if err != nil {
			if grpcStatus := grpcerrors.ParseError(err).GetStatus(); grpcStatus != codes.InvalidArgument {
				log.TestErrorf("Expected error code to be InvalidArgument, but got %s: %v", grpcStatus, err)
			}
			return
		}

		log.TestErrorf("Expected validation error for title exceeding max length, but command was created")
	})

	t.Run("Update post with empty title should fail", func(t *testing.T) {
		postID := createTestPost(t, "Original", "Content", "original-empty-title", "author-010")

		emptyTitle := ""
		_, err := updatePostCommands.NewUpdatePostCommand(
			&updatePostDto.UpdatePostRequest{
				ID:    postID,
				Title: &emptyTitle,
			},
		)

		if err != nil {
			if grpcStatus := grpcerrors.ParseError(err).GetStatus(); grpcStatus != codes.InvalidArgument {
				log.TestErrorf("Expected error code to be InvalidArgument, but got %s: %v", grpcStatus, err)
			}
			return
		}

		log.TestErrorf("Expected validation error for empty title, but command was created")
	})

	t.Run("Update post with empty slug should fail", func(t *testing.T) {
		postID := createTestPost(t, "Original", "Content", "original-empty-slug", "author-011")

		emptySlug := ""
		_, err := updatePostCommands.NewUpdatePostCommand(
			&updatePostDto.UpdatePostRequest{
				ID:   postID,
				Slug: &emptySlug,
			},
		)

		if err != nil {
			if grpcStatus := grpcerrors.ParseError(err).GetStatus(); grpcStatus != codes.InvalidArgument {
				log.TestErrorf("Expected error code to be InvalidArgument, but got %s: %v", grpcStatus, err)
			}
			return
		}

		log.TestErrorf("Expected validation error for empty slug, but command was created")
	})

	t.Run("Update post with empty content should fail", func(t *testing.T) {
		postID := createTestPost(t, "Original", "Content", "original-empty-content", "author-012")

		emptyContent := ""
		_, err := updatePostCommands.NewUpdatePostCommand(
			&updatePostDto.UpdatePostRequest{
				ID:      postID,
				Content: &emptyContent,
			},
		)

		if err != nil {
			if grpcStatus := grpcerrors.ParseError(err).GetStatus(); grpcStatus != codes.InvalidArgument {
				log.TestErrorf("Expected error code to be InvalidArgument, but got %s: %v", grpcStatus, err)
			}
			return
		}

		log.TestErrorf("Expected validation error for empty content, but command was created")
	})

	t.Run("Update post with summary exceeding max length should fail", func(t *testing.T) {
		postID := createTestPost(t, "Original", "Content", "original-summary-length", "author-013")

		// Summary too long (max 500)
		veryLongSummary := ""
		for i := 0; i < 510; i++ {
			veryLongSummary += "a"
		}

		_, err := updatePostCommands.NewUpdatePostCommand(
			&updatePostDto.UpdatePostRequest{
				ID:      postID,
				Summary: &veryLongSummary,
			},
		)

		if err != nil {
			if grpcStatus := grpcerrors.ParseError(err).GetStatus(); grpcStatus != codes.InvalidArgument {
				log.TestErrorf("Expected error code to be InvalidArgument, but got %s: %v", grpcStatus, err)
			}
			return
		}

		log.TestErrorf("Expected validation error for summary exceeding max length, but command was created")
	})
}
