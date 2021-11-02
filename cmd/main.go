package main

import (
	"fmt"
	"golang.org/x/net/context"
	"graphQL-API-PostgresDB/graph/model"
	"graphQL-API-PostgresDB/postgres"
)

const defaultPort = "8080"

type User struct {
	ID    int    `json:"id"`
	Phone string `json:"phone"`
}

type CodeUsers struct {
	ID    int    `json:"id"`
	UsersId int `json:"users_id"`
	AuthCode string `json:"code"`
}

func main() {
	DB := postgres.ConnectDB()
	defer DB.Close()

	ctx := context.Background()

	var codeUser *model.CodeUsers
	codeUser.ID = 3
	codeUser.UsersId = 2
	codeUser.AuthCode = "1234"

	_, errSaveCode := DB.NewInsert().
		Model(&codeUser).
		Exec(ctx)
	if errSaveCode != nil {
		fmt.Errorf("%v", errSaveCode)
	}
}
