package config

import (
	"fmt"
	"reflect"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/config"
	sharedCfg "github.com/phanhotboy/nien-su-viet/libs/pkg/config"
)

type BillingConfig interface {
	sharedCfg.Config
	GetZaloPayOptions() ZaloPayOptions
}

type ZaloPayOptions struct {
	AppID       int    `mapstructure:"app_id"`       // BILLING_ZALOPAY_APP_ID
	Key1        string `mapstructure:"key1"`         // BILLING_ZALOPAY_KEY1
	Key2        string `mapstructure:"key2"`         // BILLING_ZALOPAY_KEY2
	CreateURL   string `mapstructure:"create_url"`   // BILLING_ZALOPAY_CREATE_URL
	QueryURL    string `mapstructure:"query_url"`    // BILLING_ZALOPAY_QUERY_URL
	CallbackURL string `mapstructure:"callback_url"` // BILLING_ZALOPAY_CALLBACK_URL
	RedirectURL string `mapstructure:"redirect_url"` // BILLING_ZALOPAY_REDIRECT_URL
}

type Config struct {
	config.DefaultConfig `mapstructure:",squash"`
	ZaloPay              ZaloPayOptions `mapstructure:"zalopay"`
}

func ConfigType() reflect.Type {
	return reflect.TypeFor[Config]()
}

func (c ZaloPayOptions) Validate() error {
	if c.AppID == 0 {
		return fmt.Errorf("zalopay: missing AppID")
	}
	if c.Key1 == "" {
		return fmt.Errorf("zalopay: missing Key1")
	}
	if c.Key2 == "" {
		return fmt.Errorf("zalopay: missing Key2")
	}
	if c.CreateURL == "" {
		return fmt.Errorf("zalopay: missing CreateURL")
	}
	if c.QueryURL == "" {
		return fmt.Errorf("zalopay: missing QueryURL")
	}
	return nil
}

func (c Config) GetZaloPayOptions() ZaloPayOptions {
	return c.ZaloPay
}
