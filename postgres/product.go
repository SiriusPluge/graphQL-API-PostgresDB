package postgres

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"graphQL-API-PostgresDB/graph/model"
	//"github.com/go-pg/pg/v10"
)

func (u *DB) GetProducts(ctx context.Context) ([]*model.Product, error) {

	products := make([]*model.Product, 0)
	err := u.DB.NewSelect().Model(&products).OrderExpr("id ASC").Scan(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error for GetProducts")
	}
	return products, nil
}