package repositories

import (
	"github.com/jackc/pgx/v5"
	"github.com/kaolnwza/proj-blueprint/internal/domains/user/core/ports"
	"github.com/kaolnwza/proj-blueprint/pkg/database"
	"gorm.io/gorm"
)

type repo struct {
	gormDb database.RdbmsDB[*gorm.DB]
	pgxDb  database.RdbmsDB[*pgx.Conn]
}

func New(gormDb database.RdbmsDB[*gorm.DB], pgxDb database.RdbmsDB[*pgx.Conn]) ports.Repository {
	return repo{
		gormDb: gormDb,
		pgxDb:  pgxDb,
	}
}
