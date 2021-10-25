package postgres

import (
	"golang.org/x/net/context"
	"graphQL-API-PostgresDB/graph/model"
	"log"
)

func (u *DB) Autorization(ctx context.Context, input model.SignInByCodeInput) (model.SignInOrErrorPayload) {

	getUser := u.DB.NewSelect().Model(&input).OrderExpr("id ASC").Scan(ctx)
	if getUser != nil {
		token := u.GenToken()
		viewer :=
		authUser := model.SignInPayload{
			Token: token,
			Viewer: viewer,
		}
		return authUser
	} else {
		log.Fatal("there is no such User")
	}

}


