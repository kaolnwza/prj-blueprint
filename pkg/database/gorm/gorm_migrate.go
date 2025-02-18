package gormdb

// func RunGormAutoMigrate() {
// 	conf := config.New()
// 	db, err := New(conf.BDatabaseConf)
// 	if err != nil {
// 		panic(fmt.Errorf("Gorm auto migrate database connection failed: %w", err))
// 	}

// 	ctx := context.Background()
// 	if err := db.New(ctx).Transaction(func(tx *gorm.DB) error {
// 		return db.New(ctx).AutoMigrate(
// 			&entities.Users{},
// 			&entities.Wallets{},
// 		)
// 	}); err != nil {
// 		log.Print(fmt.Errorf("Gorm auto migrate failed: %w", err))
// 	}
// }
