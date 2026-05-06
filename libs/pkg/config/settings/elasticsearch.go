package settings

type ElasticsearchConfig struct {
	URL      string `mapstructure:"node_url"` // SEARCH_ELASTICSEARCH_NODE_URL
	Username string `mapstructure:"username"` // SEARCH_ELASTICSEARCH_USERNAME
	Password string `mapstructure:"password"` // SEARCH_ELASTICSEARCH_PASSWORD
}
