package model

type User struct {
	ID    int    `json:"id"`
	Phone string `json:"phone"`
}

type CodeUsers struct {
	ID    int    `json:"id"`
	UsersId int `json:"users_id"`
	AuthCode string `json:"code"`
}

//type UserUp struct {
//	ID    int    `json:"id"`
//	Phone string `json:"phone"`
//	AuthCode string `json:"code"`
//	DepatureTime time.Time
//}

type LoginInput struct {
	Phone     string
}
