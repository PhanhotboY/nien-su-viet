package elasticsearch

import (
	"go.uber.org/fx"
)

var Module = fx.Module("elasticModule",
	fx.Provide(NewElasticClient),
)
