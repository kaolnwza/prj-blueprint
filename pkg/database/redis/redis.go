package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/kaolnwza/proj-blueprint/config"
	"github.com/kaolnwza/proj-blueprint/pkg/database"
	"github.com/kaolnwza/proj-blueprint/pkg/logger"
	"github.com/sirupsen/logrus"
)

type rdb struct {
	cli *redis.Client
}

func New[T any](conf config.BaseRedisConfig[T], connConf config.RedisConnectionConfig) database.Redis {
	conn := redisConnect(conf, connConf)
	return rdb{
		cli: conn,
	}
}

func redisConnect[T any](conf config.BaseRedisConfig[T], connConf config.RedisConnectionConfig) *redis.Client {
	logger := logger.GetById("init-redis")
	options := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       conf.Db,
	}

	if !connConf.Default {
		options.MaxRetries = connConf.MaxRetries
		options.PoolSize = connConf.PoolSize
		options.MinIdleConns = connConf.MinIdleConns

		options.MinRetryBackoff = parseDurationWithLog(connConf.MinRetryBackoff, "MinRetryBackoff", logger)
		options.MaxRetryBackoff = parseDurationWithLog(connConf.MaxRetryBackoff, "MaxRetryBackoff", logger)
		options.DialTimeout = parseDurationWithLog(connConf.DialTimeout, "DialTimeout", logger)
		options.ReadTimeout = parseDurationWithLog(connConf.ReadTimeout, "ReadTimeout", logger)
		options.WriteTimeout = parseDurationWithLog(connConf.WriteTimeout, "WriteTimeout", logger)
		options.MaxConnAge = parseDurationWithLog(connConf.MaxConnAge, "MaxConnAge", logger)
		options.PoolTimeout = parseDurationWithLog(connConf.PoolTimeout, "PoolTimeout", logger)
		options.IdleTimeout = parseDurationWithLog(connConf.IdleTimeout, "IdleTimeout", logger)
		options.IdleCheckFrequency = parseDurationWithLog(connConf.IdleCheckFrequency, "IdleCheckFrequency", logger)

		logger.Infof("Setup Redis with custom settings.")
	}

	conn := redis.NewClient(options)
	if err := conn.Ping(context.Background()).Err(); err != nil {
		logger.Fatalf("Cannot connect to Redis: %s", err)
	}

	return conn
}

func parseDurationWithLog(value, field string, logger *logrus.Entry) time.Duration {
	if value == "" {
		return 0
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		logger.Errorf("Failed to parse %s=%v in Redis, err=%v", field, value, err)
		return 0
	}

	return duration
}

// SetValue - Save value to rdb using `key` and `value` which could be any any value.
// Example - SetValue("test::string", "abc", 0)
func (r rdb) SetValue(ctx context.Context, key string, value interface{}, exp time.Duration) error {
	b, err := json.Marshal(&value)

	if err != nil {
		return err
	}

	return r.cli.Set(ctx, key, string(b), exp).Err()
}

// GetValue - Get value from rdb and change type back, according to submit type using `key` and `value` which MUST be address of target type.
// Example
// var getString string
// GetValue("test::string", &getString)
func (r rdb) GetValue(ctx context.Context, key string, value interface{}) error {
	result, err := r.cli.Get(ctx, key).Result()

	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(result), value)
}

// IsExist - Check whether key is exist in rdb or not by using `key`.
func (r rdb) IsExist(ctx context.Context, key string) bool {
	val := r.cli.Exists(ctx, key).Val()
	return val == 1
}

// DeleteValue - Delete value from rdb using `key` and return error.
func (r rdb) DeleteValue(ctx context.Context, key string) error {
	return r.cli.Del(ctx, key).Err()
}

// ClearDB - Clear all cache from rdb asynchronous or synchronous using `isAsync` parameter .
func (r rdb) ClearDB(ctx context.Context, isAsync bool) error {
	if isAsync {
		return r.cli.FlushAll(ctx).Err()
	} else {
		return r.cli.FlushAllAsync(ctx).Err()
	}
}

// SetExpire - Set expire time (second) to value in rdb using `key` and exp.
func (r rdb) SetExpire(ctx context.Context, key string, exp time.Duration) error {
	return r.cli.Expire(ctx, key, exp).Err()
}

// SetExpireAt - Same behavior as SetExpire, but set expire time using time.Time (specific time in Unix clock).
func (r rdb) SetExpireAt(ctx context.Context, key string, expAt time.Time) error {
	return r.cli.ExpireAt(ctx, key, expAt).Err()
}
