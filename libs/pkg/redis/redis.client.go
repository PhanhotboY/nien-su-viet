package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/phanhotboy/nien-su-viet/libs/pkg/config/settings"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

const (
	maxRetries      = 5
	minRetryBackoff = 300 * time.Millisecond
	maxRetryBackoff = 500 * time.Millisecond
	dialTimeout     = 5 * time.Second
	readTimeout     = 5 * time.Second
	writeTimeout    = 3 * time.Second
	minIdleConns    = 20
	poolTimeout     = 6 * time.Second
	defaultTTL      = 15 * time.Minute
)

type RedisClientWithExpire struct {
	*redis.Client
}

func NewRedisClient(s settings.Config) RedisClientWithExpire {
	cfg := s.Redis
	universalClient := redis.NewClient(&redis.Options{
		Addr:            fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:        cfg.Password,
		DB:              cfg.Database,
		MaxRetries:      maxRetries,
		MinRetryBackoff: minRetryBackoff,
		MaxRetryBackoff: maxRetryBackoff,
		DialTimeout:     dialTimeout,
		ReadTimeout:     readTimeout,
		WriteTimeout:    writeTimeout,
		PoolSize:        cfg.PoolSize,
		MinIdleConns:    minIdleConns,
		PoolTimeout:     poolTimeout,
	})

	if cfg.EnableTracing {
		_ = redisotel.InstrumentTracing(universalClient)
	}

	return RedisClientWithExpire{
		Client: universalClient,
	}
}

// ============================================================================
// STRING OPERATIONS WITH EXPIRATION
// ============================================================================

// Set stores a value with expiration (default 5 minutes)
func (c RedisClientWithExpire) Set(ctx context.Context, key string, value interface{}, ttl ...time.Duration) error {
	duration := defaultTTL
	if len(ttl) > 0 && ttl[0] > 0 {
		duration = ttl[0]
	}

	serialized, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %w", err)
	}

	return c.Client.Set(ctx, key, serialized, duration).Err()
}

// Get retrieves and deserializes a value
func (c RedisClientWithExpire) Get(ctx context.Context, key string, dest interface{}) error {
	val, err := c.Client.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil
		}
		return fmt.Errorf("failed to get key: %w", err)
	}

	if len(val) == 0 {
		return nil
	}

	if err := json.Unmarshal(val, dest); err != nil {
		return fmt.Errorf("failed to unmarshal value: %w", err)
	}

	return nil
}

// GetString retrieves a string value without deserialization
func (c RedisClientWithExpire) GetString(ctx context.Context, key string) (string, error) {
	return c.Client.Get(ctx, key).Result()
}

// ============================================================================
// HASH OPERATIONS WITH EXPIRATION
// ============================================================================

// HSet stores a hash field with value and expiration (default 5 minutes)
func (c RedisClientWithExpire) HSet(ctx context.Context, key string, field string, value interface{}, ttl ...time.Duration) error {
	duration := defaultTTL
	if len(ttl) > 0 && ttl[0] > 0 {
		duration = ttl[0]
	}

	serialized, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %w", err)
	}

	err = c.Client.HSet(ctx, key, field, serialized).Err()
	if err != nil {
		return fmt.Errorf("failed to set hash field: %w", err)
	}

	// Set expiration
	return c.HExpire(ctx, key, duration, field).Err()
}

// HGet retrieves and deserializes a hash field
func (c RedisClientWithExpire) HGet(ctx context.Context, key string, field string, dest interface{}) error {
	val, err := c.Client.HGet(ctx, key, field).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil
		}
		return fmt.Errorf("failed to get hash field: %w", err)
	}

	if len(val) == 0 {
		return nil
	}

	if err := json.Unmarshal(val, dest); err != nil {
		return fmt.Errorf("failed to unmarshal hash field: %w", err)
	}

	return nil
}

// HGetAll retrieves all hash fields with deserialization
func (c RedisClientWithExpire) HGetAll(ctx context.Context, key string) (map[string]interface{}, error) {
	data, err := c.Client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get all hash fields: %w", err)
	}

	result := make(map[string]interface{})
	for field, value := range data {
		var parsed interface{}
		err := json.Unmarshal([]byte(value), &parsed)
		if err != nil {
			// If unmarshaling fails, store the raw value
			result[field] = value
			continue
		}
		result[field] = parsed
	}

	return result, nil
}

