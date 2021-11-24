package postgres

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"graphQL-API-PostgresDB/graph/model"
)

func (u *DB) SignIn(ctx context.Context, input model.SignInByCodeInput) model.SignInOrErrorPayload {

	check, User := u.UserPresence(ctx, input.Phone)
	switch check {
	case true:
		viewer := model.Viewer{
			&User,
		}

		//get the codeUser in DB
		var codeUser model.CodeUsers
		errGetCode := u.DB.NewSelect().
			Model(&codeUser).
			Where("phone = ?", User.Phone).
			Scan(ctx)
		if errGetCode != nil {
			errors.Wrap(errGetCode, "error when getting the code from the database")
		}

		codeClient := input.Code        // we receive the code from the client
		codeVerify := codeUser.AuthCode // code issued by the service

		if codeClient == codeVerify {

			token, err := u.GenerateToken(User.Phone)
			if err != nil {
				errors.Wrap(err, "Error in take token")
			}

			authUser := model.SignInPayload{
				Token:  token,
				Viewer: &viewer,
			}
			return authUser
		}

	case false:
		var NewUser model.User
		NewUser.Phone = input.Phone

		//adding a user to the database
		_, errInsert := u.DB.NewInsert().Model(&NewUser).Exec(ctx)
		if errInsert != nil {
			errors.Wrap(errInsert, "error when adding the code from the database")
		}

		viewer := model.Viewer{
			&NewUser,
		}

		//get the codeUser in DB
		var codeUser model.CodeUsers
		errGetCode := u.DB.NewSelect().
			Model(&codeUser).
			Where("phone = ?", NewUser.Phone).
			Scan(ctx)
		if errGetCode != nil {
			errors.Wrap(errGetCode, "error when getting the code from the database")
		}

		codeClient := input.Code        // we receive the code from the client
		codeVerify := codeUser.AuthCode // code issued by the service

		if codeClient == codeVerify {

			token, err := u.GenerateToken(NewUser.Phone)
			if err != nil {
				errors.Wrap(err, "Error in take token")
			}

			authUser := model.SignInPayload{
				Token:  token,
				Viewer: &viewer,
			}
			return authUser
		}

	default:
		return model.ErrorPayload{Message: "The code is incorrect"}
	}

	return model.ErrorPayload{Message: "The code is incorrect"}
}