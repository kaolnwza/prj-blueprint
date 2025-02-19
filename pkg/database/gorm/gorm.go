package gormdb

import (
	"context"
	"fmt"

	"github.com/kaolnwza/proj-blueprint/config"
	"github.com/kaolnwza/proj-blueprint/pkg/database"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type txKey struct{}

type gormDB struct {
	db *gorm.DB
}

func (g gormDB) New(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txKey{}).(*gorm.DB); ok {
		return gormDB{db: tx}.db.WithContext(ctx)
	}

	return g.db.WithContext(ctx)
}

func (g gormDB) Query(ctx context.Context, query string, dest any, args ...interface{}) error {
	return g.db.Raw(query, args...).Scan(&dest).Error
}

func (g gormDB) Exec(ctx context.Context, query string, args ...interface{}) error {
	return g.db.Raw(query, args...).Error
}

func (g gormDB) ExecReturning(ctx context.Context, query string, dest any, args ...interface{}) error {
	return g.db.Raw(query, args...).Scan(&dest).Error
}

// default
func New(conf config.BaseDatabaseConfig) database.RdbmsDB[*gorm.DB] {
	db, err := gormConnect(conf)
	if err != nil {
		panic(fmt.Errorf("gorm failed to connect database: %v, err = %w", conf.Host, err))
	}

	return gormDB{db: db}
}

const (
	driverPostgres = "postgres"
	driverMysql    = "mysql"
)

func gormConnect(conf config.BaseDatabaseConfig) (*gorm.DB, error) {
	gormConf := gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
		Logger: logger.Default.LogMode(logger.Silent),
	}

	switch conf.Driver {
	case driverMysql:
		return gorm.Open(mysql.New(mysql.Config{
			DSN:                       conf.GetMysqlDsn(), // data source name
			DefaultStringSize:         256,                // default size for string fields
			DisableDatetimePrecision:  true,               // disable datetime precision, which not supported before MySQL 5.6
			DontSupportRenameIndex:    true,               // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
			DontSupportRenameColumn:   true,               // `change` when rename column, rename column not supported before MySQL 8, MariaDB
			SkipInitializeWithVersion: false,              // auto configure based on currently MySQL version
		}), &gormConf)
	}

	return nil, fmt.Errorf("unknow driver")
}
