//go:build integration

package v1

import (
	"context"
	"testing"
	"time"

	sharedDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPopularPosts/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPopularPosts/v1/queries"
	testlogger "github.com/phanhotboy/nien-su-viet/libs/pkg/logger/test"

	createPostCmd "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/createPost/v1/commands"
	createPostDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/createPost/v1/dto"

	updatePostCmd "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/updatePost/v1/commands"
	updatePostDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/updatePost/v1/dto"

	incrementPostLikesCmd "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/incrementPostLikes/v1/commands"
	incrementPostLikesDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/incrementPostLikes/v1/dto"

	testhelper "github.com/phanhotboy/nien-su-viet/apps/post/test/integration/shared/helper"
)

func TestGetPopularPosts(t *testing.T) {
	var (
		getPostsHandler           queries.GetPopularPostsHandler
		createPostHandler         createPostCmd.CreatePostHandler
		updatePostHandler         updatePostCmd.UpdatePostHandler
		incrementPostLikesHandler incrementPostLikesCmd.IncrementPostLikesHandler
		log                       testlogger.TestLogger
		page                      uint32 = 1
		limit                     uint32 = 10
	)

	testhelper.GetDIServices(t, &getPostsHandler, &createPostHandler, &updatePostHandler, &incrementPostLikesHandler, &log)

	t.Run("Get empty list of popular posts", func(t *testing.T) {
		cmd, err := queries.NewGetPopularPostsQuery(&dto.GetPopularPostsQueryReq{
			PostListQueryDto: sharedDto.PostListQueryDto{
				Page:  page,
				Limit: limit,
			},
		})
		if err != nil {
			log.TestFatalf("Failed to create get popular posts query: %v", err)
		}

		ctx := context.Background()
		res, err := getPostsHandler.Handle(ctx, cmd)
		if err != nil {
			log.TestFatalf("Expected to get popular posts, but got error: %v", err)
		}

		if len(res.Data) != 0 {
			log.TestErrorf("Expected empty list of popular posts, but got %d posts", len(res.Data))
		}

		if res.Pagination.Total != 0 {
			log.TestErrorf("Expected total to be 0, but got %d", res.Pagination.Total)
		}
	})

	t.Run("Popular posts are sorted by likes in descending order", func(t *testing.T) {
		ctx := context.Background()
		var createdPostIds []string
		var createdPostTitles []string

		// Create 3 posts
		for i := 1; i <= 3; i++ {
			title := "Popular Post " + string(rune('0'+i))
			createdPostTitles = append(createdPostTitles, title)

			mockPost, err := createPostCmd.NewCreatePostCommand(
				&createPostDto.CreatePostRequest{
					Title:    title,
					Content:  "This is popular post " + string(rune('0'+i)),
					Slug:     "popular-post-" + string(rune('0'+i)),
					AuthorId: "123",
				},
			)
			if err != nil {
				log.TestFatalf("Failed to create post command: %v", err)
			}

			res, err := createPostHandler.Handle(ctx, mockPost)
			if err != nil {
				log.TestFatalf("Expected post to be created, but got error: %v", err)
			}
			createdPostIds = append(createdPostIds, res.Data.ID)

			// Publish the post
			isPublished := true
			publishedAt := time.Now()
			updateCmd, err := updatePostCmd.NewUpdatePostCommand(
				&updatePostDto.UpdatePostRequest{
					ID:          res.Data.ID,
					Published:   &isPublished,
					PublishedAt: &publishedAt,
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

		// Increment likes to create different like counts
		// Post 1: 3 likes
		// Post 2: 1 like
		// Post 3: 2 likes
		for i := 0; i < 3; i++ {
			likesCount := 3 - i // 3, 2, 1
			for j := 0; j < likesCount; j++ {
				likeCmd, err := incrementPostLikesCmd.NewIncrementPostLikesCommand(
					&incrementPostLikesDto.IncrementPostLikesRequest{
						ID: createdPostIds[i],
					},
				)
				if err != nil {
					log.TestFatalf("Failed to create increment likes command: %v", err)
				}
				_, err = incrementPostLikesHandler.Handle(ctx, likeCmd)
				if err != nil {
					log.TestFatalf("Expected to increment likes, but got error: %v", err)
				}
			}
		}

		// Get popular posts - should be sorted by likes descending
		getCmd, err := queries.NewGetPopularPostsQuery(&dto.GetPopularPostsQueryReq{
			PostListQueryDto: sharedDto.PostListQueryDto{
				Page:  1,
				Limit: 10,
			},
		})
		if err != nil {
			log.TestFatalf("Failed to create get popular posts query: %v", err)
		}

		res, err := getPostsHandler.Handle(ctx, getCmd)
		if err != nil {
			log.TestFatalf("Expected to get popular posts, but got error: %v", err)
		}

		if len(res.Data) != 3 {
			log.TestErrorf("Expected 3 popular posts, but got %d", len(res.Data))
		}

		// Verify sorting by likes (descending): 3 likes, 2 likes, 1 like
		expectedLikeOrder := []int{3, 2, 1}
		for i, post := range res.Data {
			if post.Likes != expectedLikeOrder[i] {
				log.TestErrorf("Expected post at index %d to have %d likes, but got %d", i, expectedLikeOrder[i], post.Likes)
			}
		}
	})

	t.Run("Get popular posts with pagination", func(t *testing.T) {
		ctx := context.Background()

		// Create 5 posts and publish them
		for i := 1; i <= 5; i++ {
			title := "Pagination Popular Post " + string(rune('0'+i))
			mockPost, err := createPostCmd.NewCreatePostCommand(
				&createPostDto.CreatePostRequest{
					Title:    title,
					Content:  "This is pagination popular post " + string(rune('0'+i)),
					Slug:     "pagination-popular-post-" + string(rune('0'+i)),
					AuthorId: "123",
				},
			)
			if err != nil {
				log.TestFatalf("Failed to create post command: %v", err)
			}

			res, err := createPostHandler.Handle(ctx, mockPost)
			if err != nil {
				log.TestFatalf("Expected post to be created, but got error: %v", err)
			}

			// Publish the post
			isPublished := true
			publishedAt := time.Now()
			updateCmd, err := updatePostCmd.NewUpdatePostCommand(
				&updatePostDto.UpdatePostRequest{
					ID:          res.Data.ID,
					Published:   &isPublished,
					PublishedAt: &publishedAt,
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

		// Get first page (limit 2)
		page = 1
		limit = 2
		getCmd, err := queries.NewGetPopularPostsQuery(&dto.GetPopularPostsQueryReq{
			PostListQueryDto: sharedDto.PostListQueryDto{
				Page:  page,
				Limit: limit,
			},
		})
		if err != nil {
			log.TestFatalf("Failed to create get popular posts query: %v", err)
		}

		res, err := getPostsHandler.Handle(ctx, getCmd)
		if err != nil {
			log.TestFatalf("Expected to get popular posts, but got error: %v", err)
		}

		if len(res.Data) != 2 {
			log.TestErrorf("Expected 2 posts on first page, but got %d", len(res.Data))
		}

		if res.Pagination.Page != 1 {
			log.TestErrorf("Expected page to be 1, but got %d", res.Pagination.Page)
		}

		// Get second page (limit 2)
		page = 2
		getCmd, err = queries.NewGetPopularPostsQuery(&dto.GetPopularPostsQueryReq{
			PostListQueryDto: sharedDto.PostListQueryDto{
				Page:  page,
				Limit: limit,
			},
		})
		if err != nil {
			log.TestFatalf("Failed to create get popular posts query: %v", err)
		}

		res, err = getPostsHandler.Handle(ctx, getCmd)
		if err != nil {
			log.TestFatalf("Expected to get popular posts, but got error: %v", err)
		}

		if len(res.Data) != 2 {
			log.TestErrorf("Expected 2 posts on second page, but got %d", len(res.Data))
		}

		if res.Pagination.Page != 2 {
			log.TestErrorf("Expected page to be 2, but got %d", res.Pagination.Page)
		}
	})

	t.Run("Only published posts appear in popular posts", func(t *testing.T) {
		ctx := context.Background()
		var publishedPostId string

		// Create 2 posts - one published, one unpublished
		for i := 1; i <= 2; i++ {
			title := "Mixed Post " + string(rune('0'+i))
			mockPost, err := createPostCmd.NewCreatePostCommand(
				&createPostDto.CreatePostRequest{
					Title:    title,
					Content:  "This is mixed post " + string(rune('0'+i)),
					Slug:     "mixed-post-" + string(rune('0'+i)),
					AuthorId: "123",
				},
			)
			if err != nil {
				log.TestFatalf("Failed to create post command: %v", err)
			}

			res, err := createPostHandler.Handle(ctx, mockPost)
			if err != nil {
				log.TestFatalf("Expected post to be created, but got error: %v", err)
			}

			// Only publish the first post
			if i == 1 {
				isPublished := true
				publishedAt := time.Now()
				updateCmd, err := updatePostCmd.NewUpdatePostCommand(
					&updatePostDto.UpdatePostRequest{
						ID:          res.Data.ID,
						Published:   &isPublished,
						PublishedAt: &publishedAt,
					},
				)
				if err != nil {
					log.TestFatalf("Failed to create update post command: %v", err)
				}
				_, err = updatePostHandler.Handle(ctx, updateCmd)
				if err != nil {
					log.TestFatalf("Expected post to be updated, but got error: %v", err)
				}
				publishedPostId = res.Data.ID

				// Add likes to the published post
				for j := 0; j < 5; j++ {
					likeCmd, err := incrementPostLikesCmd.NewIncrementPostLikesCommand(
						&incrementPostLikesDto.IncrementPostLikesRequest{
							ID: publishedPostId,
						},
					)
					if err != nil {
						log.TestFatalf("Failed to create increment likes command: %v", err)
					}
					_, err = incrementPostLikesHandler.Handle(ctx, likeCmd)
					if err != nil {
						log.TestFatalf("Expected to increment likes, but got error: %v", err)
					}
				}
			}
		}

		// Get popular posts
		getCmd, err := queries.NewGetPopularPostsQuery(&dto.GetPopularPostsQueryReq{
			PostListQueryDto: sharedDto.PostListQueryDto{
				Page:  1,
				Limit: 10,
			},
		})
		if err != nil {
			log.TestFatalf("Failed to create get popular posts query: %v", err)
		}

		res, err := getPostsHandler.Handle(ctx, getCmd)
		if err != nil {
			log.TestFatalf("Expected to get popular posts, but got error: %v", err)
		}

		// Verify only the published post is in the results
		found := false
		for _, post := range res.Data {
			if post.Id == publishedPostId {
				found = true
				if post.Likes != 5 {
					log.TestErrorf("Expected published post to have 5 likes, but got %d", post.Likes)
				}
			}
		}

		if !found {
			log.TestErrorf("Expected to find published post in popular posts, but it was not found")
		}
	})
}
