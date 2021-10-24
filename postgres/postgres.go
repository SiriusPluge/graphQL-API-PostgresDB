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

func (d DBLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d DBLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}

func ConnectDB() *bun.DB {
	dsn := "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
	// dsn := "unix://user:pass@dbname/var/run/postgresql/.s.PGSQL.5432"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	DB := bun.NewDB(sqldb, pgdialect.New())

	return DB
}

//func ConnectPostgresDB() {
//
//	// init DB
//	opt, err := pg.ParseURL("postgres://user:pass@localhost:5432/postgres")
//	if err != nil {
//		panic(err)
//	}
//
//	db := pg.Connect(opt)
//
//	ctx := context.Background()
//	if err := db.Ping(ctx); err != nil {
//		panic(err)
//	}


