//go:build integration

package v1

import (
	"context"
	"testing"

	sharedDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getAllPosts/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getAllPosts/v1/queries"
	testlogger "github.com/phanhotboy/nien-su-viet/libs/pkg/logger/test"

	createPostCmd "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/createPost/v1/commands"
	createPostDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/createPost/v1/dto"

	updatePostCmd "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/updatePost/v1/commands"
	updatePostDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/updatePost/v1/dto"

	testhelper "github.com/phanhotboy/nien-su-viet/apps/post/test/shared/helper"
)

func TestGetAllPosts(t *testing.T) {
	var (
		getAllPostsHandler queries.GetAllPostsHandler
		createPostHandler  createPostCmd.CreatePostHandler
		updatePostHandler  updatePostCmd.UpdatePostHandler
		log                testlogger.TestLogger
		limit              uint32 = 10
		page               uint32 = 1
	)

	testhelper.GetDIServices(t, &getAllPostsHandler, &createPostHandler, &updatePostHandler, &log)

	t.Run("Get all posts when empty", func(t *testing.T) {
		cmd, err := queries.NewGetAllPostsQuery(&dto.GetAllPostsQueryReq{
			PostListQueryDto: sharedDto.PostListQueryDto{Page: page, Limit: limit},
		})
		if err != nil {
			log.TestFatalf("Failed to create get all posts query: %v", err)
		}

		ctx := context.Background()
		res, err := getAllPostsHandler.Handle(ctx, cmd)
		if err != nil {
			log.TestFatalf("Expected to get all posts, but got error: %v", err)
		}

		if res.Data == nil || len(res.Data) != 0 {
			log.TestErrorf("Expected empty posts list, but got %d posts", len(res.Data))
		}
	})

	t.Run("Create multiple posts and get all posts regardless of publish status", func(t *testing.T) {
		ctx := context.Background()

		// Create first post
		mockPost1, err := createPostCmd.NewCreatePostCommand(
			&createPostDto.CreatePostRequest{
				Title:    "First Test Post",
				Content:  "This is the first test post.",
				Slug:     "first-test-post",
				AuthorId: "author-1",
			},
		)

		res1, err := createPostHandler.Handle(ctx, mockPost1)
		if err != nil {
			log.TestFatalf("Expected first post to be created, but got error: %v", err)
		}

		// Create second post
		mockPost2, err := createPostCmd.NewCreatePostCommand(
			&createPostDto.CreatePostRequest{
				Title:    "Second Test Post",
				Content:  "This is the second test post.",
				Slug:     "second-test-post",
				AuthorId: "author-2",
			},
		)

		res2, err := createPostHandler.Handle(ctx, mockPost2)
		if err != nil {
			log.TestFatalf("Expected second post to be created, but got error: %v", err)
		}

		// Create third post
		mockPost3, err := createPostCmd.NewCreatePostCommand(
			&createPostDto.CreatePostRequest{
				Title:    "Third Test Post",
				Content:  "This is the third test post.",
				Slug:     "third-test-post",
				AuthorId: "author-3",
			},
		)

		res3, err := createPostHandler.Handle(ctx, mockPost3)
		if err != nil {
			log.TestFatalf("Expected third post to be created, but got error: %v", err)
		}

		// Get all posts - should return all 3 posts even though none are published yet
		getAllCmd, err := queries.NewGetAllPostsQuery(&dto.GetAllPostsQueryReq{
			PostListQueryDto: sharedDto.PostListQueryDto{Page: page, Limit: limit},
		})
		if err != nil {
			log.TestFatalf("Failed to create get all posts query: %v", err)
		}
		allRes, err := getAllPostsHandler.Handle(ctx, getAllCmd)
		if err != nil {
			log.TestFatalf("Expected to get all posts, but got error: %v", err)
		}

		if len(allRes.Data) != 3 {
			log.TestErrorf("Expected 3 posts, but got %d posts", len(allRes.Data))
		}

		// Verify that all created posts are in the result
		postTitles := make(map[string]bool)
		for _, post := range allRes.Data {
			postTitles[post.Title] = true
		}

		if !postTitles["First Test Post"] {
			log.TestErrorf("Expected to find 'First Test Post' in results")
		}

		if !postTitles["Second Test Post"] {
			log.TestErrorf("Expected to find 'Second Test Post' in results")
		}

		if !postTitles["Third Test Post"] {
			log.TestErrorf("Expected to find 'Third Test Post' in results")
		}

		// Publish first and second posts
		published := new(bool)
		*published = true
		updateCmd1, err := updatePostCmd.NewUpdatePostCommand(
			&updatePostDto.UpdatePostRequest{
				ID:        res1.Data.ID,
				Published: published,
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create update post command: %v", err)
		}
		_, err = updatePostHandler.Handle(ctx, updateCmd1)
		if err != nil {
			log.TestFatalf("Expected first post to be updated, but got error: %v", err)
		}

		updateCmd2, err := updatePostCmd.NewUpdatePostCommand(
			&updatePostDto.UpdatePostRequest{
				ID:        res2.Data.ID,
				Published: published,
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create update post command: %v", err)
		}
		_, err = updatePostHandler.Handle(ctx, updateCmd2)
		if err != nil {
			log.TestFatalf("Expected second post to be updated, but got error: %v", err)
		}

		// Get all posts again - should still return all 3 posts (including unpublished)
		allRes, err = getAllPostsHandler.Handle(ctx, getAllCmd)
		if err != nil {
			log.TestFatalf("Expected to get all posts, but got error: %v", err)
		}

		if len(allRes.Data) != 3 {
			log.TestErrorf("Expected 3 posts (including unpublished), but got %d posts", len(allRes.Data))
		}

		// Publish third post
		updateCmd3, err := updatePostCmd.NewUpdatePostCommand(
			&updatePostDto.UpdatePostRequest{
				ID:        res3.Data.ID,
				Published: published,
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create update post command: %v", err)
		}
		_, err = updatePostHandler.Handle(ctx, updateCmd3)
		if err != nil {
			log.TestFatalf("Expected third post to be updated, but got error: %v", err)
		}

		// Get all posts - should still return all 3 posts
		allRes, err = getAllPostsHandler.Handle(ctx, getAllCmd)
		if err != nil {
			log.TestFatalf("Expected to get all posts, but got error: %v", err)
		}

		if len(allRes.Data) != 3 {
			log.TestErrorf("Expected 3 posts, but got %d posts", len(allRes.Data))
		}
	})
}
