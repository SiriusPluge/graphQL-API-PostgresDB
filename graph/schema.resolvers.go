package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"graphQL-API-PostgresDB/graph/generated"
	"graphQL-API-PostgresDB/graph/model"
	"graphQL-API-PostgresDB/scripts"
)

func (r *mutationResolver) RequestSignInCode(ctx context.Context, input model.RequestSignInCodeInput) (*model.ErrorPayload, error) {

	lastIndex, err := r.Domain.DB.LastIndexUsers(ctx)
	if err != nil {
		fmt.Println("Ошибка")
	}

	check := r.Domain.DB.UserPresence(ctx, input.Phone)
	if check == false {
		errors.New("a user with such a phone was not found")
		//if errGetUser := r.Domain.DB.DB.NewSelect().Model(user1).Where("phone = ?", input.Phone).Scan(ctx); err != nil {
		var u1 model.User
		u1.Phone = input.Phone
		u1.ID = lastIndex + 1
		_, errInsert := r.Domain.DB.DB.NewInsert().Model(&u1).Exec(ctx)
		if errInsert != nil {
			errors.New("Error in errInsert")
		}

		fmt.Println("вариант с вставкой")
		_, err := r.Domain.DB.GetInCode(u1.Phone)
		if err != nil {
			fmt.Errorf("Error in RequestSignInCode: %s", err)
		}

		var msg *model.ErrorPayload
		return msg, nil
	} else {
		fmt.Println("вараинт без вставки")
		_, err := r.Domain.DB.GetInCode(input.Phone)
		if err != nil {
			fmt.Errorf("Error in RequestSignInCode: %s", err)
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

	fmt.Println(ctx)

	headerToken, err := scripts.ParseAuthHeader(ctx)
	if err != nil {
		return nil, errors.New("missing jwt-token in headers request")
	}

	fmt.Println(headerToken)

	phone, err := scripts.Parse(headerToken)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	fmt.Println(phone)

	user := new(model.User)
	if err := r.Domain.DB.DB.NewSelect().Model(&user).Where("phone = ?", phone).Scan(ctx); err != nil {
		panic(err)
	}

	fmt.Println(user)
	return &model.Viewer{User: user}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
