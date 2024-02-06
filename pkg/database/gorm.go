package database

import (
	"fmt"

	"github.com/kaolnwza/proj-blueprint/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormDB struct {
	*gorm.DB
}

func NewGormDatabase(conf config.DatabaseConfig) (GormDB, error) {
	db, err := GormConnect(conf)
	if err != nil {
		return GormDB{}, err
	}

	return GormDB{DB: db}, nil
}

func GormConnect(conf config.DatabaseConfig) (*gorm.DB, error) {
	dsn := newGormDSN(conf.Host, conf.Username, conf.Password, conf.Database, conf.Port)
	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

func newGormDSN(host, user, password, db, port string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", host, user, password, db, port)
}
