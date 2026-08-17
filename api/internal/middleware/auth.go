package middleware

import (
	"context"
	"net/http"

	"github.com/YomikaiHub/ohara/api/internal/utils"
)

func AuthMiddleware(secret []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			cookie, err := req.Cookie("yomikai_access")
			if err != nil {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			claims, err := utils.ValidateAccessToken(
				cookie.Value,
				secret,
			)
			if err != nil {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(
				req.Context(),
				utils.UserIDKey,
				claims.UserID,
			)

			ctx = context.WithValue(
				ctx,
				utils.SessionIDKey,
				claims.SessionID,
			)

			next.ServeHTTP(w, req.WithContext(ctx))
		})
	}
}
