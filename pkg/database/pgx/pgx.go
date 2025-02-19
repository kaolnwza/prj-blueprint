package pgxdb

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/kaolnwza/proj-blueprint/config"
	"github.com/kaolnwza/proj-blueprint/pkg/database"
)

type pgxDb struct {
	conn *pgx.Conn
}

func New(conf config.BaseDatabaseConfig) database.RdbmsDB[*pgx.Conn] {
	conn, err := pgxConnect(conf)
	if err != nil {
		panic(fmt.Errorf("pgx failed to connect database: %v, err = %w", conf.Host, err))
	}

	return pgxDb{
		conn: conn,
	}
}

func pgxConnect(conf config.BaseDatabaseConfig) (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), conf.GetMysqlDsn())
}

func (d pgxDb) New(ctx context.Context) *pgx.Conn {
	return d.conn
}

func (d pgxDb) Query(ctx context.Context, query string, dest any, args ...interface{}) error {
	err := d.conn.QueryRow(ctx, query, args...).Scan(&dest)
	if err != nil {
		return fmt.Errorf("Query failed: %v", err)
	}
	return nil
}

func (d pgxDb) Exec(ctx context.Context, query string, args ...interface{}) error {
	_, err := d.conn.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("Exec failed: %v", err)
	}
	return nil
}

// execute with returning.
func (d pgxDb) ExecReturning(ctx context.Context, query string, dest any, args ...interface{}) error {
	if err := d.conn.QueryRow(ctx, query, args...).Scan(&dest); err != nil {
		return fmt.Errorf("ExecReturning failed: %v", err)
	}
	return nil
}
