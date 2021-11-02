package postgres

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func (u *DB) GetInCode(phone string) (string, error) {

	rand.Seed(time.Now().UnixNano())
	c := randomInt(1000, 9999)
	sms := strconv.Itoa(c)
	fmt.Printf("На Ваш номер телефона: %s отправлен код для входа: %s \n", phone, sms)
	return sms, nil
}
