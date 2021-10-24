package postgres

import (
	"fmt"
	"github.com/uptrace/bun"
	"golang.org/x/net/context"
	"graphQL-API-PostgresDB/graph/model"
)

type DB struct {
	DB *bun.DB
}

func (u *DB) GetProducts(ctx context.Context) ([]*model.Product, error) {

	var products []*model.Product
	err := u.DB.NewSelect().Model(&products).OrderExpr("id ASC").Scan(ctx)
	if err != nil {
		fmt.Errorf("Error for GetProducts: %v", err)
	}
	return products, nil
}