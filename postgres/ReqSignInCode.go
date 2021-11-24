package postgres

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"graphQL-API-PostgresDB/graph/model"
)

func (u *DB) ReqSignInCode(ctx context.Context, input model.RequestSignInCodeInput) (*model.ErrorPayload, error) {

	//checking availability in the database
	check := u.UserPresencePhone(ctx, input.Phone)
	if check == false {

		//sending the code to the user
		code := u.GetInCode(input.Phone)

		//adding code to the database
		var codeUsers model.CodeUsers
		codeUsers.Phone = input.Phone
		codeUsers.AuthCode = code

		_, errSaveCode := u.DB.NewInsert().
			Model(&codeUsers).
			Exec(ctx)
		if errSaveCode != nil {
			msgErr := model.ErrorPayload{Message: "error when adding code to the database"}
			return &msgErr, errors.Wrap(errSaveCode, "error when adding code to the database")
		}

		return nil, nil

	} else {

		//sending the code to the user
		code := u.GetInCode(input.Phone)

		//adding code to the database
		var codeUsers model.CodeUsers
		codeUsers.Phone = input.Phone
		codeUsers.AuthCode = code

		_, errSaveCode := u.DB.NewUpdate().
			Model(&codeUsers).
			Column("auth_code").
			Where("phone = ?", &codeUsers.Phone).
			Exec(ctx)
		if errSaveCode != nil {
			msgErr := model.ErrorPayload{Message: "error when updating the code in the database"}
			return &msgErr, errors.Wrap(errSaveCode, "error when updating the code in the database")
		}

		return nil, nil
	}
}
