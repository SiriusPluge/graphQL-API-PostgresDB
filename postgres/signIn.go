package postgres

import (
	"fmt"
	"golang.org/x/net/context"
	"graphQL-API-PostgresDB/graph/model"
	"log"
)

func (u *DB) SignIn(ctx context.Context, input model.SignInByCodeInput) (model.SignInOrErrorPayload) {

	user1 := new(model.Viewer)
	if err := u.DB.NewSelect().Model(user1).Where("id = ?", 1).Scan(ctx); err != nil {
		panic(err)
	}
	fmt.Println(user1)

	getUserToken := new(model.User)
	err := u.DB.NewSelect().Model(&getUserToken).OrderExpr("id ASC")
	fmt.Println(getUserToken)

	codeClient := "0000" // we receive the code from the client
	codeVerify := "0000" // code issued by the service
	if codeClient != codeVerify {
		return model.ErrorPayload{Message: "The code is incorrect"}
	}

	token, err := u.GenerateToken(getUserToken.Phone)
	if err != nil {
		log.Fatalf("Error in take token: %v", err)
	}

	authUser := model.SignInPayload{
		Token: token,
		Viewer: user1,
	}
	return authUser
}