// HMSet stores multiple hash fields with expiration (default 5 minutes)
func (c RedisClientWithExpire) HMSet(ctx context.Context, key string, data map[string]interface{}, ttl ...time.Duration) error {
	duration := defaultTTL
	if len(ttl) > 0 && ttl[0] > 0 {
		duration = ttl[0]
	}

	for field, value := range data {
		marshaled, err := json.Marshal(value)
		if err != nil {
			return fmt.Errorf("failed to marshal field %s: %w", field, err)
		}

		err = c.HSet(ctx, key, field, string(marshaled))
		if err != nil {
			return fmt.Errorf("failed to set hash fields: %w", err)
		}
	}

	// Set expiration
	return c.Expire(ctx, key, duration)
}

// HDel deletes one or more hash fields
func (c RedisClientWithExpire) HDel(ctx context.Context, key string, fields ...string) (int64, error) {
	count, err := c.Client.HDel(ctx, key, fields...).Result()
	if err != nil {
		return 0, fmt.Errorf("failed to delete hash fields: %w", err)
	}
	return count, nil
}

// HVals retrieves all values from a hash with deserialization
func (c RedisClientWithExpire) HVals(ctx context.Context, key string) ([]interface{}, error) {
	vals, err := c.Client.HVals(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get hash values: %w", err)
	}

	result := make([]interface{}, len(vals))
	for i, val := range vals {
		if val == "null" || val == "" {
			result[i] = nil
			continue
		}

		var parsed interface{}
		err := json.Unmarshal([]byte(val), &parsed)
		if err != nil {
			// If unmarshaling fails, store the raw value
			result[i] = val
			continue
		}
		result[i] = parsed
	}

	return result, nil
}

// HExists checks if a hash field exists
func (c RedisClientWithExpire) HExists(ctx context.Context, key string, field string) (bool, error) {
	exists, err := c.Client.HExists(ctx, key, field).Result()
	if err != nil {
		return false, fmt.Errorf("failed to check hash field existence: %w", err)
	}
	return exists, nil
}

// ============================================================================
// KEY EXPIRATION OPERATIONS
// ============================================================================

// Expire sets the expiration time for a key
func (c RedisClientWithExpire) Expire(ctx context.Context, key string, ttl time.Duration) error {
	if ttl <= 0 {
		return nil
	}

	return c.Client.Expire(ctx, key, ttl).Err()
}

// ExpireWithMode sets expiration with mode (NX, XX, GT, LT)
// NX: Only set expiry when the key has no expiry
// XX: Only set expiry when the key has an existing expiry
// GT: Only set expiry when the new expiry is greater than current one
// LT: Only set expiry when the new expiry is less than current one
func (c RedisClientWithExpire) ExpireWithMode(ctx context.Context, key string, ttl time.Duration, mode string) (bool, error) {
	if ttl <= 0 {
		return false, nil
	}

	result, err := c.Client.Expire(ctx, key, ttl).Result()
	if err != nil {
		return false, fmt.Errorf("failed to set expiration: %w", err)
	}

	return result, nil
}

// TTL gets the remaining time to live of a key
func (c RedisClientWithExpire) TTL(ctx context.Context, key string) (time.Duration, error) {
	ttl, err := c.Client.TTL(ctx, key).Result()
	if err != nil {
		return 0, fmt.Errorf("failed to get TTL: %w", err)
	}
	return ttl, nil
}

// PTTL gets the remaining time to live of a key in milliseconds
func (c RedisClientWithExpire) PTTL(ctx context.Context, key string) (time.Duration, error) {
	pttl, err := c.Client.PTTL(ctx, key).Result()
	if err != nil {
		return 0, fmt.Errorf("failed to get PTTL: %w", err)
	}
	return pttl, nil
}

// Persist removes the expiration from a key
func (c RedisClientWithExpire) Persist(ctx context.Context, key string) (bool, error) {
	result, err := c.Client.Persist(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("failed to persist key: %w", err)
	}
	return result, nil
}

// ============================================================================
// MULTI-KEY OPERATIONS
// ============================================================================

// MDel deletes multiple keys matching a pattern
func (c RedisClientWithExpire) MDel(ctx context.Context, keyPattern string) (int64, error) {
	keys, err := c.Client.Keys(ctx, keyPattern).Result()
	if err != nil {
		return 0, fmt.Errorf("failed to find keys matching pattern: %w", err)
	}

	if len(keys) == 0 {
		return 0, nil
	}

	return c.Client.Del(ctx, keys...).Result()
}

