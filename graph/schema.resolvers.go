package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphQL-API-PostgresDB/graph/generated"
	"graphQL-API-PostgresDB/graph/model"
	"graphQL-API-PostgresDB/scripts"
)

func (r *mutationResolver) RequestSignInCode(ctx context.Context, input model.RequestSignInCodeInput) (*model.ErrorPayload, error) {
	return r.Domain.DB.ReqSignInCode(ctx, input)
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
	return r.Domain.DB.GetProducts(ctx)
}

func (r *queryResolver) Viewer(ctx context.Context) (*model.Viewer, error) {

	//
	getToken := scripts.GetTokenFromCTX(ctx)

	claims := scripts.DecodeToken(getToken)

	var user model.User
	if err := r.Domain.DB.DB.NewSelect().Model(&user).Where("phone = ?", claims.Phone).Scan(ctx); err != nil {
		panic(err)
	}

	return &model.Viewer{User: &user}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
