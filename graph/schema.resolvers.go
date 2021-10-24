package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"fmt"

	"github.com/uptrace/bun"
	"graphQL-API-PostgresDB/graph/generated"
	"graphQL-API-PostgresDB/graph/model"
)

type db struct {
	db *bun.DB
}

func (r *mutationResolver) RequestSignInCode(ctx context.Context, input model.RequestSignInCodeInput) (*model.ErrorPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SignInByCode(ctx context.Context, input model.SignInByCodeInput) (model.SignInOrErrorPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	result, err := r.Domain.DB.GetProducts(ctx)
	if err != nil {
		panic(fmt.Errorf("Error in resolvers Products"))
	}
	return result, nil
}

func (r *queryResolver) Viewer(ctx context.Context) (*model.Viewer, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
