package postgres

import (
	"fmt"
	"graphQL-API-PostgresDB/graph/model"
)

func (u *DB) GetSignInCode(input model.RequestSignInCodeInput) (phone *model.RequestSignInCodeInput, err error) {
	sms := "0000"
	fmt.Printf("На Ваш номер телефона %s отправлен код для входа: %s", input, sms)
	return phone
}
