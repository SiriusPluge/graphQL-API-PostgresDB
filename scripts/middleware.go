package scripts

import (
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type MyCustomClaims struct {
	Phone interface{}
	jwt.MapClaims
}

func GetTokenFromCTX(ctx context.Context) string {
	errNoUserInContext := errors.New("no user in context")

	if ctx.Value(AuthorizationTokenKey) == nil {
		panic(errNoUserInContext)
	}

	token, ok := ctx.Value(AuthorizationTokenKey).(string)
	if !ok || token == "" {
		errors.Wrapf(errNoUserInContext, "error when receiving the token")
	}

	return token
}

func DecodeToken(getToken string) *MyCustomClaims {

	// Parse the token
	token, err := jwt.ParseWithClaims(getToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		var verifyKey = []byte("secret")
		return verifyKey, nil
	})
	if err != nil {
		errors.Wrap(err, "token decoding error")
	}

	claims := token.Claims.(*MyCustomClaims)

	return claims
}