package postgres

import (
	"golang.org/x/net/context"
	"graphQL-API-PostgresDB/graph/model"
	"log"
)

func (u *DB) SignIn(ctx context.Context, input model.SignInByCodeInput) (model.SignInOrErrorPayload) {

	user1 := new(model.User)
	if err := u.DB.NewSelect().Model(user1).Where("phone = ?", input.Phone).Scan(ctx); err != nil {
		panic(err)
	}
	viewer := model.Viewer{
		user1,
	}

	//get the codeUser in DB
	var codeUser model.CodeUsers
	errGetCode := u.DB.NewSelect().
		Model(&codeUser).
		Where("users_id = ?", user1.ID).
		Scan(ctx)
	if errGetCode != nil {
		panic(errGetCode)
	}

	codeClient := input.Code // we receive the code from the client
	codeVerify := codeUser.AuthCode // code issued by the service

	if codeClient == codeVerify && user1 != nil {

		token, err := u.GenerateToken(input.Phone)
		if err != nil {
			log.Fatalf("Error in take token: %v", err)
		}

		authUser := model.SignInPayload{
			Token: token,
			Viewer: &viewer,
		}
		return authUser

	} else {
		return model.ErrorPayload{Message: "The code is incorrect"}
	}
}


