package postgres

import (
	"log"
	"time"
	"github.com/golang-jwt/jwt"
)

func (u *DB) GenToken() (string) {
	var hmacSampleSecret []byte

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		log.Fatal("error while generating the token: %v", err)
	}
	return tokenString
}
