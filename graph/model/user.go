package model

type User struct {
	ID    int    `json:"id"`
	Phone string `json:"phone"`
}

type LoginInput struct {
	Phone     string
}
