package postgres

import (
	"database/sql"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"golang.org/x/net/context"
)

type DBLogger struct {}

type DB struct {
	DB *bun.DB
}

func (d DBLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d DBLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}

func ConnectDB() *bun.DB {
	dsn := "postgres://postgres:postgres@localhost:7323/postgres?sslmode=disable"
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// Create a Bun db on top of it.
	DB := bun.NewDB(pgdb, pgdialect.New())

	return DB
}


