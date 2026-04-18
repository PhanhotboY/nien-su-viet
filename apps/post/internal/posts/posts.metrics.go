package posts

import (
	"go.opentelemetry.io/otel/metric"

	"github.com/phanhotboy/nien-su-viet/apps/post/internal/posts/infrastructure/contracts"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/settings"
)

func configPostsMetrics(
	cfg settings.Config,
	meter metric.Meter,
) (*contracts.PostsMetrics, error) {
	if meter == nil {
		return nil, nil
	}

	requestsTotalCounter, err := meter.Int64Counter(
		"requests_total",
	)
	if err != nil {
		return nil, err
	}

	requestsDurationHistogram, err := meter.Float64Histogram(
		"requests_duration",
	)
	if err != nil {
		return nil, err
	}

	return &contracts.PostsMetrics{
		RequestsTotal:    requestsTotalCounter,
		RequestsDuration: requestsDurationHistogram,
	}, nil
}
