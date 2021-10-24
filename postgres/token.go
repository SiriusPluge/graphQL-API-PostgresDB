package postgres

//func (u *model.User) GenToken() (*AuthToken, error) {
//	expiredAt := time.Now().Add(time.Hour * 24 * 7) // a week
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
//		ExpiresAt: expiredAt.Unix(),
//		Id:        u.ID,
//		IssuedAt:  time.Now().Unix(),
//		Issuer:    "meetmeup",
//	})
//
//	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
//	if err != nil {
//		return nil, err
//	}
//
//	return &AuthToken{
//		AccessToken: accessToken,
//		ExpiredAt:   expiredAt,
//	}, nil
//}
