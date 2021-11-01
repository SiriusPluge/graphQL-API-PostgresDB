package postgres

import (
	"fmt"
	"golang.org/x/net/context"
	"graphQL-API-PostgresDB/graph/model"
	//"github.com/go-pg/pg/v10"
)

func (u *DB) GetProducts(ctx context.Context) ([]*model.Product, error) {

	products := make([]*model.Product, 0)
	err := u.DB.NewSelect().Model(&products).OrderExpr("id ASC").Scan(ctx)
	//err := u.DB.NewSelect().Model(&products).Scan(ctx)
	if err != nil {
		fmt.Errorf("Error for GetProducts: %v", err)
	}
	return products, nil
}