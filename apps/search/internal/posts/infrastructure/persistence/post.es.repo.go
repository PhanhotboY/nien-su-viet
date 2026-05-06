package persistence

import (
	"context"
	"math"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/esdsl"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/domain/entity"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/posts/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type postEsRepository struct {
	log      logger.Logger
	esClient *elasticsearch.TypedClient
}

func NewPostEsRepository(
	log logger.Logger,
	esClient *elasticsearch.TypedClient,
) repository.PostEsRepository {
	postEntity := entity.Post{}
	if exists, err := esClient.Indices.Exists(postEntity.IndexName()).Do(context.Background()); !exists || err != nil {
		log.Infof("Index '%s' is not existed. Create new index.", postEntity.IndexName())
		esClient.Indices.Create(postEntity.IndexName()).Mappings(&types.TypeMapping{
			Properties: postEntity.ToTypeMapping(),
		}).Do(context.Background())
	}

	return &postEsRepository{
		log:      log,
		esClient: esClient,
	}
}

func (r *postEsRepository) IndexPost(ctx context.Context, post entity.Post) error {
	_, err := r.esClient.Index(post.IndexName()).
		Id(post.Id).
		Document(post).
		Do(ctx)
	return err
}

func (r *postEsRepository) SearchPosts(ctx context.Context, query repository.PostQuery) (*repository.PostSearchResponse, error) {
	sort := esdsl.NewSortOptions()
	if query.SortBy != "" {
		sortOrder := sortorder.Desc
		if query.SortOrder == "asc" {
			sortOrder = sortorder.Asc
		}
		sort = sort.AddSortOption(query.SortBy, esdsl.NewFieldSort(sortOrder))
	}

	res, err := r.esClient.Search().
		Index(entity.Post{}.IndexName()).
		Query(query).
		From(int(math.Max(float64(query.Page-1), 0)*float64(query.Limit))).
		Size(int(query.Limit)).
		Sort(sort, esdsl.NewScoreSort().Order(sortorder.Desc).SortOptionsCaster()).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	posts := make([]entity.PostBrief, 0, len(res.Hits.Hits))
	for _, hit := range res.Hits.Hits {
		post := new(entity.PostBrief)
		if err = dtoUtil.ValidateStruct(hit.Source_, post); err != nil {
			r.log.Errorf("failed to unmarshal post from search result: %v", err)
			continue
		}
		posts = append(posts, *post)
	}

	total := int64(0)
	if res.Hits.Total != nil {
		total = res.Hits.Total.Value
	}

	return &repository.PostSearchResponse{
		Posts:      posts,
		TotalCount: total,
	}, nil
}

func (r *postEsRepository) GetPostByID(ctx context.Context, postId entity.PostId) (*entity.Post, error) {
	res, err := r.esClient.Get(entity.Post{}.IndexName(), postId).Do(ctx)
	if err != nil {
		return nil, err
	}

	if !res.Found {
		return nil, nil
	}

	post := new(entity.Post)
	if err = dtoUtil.ValidateStruct(res.Source_, post); err != nil {
		return nil, err
	}
	return post, nil
}

func (r *postEsRepository) GetPostBySlug(ctx context.Context, slug string) (*entity.Post, error) {
	res, err := r.esClient.Search().Index(entity.Post{}.IndexName()).Query(esdsl.NewTermQuery("slug", esdsl.NewFieldValue().String(slug))).Do(ctx)
	if err != nil {
		return nil, err
	}

	if len(res.Hits.Hits) == 0 {
		return nil, nil
	}

	post := new(entity.Post)
	if err = dtoUtil.ValidateStruct(res.Hits.Hits[0].Source_, post); err != nil {
		return nil, err
	}
	return post, nil
}

func (r *postEsRepository) UpdatePost(ctx context.Context, postId entity.PostId, post *entity.Post) error {
	r.log.Debugf("Updating post with ID %s: %+v", postId, post)
	_, err := r.esClient.Update(post.IndexName(), postId).Doc(post).Do(ctx)
	return err
}

func (r *postEsRepository) DeletePost(ctx context.Context, postId entity.PostId) error {
	_, err := r.esClient.Delete(entity.Post{}.IndexName(), postId).Do(ctx)
	return err
}

func (r *postEsRepository) DeleteAllPosts(ctx context.Context) error {
	_, err := r.esClient.DeleteByQuery(entity.Post{}.IndexName()).
		Query(esdsl.NewMatchAllQuery()).
		Do(ctx)
	return err
}
