package main

import (
	"github.com/kaolnwza/proj-blueprint/config"
	gormdb "github.com/kaolnwza/proj-blueprint/pkg/database/gorm"
)

func main() {
	conf := config.New()
	myDbKubDb := gormdb.New(conf.DatabaseConf.MyDbKub)
}
