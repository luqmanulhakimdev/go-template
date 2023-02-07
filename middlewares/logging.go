package middlewares

import (
	"context"
	"go-template/logger"
	"go-template/response"
	"net/http"
	"time"

	"github.com/google/uuid"
)

const (
	X_CORRELATION_ID = "X-Correlation-ID"
)

func Logging() func(nextHandler http.Handler) http.Handler {
	return func(nextHandler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			corrID := uuid.New()
			corrIDString := r.Header.Get(X_CORRELATION_ID)
			if corrIDString != "" {
				var err error
				corrID, err = uuid.Parse(corrIDString)
				if err != nil {
					logger.Error(r.Context(), "Unauthorized error, invalid correlation ID %s", corrIDString)
					response.RespondError(w, http.StatusUnauthorized, ErrUnauthorizedError)
					return
				}
			}

			ctx := context.WithValue(r.Context(), X_CORRELATION_ID, corrID.String())
			logger.Info(ctx, "%s: %s", r.Method, r.URL.Path)
			start := time.Now()
			r = r.WithContext(ctx)
			w.Header().Set(X_CORRELATION_ID, corrID.String())
			nextHandler.ServeHTTP(w, r)
			logger.Info(r.Context(), "%s: %s took %s", r.Method, r.URL.Path, time.Since(start))
		})
	}
}
