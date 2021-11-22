package postgres

import (
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

func (u *DB) UserPresencePhone(ctx context.Context, phone string) bool {
	var codeUsers model.CodeUsers
	err := u.DB.NewSelect().Model(&codeUsers).Where("phone = ?", phone).Scan(ctx)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (u *DB) UserPresence(ctx context.Context, phone string) (bool, model.User) {
	var User model.User
	err := u.DB.NewSelect().Model(&User).Where("phone = ?", phone).Scan(ctx)
	if err != nil {
		return false, User
	} else {
		return true, User
	}
}
