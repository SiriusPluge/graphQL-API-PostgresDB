package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphQL-API-PostgresDB/graph/generated"
	"graphQL-API-PostgresDB/graph/model"
	"graphQL-API-PostgresDB/scripts"
)

func (r *mutationResolver) RequestSignInCode(ctx context.Context, input model.RequestSignInCodeInput) (*model.ErrorPayload, error) {
	_, err := r.Domain.DB.GetInCode(input)
	if err != nil {
		fmt.Errorf("Error in RequestSignInCode: %s", err)
	}
	var msg *model.ErrorPayload
	return msg, nil
}

func (r *mutationResolver) SignInByCode(ctx context.Context, input model.SignInByCodeInput) (model.SignInOrErrorPayload, error) {
	getUserByToken := r.Domain.DB.SignIn(ctx, input)
	if getUserByToken != nil {
		return getUserByToken, nil
	} else {
		return &model.ErrorPayload{Message: "Authorization error"}, nil
	}
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	result, err := r.Domain.DB.GetProducts(ctx)
	if err != nil {
		panic(fmt.Errorf("Error in resolvers Products"))
	}
	return result, nil
}

func (r *queryResolver) Viewer(ctx context.Context) (*model.Viewer, error) {

	user1 := new(model.User)
	if err := r.Domain.DB.DB.NewSelect().Model(user1).Where("id = ?", 1).Scan(ctx); err != nil {
		panic(err)
	}

	token := scripts.AuthorizationTokenKey
	if token != "" {
		user := user1
		return &model.Viewer{User: user}, nil
	}

	return &model.Viewer{User: nil}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
