package postgres

import (
	"github.com/go-pg/pg/v10"
	"golang.org/x/net/context"
)

func NewPostgresDB() {

	// init DB
	opt, err := pg.ParseURL("postgres://user:pass@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)

	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		panic(err)
	}


}

