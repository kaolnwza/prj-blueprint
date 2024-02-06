package database

import (
	"fmt"
	"log"

	"github.com/kaolnwza/proj-blueprint/config"
	"github.com/kaolnwza/proj-blueprint/libs/schemes"
	"gorm.io/gorm"
)

func RunGormAutoMigrate() {
	conf := config.NewConfig()
	db, err := NewGormDatabase(conf.DatabaseConf)
	if err != nil {
		panic(fmt.Errorf("Gorm auto migrate database connection failed: %w", err))
	}

	if err := db.Transaction(func(tx *gorm.DB) error {
		return db.AutoMigrate(
			&schemes.Users{},
			&schemes.Wallets{},
		)
	}); err != nil {
		log.Print(fmt.Errorf("Gorm auto migrate failed: %w", err))
	}
}
