package utils

import (
	"context"
	"fmt"
)

func MustGetContext[T any](ctx context.Context, key string) T {
	return ctx.Value(key).(T)
}

func SetContext[T any](ctx context.Context, key string, value T) context.Context {
	return context.WithValue(ctx, key, value)
}

func GetContext[T any](ctx context.Context, key string) (T, error) {
	v := ctx.Value(key)

	res, ok := v.(T)
	if !ok {
		var t T
		return t, fmt.Errorf("get context: failed to convert context value type")
	}

	return res, nil
}
