package config

import (
	"reflect"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/config"
	sharedCfg "github.com/phanhotboy/nien-su-viet/libs/pkg/config"
)

type BillingConfig interface {
	sharedCfg.Config
	GetZaloPayOptions() ZaloPayOptions
}

type ZaloPayOptions struct {
	AppID     string `mapstructure:"app_id"`     // BILLING_ZALOPAY_APP_ID
	Key1      string `mapstructure:"key1"`       // BILLING_ZALOPAY_KEY1
	Key2      string `mapstructure:"key2"`       // BILLING_ZALOPAY_KEY2
	CreateURL string `mapstructure:"create_url"` // BILLING_ZALOPAY_CREATE_URL
	QueryURL  string `mapstructure:"query_url"`  // BILLING_ZALOPAY_QUERY_URL
}

type Config struct {
	config.DefaultConfig `mapstructure:",squash"`
	ZaloPay              ZaloPayOptions `mapstructure:"zalopay"`
}

func ConfigType() reflect.Type {
	return reflect.TypeOf(Config{})
}
