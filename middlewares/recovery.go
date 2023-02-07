package middlewares

import (
	"go-template/logger"
	"go-template/response"
	"net/http"
	"runtime/debug"
)

func Recovery() func(nextHandler http.Handler) http.Handler {
	return func(nextHandler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				err := recover()
				if err != nil {
					logger.Error(r.Context(), "%s", debug.Stack())
					response.RespondError(w, http.StatusInternalServerError, ErrUnknownError)
					return
				}
			}()
			nextHandler.ServeHTTP(w, r)
		})
	}
}
