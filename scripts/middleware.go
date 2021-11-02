package scripts

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"golang.org/x/net/context"
)

type MyCustomClaims struct {
	Phone interface{}
	jwt.MapClaims
}

func GetTokenFromCTX(ctx context.Context) (string, error) {
	errNoUserInContext := errors.New("no user in context")

	if ctx.Value(AuthorizationTokenKey) == nil {
		panic(errNoUserInContext)
	}

	token, ok := ctx.Value(AuthorizationTokenKey).(string)
	if !ok || token == "" {
		panic(errNoUserInContext)
	}

	return token, nil
}

func DecodeToken(getToken string) *MyCustomClaims {

	// Parse the token
	token, err := jwt.ParseWithClaims(getToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		var verifyKey = []byte("secret")
		return verifyKey, nil
	})
	if err != nil {
		panic(err)
	}

	claims := token.Claims.(*MyCustomClaims)

	return claims
}