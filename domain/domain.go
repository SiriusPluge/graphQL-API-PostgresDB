package domain

import (
	"graphQL-API-PostgresDB/postgres"
)

type Domain struct {
	DB postgres.DB
}

func NewDomain(DB postgres.DB) *Domain {
	return &Domain{DB: DB}
}