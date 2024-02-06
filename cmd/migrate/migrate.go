package main

import "github.com/kaolnwza/proj-blueprint/pkg/database"

func main() {
	database.RunGormAutoMigrate()
}
