package queries

import (
	"context"

	eventDto "github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/application/dto"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/application/query/getAllEvents/v1/dto"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/domain/entity"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/domain/repository"
	"github.com/phanhotboy/nien-su-viet/apps/search/internal/shared/helper"
	grpcerrors "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/grpcErrors"
	grpcTypes "github.com/phanhotboy/nien-su-viet/libs/pkg/grpc/types"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/logger"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/utils"
	"go.opentelemetry.io/otel/trace"
)

type getAllEventsHandler struct {
	log          logger.Logger
	tracer       trace.Tracer
	heventEsRepo repository.HistoricalEventEsRepository
	cacheRepo    repository.HistoricalEventCacheRepository
}

type IGetAllEventsHandler interface {
	grpcTypes.GrpcHandler[*GetAllEventsQuery, *dto.GetAllEventsRes]
}

func NewGetAllEventsHandler(
	log logger.Logger,
	heventEsRepo repository.HistoricalEventEsRepository,
	cacheRepo repository.HistoricalEventCacheRepository,
	tracer trace.Tracer,
) IGetAllEventsHandler {
	return getAllEventsHandler{
		log:          log,
		tracer:       tracer,
		heventEsRepo: heventEsRepo,
		cacheRepo:    cacheRepo,
	}
}

func (c getAllEventsHandler) Handle(
	ctx context.Context,
	query *GetAllEventsQuery,
) (*dto.GetAllEventsRes, error) {
	var cached *utils.PaginatedResponse[entity.HistoricalEventBrief]

	queryEntity := query.MapToQuery()
	queryKey := helper.GenerateCacheKey(queryEntity)
	cached, err := c.cacheRepo.GetHistoricalEvents(ctx, string(queryKey))
	if err != nil {
		c.log.Warnf("failed to get all events from cache: %v, fallback to db", err)
	}

	if cached == nil {
		res, err := c.heventEsRepo.SearchHistoricalEvents(ctx, queryEntity)
		if err != nil {
			c.log.Errorf("failed to get all events: %v", err)
			return nil, grpcerrors.ParseError(err)
		}

		cached = &utils.PaginatedResponse[entity.HistoricalEventBrief]{
			Data:       res.HistoricalEvents,
			Pagination: utils.NewPagination(queryEntity.Limit, queryEntity.Page, uint32(res.TotalCount)),
		}

		if err := c.cacheRepo.PutHistoricalEvents(ctx, string(queryKey), cached); err != nil {
			c.log.Warnf("failed to cache all events: %v", err)
		}
	}

	eventBriefs := make([]eventDto.HistoricalEventBriefDto, 0)
	for _, event := range cached.Data {
		eventBrief := eventDto.HistoricalEventBriefDto{}
		eventBrief.FromEntity(&event)
		eventBriefs = append(eventBriefs, eventBrief)
	}

	return dto.NewGetAllEventsRes(
		eventBriefs,
		*cached.Pagination,
	), nil
}
