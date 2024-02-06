package utils

import "context"

func MustGetContext[T any](ctx context.Context, key string) T {
	return ctx.Value(key).(T)
}
