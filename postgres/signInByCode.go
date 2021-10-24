package postgres

import (
	"fmt"
	"graphQL-API-PostgresDB/graph/model"
)

func (u *DB) GetInCode(input model.RequestSignInCodeInput) (string, error) {
	sms := "0000"
	fmt.Printf("На Ваш номер телефона: %s отправлен код для входа: %s", input, sms)
	return sms, nil
}
