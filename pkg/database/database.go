package database

import (
	"context"
	"time"
)

type RdbmsDB[T any] interface {
	New(ctx context.Context) T // call driver, use in scenarios of using specific method of driver.
	Query(ctx context.Context, query string, dest any, args ...interface{}) error
	Exec(ctx context.Context, query string, args ...interface{}) error
	ExecReturning(ctx context.Context, query string, dest any, args ...interface{}) error
}

type Redis interface {
	SetValue(ctx context.Context, key string, value interface{}, exp time.Duration) error
	GetValue(ctx context.Context, key string, value interface{}) error
	IsExist(ctx context.Context, key string) bool
	DeleteValue(ctx context.Context, key string) error
	SetExpire(ctx context.Context, key string, exp time.Duration) error
	SetExpireAt(ctx context.Context, key string, expAt time.Time) error
}
