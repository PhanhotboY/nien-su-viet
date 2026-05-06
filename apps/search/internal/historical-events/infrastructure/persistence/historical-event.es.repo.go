package persistence

import (
	"context"
	"math"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/esdsl"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/domain/entity"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/domain/repository"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	dtoUtil "github.com/phanhotboy/nien-su-viet/libs/pkg/utils/dto"
)

type historicalEventEsRepository struct {
	log      logger.Logger
	esClient *elasticsearch.TypedClient
}

func NewHistoricalEventEsRepository(log logger.Logger, esClient *elasticsearch.TypedClient) repository.HistoricalEventEsRepository {
	eventEntity := entity.HistoricalEvent{}
	if exists, err := esClient.Indices.Exists(eventEntity.IndexName()).Do(context.Background()); !exists || err != nil {
		log.Infof("Index '%s' is not existed. Create new index.", eventEntity.IndexName())
		esClient.Indices.Create(eventEntity.IndexName()).Mappings(&types.TypeMapping{
			Properties: eventEntity.ToTypeMapping(),
		}).Do(context.Background())
	}

	return &historicalEventEsRepository{
		log:      log,
		esClient: esClient,
	}
}

func (r *historicalEventEsRepository) IndexHistoricalEvent(ctx context.Context, event entity.HistoricalEvent) error {
	_, err := r.esClient.Index(event.IndexName()).
		Id(event.Id).
		Document(event).
		Do(ctx)
	return err
}

func (r *historicalEventEsRepository) GetHistoricalEventById(ctx context.Context, id string) (*entity.HistoricalEvent, error) {
	res, err := r.esClient.Get(entity.HistoricalEvent{}.IndexName(), id).Do(ctx)
	if err != nil {
		return nil, err
	}
	if !res.Found {
		return nil, nil
	}

	event := new(entity.HistoricalEvent)
	if err = dtoUtil.ValidateStruct(res.Source_, event); err != nil {
		return nil, err
	}
	return event, nil
}

func (r *historicalEventEsRepository) SearchHistoricalEvents(ctx context.Context, query repository.HistoricalEventsQuery) (*repository.HistoricalEventsSearchResponse, error) {
	sort := esdsl.NewSortOptions()
	if query.SortBy != "" {
		sortOrder := sortorder.Desc
		if query.SortOrder == "asc" {
			sortOrder = sortorder.Asc
		}
		sort = sort.AddSortOption(query.SortBy, esdsl.NewFieldSort(sortOrder))
	}

	res, err := r.esClient.Search().
		Index(entity.HistoricalEvent{}.IndexName()).
		Query(query.QueryVariant).
		From(int(math.Max(float64(query.Page-1), 0)*float64(query.Limit))). // Assure that from is not negative
		Size(int(query.Limit)).
		Sort(sort, esdsl.NewScoreSort().Order(sortorder.Desc).SortOptionsCaster()).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	historicalEvents := make([]entity.HistoricalEventBrief, 0, len(res.Hits.Hits))

	for _, hit := range res.Hits.Hits {
		item := new(entity.HistoricalEventBrief)
		if err = dtoUtil.ValidateStruct(hit.Source_, item); err != nil {
			r.log.Errorf("failed to unmarshal historical event from search result: %v", err)
			continue
		}
		historicalEvents = append(historicalEvents, *item)
	}

	total := int64(0)
	if res.Hits.Total != nil {
		total = res.Hits.Total.Value
	}

	return &repository.HistoricalEventsSearchResponse{
		HistoricalEvents: historicalEvents,
		TotalCount:       total,
	}, nil
}

func (r *historicalEventEsRepository) UpdateHistoricalEvent(ctx context.Context, id entity.HistoricalEventId, event *entity.HistoricalEvent) error {
	_, err := r.esClient.Update(event.IndexName(), id).
		Doc(event).
		Do(ctx)
	return err
}

func (r *historicalEventEsRepository) DeleteHistoricalEvent(ctx context.Context, id entity.HistoricalEventId) error {
	_, err := r.esClient.Delete(entity.HistoricalEvent{}.IndexName(), id).Do(ctx)
	return err
}
