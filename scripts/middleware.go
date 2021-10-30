package scripts

import (
	"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/request"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"graphQL-API-PostgresDB/graph/model"
	"graphQL-API-PostgresDB/postgres"
	"net/http"
	"os"
	"strings"
)

const CurrentUserKey = "currentUser"

func AuthMiddleware(u postgres.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := parseToken(r)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)

			if !ok || !token.Valid {
				next.ServeHTTP(w, r)
				return
			}

			user, err := u.GetUserByPhone(claims["phone"].(string))
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), CurrentUserKey, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

var authHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromToken,
}

func stripBearerPrefixFromToken(token string) (string, error) {
	bearer := "BEARER"

	if len(token) > len(bearer) && strings.ToUpper(token[0:len(bearer)]) == bearer {
		return token[len(bearer)+1:], nil
	}

	return token, nil
}

var authExtractor = &request.MultiExtractor{
	authHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

func parseToken(r *http.Request) (*jwt.Token, error) {
	jwtToken, err := request.ParseFromRequest(r, authExtractor, func(token *jwt.Token) (interface{}, error) {
		t := []byte(os.Getenv("JWT_SECRET"))
		return t, nil
	})

	return jwtToken, errors.Wrap(err, "parseToken error: ")
}

func GetCurrentUserFromCTX(ctx context.Context) (*model.User, error) {
	errNoUserInContext := errors.New("no user in context")

	if ctx.Value(CurrentUserKey) == nil {
		return nil, errNoUserInContext
	}

	user, ok := ctx.Value(CurrentUserKey).(*model.User)
	if !ok || user.Phone == "" {
		return nil, errNoUserInContext
	}

	return user, nil
}