// Del deletes one or more keys
func (c RedisClientWithExpire) Del(ctx context.Context, keys ...string) (int64, error) {
	if len(keys) == 0 {
		return 0, nil
	}

	return c.Client.Del(ctx, keys...).Result()
}

// Keys finds all keys matching a pattern
func (c RedisClientWithExpire) Keys(ctx context.Context, pattern string) ([]string, error) {
	return c.Client.Keys(ctx, pattern).Result()
}

// Exists checks if one or more keys exist
func (c RedisClientWithExpire) Exists(ctx context.Context, keys ...string) (int64, error) {
	if len(keys) == 0 {
		return 0, nil
	}

	return c.Client.Exists(ctx, keys...).Result()
}

// ============================================================================
// UTILITY OPERATIONS
// ============================================================================

// Ping checks the Redis connection
func (c RedisClientWithExpire) Ping(ctx context.Context) (string, error) {
	return c.Client.Ping(ctx).Result()
}

// FlushDB flushes all keys from the current database
func (c RedisClientWithExpire) FlushDB(ctx context.Context) error {
	return c.Client.FlushDB(ctx).Err()
}

// FlushAll flushes all keys from all databases
func (c RedisClientWithExpire) FlushAll(ctx context.Context) error {
	return c.Client.FlushAll(ctx).Err()
}

// DBSize returns the number of keys in the current database
func (c RedisClientWithExpire) DBSize(ctx context.Context) (int64, error) {
	return c.Client.DBSize(ctx).Result()
}

// ============================================================================
// ADVANCED OPERATIONS
// ============================================================================

// SetWithNX sets a key only if it does not exist with expiration
func (c RedisClientWithExpire) SetWithNX(ctx context.Context, key string, value interface{}, ttl ...time.Duration) (bool, error) {
	duration := defaultTTL
	if len(ttl) > 0 && ttl[0] > 0 {
		duration = ttl[0]
	}

	serialized, err := json.Marshal(value)
	if err != nil {
		return false, fmt.Errorf("failed to marshal value: %w", err)
	}

	result, err := c.Client.SetNX(ctx, key, serialized, duration).Result()
	if err != nil {
		return false, fmt.Errorf("failed to set key with NX: %w", err)
	}

	return result, nil
}

// SetWithXX sets a key only if it already exists with expiration
func (c RedisClientWithExpire) SetWithXX(ctx context.Context, key string, value interface{}, ttl ...time.Duration) (bool, error) {
	duration := defaultTTL
	if len(ttl) > 0 && ttl[0] > 0 {
		duration = ttl[0]
	}

	serialized, err := json.Marshal(value)
	if err != nil {
		return false, fmt.Errorf("failed to marshal value: %w", err)
	}

	result, err := c.Client.SetXX(ctx, key, serialized, duration).Result()
	if err != nil {
		return false, fmt.Errorf("failed to set key with XX: %w", err)
	}

	return result, nil
}

// Incr increments the integer value of a key
func (c RedisClientWithExpire) Incr(ctx context.Context, key string) (int64, error) {
	return c.Client.Incr(ctx, key).Result()
}

// Decr decrements the integer value of a key
func (c RedisClientWithExpire) Decr(ctx context.Context, key string) (int64, error) {
	return c.Client.Decr(ctx, key).Result()
}

// IncrBy increments the integer value of a key by increment
func (c RedisClientWithExpire) IncrBy(ctx context.Context, key string, increment int64) (int64, error) {
	return c.Client.IncrBy(ctx, key, increment).Result()
}

// DecrBy decrements the integer value of a key by decrement
func (c RedisClientWithExpire) DecrBy(ctx context.Context, key string, decrement int64) (int64, error) {
	return c.Client.DecrBy(ctx, key, decrement).Result()
}

// Append appends a value to a key
func (c RedisClientWithExpire) Append(ctx context.Context, key string, value string) (int64, error) {
	return c.Client.Append(ctx, key, value).Result()
}

// ============================================================================
// LIST OPERATIONS
// ============================================================================

