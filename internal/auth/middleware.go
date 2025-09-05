package auth

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/htanmo/hackernews/internal/jwt"
	"github.com/htanmo/hackernews/internal/users"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			if header == "" {
				http.Error(w, "Authorization header must be provided", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(header, " ")
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
				return
			}

			tokenStr := parts[1]
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			user := users.User{Username: username}
			id, err := users.GetUserIDByUsername(r.Context(), username)
			if err != nil {
				http.Error(w, "User not found", http.StatusUnauthorized)
				return
			}
			user.ID = strconv.Itoa(id)
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *users.User {
	raw, _ := ctx.Value(userCtxKey).(*users.User)
	return raw
}
