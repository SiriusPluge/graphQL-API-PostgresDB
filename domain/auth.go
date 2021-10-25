package domain

import (
	"errors"
	"golang.org/x/net/context"
	"graphQL-API-PostgresDB/graph/model"
	"graphQL-API-PostgresDB/postgres"
)

func (d *Domain) Login(ctx context.Context, input model.RequestSignInCodeInput) (*model.SignInPayload, error) {
	user := d.DB.DB.NewSelect().Model(&input).OrderExpr("id ASC").Scan(ctx)
	codeClient, _ := d.DB.GetInCode(input)
	codeVerify := "0000"

	if user != nil {
		if codeClient == codeVerify {
			token, err := postgres.GenerateToken(user)
			if err != nil {
				return nil, errors.New("something went wrong")
			}
		}
	}


	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}
