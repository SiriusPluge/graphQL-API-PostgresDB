package postgres

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type MyCustomClaims struct {
	Phone string
	jwt.MapClaims
}

//PublicKey set secret public key
var PublicKey = []byte("secret")

// GenerateToken generates JWT token en returns it
func (u *DB) GenerateToken(phone string) (string, error) {

	// Create the Claims
	claims := MyCustomClaims{
		phone,
		jwt.MapClaims{
			"exp": time.Now().Add(time.Hour * time.Duration(24)).Unix(),
			"iat": time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, errGetToken := token.SignedString(PublicKey)

	return tokenString, errGetToken
}