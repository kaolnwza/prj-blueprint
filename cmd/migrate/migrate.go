package main

import (
	gormdb "github.com/kaolnwza/proj-blueprint/pkg/database/gorm"
)

func main() {
	gormdb.RunGormAutoMigrate()
}
