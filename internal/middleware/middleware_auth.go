package middleware

import (
	"fmt"
	"net/http"

	"github.com/fakelozic/go-todo/internal/auth"
	"github.com/fakelozic/go-todo/internal/database"
	"github.com/fakelozic/go-todo/internal/handler"
)

type AuthedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *ApiConfig) MiddlewareAuth(authedHandler AuthedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			handler.ResponseWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey[2])
		if err != nil {
			handler.ResponseWithError(w, 404, fmt.Sprintf("couldn't get user: %v", err))
			return
		}
		if user.Username != apiKey[1] {
			handler.ResponseWithError(w, 403, "api and username didn't match")
			return
		}
		authedHandler(w, r, user)
	}
}
