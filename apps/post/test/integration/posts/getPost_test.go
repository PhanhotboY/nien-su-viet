//go:build integration

package v1

import (
	"context"
	"testing"

	"google.golang.org/grpc/codes"

	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	testlogger "github.com/phanhotboy/nien-su-viet/libs/pkg/logger/test"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPost/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/query/getPost/v1/queries"

	createPostCmd "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/createPost/v1/commands"
	createPostDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/createPost/v1/dto"

	updatePostCmd "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/updatePost/v1/commands"
	updatePostDto "github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/application/command/updatePost/v1/dto"

	testhelper "github.com/phanhotboy/nien-su-viet/apps/post/test/integration/shared/helper"
)

func TestGetPost(t *testing.T) {
	var (
		getPostHandler    queries.GetPostHandler
		createPostHandler createPostCmd.CreatePostHandler
		updatePostHandler updatePostCmd.UpdatePostHandler
		log               testlogger.TestLogger
	)

	testhelper.GetDIServices(t, &getPostHandler, &createPostHandler, &updatePostHandler, &log)

	t.Run("Get not found post by slug", func(t *testing.T) {
		cmd, err := queries.NewGetPostQuery(&dto.GetPostQueryReq{
			IDOrSlug: "slug",
		})
		if err != nil {
			log.TestFatalf("Failed to create get post query: %v", err)
		}
		expect := "slug"

		if cmd.IDOrSlug != expect {
			log.TestFatalf("Expected IDOrSlug to be %s, but got %s", expect, cmd.IDOrSlug)
		}

		ctx := context.Background()
		if _, err := getPostHandler.Handle(ctx, cmd); err != nil {
			if grpcStatus := grpcerrors.ParseError(err).GetStatus(); grpcStatus != codes.NotFound {
				log.TestErrorf("Expected error code to be NotFound, but got %s: %v", grpcStatus, err)
			}
		}
	})

	t.Run("Create Post", func(t *testing.T) {
		mockPost, err := createPostCmd.NewCreatePostCommand(
			&createPostDto.CreatePostRequest{
				Title:    "Mocking Post",
				Content:  "This is a test post.",
				Slug:     "mocking-post",
				AuthorId: "123",
			},
		)

		ctx := context.Background()
		res, err := createPostHandler.Handle(ctx, mockPost)
		if err != nil {
			log.TestFatalf("Expected post to be created, but got error: %v", err)
		}

		getCmd, err := queries.NewGetPostQuery(&dto.GetPostQueryReq{
			IDOrSlug: res.Data.ID,
		})
		if err != nil {
			log.TestFatalf("Failed to create get post query: %v", err)
		}

		getRes, err := getPostHandler.Handle(ctx, getCmd)
		if err != nil {
			log.TestFatalf("Expected to get post, but got error: %v", err)
		}

		if getRes.Data.Title != mockPost.Title {
			log.TestErrorf("Expected title to be %s, but got %s", mockPost.Title, getRes.Data.Title)
		}

		// the created post is not published yet, get post by slug should return not found
		getCmd.IDOrSlug = mockPost.Slug
		getRes, err = getPostHandler.Handle(ctx, getCmd)
		if err != nil {
			if grpcStatus := grpcerrors.ParseError(err).GetStatus(); grpcStatus != codes.NotFound {
				log.TestErrorf("Expected error code to be NotFound, but got %s: %v", grpcStatus, err)
			}
			return
		}

		mockPost.Published = new(bool)
		*mockPost.Published = true
		updateCmd, err := updatePostCmd.NewUpdatePostCommand(
			&updatePostDto.UpdatePostRequest{
				ID:        res.Data.ID,
				Published: mockPost.Published,
			},
		)
		if err != nil {
			log.TestFatalf("Failed to create update post command: %v", err)
		}
		_, err = updatePostHandler.Handle(ctx, updateCmd)
		if err != nil {
			log.TestFatalf("Expected post to be updated, but got error: %v", err)
		}

		// now the post is published, get post by slug should return the post
		getRes, err = getPostHandler.Handle(ctx, getCmd)
		if err != nil {
			log.TestFatalf("Expected to get post, but got error: %v", err)
		}
	})
}
