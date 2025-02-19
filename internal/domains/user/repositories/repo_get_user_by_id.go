package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/kaolnwza/proj-blueprint/libs/entities"
)

func (r repo) GetUserById(ctx context.Context, userId uuid.UUID) (*entities.Users, error) {
	resp := entities.Users{}
	q := `
		SELECT
			id,
			firstname,
			lastname
		FROM users
		WHERE id = ?`

	return &resp, r.pgxDb.Query(ctx, q, &resp, q)
}
