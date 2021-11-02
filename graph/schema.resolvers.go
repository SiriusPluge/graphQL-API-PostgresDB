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
	//getting the latest index
	lastIndex, err := r.Domain.DB.LastIndexUsers(ctx)
	if err != nil {
		panic(err)
	}

	var u1 model.User
	u1.Phone = input.Phone
	u1.ID = lastIndex + 1

	//checking availability in the database
	check, User := r.Domain.DB.UserPresence(ctx, input.Phone)
	if check == false {

		//adding a user to the database
		_, errInsert := r.Domain.DB.DB.NewInsert().Model(&u1).Exec(ctx)
		if errInsert != nil {
			panic(errInsert)
		}

		//sending the code to the user
		code, errCode := r.Domain.DB.GetInCode(u1.Phone)
		if errCode != nil {
			panic(errCode)
		}

		//adding code to the database
		var codeUsers model.CodeUsers
		codeUsers.UsersId = u1.ID
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
		if err != nil {
			panic(errGet)
		}

		//adding code to the database
		var codeUsers model.CodeUsers
		codeUsers.UsersId = User.ID
		codeUsers.AuthCode = code

		_, errSaveCode := r.Domain.DB.DB.NewInsert().
			Model(&codeUsers).
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
	result, err := r.Domain.DB.GetProducts(ctx)
	if err != nil {
		panic(fmt.Errorf("Error in resolvers Products"))
	}
	return result, nil
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
