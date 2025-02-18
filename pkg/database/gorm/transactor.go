package gormdb

import (
	"context"
	"fmt"

	"github.com/kaolnwza/proj-blueprint/pkg/logger"
	"gorm.io/gorm"
)

type transactor struct {
	db GormDB
}

type Transactor interface {
	WithinTransaction(ctx context.Context, tFunc func(fnCtx context.Context) error) error
}

// made for handle ACID properties on biz layer.
func NewTransactor(db GormDB) Transactor {
	return &transactor{db: db}
}

func injectTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func extractTx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txKey{}).(*gorm.DB); ok {
		return tx
	}

	return nil
}

// working on Tx: begin -> all_flows -> commit
func (t transactor) WithinTransaction(ctx context.Context, tFunc func(fnCtx context.Context) error) error {
	log := logger.GetByContext(ctx)
	tx := t.db.db.Begin()

	if err := tFunc(injectTx(ctx, tx)); err != nil {
		if err := tx.Rollback().Error; err != nil {
			log.Error(ctx, fmt.Errorf("rollback transaction error: %v", err))
			return err
		}

		return err
	}

	if err := tx.Commit().Error; err != nil {
		log.Error(ctx, fmt.Errorf("commit transaction error: %v", err))
		return err
	}

	return nil
}

// calling before execute database.
func (t transactor) GetTx(ctx context.Context) *gorm.DB {
	tx := extractTx(ctx)
	if tx != nil {
		return tx // with tx
	}

	return t.db.db // without tx
	// return t.db.driver.DB // without tx
}
