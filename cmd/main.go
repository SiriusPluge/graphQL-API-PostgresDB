package cmd

import (
	"fmt"
	"graphQL-API-PostgresDB/graph/model"
	"graphQL-API-PostgresDB/postgres"
	"os"
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

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	var codeUser model.CodeUsers
	codeUser.UsersId = 2
	codeUser.AuthCode = "1234"
	res2, errSaveCode := DB.NewInsert().
		Model(&codeUser).
		Exec()
	if errSaveCode != nil {
		fmt.Errorf("%v", errSaveCode)
	}
	fmt.Println(res2)
}
