package postgres

import (
	"fmt"
	"graphQL-API-PostgresDB/graph/model"
	"log"
)

func (u *DB) GetUserByField(field, value string) (*model.User, error) {
	var user model.User
	err := u.DB.NewSelect().Model(&user).Where(fmt.Sprintf("%v = ?", field), value)
	if err != nil {
		log.Fatalf("error in getting user from DB")
	}
	return &user, nil
}

func (u *DB) GetUserByPhone(phone string) (*model.User, error) {
	return u.GetUserByField("phone", phone)
}
