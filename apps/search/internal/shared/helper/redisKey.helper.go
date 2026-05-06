package helper

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/phanhotboy/nien-su-viet/apps/search/internal/historical-events/domain/repository"
)

func GenerateCacheKey(query repository.HistoricalEventsQuery) string {
	queryByte, err := query.QueryVariant.QueryCaster().MarshalJSON()
	if err != nil {
		return ""
	}
	pureJson, err := json.Marshal(query)
	if err != nil {
		return ""
	}

	// replace the empty QueryVariant with the actual query variant
	// because the json.Marshal can't marshal QueryVariant correctly
	return strings.ReplaceAll(string(pureJson), "\"QueryVariant\":{}", fmt.Sprintf("\"QueryVariant\":%s", string(queryByte)))
}
