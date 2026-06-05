package zalopay

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	SandboxCreateOrderURL = "https://sb-openapi.zalopay.vn/v2/create"
	SandboxQueryOrderURL  = "https://sb-openapi.zalopay.vn/v2/query"
	ProdCreateOrderURL    = "https://openapi.zalopay.vn/v2/create"
	ProdQueryOrderURL     = "https://openapi.zalopay.vn/v2/query"
)

type Config struct {
	AppID      string
	Key1       string
	Key2       string
	CreateURL  string
	QueryURL   string
	HTTPClient *http.Client
}

func (c Config) validate() error {
	if c.AppID == "" {
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

type Client struct {
	appID      string
	key1       string
	key2       string
	createURL  string
	queryURL   string
	httpClient *http.Client
}

func New(cfg Config) (*Client, error) {
	if err := cfg.validate(); err != nil {
		return nil, err
	}
	hc := cfg.HTTPClient
	if hc == nil {
		hc = &http.Client{Timeout: 15 * time.Second}
	}
	return &Client{
		appID:      cfg.AppID,
		key1:       cfg.Key1,
		key2:       cfg.Key2,
		createURL:  cfg.CreateURL,
		queryURL:   cfg.QueryURL,
		httpClient: hc,
	}, nil
}

func HMACSHA256Hex(data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	_, _ = mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}

func doJSON[T any](ctx context.Context, hc *http.Client, req *http.Request) (*T, error) {
	resp, err := hc.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("zalopay: http %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	var out T
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, fmt.Errorf("zalopay: decode response: %w", err)
	}
	return &out, nil
}

func (c *Client) CreateOrder(ctx context.Context, req CreateOrderRequest) (*CreateOrderResponse, error) {
	req.AppID = c.appID
	req.Mac = req.computeMac(c.key1)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("zalopay: marshal create order: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.createURL, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("zalopay: build create order request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	return doJSON[CreateOrderResponse](ctx, c.httpClient, httpReq)
}

func (c *Client) QueryOrder(ctx context.Context, req QueryOrderRequest) (*QueryOrderResponse, error) {
	req.AppID = c.appID
	req.Mac = req.computeMac(c.key1)

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("zalopay: marshal query order: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.queryURL, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("zalopay: build query order request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	return doJSON[QueryOrderResponse](ctx, c.httpClient, httpReq)
}

func (c *Client) VerifyCallback(body []byte, receivedMac string) error {
	expected := HMACSHA256Hex(string(body), c.key2)
	if !hmac.Equal([]byte(expected), []byte(receivedMac)) {
		return fmt.Errorf("zalopay: invalid callback signature")
	}
	return nil
}
