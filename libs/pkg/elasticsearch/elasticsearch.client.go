package elasticsearch

import (
	"emperror.dev/errors"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/settings"
)

func NewElasticClient(settings settings.Config) (*elasticsearch.TypedClient, error) {
	cfg := settings.Elasticsearch
	if cfg.URL == "" || cfg.Username == "" || cfg.Password == "" {
		return nil, errors.New("missing elasticsearch configuration")
	}
	es, err := elasticsearch.NewTyped(
		elasticsearch.WithAddresses(cfg.URL),
		elasticsearch.WithBasicAuth(cfg.Username, cfg.Password),
		elasticsearch.WithInstrumentation(elasticsearch.NewOpenTelemetryInstrumentation(nil, false)),
	)
	if err != nil {
		return nil, errors.WrapIf(err, "v9.elasticsearch")
	}

	return es, nil
}
