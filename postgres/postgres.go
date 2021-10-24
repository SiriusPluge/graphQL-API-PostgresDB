package postgres

import (
	"fmt"
	"github.com/go-pg/pg/v10"
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

func New(opts *pg.Options) *pg.DB {
	return pg.Connect(opts)
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


