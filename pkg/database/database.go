package database

import "context"

type RdbmsDB[T any] interface {
	New(ctx context.Context) T
}
