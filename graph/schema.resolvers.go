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

	//checking availability in the database
	check := r.Domain.DB.UserPresencePhone(ctx, input.Phone)
	if check == false {

		//sending the code to the user
		code, errCode := r.Domain.DB.GetInCode(input.Phone)
		if errCode != nil {
			panic(errCode)
		}

		//adding code to the database
		var codeUsers model.CodeUsers
		codeUsers.Phone = input.Phone
		codeUsers.AuthCode = code

		_, errSaveCode := r.Domain.DB.DB.NewInsert().
			Model(&codeUsers).
			Exec(ctx)
		if errSaveCode != nil {
			panic(errSaveCode)
		}

		var msg *model.ErrorPayload
		return msg, nil

	} else {

		//sending the code to the user
		code, errGet := r.Domain.DB.GetInCode(input.Phone)
		if errGet != nil {
			panic(errGet)
		}

		//adding code to the database
		var codeUsers model.CodeUsers
		codeUsers.Phone = input.Phone
		codeUsers.AuthCode = code

		_, errSaveCode := r.Domain.DB.DB.NewUpdate().
			Model(&codeUsers).
			Column("auth_code").
			Where("phone = ?", &codeUsers.Phone).
			Exec(ctx)
		if errSaveCode != nil {
			panic(errSaveCode)
		}

		var msg *model.ErrorPayload
		return msg, nil
	}
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
	getToken, errGetPhoneFromCTX := scripts.GetTokenFromCTX(ctx)
	if errGetPhoneFromCTX != nil {
		panic(errGetPhoneFromCTX)
	}

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
