package postgres

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"graphQL-API-PostgresDB/graph/model"
)

func (u *DB) GetUserByField(field, value string) *model.User {
	var user model.User
	err := u.DB.NewSelect().Model(&user).Where(fmt.Sprintf("%v = ?", field), value)
	if err != nil {
		errors.New("error in getting user from DB")
	}
	return &user
}

func (u *DB) GetUserByPhone(phone string) *model.User {
	return u.GetUserByField("phone", phone)
}

func (u *DB) LastIndexUsers(ctx context.Context) (int, error) {
	var user []model.User
	err := u.DB.NewSelect().Model(&user).Scan(ctx)
	if err != nil {
		errors.New("error in getting user from DB")
	}
	index := len(user)
	return index, nil
}

func (u *DB) UserPresence(ctx context.Context, phone string) (bool, *model.User) {
	var user model.User
	err := u.DB.NewSelect().Model(&user).Where("phone = ?", phone).Scan(ctx)
	if err != nil {
		return false, nil
	} else {
		return true, &user
	}
}

func (u *DB) InsertCodeID(ctx context.Context, codeUser *model.CodeUsers) (*sql.Result, error) {
	res, err := u.DB.NewInsert().
		Model(&codeUser).
		Exec(ctx)

	if err != nil {
		errors.New("the code don`t saved")
		return &res, err
	}
	fmt.Println("the code has been saved successfully")
	return &res, nil
}