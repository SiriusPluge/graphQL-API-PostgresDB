package postgres

import (
	"fmt"
	"golang.org/x/net/context"
	"graphQL-API-PostgresDB/graph/model"
	"log"
)

func (u *DB) SignIn(ctx context.Context, input model.SignInByCodeInput) (model.SignInOrErrorPayload) {

	user1 := new(model.User)
	if err := u.DB.NewSelect().Model(user1).Where("phone = ?", input.Phone).Scan(ctx); err != nil {
		fmt.Println(err)
		panic(err)
	}

	viewer := model.Viewer{
		user1,
	}

	codeClient := "0000" // we receive the code from the client
	codeVerify := "0000" // code issued by the service

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