// LPush inserts one or more values at the head of a list with expiration
func (c RedisClientWithExpire) LPush(ctx context.Context, key string, values []interface{}, ttl ...time.Duration) error {
	duration := defaultTTL
	if len(ttl) > 0 && ttl[0] > 0 {
		duration = ttl[0]
	}

	// Serialize values
	serialized := make([]interface{}, len(values))
	for i, v := range values {
		data, err := json.Marshal(v)
		if err != nil {
			return fmt.Errorf("failed to marshal value: %w", err)
		}
		serialized[i] = data
	}

	err := c.Client.LPush(ctx, key, serialized...).Err()
	if err != nil {
		return fmt.Errorf("failed to push to list: %w", err)
	}

	return c.Expire(ctx, key, duration)
}

// RPush inserts one or more values at the tail of a list with expiration
func (c RedisClientWithExpire) RPush(ctx context.Context, key string, values []interface{}, ttl ...time.Duration) error {
	duration := defaultTTL
	if len(ttl) > 0 && ttl[0] > 0 {
		duration = ttl[0]
	}

	// Serialize values
	serialized := make([]interface{}, len(values))
	for i, v := range values {
		data, err := json.Marshal(v)
		if err != nil {
			return fmt.Errorf("failed to marshal value: %w", err)
		}
		serialized[i] = data
	}

	err := c.Client.RPush(ctx, key, serialized...).Err()
	if err != nil {
		return fmt.Errorf("failed to push to list: %w", err)
	}

	return c.Expire(ctx, key, duration)
}

// LRange retrieves a range of elements from a list
func (c RedisClientWithExpire) LRange(ctx context.Context, key string, start int64, stop int64) ([]interface{}, error) {
	vals, err := c.Client.LRange(ctx, key, start, stop).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get list range: %w", err)
	}

	result := make([]interface{}, len(vals))
	for i, val := range vals {
		if val == "null" || val == "" {
			result[i] = nil
			continue
		}

		var parsed interface{}
		err := json.Unmarshal([]byte(val), &parsed)
		if err != nil {
			result[i] = val
			continue
		}
		result[i] = parsed
	}

	return result, nil
}

// LLen returns the length of a list
func (c RedisClientWithExpire) LLen(ctx context.Context, key string) (int64, error) {
	return c.Client.LLen(ctx, key).Result()
}

// ============================================================================
// SET OPERATIONS
// ============================================================================

// SAdd adds one or more members to a set with expiration
func (c RedisClientWithExpire) SAdd(ctx context.Context, key string, members []interface{}, ttl ...time.Duration) error {
	duration := defaultTTL
	if len(ttl) > 0 && ttl[0] > 0 {
		duration = ttl[0]
	}

	// Serialize members
	serialized := make([]interface{}, len(members))
	for i, m := range members {
		data, err := json.Marshal(m)
		if err != nil {
			return fmt.Errorf("failed to marshal member: %w", err)
		}
		serialized[i] = data
	}

	err := c.Client.SAdd(ctx, key, serialized...).Err()
	if err != nil {
		return fmt.Errorf("failed to add to set: %w", err)
	}

	return c.Expire(ctx, key, duration)
}

// SMembers returns all members of a set
func (c RedisClientWithExpire) SMembers(ctx context.Context, key string) ([]interface{}, error) {
	vals, err := c.Client.SMembers(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get set members: %w", err)
	}

	result := make([]interface{}, len(vals))
	for i, val := range vals {
		if val == "null" || val == "" {
			result[i] = nil
			continue
		}

		var parsed interface{}
		err := json.Unmarshal([]byte(val), &parsed)
		if err != nil {
			result[i] = val
			continue
		}
		result[i] = parsed
	}

	return result, nil
}

// SCard returns the cardinality (number of elements) of a set
func (c RedisClientWithExpire) SCard(ctx context.Context, key string) (int64, error) {
	return c.Client.SCard(ctx, key).Result()
}

// SIsMember returns whether a member exists in a set
func (c RedisClientWithExpire) SIsMember(ctx context.Context, key string, member interface{}) (bool, error) {
	serialized, err := json.Marshal(member)
	if err != nil {
		return false, fmt.Errorf("failed to marshal member: %w", err)
	}

	return c.Client.SIsMember(ctx, key, serialized).Result()
}

// SRem removes one or more members from a set
func (c RedisClientWithExpire) SRem(ctx context.Context, key string, members ...interface{}) (int64, error) {
	if len(members) == 0 {
		return 0, nil
	}

	return c.Client.SRem(ctx, key, members...).Result()
}
