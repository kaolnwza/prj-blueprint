package gormdb

import (
	"context"
	"fmt"

	"github.com/kaolnwza/proj-blueprint/config"
	"github.com/kaolnwza/proj-blueprint/pkg/database"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type txKey struct{}

type GormDB struct {
	db *gorm.DB
}

func (g GormDB) New(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txKey{}).(*gorm.DB); ok {
		return GormDB{db: tx}.db.WithContext(ctx)
	}

	return g.db.WithContext(ctx)
}

func New(conf config.DatabaseConfig) (database.RdbmsDB[*gorm.DB], error) {
	db, err := GormConnect(conf)
	if err != nil {
		return nil, err
	}

	return GormDB{db: db}, nil
}

const (
	DriverPostgres = "postgres"
)

func GormConnect(conf config.DatabaseConfig) (*gorm.DB, error) {
	switch conf.Driver {
	case DriverPostgres:
		dsn := newGormDSN(conf.Host, conf.Username, conf.Password, conf.Database, conf.Port)
		return gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{
			// Logger: logger.Default.LogMode(logger.Info),
			Logger: logger.Default.LogMode(logger.Silent),
		})
	}

	return nil, fmt.Errorf("unknow driver")
}

func newGormDSN(host, user, password, db, port string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", host, user, password, db, port)
}
