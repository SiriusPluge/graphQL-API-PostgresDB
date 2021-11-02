package scripts

import (
	"golang.org/x/net/context"
	"net/http"
	"strings"
)


var AuthorizationTokenKey string

func AuthorizationTokenContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw := r.Header.Get("Authorization")
		if raw != "" {
			fields := strings.Fields(raw)
			token := fields[len(fields)-1]
			ctx := context.WithValue(
				r.Context(),
				AuthorizationTokenKey,
				token,
			)

			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}

	})
}
