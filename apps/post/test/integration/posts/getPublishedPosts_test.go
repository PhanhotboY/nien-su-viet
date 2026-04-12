//go:build integration

package v1

import (
	"context"
	"testing"

	sharedDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPublishedPosts/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPublishedPosts/v1/queries"
	testlogger "github.com/phanhotboy/nien-su-viet/libs/pkg/logger/test"

	createPostCmd "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/createPost/v1/commands"
	createPostDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/createPost/v1/dto"

	updatePostCmd "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/updatePost/v1/commands"
	updatePostDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/updatePost/v1/dto"

	testhelper "github.com/phanhotboy/nien-su-viet/apps/post/test/shared/helper"
)

func TestGetPublishedPosts(t *testing.T) {
	var (
		getPostsHandler   queries.GetPublishedPostsHandler
		createPostHandler createPostCmd.CreatePostHandler
		updatePostHandler updatePostCmd.UpdatePostHandler
		log               testlogger.TestLogger
		page              uint32 = 1
		limit             uint32 = 10
		published         bool   = true
	)

	testhelper.GetDIServices(t, &getPostsHandler, &createPostHandler, &updatePostHandler, &log)

	t.Run("Get empty list of published posts", func(t *testing.T) {
		cmd, err := queries.NewGetPublishedPostsQuery(&dto.GetPublicPostsQueryReq{
			PostListQueryDto: sharedDto.PostListQueryDto{
				Page:  page,
				Limit: limit,
			},
		})
		if err != nil {
			log.TestFatalf("Failed to create get published posts query: %v", err)
		}

		ctx := context.Background()
		res, err := getPostsHandler.Handle(ctx, cmd)
		if err != nil {
			log.TestFatalf("Expected to get published posts, but got error: %v", err)
		}

		if len(res.Data) != 0 {
			log.TestErrorf("Expected empty list of published posts, but got %d posts", len(res.Data))
		}

		if res.Pagination.Total != 0 {
			log.TestErrorf("Expected total to be 0, but got %d", res.Pagination.Total)
		}
	})

	t.Run("Create multiple posts and get only published ones", func(t *testing.T) {
		ctx := context.Background()
		var createdPostIds []string

		// Create 3 posts
		for i := 1; i <= 3; i++ {
			title := "Post " + string(rune('0'+i))
			mockPost, err := createPostCmd.NewCreatePostCommand(
				&createPostDto.CreatePostRequest{
					Title:    title,
					Content:  "This is test post " + string(rune('0'+i)),
					Slug:     "test-post-" + string(rune('0'+i)),
					AuthorId: "123",
				},
			)

			res, err := createPostHandler.Handle(ctx, mockPost)
			if err != nil {
				log.TestFatalf("Expected post to be created, but got error: %v", err)
			}
			createdPostIds = append(createdPostIds, res.Data.ID)
		}

		// Get published posts - should be empty because posts are not published yet
		page = 1
		limit = 10
		getCmd, err := queries.NewGetPublishedPostsQuery(&dto.GetPublicPostsQueryReq{
			PostListQueryDto: sharedDto.PostListQueryDto{
				Page:  page,
				Limit: limit,
			},
		})
		if err != nil {
			log.TestFatalf("Failed to create GetPublishedPostsQuery: %v", err)
		}
		res, err := getPostsHandler.Handle(ctx, getCmd)
		if err != nil {
			log.TestFatalf("Expected to get published posts, but got error: %v", err)
		}

		if len(res.Data) != 0 {
			log.TestErrorf("Expected 0 published posts before publishing, but got %d", len(res.Data))
		}

		// Publish first 2 posts
		for i := range 2 {
			published := true
			updateCmd, err := updatePostCmd.NewUpdatePostCommand(
				&updatePostDto.UpdatePostRequest{
					ID:        createdPostIds[i],
					Published: &published,
				},
			)
			if err != nil {
				log.TestFatalf("Failed to create update post command: %v", err)
			}
			_, err = updatePostHandler.Handle(ctx, updateCmd)
			if err != nil {
				log.TestFatalf("Expected post to be updated, but got error: %v", err)
			}
		}

		// Get published posts - should return 2 posts
		res, err = getPostsHandler.Handle(ctx, getCmd)
		if err != nil {
			log.TestFatalf("Expected to get published posts, but got error: %v", err)
		}

		if len(res.Data) != 2 {
			log.TestErrorf("Expected 2 published posts, but got %d", len(res.Data))
		}

		if res.Pagination.Total != 2 {
			log.TestErrorf("Expected total to be 2, but got %d", res.Pagination.Total)
		}

		// Verify the published posts have correct titles
		expectedTitles := map[string]bool{"Post 1": true, "Post 2": true}
		for _, post := range res.Data {
			if !expectedTitles[post.Title] {
				log.TestErrorf("Unexpected post title: %s", post.Title)
			}
		}

		// Publish the third post
		updateCmd, err := updatePostCmd.NewUpdatePostCommand(
			&updatePostDto.UpdatePostRequest{
				ID:        createdPostIds[2],
				Published: &published,
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create update post command: %v", err)
		}
		_, err = updatePostHandler.Handle(ctx, updateCmd)
		if err != nil {
			log.TestFatalf("Expected post to be updated, but got error: %v", err)
		}

		// Get published posts again - should return 3 posts
		res, err = getPostsHandler.Handle(ctx, getCmd)
		if err != nil {
			log.TestFatalf("Expected to get published posts, but got error: %v", err)
		}

		if len(res.Data) != 3 {
			log.TestErrorf("Expected 3 published posts, but got %d", len(res.Data))
		}

		if res.Pagination.Total != 3 {
			log.TestErrorf("Expected total to be 3, but got %d", res.Pagination.Total)
		}
	})

	t.Run("Get published posts with pagination", func(t *testing.T) {
		ctx := context.Background()

		// Create 5 posts and publish them
		for i := range 5 {
			title := "Pagination Post " + string(rune('0'+i))
			mockPost, err := createPostCmd.NewCreatePostCommand(
				&createPostDto.CreatePostRequest{
					Title:    title,
					Content:  "This is pagination test post " + string(rune('0'+i)),
					Slug:     "pagination-post-" + string(rune('0'+i)),
					AuthorId: "123",
				},
			)

			res, err := createPostHandler.Handle(ctx, mockPost)
			if err != nil {
				log.TestFatalf("Expected post to be created, but got error: %v", err)
			}

			published := true
			updateCmd, err := updatePostCmd.NewUpdatePostCommand(
				&updatePostDto.UpdatePostRequest{
					ID:        res.Data.ID,
					Published: &published,
				},
			)
			if err != nil {
				log.TestFatalf("Failed to create update post command: %v", err)
			}
			_, err = updatePostHandler.Handle(ctx, updateCmd)
			if err != nil {
				log.TestFatalf("Expected post to be updated, but got error: %v", err)
			}
		}

		page = 1
		limit = 2
		getCmd, err := queries.NewGetPublishedPostsQuery(&dto.GetPublicPostsQueryReq{
			PostListQueryDto: sharedDto.PostListQueryDto{
				Page:  page,
				Limit: limit,
			},
		})
		if err != nil {
			log.TestFatalf("Failed to create get published posts query: %v", err)
		}

		res, err := getPostsHandler.Handle(ctx, getCmd)
		if err != nil {
			log.TestFatalf("Expected to get published posts, but got error: %v", err)
		}

		if len(res.Data) != 2 {
			log.TestErrorf("Expected 2 posts on first page, but got %d", len(res.Data))
		}

		// 3 posts from previous test + 5 posts from this test = 8 published posts in total
		if res.Pagination.Total != 8 {
			log.TestErrorf("Expected total to be 8, but got %d", res.Pagination.Total)
		}

		if res.Pagination.Page != 1 {
			log.TestErrorf("Expected page to be 1, but got %d", res.Pagination.Page)
		}

		// Get second page
		limit = 2
		page = 2
		getCmd, err = queries.NewGetPublishedPostsQuery(&dto.GetPublicPostsQueryReq{
			PostListQueryDto: sharedDto.PostListQueryDto{
				Page:  page,
				Limit: limit,
			},
		})
		if err != nil {
			log.TestFatalf("Failed to create get published posts query: %v", err)
		}

		res, err = getPostsHandler.Handle(ctx, getCmd)
		if err != nil {
			log.TestFatalf("Expected to get published posts, but got error: %v", err)
		}

		if len(res.Data) != 2 {
			log.TestErrorf("Expected 2 posts on second page, but got %d", len(res.Data))
		}

		if res.Pagination.Page != 2 {
			log.TestErrorf("Expected page to be 2, but got %d", res.Pagination.Page)
		}
	})
}
