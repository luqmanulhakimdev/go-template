package middlewares

import (
	"go-template/logger"
	"go-template/response"
	"net/http"

	"github.com/gorilla/mux"
)

func VerifyAPIToken(apiToken string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")

			if token != apiToken {
				logger.Error(r.Context(), "Invalid token ")
				response.RespondError(w, http.StatusUnauthorized, ErrUnauthorizedError)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
